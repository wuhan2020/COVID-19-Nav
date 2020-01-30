package conf;

import (
	"github.com/astaxie/beego/config"
)
var (
	Conf config.Configer
)
func LoadConfig(path string) {
	var err error
	Conf, err = config.NewConfig("ini", path)
	if err != nil {
		panic(err)
	}
}