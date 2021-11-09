# holidayGin
Get the day is holiday or not.

[Only support Chinese Mainland.]

[中文文档](https://github.com/xfaifai/holidayGin/blob/main/README_ZH.md)

- ## Usage  
Running (default port 8080)
````
git clone https://github.com/xfaifai/holidayGin.git

go mod tidy

go run main.go
````

Request
````
http://localhost:8080/YYYYMMDD 

YYYYMMDD is which day you need.

Example:
http://localhost:8080/20220101 
````

Response
````
{
    "data": {
        "20220104": 0
    },
    "status": "success"
}

Type:
0 Working day.
1 Double rest day
2 Triple rest day
````


