package config

import (
	"github.com/BurntSushi/toml"
	"os"
	"strings"
	"sync"
)

type Config struct {
	App     app
	Mysql   mysql
	Mongodb mongodb
	Redis   redis
}

type app struct {
	Addr string
}

type mysql struct {
	Connect string
	MaxIdle int
	MaxOpen int
}

type mongodb struct {
	Connect string
}

type redis struct {
	Connect string
	DB      int
	MaxIdle int
	MaxOpen int
}

var (
	c    *Config
	once sync.Once
)

/**
 * 返回单例实例
 * @method New
 */
func New() *Config {
	once.Do(func() { //只执行一次
		file, _ := os.Getwd()
		file = strings.Replace(file, "restest", "", -1)
		if _, err := toml.DecodeFile(file + "/config.toml", &c); err != nil {
			panic(err.Error())
		}
	})
	return c
}

// /**
//  * 获取value
//  * @method func
//  * @param  {[type]} c *Config       [description]
//  * @return {[type]}   [description]
//  */
// func (c *Config) Get(key string) string {
// 	return c[key]
// }
