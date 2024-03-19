package config

import "log"

type ConfigList struct {
	Port string
	SQLDriver string
	DbName string
	LogFile string
}
//create instance
var Config ConfigList

//want to load 

//Load config.ini and implement it to ConfigList struct
func LoadConfig(){
	cfg, err := ini.Load("config.ini")
	if err != nil{
		log.Fatalln(err)
	}
	//load contents of ini-file in ConfigList scope and set its data to ConfigList
	Config = ConfigList{
		Port: cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName: cfg.Section("db").Key("name").String(),
		LogFile: cfg.Section("web").Key("logfile").MustString(),
	}
}
