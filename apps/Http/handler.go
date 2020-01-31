package Http

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"nCoV-API/apps/Service"
	"net/http"
	"time"
)

// 获取最新数据
func handlerLatest(writer http.ResponseWriter, request *http.Request) {
	var resp Resp
	defer MarshalJson(writer, &resp)
	data := Service.GetLatestData()
	log.Print("get latest api request")
	resp.Data = data
}

// 处理原版数据请求
func handlerOrigin(writer http.ResponseWriter, request *http.Request) {
	data := Service.GetOriginalLatestData()
	log.Print("get origin latest api request")
	MarshalJson(writer, data)
}

func handlerTxLatest(writer http.ResponseWriter, request *http.Request) {
	var resp Resp
	defer MarshalJson(writer, &resp)
	data := Service.GetTxApiData()
	log.Print("get tx api request")
	resp.Data = data
}

func handlerTxOrigin(writer http.ResponseWriter, request *http.Request) {
	data := Service.GetOriginalTxApiData()
	log.Print("get origin txapi api request")
	MarshalJson(writer, data)
}

func handlerTogether(writer http.ResponseWriter, request *http.Request) {
	var resp Resp
	defer MarshalJson(writer, &resp)
	request.ParseForm()
	form := request.Form
	data := Service.GetTogetherData(form)
	log.Print("get together api request")
	resp.Data = data
}

// 获取导航信息 移动端
func handlerNavApp(writer http.ResponseWriter, request *http.Request) {
	data := Service.GetNavData()
	log.Print("get origin nav api request")
	MarshalJson(writer, data)
}

// 获取导航信息 pc端
func handlerNavPc(writer http.ResponseWriter, request *http.Request) {
	data := Service.GetNavData()
	log.Print("request nav api ")
	MarshalJson(writer, data)
}

func handlerDailySum(writer http.ResponseWriter, request *http.Request) {
	var resp Resp
	defer MarshalJson(writer, &resp)
	request.ParseForm()
	form := request.Form
	data, err := Service.GetLatestDailySumInfoByCache(form)
	log.Print("get together api request")
	if err != nil {
		resp.Data = err
		return
	}
	resp.Data = data
}

/**
 * 生成一个md5
 */
func CreateMd5(str string) string {
	//sessionId
	m := md5.New()
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	io.WriteString(m, str+"_image_crop_"+timestamp+"_"+str)
	return fmt.Sprintf("%x", m.Sum(nil))
}
