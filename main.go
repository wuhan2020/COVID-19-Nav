package main

import (
	"nCoV-API/apps/Http"
	"nCoV-API/apps/Service"
	"nCoV-API/lib/conf"
)



func main() {
	conf.LoadConfig("./config/app.conf")
	Service.Crond()
	Http.HttpService()
}