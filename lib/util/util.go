package util;

import (
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
	"crypto/md5"
	"encoding/hex"
)

var c = &http.Client{
	Transport: &http.Transport{
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, time.Second*1)
		},
		MaxIdleConns:        500,
		MaxIdleConnsPerHost: 800,
	},
}

func NewRequest(method, path string, headers, params map[string]string) ([]byte, error) {
	req, _ := http.NewRequest(method, path, nil)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	p := make(url.Values, len(params))
	for k, v := range params {
		p[k] = []string{v}
	}
	req.Body = ioutil.NopCloser(strings.NewReader(p.Encode()))
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func GetMd5(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}


func SliceToMap(keys, values []string) map[string]string {
	mapObj := map[string]string{}
	for index := range keys {
		mapObj[keys[index]] = values[index]
	}

	return mapObj
}


func GetParam(params url.Values, name string, defaultValue interface{}) interface{} {
	value, ok := params[name]
	if ok == false {
		return defaultValue
	}
	return value[0]
}