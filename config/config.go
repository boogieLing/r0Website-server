// Package config
/**
 * @Author: r0
 * @Mail: boogieLing_o@qq.com
 * @Description: 配置项
 * @File:  config
 * @Version: 1.0.0
 * @Date: 2022/7/3 18:21
 */
package config

type SystemConfig struct {
	System `yaml:"system"`
	Logger `yaml:"logger"`
	JWT    `yaml:"jwt"`
	Mongo  `yaml:"mongo"`
	Author `yaml:"author"`
}

type System struct {
	Part   string `yaml:"part"`   // 地址
	Status string `yaml:"status"` // 状态
}

type Logger struct {
	Path string `yml:"path"` // 日志文件路径
}

type JWT struct {
	SignKey     string `yaml:"sign-key"`
	ExpiresTime int64  `yaml:"expires-time"`
	Issuer      string `yaml:"issuer"`
}

type Mongo struct {
	Address        string `yaml:"address"`
	Part           string `yaml:"part"`
	DB             string `yaml:"db"`
	SuperAdminName string `yaml:"super-admin-name"`
	SuperAdminPswd string `yaml:"super-admin-pswd"`
}

type Author struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}
