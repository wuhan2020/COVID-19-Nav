package Service

import (
	"encoding/json"
	"fmt"
	"log"
	"nCoV-API/lib/conf"
	"nCoV-API/lib/util"
)

type LatestData struct {
	Count struct {
		Confirmed     int `json:"confirmed"`
		Suspected     int `json:"suspected"`
		Cure          int `json:"cure"`
		Death         int `json:"death"`
		Serious       int `json:"serious"`
		SuspectedIncr int `json:"suspected_incr"`
		ConfirmedIncr int `json:"confirmed_incr"`
		CuredIncr     int `json:"cured_incr"`
		DeadIncr      int `json:"dead_incr"`
		SeriousIncr   int `json:"serious_incr"`
	} `json:"count"`
	Info struct {
		// 疫情名称
		Name string `json:"name"`
		// 易感人群
		VulnerableGroup string `json:"vulnerable_group"`
		// 潜伏期
		IncubationPeriod string `json:"incubation_period"`
		// 数据源
		DataSource string `json:"data_source"`
		// 传染源
		InfectionSource string `json:"infection_source"`
		// 传播路径
		Transmission string `json:"transmission"`
	} `json:"info"`
	// 数据更新时间
	UpdateTime int64 `json:"update_time"`
}

type NcovApiRes struct {
	Results []struct {
		AbroadRemark   string `json:"abroadRemark"`
		Confirmed      int    `json:"confirmed"`
		ConfirmedCount int    `json:"confirmedCount"`
		CountRemark    string `json:"countRemark"`
		Cured          int    `json:"cured"`
		CuredCount     int    `json:"curedCount"`
		DailyPic       string `json:"dailyPic"`
		DeadCount      int    `json:"deadCount"`
		Death          int    `json:"death"`
		GeneralRemark  string `json:"generalRemark"`
		InfectSource   string `json:"infectSource"`
		PassWay        string `json:"passWay"`
		Remark1        string `json:"remark1"`
		Remark2        string `json:"remark2"`
		Remark3        string `json:"remark3"`
		Remark4        string `json:"remark4"`
		Remark5        string `json:"remark5"`
		Summary        string `json:"summary"`
		Suspect        int    `json:"suspect"`
		SuspectedCount int    `json:"suspectedCount"`
		UpdateTime     int64  `json:"updateTime"`
		Virus          string `json:"virus"`
	} `json:"results"`
	Success bool `json:"success"`
}

type TxApiRes struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	Newslist []struct {
		News []struct {
			ID               int    `json:"id"`
			PubDate          int64  `json:"pubDate"`
			PubDateStr       string `json:"pubDateStr"`
			Title            string `json:"title"`
			Summary          string `json:"summary"`
			InfoSource       string `json:"infoSource"`
			SourceURL        string `json:"sourceUrl"`
			ProvinceID       string `json:"provinceId"`
			ProvinceName     string `json:"provinceName"`
			CreateTime       int64  `json:"createTime"`
			ModifyTime       int64  `json:"modifyTime"`
			EntryWay         int    `json:"entryWay"`
			AdoptType        int    `json:"adoptType"`
			InfoType         int    `json:"infoType"`
			DataInfoState    int    `json:"dataInfoState"`
			DataInfoOperator string `json:"dataInfoOperator"`
			DataInfoTime     int64  `json:"dataInfoTime"`
		} `json:"news"`
		Case []struct {
			ID                int    `json:"id"`
			CreateTime        int64  `json:"createTime"`
			ModifyTime        int64  `json:"modifyTime"`
			Tags              string `json:"tags"`
			CountryType       int    `json:"countryType"`
			Continents        string `json:"continents"`
			ProvinceID        string `json:"provinceId"`
			ProvinceName      string `json:"provinceName"`
			ProvinceShortName string `json:"provinceShortName"`
			CityName          string `json:"cityName"`
			ConfirmedCount    int    `json:"confirmedCount"`
			SuspectedCount    int    `json:"suspectedCount"`
			CuredCount        int    `json:"curedCount"`
			DeadCount         int    `json:"deadCount"`
			Comment           string `json:"comment"`
			Sort              int    `json:"sort"`
			Operator          string `json:"operator"`
			LocationID        int    `json:"locationId"`
		} `json:"case"`
		Desc struct {
			ID             int           `json:"id"`
			CreateTime     int64         `json:"createTime"`
			ModifyTime     int64         `json:"modifyTime"`
			InfectSource   string        `json:"infectSource"`
			PassWay        string        `json:"passWay"`
			ImgURL         string        `json:"imgUrl"`
			DailyPic       string        `json:"dailyPic"`
			DailyPics      []string      `json:"dailyPics"`
			Summary        string        `json:"summary"`
			Deleted        bool          `json:"deleted"`
			CountRemark    string        `json:"countRemark"`
			ConfirmedCount int           `json:"confirmedCount"`
			SuspectedCount int           `json:"suspectedCount"`
			CuredCount     int           `json:"curedCount"`
			DeadCount      int           `json:"deadCount"`
			SeriousCount   int           `json:"seriousCount"`
			SuspectedIncr  int           `json:"suspectedIncr"`
			ConfirmedIncr  int           `json:"confirmedIncr"`
			CuredIncr      int           `json:"curedIncr"`
			DeadIncr       int           `json:"deadIncr"`
			SeriousIncr    int           `json:"seriousIncr"`
			Virus          string        `json:"virus"`
			Remark1        string        `json:"remark1"`
			Remark2        string        `json:"remark2"`
			Remark3        string        `json:"remark3"`
			Remark4        string        `json:"remark4"`
			Remark5        string        `json:"remark5"`
			Note1          string        `json:"note1"`
			Note2          string        `json:"note2"`
			Note3          string        `json:"note3"`
			GeneralRemark  string        `json:"generalRemark"`
			AbroadRemark   string        `json:"abroadRemark"`
			Marquee        []interface{} `json:"marquee"`
		} `json:"desc"`
	} `json:"newslist"`
}
type NavItemType struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Image string `json:"image"`
	URL   string `json:"url"`
}

type RecommendationApiType struct {
	Code int           `json:"code"`
	Data []NavItemType `json:"data"`
}

type NavInfoType struct {
	Code int `json:"code"`
	Data []struct {
		Title string        `json:"title"`
		Item  []NavItemType `json:"item"`
	} `json:"data"`
}

// 全局缓存map
var Latest = make(map[string]LatestData)
var Original = make(map[string]interface{})
var Nav = make(map[string]NavInfoType)
var Recommendation = make(map[string]RecommendationApiType)

// 获取最新数据
func GetLatestData() LatestData {
	_, ok := Latest["latest"]
	if ok == false {
		RequestLatestData()
	}
	return Latest["latest"]
}

// 获取原版输出
func GetOriginalLatestData() interface{} {
	_, ok := Original["latest"]
	if ok == false {
		RequestLatestData()
	}
	return Original["latest"]
}

func GetTxApiData() LatestData {
	_, ok := Latest["txApi"]

	if ok == false {
		RequestTxApiData()
	}
	return Latest["txApi"]
}

func GetOriginalTxApiData() interface{} {
	_, ok := Original["txApi"]
	if ok == false {
		RequestTxApiData()
	}
	return Original["txApi"]
}

// 请求最新数据
func RequestLatestData() error {
	url := fmt.Sprintf(conf.Conf.String("api::ncovapi"))
	ret, err := util.NewRequest("GET", url, map[string]string{}, nil)
	var resp NcovApiRes
	json.Unmarshal(ret, &resp)
	if err != nil {
		return err
	}
	if resp.Success == false {
		return fmt.Errorf("接口请求失败")
	}
	var latestData LatestData
	//make()
	latestData.Count.Confirmed = resp.Results[0].ConfirmedCount
	latestData.Count.Cure = resp.Results[0].CuredCount
	latestData.Count.Suspected = resp.Results[0].SuspectedCount
	latestData.Count.Death = resp.Results[0].DeadCount
	latestData.Info.DataSource = resp.Results[0].GeneralRemark
	latestData.Info.IncubationPeriod = resp.Results[0].Remark2
	latestData.Info.VulnerableGroup = resp.Results[0].Remark1
	latestData.Info.InfectionSource = resp.Results[0].InfectSource
	latestData.Info.Name = resp.Results[0].Virus
	latestData.UpdateTime = resp.Results[0].UpdateTime
	Latest["latest"] = latestData
	Original["latest"] = resp

	return nil
}

func RequestTxApiData() error {
	url := fmt.Sprintf(conf.Conf.String("api::txApi"))
	ret, err := util.NewRequest("GET", url, map[string]string{}, nil)
	var resp TxApiRes
	json.Unmarshal(ret, &resp)
	if err != nil {
		log.Println(err)
		return err
	}
	if resp.Code != 200 {
		return nil
	}
	var latestData LatestData
	//make()
	latestData.Count.Confirmed = resp.Newslist[0].Desc.ConfirmedCount
	latestData.Count.Cure = resp.Newslist[0].Desc.CuredCount
	latestData.Count.Suspected = resp.Newslist[0].Desc.SuspectedCount
	latestData.Count.Death = resp.Newslist[0].Desc.DeadCount
	latestData.Count.Serious = resp.Newslist[0].Desc.SeriousCount
	latestData.Count.SeriousIncr = resp.Newslist[0].Desc.SeriousIncr
	latestData.Count.SuspectedIncr = resp.Newslist[0].Desc.SuspectedIncr
	latestData.Count.ConfirmedIncr = resp.Newslist[0].Desc.ConfirmedIncr
	latestData.Count.CuredIncr = resp.Newslist[0].Desc.CuredIncr
	latestData.Count.DeadIncr = resp.Newslist[0].Desc.DeadIncr
	latestData.Info.DataSource = resp.Newslist[0].Desc.GeneralRemark
	latestData.Info.IncubationPeriod = resp.Newslist[0].Desc.Remark2
	latestData.Info.VulnerableGroup = resp.Newslist[0].Desc.Remark1
	latestData.Info.InfectionSource = resp.Newslist[0].Desc.Note2
	latestData.Info.Transmission = resp.Newslist[0].Desc.Note3
	latestData.Info.Name = resp.Newslist[0].Desc.Note1
	latestData.UpdateTime = resp.Newslist[0].Desc.ModifyTime

	Latest["txApi"] = latestData
	Original["txApi"] = resp
	//UpdateTodaySumCountData()
	return nil
}

func RequestNavData() error {
	url := fmt.Sprintf(conf.Conf.String("api::nav"))
	ret, err := util.NewRequest("GET", url, map[string]string{}, nil)
	var resp NavInfoType
	json.Unmarshal(ret, &resp)
	if err != nil {
		log.Println(err)
		return err
	}

	if resp.Code != 1 {
		return nil
	}

	Nav["latest"] = resp
	return nil
}

// 获取导航信息数据
func GetNavData() interface{} {
	_, ok := Nav["latest"]
	if ok == false {
		RequestNavData()
	}
	return Nav["latest"]
}

func GetNavRecommendationAppData() interface{} {
	_, ok := Recommendation["app"]
	if ok == false {
		RequestRecommendationData()
	}
	return Recommendation["app"]
}

//读取推荐信息数据到内存
func RequestRecommendationData() error {
	url := fmt.Sprintf(conf.Conf.String("api::recommendation"))
	ret, err := util.NewRequest("GET", url, map[string]string{}, nil)
	var resp RecommendationApiType
	json.Unmarshal(ret, &resp)
	if err != nil {
		log.Println(err)
		return err
	}
	if resp.Code != 1 {
		return nil
	}
	Recommendation["app"] = resp
	return nil
}
