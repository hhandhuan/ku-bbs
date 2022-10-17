package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var Conf *conf

type db struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Name string `yaml:"name"`
	Pass string `yaml:"pass"`
	DB   string `yaml:"db"`
}

type app struct {
	Version   string `yaml:"version"`
	Name      string `yaml:"name"`
	Desc      string `yaml:"desc"`
	Keywords  string `yaml:"keywords"`
	VisitMode string `yaml:"visitMode"`
}

type session struct {
	Name   string `yaml:"name"`
	Secret string `yaml:"secret"`
}

type upload struct {
	Path           string   `yaml:"path"`
	ImageExt       []string `yaml:"imageExt"`
	AvatarFileSize int64    `yaml:"avatarFileSize"`
	TopicFileSize  int64    `yaml:"topicFileSize"`
}

type redis struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	Pass        string `yaml:"pass"`
	DB          string `yaml:"db"`
	IdleTimeout string `yaml:"idleTimeout"`
}

type system struct {
	Env  string `yaml:"env"`
	Addr string `yaml:"addr"`
}

type conf struct {
	App     app     `yaml:"app"`
	System  system  `yaml:"system"`
	DB      db      `yaml:"db"`
	Session session `yaml:"session"`
	Upload  upload  `yaml:"upload"`
	Redis   redis   `yaml:"redis"`
}

func init() {
	b, err := os.ReadFile("../config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	var c *conf
	err = yaml.Unmarshal(b, &c)
	if err != nil {
		log.Fatal(err)
	}
	Conf = c
}
