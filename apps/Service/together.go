package Service

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"nCoV-API/lib/conf"
	"nCoV-API/lib/util"
	"net/url"
	"strconv"
	"strings"
)

type TogetherApiRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		ID        int    `json:"id"`
		TDate     string `json:"t_date"`
		TStart    string `json:"t_start"`
		TEnd      string `json:"t_end"`
		TType     int    `json:"t_type"`
		TNo       string `json:"t_no"`
		TMemo     string `json:"t_memo"`
		TNoSub    string `json:"t_no_sub"`
		TPosStart string `json:"t_pos_start"`
		TPosEnd   string `json:"t_pos_end"`
		Source    string `json:"source"`
		Who       string `json:"who"`
		Verified  int    `json:"verified"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	} `json:"data"`
}

type TripsInfoType struct {
	ID                int    `json:"id"`
	Date              string `json:"date"`
	StartTime         string `json:"start"`
	EndTime           string `json:"end"`
	Type              int    `json:"type"`
	TrainNumber       string `json:"train_number"`
	Remark            string `json:"remark"`
	CompartmentNumber string `json:"crriage_number"`
	StartStation      string `json:"start_station"`
	EndStation        string `json:"end_station"`
	Source            string `json:"source"`
	PublishMedia      string `json:"publish_media"`
	Verified          int    `json:"verified"`
	CreateTime        string `json:"create_time"`
	UpdateTime        string `json:"create_time"`
}

var Trips = make(map[string][]TripsInfoType)

func RequestTogetherData() error {
	url := fmt.Sprintf(conf.Conf.String("api::togetherApi"))
	ret, err := util.NewRequest("GET", url, map[string]string{}, nil)
	log.Println()
	var resp TogetherApiRes

	json.Unmarshal(ret, &resp)

	var _resp interface{}
	json.Unmarshal(ret, &_resp)
	if err != nil {
		return err
	}
	if resp.Code != 0 {
		return fmt.Errorf("接口请求失败")
	}

	var together []TripsInfoType
	for _, e := range resp.Data { //e是值拷贝
		var info TripsInfoType
		info.ID = e.ID
		info.Date = e.TDate
		info.StartTime = e.TStart
		info.EndTime = e.TEnd
		info.Type = e.TType
		info.TrainNumber = e.TNo
		info.Remark = e.TMemo
		info.CompartmentNumber = e.TNoSub
		info.StartStation = e.TPosStart
		info.EndStation = e.TPosEnd
		info.Source = e.Source
		info.PublishMedia = e.Who
		info.Verified = e.Verified
		info.CreateTime = e.CreatedAt
		info.UpdateTime = e.UpdatedAt
		together = append(together, info)
	}
	Trips["temp"] = together
	return nil
}

func GetTogetherData(params url.Values) []TripsInfoType {
	v, ok := Trips["temp"]
	if ok == false || v == nil {
		log.Println("请求缓存数据")
		RequestTogetherData()
	}
	// 获取分页参数
	_page := util.GetParam(params, "page", "1").(string)
	log.Println("_page:", _page)
	page, err := strconv.Atoi(_page)
	if err != nil {
		page = 1
	}
	// 每页长度
	_limit := util.GetParam(params, "limit", "25").(string)
	limit, err := strconv.Atoi(_limit)
	if err != nil {
		limit = 25
	}

	var searchParams = make(map[string]string)
	// 获取类型
	_type := util.GetParam(params, "type", "").(string)
	log.Print("type:", _type)
	searchParams["type"] = _type
	// 获取车次
	searchParams["train_number"] = util.GetParam(params, "train_number", "").(string)
	// 获取地区
	searchParams["station"] = util.GetParam(params, "station", "").(string)
	// date
	searchParams["date"] = util.GetParam(params, "date", "").(string)
	trips := searchTrips(Trips["temp"], searchParams)
	RequestTogetherData()
	// 获取数据长度
	count := len(trips)
	log.Println("trips:", Trips["temp"])
	//log.Println("count", count)
	// 获取总页数
	ceil := math.Ceil(util.IntDivCeil(count, limit))
	pageCount := int(ceil)
	var nilTrips []TripsInfoType
	if page > pageCount {
		//page = pageCount
		log.Println("page > pageCount", page, pageCount, trips)
		return nilTrips
	}

	if page < 1 {
		page = 1
		return nilTrips
	}
	start := (page - 1) * limit

	end := start + limit
	if count-start < limit {
		end = count
	}
	// 获取总页数
	return trips[start:end]
}

func searchTrips(trips []TripsInfoType, where map[string]string) []TripsInfoType {
	res := _searchTripsType(trips, where["type"])
	res = _searchTripsTrainNumber(res, where["train_number"])
	res = _searchTripsStation(res, where["station"])
	res = _searchTripsDate(res, where["date"])
	return res
}

func _searchTripsDate(trips []TripsInfoType, s string) []TripsInfoType {
	if s == "" {
		return trips
	}
	var res []TripsInfoType
	for _, item := range trips {
		if item.Date == s {
			res = append(res, item)
		}
	}
	return res
}

func _searchTripsTrainNumber(trips []TripsInfoType, s string) []TripsInfoType {
	if s == "" {
		return trips
	}
	var res []TripsInfoType
	for _, item := range trips {
		if strings.Contains(item.TrainNumber, s) {
			res = append(res, item)
		}
	}
	return res
}

// 站点搜索
func _searchTripsStation(trips []TripsInfoType, s string) []TripsInfoType {
	if s == "" {
		return trips
	}
	var res []TripsInfoType
	for _, item := range trips {
		if strings.Contains(item.StartStation, s) || strings.Contains(item.EndStation, s) {
			res = append(res, item)
		}
	}
	return res
}

func _searchTripsType(trips []TripsInfoType, _type string) []TripsInfoType {
	if _type == "" {
		return trips
	}
	_typeInt, _ := strconv.Atoi(_type)
	var res []TripsInfoType
	for _, item := range trips {
		if item.Type == _typeInt {
			res = append(res, item)
		}
	}
	log.Println("res type search :", res)
	return res

}

//func util.GetParam(params url.Values, name string, defaultValue interface{}) interface{} {
//	value, ok := params[name]
//	if ok == false {
//		return defaultValue
//	}
//	return value[0]
//}
