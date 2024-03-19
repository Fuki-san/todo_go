package config

import "log"

type ConfigList struct {
	Port string
	SQLDriver string
	DbName string
	LogFile string
}
//var instance-name 
var Config ConfigList
//Load config.ini and implement it to ConfigList struct
func LoadConfig(){
	cfg, err := ini.Load("config.ini")
	if err != nil{
		log.Fatalln(err)
	}
	Config = ConfigList{

	}
}
