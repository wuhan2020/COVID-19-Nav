package Service

import (
	"log"
	"nCoV-API/apps/Model"
	"nCoV-API/lib/util"
	"net/url"

	//"net/url"
	"time"
)

type DailySumBson struct {
	ID           string `bson:"_id"`
	AddConfirmed int    `json:"add_confirmed" bson:"add_confirmed"`
	AddSuspect   int    `json:"add_suspect" bson:"add_suspect"`
	AddDeath     int    `json:"add_death" bson:"add_death"`
	AddCure      int    `json:"add_cure" bson:"add_cure"`
	SumConfirmed int    `json:"sum_confirmed" bson:"sum_confirmed"`
	SumSuspect   int    `json:"sum_suspect" bson:"sum_suspect"`
	SumDeath     int    `json:"sum_death" bson:"sum_death"`
	SumCure      int    `json:"sum_cure" bson:"sum_cure"`
	CountDate    string `json:"count_date" bson:"count_date"`
	UpdateTime   string `json:"update_time" bson:"update_time"`
}

var DailySum = make(map[string]DailySumBson)

func UpdateDailySumCountCache(date string) (DailySumBson, error) {
	//idStr := "2020-01-24"
	//objectId := bson.ObjectIdHex(idStr)
	var dailySum DailySumBson
	err := Model.GetCollection("daily_sum_data").FindId(date).One(&dailySum)
	if err != nil {
		log.Println("get db error : ", err)
		return dailySum, err
	}
	log.Println("daily sum date:", date, " date:", dailySum)
	DailySum[date] = dailySum
	return dailySum, nil
}

func GetLatestDailySumInfoByCache(params url.Values) (DailySumBson, error) {
	date := util.GetParam(params, "date", "").(string)
	if date == "" {
		date = getDefaultDate()
	}
	log.Println("nowTime:", date)

	rebuild := util.GetParam(params, "rebuild", "").(string)
	_, ok := DailySum[date]
	if ok == false || rebuild == "1" {
		_, err := UpdateDailySumCountCache(date)
		return DailySum[date], err
	}
	return DailySum[date], nil
}

// 获取今日汇总统计数据（实时变动）
func UpdateTodaySumCountData() (DailySumBson, error) {
	// 获取前一日数据
	nowTime := getDefaultTime()
	todayDate := nowTime.Format("2006-01-02")

	nowTime = nowTime.Add(-time.Hour * 24)
	lastDate := nowTime.Format("2006-01-02")
	lastData, err := UpdateDailySumCountCache(lastDate)
	var today DailySumBson
	if err != nil {
		return today, err
	}

	// 获取实时数据
	realTimeData := GetTxApiData()
	log.Println("lastData ", lastData)
	log.Println("real time data", realTimeData)

	// 计算今日实时统计
	today.AddConfirmed = realTimeData.Count.Confirmed - lastData.SumConfirmed
	today.AddDeath = realTimeData.Count.Death - lastData.SumDeath
	today.AddSuspect = realTimeData.Count.Suspected - lastData.SumSuspect
	today.AddCure = 0
	today.SumConfirmed = realTimeData.Count.Confirmed
	today.SumDeath = realTimeData.Count.Death
	today.SumCure = realTimeData.Count.Cure
	today.SumSuspect = realTimeData.Count.Suspected
	today.CountDate = todayDate
	//today.UpdateTime =
	//today.UpdateTime = time.Now().Format("20060102")

	//today.UpdateTime =
	//log.Println("today data", today)
	DailySum[todayDate] = today
	return today, err
}

func getDefaultDate() string {
	nowTime := getDefaultTime()
	defaultDate := nowTime.Format("2006-01-02")
	return defaultDate
}

func getDefaultTime() time.Time {
	nowTime := time.Now()
	if nowTime.Hour() <= 7 {
		nowTime = nowTime.Add(-time.Hour * 24)
	}
	return nowTime
}
