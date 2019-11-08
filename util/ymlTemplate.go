package util

import (
	"bytes"
	"fmt"
	"github.com/hoisie/mustache"
	"os"
	"os/exec"
	"strings"
)

type PVC struct {
	PVname           string `json:"pvname" form:"pvname"`
	Kind             string `json:"kind" form:"kind"`
	KindStr          string `json:"kindStr" form:"kindStr"`
	Labels           string `json:"labels" form:"labels"`
	Storage          string `json:"storage" form:"storage"`
	StorageClassName string `json:"storageClassName" form:"storageClassName"`
	HostPath         string `json:"hostPath" form:"hostPath"`
	MatchLabels      string `json:"matchLabels" form:"matchLabels"`
}

type PREDICT struct {
	PVCname      string `json:"pvname" form:"pvname"`
	PVCMountPath string `json:"pvcMountPath" form:"pvcMountPath"`
	ModelBsePath string `json:"modelBsePath" form:"modelBsePath"`
	Image        string `json:"image"`
}

type TFJ struct {
	TrainName string    `json:"trainName"`
	Count     string    `json:"count"`
	Replicas  []Replica `json:"replicas"`
}

type Replica struct {
	Type    string `json:"type"`
	Num     string `json:"num"`
	Image   string `json:"image"`
	Command string `json:"command"`
	Gpu     string `json:"gpu"`
	Args    []Arg  `json:"args"`
	VolumeMountname    string  `json:"volumeMountname"`
	MountPath    string  `json:"mountPath"`
	VolumesName    string  `json:"volumesName"`
}

type Arg struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func RenderPVFromFile(pvMap PVC) {
	data := mustache.RenderFile("template/pv.yml.template",
		map[string]string{
			"pvname":           pvMap.PVname,
			"labels":           pvMap.Labels,
			"storage":          pvMap.Storage,
			"storageClassName": pvMap.StorageClassName,
			"hostPath":         pvMap.HostPath,
		})
	print(data)
	pvFileName := fmt.Sprintf("yml/%s.yaml", pvMap.PVname)
	WriteYml(pvFileName, data)
}

func RenderPVCFromFile(pvMap PVC) {
	data := mustache.RenderFile("template/pvc.yml.template",
		map[string]string{
			"pvname":           pvMap.PVname,
			"labels":           pvMap.Labels,
			"storage":          pvMap.Storage,
			"storageClassName": pvMap.StorageClassName,
			"matchLabels":      pvMap.MatchLabels,
		})
	print(data)
	pvFileName := fmt.Sprintf("yml/%s.yaml", pvMap.PVname)
	WriteYml(pvFileName, data)
}

func RenderDeployFromFile(pvMap PREDICT) (string, error) {
	data := mustache.RenderFile("template/deployment.yml.template",
		map[string]string{
			"image":           pvMap.Image,
		})
	print(data)
	pvFileName := fmt.Sprintf("yml/deployment.yaml")
	WriteYml(pvFileName, data)

	dir := "cp template/kustomization.yaml " + WorkPath +"/examples/mnist/serving/local"
	cmd := exec.Command("/bin/bash", "-c", "`"+dir+"`")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return "", err
	}
	fmt.Println("Result: " + out.String())

	return "OK", nil
}


func RenderJob(tfj TFJ)(string, error) {
	rep := tfj.Replicas[0]
	comms := strings.Split(rep.Command, ",")
	data := mustache.RenderFile("template/multy-gpu.yml.template",
		map[string]interface{}{
			"job_name": tfj.TrainName,
			"image": rep.Image,
			"comm":   comms[0],
			"arg": comms[1],
			"gpu": rep.Gpu,
			"args": rep.Args,
			"volumeMountname": rep.VolumeMountname,
			"mountPath": rep.MountPath,
			"valumesNameClaim": rep.VolumesName,
			"valumesName": rep.VolumesName})
	print(data)
	WriteYml("yml/"+tfj.TrainName+".yaml", data)

	cmd := exec.Command("kubectl","apply", "-f","yml/"+tfj.TrainName+".yaml")
	//cmd := exec.Command("/bin/bash", "-c", `cat template/podlog.log`)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return "", err
	}
	fmt.Println("Result: " + out.String())
	return out.String(), nil
}

func RenderChiefromFile(tfjMap TFJ) (string, error) {

	reps := tfjMap.Replicas
	for _, rep := range reps {
		if rep.Type == "1" {
			data := mustache.RenderFile("template/Chief.yml.template",
				map[string]interface{}{
					"count": rep.Num,
					"image": rep.Image,
					"gpu":   rep.Gpu})
			print(data)
			WriteYml("yml/Chief.yaml", data)
		} else if rep.Type == "2" {
			data := mustache.RenderFile("template/Worker.yml.template",
				map[string]interface{}{
					"count": rep.Num,
					"image": rep.Image,
					"gpu":   rep.Gpu})
			print(data)
			WriteYml("yml/Worker.yaml", data)
		} else {
			data := mustache.RenderFile("template/Ps.yml.template",
				map[string]interface{}{
					"count": rep.Num,
					"image": rep.Image})
			print(data)
			WriteYml("yml/Ps.yaml", data)
		}
	}
	//复制文件
	if _, err := MakeTrainDir(tfjMap.TrainName); err != nil {
		return "", err
	}
	if _, err := CpTrainDir(tfjMap.TrainName); err != nil {
		return "", err
	}
	if _, err := CpTrainYaml(tfjMap.TrainName); err != nil {
		return "", err
	}
	if _, err := Kustomize(tfjMap.TrainName, tfjMap.Replicas[0].Args); err != nil {
		return "", err
	}
	return "OK", nil
}

//func RenderFromFileDemo() {
//	data := mustache.RenderFile("template/pv.yml.template",
//
//		map[string]string{
//			"pvname":           pvMap.PVname,
//			"labels":           pvMap.Labels,
//			"storage":          pvMap.Storage,
//			"storageClassName": pvMap.StorageClassName,
//			"hostPath":         pvMap.HostPath,
//		})
//	print(data)
//	pvFileName := fmt.Sprintf("yml/%s.yaml", "yamlDemo.yaml")
//	WriteYml(pvFileName, data)
//}

func WriteYml(fileName string, dataStr string) {
	dstFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	dstFile.WriteString(dataStr)
}
