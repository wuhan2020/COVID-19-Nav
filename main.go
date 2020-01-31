package main

import (
	"flag"
	"log"
	"nCoV-API/apps/Http"
	"nCoV-API/apps/Model"
	"nCoV-API/apps/Service"
	"nCoV-API/lib/conf"
)


var ConfigPath string


func main() {
	flag.StringVar(&ConfigPath, "config", "./config/app.conf", "配置文件路径")
	flag.Parse()
	log.Println(ConfigPath)
	conf.LoadConfig(ConfigPath)
	Model.Init()
	//Service.UpdateDailySumCountCache("2020-01-30")
	Service.Crond()
	Http.HttpService()
}