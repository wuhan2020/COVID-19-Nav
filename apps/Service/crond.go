package Service

import (
	"log"
	"time"
)

//定时任务
func CrontabFunc(d time.Duration, handler func() error, desc string) {
	for {
		log.Println("crontab func runing ", desc)
		handler()
		time.Sleep(d)
	}
}

func Crond() {
	go CrontabFunc(time.Second*30, RequestLatestData, " 刷新开源数据接口缓存")
	go CrontabFunc(time.Second*30, RequestTxApiData, " 刷新天行数据接口缓存")
	go CrontabFunc(time.Second*300, RequestTogetherData, " 刷新患者同行数据缓存")
	go CrontabFunc(time.Second*30, RequestNavData, " 刷新导航信息缓存")
}
