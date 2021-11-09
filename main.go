package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/:date", func(c *gin.Context) {
		date := c.Param("date")
		t,err := time.Parse("20060102", date)
		if err != nil {
			c.JSON(200,gin.H{"status":"error", "message":"时间格式不对"})
			return
		}

		year := t.Format("2006")
		jsonPath := "./data/" + year + ".json"
		jsonFile,err := os.Open(jsonPath)
		if err != nil {
			c.JSON(200,gin.H{"status":"error", "message":"暂未支持的年份"})
			return
		}

		defer jsonFile.Close()

		byteValue,_ := ioutil.ReadAll(jsonFile)
		var jsonData map[string]interface{}
		json.Unmarshal([]byte(byteValue),&jsonData)

		day := t.Format("0102")
		value,ok := jsonData[day]
		if ok {
			c.JSON(200,gin.H{"status":"success", "data": gin.H{date:value}})
			return
		}

		w := int(t.Weekday())
		if w==6 || w==0 {
			c.JSON(200,gin.H{"status":"success", "data": gin.H{date:1}})
		} else {
			c.JSON(200,gin.H{"status":"success", "data": gin.H{date:0}})
		}
	})

	r.Run(":8080")
}
