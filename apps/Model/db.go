package Model

import (
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	//"github.com/go-mgo/mgo"
	"gopkg.in/mgo.v2"
	"nCoV-API/lib/conf"
	//"time"
)

//var Col *mgo.Collection
var Session *mgo.Session
var Db *mgo.Database

//func (user User) ToString() string {
//	return fmt.Sprintf("%#v", user)
//}

func Init() {
	Session, _ = mgo.Dial(conf.Conf.String("mongodb::host"))
	//切换到数据库
	Db = Session.DB(conf.Conf.String("mongodb::database"))
	//切换到collection
	//Col = Db.C(conf.Conf.String("mongodb::collection"))
}

func GetCollection(col string) *mgo.Collection {
	return Db.C(col)
}

//新增数据
func add() {
	//    defer Session.Close()
	//stu1 := new(User)
	//stu1.Id = bson.NewObjectId()
	//stu1.Username = "stu1_name"
	//stu1.Pass = "stu1_pass"
	//stu1.Regtime = time.Now().Unix()
	//stu1.Interests = []string{"象棋", "游泳", "跑步"}
	//err := Col.Insert(stu1)
	//if err == nil {
	//	fmt.Println("插入成功")
	//} else {
	//	fmt.Println(err.Error())
	//	defer panic(err)
	//}
}

//查询
func find() {
	//    defer Session.Close()
	//var users []User

	//    Col.Find(nil).All(&users)
	//Col.Find(bson.M{"name": "stu1_name"}).All(&users)
	//for _, value := range users {
	//	fmt.Println(value.ToString())
	//}
	////根据ObjectId进行查询
	//idStr := "577fb2d1cde67307e819133d"
	//objectId := bson.ObjectIdHex(idStr)
	//user := new(User)
	//Col.Find(bson.M{"_id": objectId}).One(user)
	//fmt.Println(user)
}

//根据id进行修改
func update() {
	//interests := []string{"象棋", "游泳", "跑步"}
	//err := Col.Update(bson.M{"_id": bson.ObjectIdHex("577fb2d1cde67307e819133d")}, bson.M{"$set": bson.M{
	//	"name":      "修改后的name",
	//	"pass":      "修改后的pass",
	//	"regtime":   time.Now().Unix(),
	//	"interests": interests,
	//}})
	//if err != nil {
	//	fmt.Println("修改失败")
	//} else {
	//	fmt.Println("修改成功")
	//}
}

//删除
func del() {

	//err := Col.Remove(bson.M{"_id": bson.ObjectIdHex("577fb2d1cde67307e819133d")})
	//if err != nil {
	//	fmt.Println("删除失败" + err.Error())
	//} else {
	//	fmt.Println("删除成功")
	//}
}
