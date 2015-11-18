package db

import (
	"code.google.com/p/gcfg"
	"github.com/jinzhu/gorm"
	"net/url"
)

type Config struct {
	DbConnection string
	Charset      string
	ParseTime    string
	Loc          string
}

var (
	db *gorm.DB
)

type configFile struct {
	Server Config
}

func loadConfiguration(cfgFile string) (Config, error) {
	var cfg configFile
	err := gcfg.ReadFileInto(&cfg, cfgFile)
	return cfg.Server, err
}

func Db() *gorm.DB {
	conf, err := loadConfiguration("data.conf")

	if err != nil {
		panic(err)
	}

	if db == nil {
		urlVals := url.Values{
			"charset":   {conf.Charset},
			"parseTime": {conf.ParseTime},
			"loc":       {conf.Loc},
		}
		args := conf.DbConnection + "?" + urlVals.Encode()
		dbl, err := gorm.Open("mysql", args)
		if err != nil {
			panic(err)
		}
		db = &dbl
		db.SingularTable(true)
	}

	return db
}
