# holidayGin
判断指定日期是否节假日。

[只支持中国大陆]

- ## Usage  
运行 (默认端口：8080)
````
git clone https://github.com/xfaifai/holidayGin.git

go mod tidy

go run main.go
````

请求
````
http://localhost:8080/YYYYMMDD 

「YYYYMMDD」是你希望得到结果的日期。

例子:
http://localhost:8080/20220101 
````

响应
````
{
    "data": {
        "20220104": 0
    },
    "status": "success"
}

类型:
0 工作日(正常工资)
1 休息日(双倍工资)
2 节假日(三倍工资)
````


