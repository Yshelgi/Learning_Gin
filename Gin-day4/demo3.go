package main

import (
	"Learn_Gin/Gin-day4/routers"
)

//从Routers包中引入路由

func main() {
	r := routers.SetupRouter()
	r.Run()
}
