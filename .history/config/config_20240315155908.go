package config

import "log"

type ConfigList struct {
	Port string
	SQLDriver string
	DbName string
	LogFile string
}

var Config ConfigList

func LoadConfig(){
	cfg, err := ini.Load("config.ini")
	if err != nil{
		log.Fatalln(err)
	}
	Config = ConfigList{
		
	}
}
