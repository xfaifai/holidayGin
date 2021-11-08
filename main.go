package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/:date", func(c *gin.Context) {
		date := c.Param("date")
		t,err := time.Parse("20060102", date)
		if err != nil {
			c.JSON(200,gin.H{
				"status":"error",
				"message":"时间格式不对",
			})
			return
		}

		year,_ := strconv.Atoi(t.Format("2006"))
		if year < 2010 {
			c.JSON(200,gin.H{
				"status":"error",
				"message":"暂未支持的年份",
			})
			return
		}

		jsonPath := "./data/"+strconv.Itoa(year)+".json"
		jsonFile,err := os.Open(jsonPath)
		if err != nil {
			c.JSON(200,gin.H{
				"status":"error",
				"message":"暂未支持的年份",
			})
			return
		}

		defer jsonFile.Close()

		byteValue,_ := ioutil.ReadAll(jsonFile)

		var jsonData map[string]interface{}
		json.Unmarshal([]byte(byteValue),&jsonData)

		day := t.Format("0102")
		value,ok := jsonData[day]
		if ok {
			var res map[string]interface{}
			res[date] = value
			c.JSON(200,gin.H{
				"status":"success",
				"data": res,
			})
			return
		}
	})
	r.Run()
}
