package main

import (
	"myproject/database"
	"myproject/routers"
)

func main() {
	database.InitMysql()
	router := routers.InitRouter()

	//静态资源
	router.Static("/static", "static")
	//router.StaticFS("/static",http.Dir("D:/awesomeProject/awesomeProject2/ginweb-master/blogweb_gin/static/upload"))
	//router.StaticFS("/static",http.Dir("D:/awesomeProject/awesomeProject2/ginweb-master/blogweb_gin/static"))
	//router.StaticFS("/static",http.Dir("D:/awesomeProject/awesomeProject2/ginweb-master/blogweb_gin/static/upload"))

	router.Run(":8080")
}
