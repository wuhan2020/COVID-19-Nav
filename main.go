package main

import (
	"flag"
	"log"
	"nCoV-API/lib/conf"
	"nCoV-API/apps/Http"
	"nCoV-API/apps/Model"
	"nCoV-API/apps/Service"
)


var ConfigPath string


func main() {
	flag.StringVar(&ConfigPath, "config", "./config/app.conf", "配置文件路径")
	flag.Parse()
	log.Println(ConfigPath)
	conf.LoadConfig(ConfigPath)
	Model.Init()
	Service.Crond()
	Http.HttpService()
}