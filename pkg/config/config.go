package config

import (
	"flag"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Mysql struct {
	Host  string `yaml:"host"`
	Port  string `yaml:"port"`
	Name  string `yaml:"name"`
	Pass  string `yaml:"pass"`
	DB    string `yaml:"db"`
	Debug bool   `yaml:"debug"`
}

type App struct {
	Version   string `yaml:"version"`
	Name      string `yaml:"name"`
	Desc      string `yaml:"desc"`
	Keywords  string `yaml:"keywords"`
	VisitMode string `yaml:"visitMode"`
}

type Session struct {
	Name   string `yaml:"name"`
	Secret string `yaml:"secret"`
}

type Upload struct {
	Path           string   `yaml:"path"`
	ImageExt       []string `yaml:"imageExt"`
	AvatarFileSize int64    `yaml:"avatarFileSize"`
	TopicFileSize  int64    `yaml:"topicFileSize"`
}

type Redis struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	Pass        string `yaml:"pass"`
	DB          string `yaml:"db"`
	IdleTimeout string `yaml:"idleTimeout"`
}

type System struct {
	Env              string `yaml:"env"`
	Addr             string `yaml:"addr"`
	ShutdownWaitTime int    `yaml:"shutdownWaitTime"`
}

type Logger struct {
	Path       string `yaml:"path"`
	Level      int    `yaml:"level"`
	MaxSize    int    `yaml:"maxSize"`
	MaxBackups int    `yaml:"maxBackups"`
	MaxAge     int    `yaml:"maxAge"`
	Compress   bool   `yaml:"compress"`
}

type Config struct {
	App     *App     `yaml:"app"`
	System  *System  `yaml:"system"`
	Mysql   *Mysql   `yaml:"db"`
	Session *Session `yaml:"session"`
	Upload  *Upload  `yaml:"upload"`
	Redis   *Redis   `yaml:"redis"`
	Logger  *Logger  `yaml:"logger"`
}

var instance *Config

func GetInstance() *Config {
	return instance
}

var (
	path = flag.String("cfg", "../config/config.yaml", "config file path")
)

func Initialize() {
	flag.Parse()
	buf, err := os.ReadFile(*path)
	if err != nil {
		log.Fatal(err)
	}

	var c Config
	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		log.Fatal(err)
	}

	instance = &c
}
