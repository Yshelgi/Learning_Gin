package main

import (
	"Learn_Gin/Gin-day4/APP/demo"
	"Learn_Gin/Gin-day4/routers"
)

func main() {
	routers.Include(demo.Routers)
	r := routers.Init()
	r.Run()
}
