package main

import (
	"github.com/gin-gonic/gin"
	"kbf-go/util"
	"net/http"
	"os"
)

func execRouter() {
	engine := gin.Default()

	engine.GET("/uipod", func(c *gin.Context) {
		res, err := util.UIpod()
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}
		print(res)
		c.JSON(200, gin.H{
			"message": res,
		})
	})

	engine.GET("/get-job", func(c *gin.Context) {
		res, err := util.GetJob()
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}
		print(res)
		c.JSON(200, gin.H{
			"message": res,
		})
	})

	engine.GET("/gpu", func(c *gin.Context) {
		res, err := util.LookGpu()
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}
		print(res)
		c.JSON(200, gin.H{
			"message": res,
		})
	})

	engine.GET("/configmap", func(c *gin.Context) {
		res, err := util.Configmap()
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}
		print(res)
		c.JSON(200, gin.H{
			"message": res,
		})
	})

	engine.GET("/configmap-remove/:configName", func(c *gin.Context) {
		configName := c.Param("configName")
		res, err := util.RemoveConfigmap(configName)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}
		print(res)
		c.JSON(200, gin.H{
			"message": res,
		})
	})

	engine.GET("/get-pv-pro", func(c *gin.Context) {
		res, err := util.ExecGetPV()
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}
		print(res)
		c.JSON(200, gin.H{
			"message": res,
		})
	})

	engine.GET("/exec-pv/:pvname", func(c *gin.Context) {
		pvname := c.Param("pvname")
		res, err := util.ExecCreatePV(pvname)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}
		print(res)
		c.JSON(200, gin.H{
			"message": res,
		})
	})

	engine.GET("/tfjob", func(c *gin.Context) {
		res, err := util.GetTfjob()
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}
		print(res)
		c.JSON(200, gin.H{
			"message": res,
		})
	})

	engine.GET("/get-tfj-log/:tfj/:namespace", func(c *gin.Context) {
		podname := c.Param("tfj")
		namespace := c.Param("namespace")
		res, err := util.GetTfjLog(podname, namespace)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}
		print(res)
		c.JSON(200, gin.H{
			"message": res,
		})
	})

	engine.GET("/get-tfj-pod/:tfj/:namespace", func(c *gin.Context) {
		podname := c.Param("tfj")
		namespace := c.Param("namespace")
		res, err := util.GetTfjPod(podname, namespace)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}
		print(res)
		c.JSON(200, gin.H{
			"message": res,
		})
	})

	engine.GET("/get-job-pod/:tfj/:namespace", func(c *gin.Context) {
		podname := c.Param("tfj")
		namespace := c.Param("namespace")
		res, err := util.GetJobPod(podname, namespace)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}
		print(res)
		c.JSON(200, gin.H{
			"message": res,
		})
	})

	engine.GET("/delete-tfj/:tfj/:namespace", func(c *gin.Context) {
		podname := c.Param("tfj")
		namespace := c.Param("namespace")
		res, err := util.DeleteTfj(podname, namespace)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}
		print(res)
		c.JSON(200, gin.H{
			"message": res,
		})
	})

	engine.GET("/delete-job/:tfj/:namespace", func(c *gin.Context) {
		podname := c.Param("tfj")
		namespace := c.Param("namespace")
		res, err := util.DeleteJob(podname, namespace)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}
		print(res)
		c.JSON(200, gin.H{
			"message": res,
		})
	})

	engine.GET("/delete-predict/:namespace", func(c *gin.Context) {
		namespace := c.Param("namespace")
		res, err := util.DeletePreDict(namespace)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}
		print(res)
		c.JSON(200, gin.H{
			"message": res,
		})
	})

	engine.POST("/create-pv", func(c *gin.Context) {
		var jsonPVC util.PVC
		if err := c.BindJSON(&jsonPVC); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		if jsonPVC.Kind == "1" {
			util.RenderPVFromFile(jsonPVC)
		} else {
			util.RenderPVCFromFile(jsonPVC)
		}
		c.JSON(http.StatusOK, gin.H{"status": "pv/pvc yaml file create success!!!"})
	})

	engine.POST("/create-predict", func(c *gin.Context) {
		var jsonPREDICT util.PREDICT
		if err := c.BindJSON(&jsonPREDICT); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}

		if _, err := util.RenderDeployFromFile(jsonPREDICT); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}

		if _, err := util.KustomizePredict(jsonPREDICT); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "pv/pvc yaml file create success!!!"})
	})

	engine.POST("/deploy-tfj", func(c *gin.Context) {
		var jsonTFJ util.TFJ
		if err := c.BindJSON(&jsonTFJ); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		println("#####################")
		if _, err := util.RenderChiefromFile(jsonTFJ); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "TFJ yaml file create success!!!"})
	})

	engine.POST("/create-job", func(c *gin.Context) {
		var jsonTFJ util.TFJ
		if err := c.BindJSON(&jsonTFJ); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		println("#####################")
		if _, err := util.RenderJob(jsonTFJ); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "TFJ yaml file create success!!!"})
	})

	engine.Static("/html", "./html")

	engine.Run(":9311") // listen and serve on 0.0.0.0:8080
}

func main() {

	if len(os.Args) <= 1 {
		util.WorkPath = "/u01/yangyue5"
		println("==================================" + util.WorkPath)
	} else {
		util.WorkPath = os.Args[1]
		println("###################################" + util.WorkPath)
	}


	execRouter()
}
