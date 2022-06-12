package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

type tomlConfig struct {
	Viewer Viewer
	System SystemConfig
}
type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Avatar      string
	UserName    string
	UserDesc    string
}
type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QiniuAccessKey  string
	QiniuSecretKey  string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
}

var Cfg *tomlConfig

func init() {
	// 程序启动时候，就会执行init方法
	Cfg = new(tomlConfig)
	// 手动为一些数据赋值
	Cfg.System.AppName = "ms-go-blog"
	Cfg.System.Version = 1.0 //版本
	currentDir, _ := os.Getwd()
	Cfg.System.CurrentDir = currentDir                    // 目录
	_, err := toml.DecodeFile("config/config.toml", &Cfg) // 为toml文件进行解码，就是主页面照片等
	if err != nil {
		panic(err)
	}
}
