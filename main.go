package main

import "fmt"
import "github.com/StefanAtanaskovic/gocpu/gocpu"
import "github.com/gin-gonic/gin"

func main() {
	fmt.Println("hello")

	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		val, err := gocpu.GetCpuUsage()

		if err != nil {
			ctx.JSON(200, gin.H{
				"error": "Failed to get data",
			})
		}

		ctx.JSON(200, gin.H{
			"cpuUsage": val,
		})
	})

	r.Run()
}
