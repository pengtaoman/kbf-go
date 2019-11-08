package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

var WorkPath string

func ExecGetPV() (string, error) {
	//production
	cmd := exec.Command("kubectl", "get", "pv", "-n" ,"kubeflow","-o" ,"json")
	//cmd := exec.Command("/bin/bash", "-c", `cat template/demoData.json`)
	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		return "", err
	}

	//执行命令
	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
		return "", err
	}

	//读取所有输出
	res, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll Stdout:", err.Error())
		return "", err
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("wait:", err.Error())
		return "", err
	}

	fmt.Printf("stdout:\n\n %s", res)
	return string(res), nil
}

func UIpod() (string, error) {
	//production
	cmd := exec.Command("kubectl","get", "pods", "--all-namespaces","-l","app=web-ui", "-o","json")
	//cmd := exec.Command("/bin/bash", "-c", `cat template/uipod.json`)
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

func GetJob() (string, error) {
	//production
	cmd := exec.Command("kubectl","get", "job", "-o","json")
	//cmd := exec.Command("/bin/bash", "-c", `cat template/getJob.json`)
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

func LookGpu() (string, error) {
	//production
	cmd := exec.Command("nvidia-smi")
	//cmd := exec.Command("/bin/bash", "-c", `cat template/gpu.info`)

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

func Configmap() (string, error) {
	//kubectl get configmap mnist-map-training-576dc56kg7 -o json
	cmd := exec.Command("kubectl","get","configmap","-o","json")
	//cmd := exec.Command("/bin/bash", "-c", `cat template/gpu.info`)

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

func RemoveConfigmap(cname string) (string, error) {

	cmd := exec.Command("kubectl", "delete", "configmap", cname)
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

func ExecCreatePV(pvname string) (string, error) {

	pvFileName := fmt.Sprintf("yml/%s.yaml", pvname)
	println(pvFileName)
	cmd := exec.Command("kubectl", "apply", "-f", pvFileName)

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

func GetTfjob() (string, error) {
	//production
	cmd := exec.Command("kubectl","get", "tfjob", "--all-namespaces","-o","json")
	//cmd := exec.Command("/bin/bash", "-c", `cat template/tfj.json`)
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

func GetTfjPod(podname string, namespace string) (string, error) {
	//production
	sele := fmt.Sprintf("tf-job-name=%s", podname)
	println(sele)
	cmd := exec.Command("kubectl","get", "pods", "--namespace",namespace, "-l",sele, "-o","json")
	//cmd := exec.Command("/bin/bash", "-c", `cat template/tfjpod.json`)
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


func GetJobPod(podname string, namespace string) (string, error) {
	//production
	sele := fmt.Sprintf("job-name=%s", podname)
	println(sele)
	cmd := exec.Command("kubectl","get", "pods", "--namespace",namespace, "-l",sele, "-o","json")
	//cmd := exec.Command("/bin/bash", "-c", `cat template/job-pod.json`)
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

func GetTfjLog(podname string, namespace string) (string, error) {
	//production
	cmd := exec.Command("kubectl","logs", podname)
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

func DeleteTfj(podname string, namespace string) (string, error) {
	//production
	println(podname)
	println(namespace)
	cmd := exec.Command("kubectl", "delete", "tfjob", podname)
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

func DeleteJob(podname string, namespace string) (string, error) {
	//production
	println(podname)
	println(namespace)
	cmd := exec.Command("kubectl", "delete", "job", podname)
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

func DeletePreDict(nameSpace string) (string, error) {

	cmds := [][]string{
		{"delete", "deployment","web-ui"},
		{"delete","deployment","mnist-service-local"},
		{"delete", "svc","web-ui"},
		{"delete", "svc","mnist-service-local"},
		{"delete", "configmap","mnist-deploy-config"},
	}

	for inx := range cmds {
		cmdEve := exec.Command("kubectl", cmds[inx]...)
		var out1 bytes.Buffer
		var stderr1 bytes.Buffer
		cmdEve.Stdout = &out1
		cmdEve.Stderr = &stderr1
		err1 := cmdEve.Run()
		if err1 != nil {
			fmt.Println(fmt.Sprint(err1) + ": " + stderr1.String())
			return "", err1
		}
		fmt.Println("Result: " + out1.String())
	}
	return "OK", nil
}

func MakeTrainDir(trainName string) (string, error) {
	//production
	dir := "mkdir -p "+WorkPath+"/examples/mnist/" + trainName
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
	return out.String(), nil
}

func CpTrainDir(trainName string) (string, error) {
	//production
	dir := "cp -r "+WorkPath+"/examples/mnist/training/{base,local} " + WorkPath+"/examples/mnist/" + trainName + "/"
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
	return out.String(), nil
}

func CpTrainYaml(trainName string) (string, error) {
	//production
	dir := WorkPath+"/examples/mnist/" + trainName
	cmd := exec.Command("/bin/bash", "-c", "`cp yml/{Chief.yaml,Worker.yaml,Ps.yaml} "+dir+"/base`")
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

func Kustomize(trainName string, kvs []Arg) (string, error) {

	cmd := exec.Command("kustomize", "edit", "add", "configmap", "mnist-map-training", "--from-literal=name="+trainName)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	cmd.Dir = WorkPath+"/examples/mnist/" + trainName + "/local"
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return "", err
	}
	fmt.Println("Result: " + out.String())

	for _, kv := range kvs {
		cmd := exec.Command("kustomize", "edit", "add", "configmap", "mnist-map-training", "--from-literal="+kv.Key+"="+kv.Value)
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		cmd.Dir = WorkPath+"/examples/mnist/" + trainName + "/local"
		err := cmd.Run()
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
			return "", err
		}
		fmt.Println("Result: " + out.String())
	}

	cmd22 := exec.Command("pwd")
	var out22 bytes.Buffer
	var stderr22 bytes.Buffer
	cmd22.Stdout = &out22
	cmd22.Stderr = &stderr22
	cmd22.Dir = WorkPath+"/examples/mnist/" + trainName + "/local"
	err22 := cmd22.Run()
	if err22 != nil {
		fmt.Println(fmt.Sprint(err22) + ": " + stderr22.String())
		return "", err22
	}
	fmt.Println("Result: " + out22.String())
	//kustomize build . |kubectl apply -f -  ///   "|kubectl apply -f -"
	cmd11 := exec.Command("kustomize", "build", WorkPath+"/examples/mnist/"+trainName+"/local")
	var out11 bytes.Buffer
	var stderr11 bytes.Buffer
	cmd11.Stdout = &out11
	cmd11.Stderr = &stderr11
	cmd11.Dir = WorkPath+"/examples/mnist/" + trainName + "/local"
	err11 := cmd11.Run()
	if err11 != nil {
		fmt.Println(fmt.Sprint(err11) + ": " + stderr11.String())
		return "", err11
	}
	fmt.Println("Result: " + out11.String())

	dstFile, err555 := os.Create(WorkPath+"/examples/mnist/" + trainName + "/local/output.yaml")
	if err555 != nil {
		fmt.Println(err555.Error())
		return "", err555
	}
	defer dstFile.Close()
	dstFile.WriteString(out11.String())

	//|kubectl apply -f -
	cmd33 := exec.Command("kubectl", "apply", "-f", "output.yaml")
	var out33 bytes.Buffer
	var stderr33 bytes.Buffer
	cmd33.Stdout = &out33
	cmd33.Stderr = &stderr33
	cmd33.Dir = WorkPath+"/examples/mnist/" + trainName + "/local"
	err33 := cmd33.Run()
	if err33 != nil {
		fmt.Println(fmt.Sprint(err33) + ": " + stderr33.String())
		return "", err33
	}
	fmt.Println("Result: " + out33.String())

	return "OK", nil
}

func KustomizePredict(predict PREDICT) (string, error) {

	//add example/mnist/serving/local/ kustomization.yaml
	cmd := exec.Command("kustomize", "edit", "add", "configmap", "mnist-map-serving", "--from-literal=name=mnist-service-local")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	cmd.Dir = WorkPath+"/examples/mnist/serving/local"
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return "", err
	}
	fmt.Println("Result: " + out.String())

	cmd11 := exec.Command("kustomize", "edit", "add", "configmap", "mnist-map-serving", "--from-literal=pvcName="+predict.PVCname)
	var out11 bytes.Buffer
	var stderr11 bytes.Buffer
	cmd11.Stdout = &out11
	cmd11.Stderr = &stderr11
	cmd11.Dir = WorkPath+"/examples/mnist/serving/local"
	err11 := cmd11.Run()
	if err11 != nil {
		fmt.Println(fmt.Sprint(err11) + ": " + stderr11.String())
		return "", err11
	}
	fmt.Println("Result: " + out11.String())

	cmd22 := exec.Command("kustomize", "edit", "add", "configmap", "mnist-map-serving", "--from-literal=pvcMountPath="+predict.PVCMountPath)
	var out22 bytes.Buffer
	var stderr22 bytes.Buffer
	cmd22.Stdout = &out22
	cmd22.Stderr = &stderr22
	cmd22.Dir = WorkPath+"/examples/mnist/serving/local"
	err22 := cmd22.Run()
	if err22 != nil {
		fmt.Println(fmt.Sprint(err22) + ": " + stderr22.String())
		return "", err22
	}
	fmt.Println("Result: " + out22.String())

	cmd33 := exec.Command("kustomize", "edit", "add", "configmap", "mnist-map-serving", "--from-literal=modelBasePath="+predict.ModelBsePath)
	var out33 bytes.Buffer
	var stderr33 bytes.Buffer
	cmd33.Stdout = &out33
	cmd33.Stderr = &out33
	cmd33.Dir = WorkPath+"/examples/mnist/serving/local"
	err33 := cmd33.Run()
	if err33 != nil {
		fmt.Println(fmt.Sprint(err33) + ": " + stderr33.String())
		return "", err33
	}
	fmt.Println("Result: " + out33.String())

	cmd44 := exec.Command("kustomize", "build", ".")
	var out44 bytes.Buffer
	var stderr44 bytes.Buffer
	cmd44.Stdout = &out44
	cmd44.Stderr = &out44
	cmd44.Dir = WorkPath+"/examples/mnist/serving/local"
	err44 := cmd44.Run()
	if err44 != nil {
		fmt.Println(fmt.Sprint(err44) + ": " + stderr44.String())
		return "", err44
	}
	fmt.Println("Result: " + out44.String())

	//,"|","kubectl", "apply","-f","-"
	dstFile, err555 := os.Create(WorkPath+"/examples/mnist/serving/local/output.yaml")
	if err555 != nil {
		fmt.Println(err555.Error())
		return "", err555
	}
	defer dstFile.Close()
	dstFile.WriteString(out44.String())

	cmd333 := exec.Command("kubectl", "apply", "-f", "output.yaml")
	var out333 bytes.Buffer
	var stderr333 bytes.Buffer
	cmd333.Stdout = &out333
	cmd333.Stderr = &stderr333
	cmd333.Dir = WorkPath+"/examples/mnist/serving/local"
	err333 := cmd333.Run()
	if err33 != nil {
		fmt.Println(fmt.Sprint(err333) + ": " + stderr333.String())
		return "", err333
	}
	fmt.Println("Result: " + out333.String())

	cmd55 := exec.Command("kubectl", "apply", "-f", "service.yaml")
	var out55 bytes.Buffer
	var stderr55 bytes.Buffer
	cmd55.Stdout = &out55
	cmd55.Stderr = &stderr55
	cmd55.Dir = WorkPath+"/examples/mnist/front"
	err55 := cmd55.Run()
	if err55 != nil {
		fmt.Println(fmt.Sprint(err55) + ": " + stderr55.String())
		return "", err55
	}
	fmt.Println("Result: " + out55.String())


	cmd666 := exec.Command("/bin/bash", "-c", "`cp yml/deployment.yaml "+WorkPath+"/examples/mnist/front`")
	var out666 bytes.Buffer
	var stderr666 bytes.Buffer
	cmd666.Stdout = &out666
	cmd666.Stderr = &stderr666
	err666 := cmd666.Run()
	if err666 != nil {
		fmt.Println(fmt.Sprint(err666) + ": " + stderr666.String())
		return "", err666
	}
	fmt.Println("Result: " + out666.String())

	cmd66 := exec.Command("kubectl", "apply", "-f", "deployment.yaml")
	var out66 bytes.Buffer
	var stderr66 bytes.Buffer
	cmd66.Stdout = &out66
	cmd66.Stderr = &stderr66
	cmd66.Dir = WorkPath+"/examples/mnist/front"
	err66 := cmd66.Run()
	if err66 != nil {
		fmt.Println(fmt.Sprint(err66) + ": " + stderr66.String())
		return "", err66
	}
	fmt.Println("Result: " + out66.String())

	return "OK", nil
}

//func TransJson(res []byte) string {
//	lines := bytes.Split(res, []byte("\n"))
//	println("?????????????????? %s", len(lines))
//	println("?????????????????? %s", string(lines[0]))
//	println()
//	println()
//    inx := 0
//	for _,line := range lines {
//		if string(line) != "" {
//			println(">>>>>>>>>>>>>>>>>>", string(line))
//			if inx == 0 {
//				titleArr := bytes.Split(line, []byte(","))
//				println(">>> titles len", len(titleArr))
//				println(">>> titles", string(titleArr[1]))
//			}
//			inx++
//		}
//	}
//
//	return ""
//}
