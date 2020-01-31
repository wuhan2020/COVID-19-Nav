package Http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"nCoV-API/lib/conf"
	"net/http"
)

/**
 * 定义Http接口JOSN响应数据结构
 */
type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

/**
 * 启动一个Http服务
 */
func HttpService() {
	fmt.Println()
	mux := http.NewServeMux()
	// 定义路由
	fmt.Println(conf.Conf.String("http::publicDir"))
	fmt.Println(conf.Conf.String("http::uploadDir"))

	// 最新数据 丁香园
	mux.HandleFunc("/latest/dxy/", handlerLatest)
	mux.HandleFunc("/latest/tx/", handlerTxLatest)
	mux.HandleFunc("/original/dxy/", handlerOrigin)
	mux.HandleFunc("/original/tx/", handlerTxOrigin)
	mux.HandleFunc("/together/", handlerTogether)
	mux.HandleFunc("/nav/pc/", handlerNavPc)
	mux.HandleFunc("/nav/app/", handlerNavApp)

	staticFiles := http.FileServer(http.Dir(conf.Conf.String("http::publicDir")))
	mux.Handle("/", http.StripPrefix("/", staticFiles))

	// 启动http服务
	server := &http.Server{
		Addr:    conf.Conf.String("http::listener"),
		Handler: mux,
	}

	//go handleMessages()
	log.Println("web服务即将启动...")
	var err error
	// 判断是否开启https
	if conf.Conf.String("http::SSLEnable") == "ON" {
		err = server.ListenAndServeTLS(conf.Conf.String("http::SSLCertFile"), conf.Conf.String("http::SSLKeyFile"))
	} else {
		err = server.ListenAndServe()
	}
	if err != nil {
		log.Println("web服务启动失败")
		log.Fatal("listenAndServer : ", err)
	}
	log.Println("web服务启动成功")
}

/**
 * 对象转JSON并响应
 */
func MarshalJson(w http.ResponseWriter, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	// 设置响应头
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, string(data))
	return nil
}

// UnMarshalJson 从request中取出对象
func UnMarshalJson(req *http.Request, v interface{}) error {
	result, err := ioutil.ReadAll(req.Body)
	fmt.Println(req)
	if err != nil {
		return err
	}
	json.Unmarshal([]byte(bytes.NewBuffer(result).String()), v)
	return nil
}
