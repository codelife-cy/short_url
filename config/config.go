package config

import (
	"gin-template/util/env"
	"github.com/fsnotify/fsnotify"
	"log"
	"time"

	"github.com/spf13/viper"
)

//	初始化config对象
var config = new(Config)

// Config 配置文件结构体
type Config struct {
	// mysql
	MySql struct {
		Read struct {
			Addr string
			User string
			Pass string
			Name string
		}
		Write struct {
			Addr string
			User string
			Pass string
			Name string
		}
		Base struct {
			MaxOpenConn     int           `mapstructure:"max_open_conn"`
			MaxIdleConn     int           `mapstructure:"max_idle_conn"`
			ConnMaxLifeTime time.Duration `mapstructure:"conn_max_life_time"`
		}
	}
	// redis
	Redis struct {
		Addr        string
		Pass        string
		Db          int
		MaxRetries  int `mapstructure:"max_retries"`
		PoolSize    int `mapstructure:"pool_size"`
		MinIdleConn int `mapstructure:"min_idle_conn"`
	}
	// jwt
	JWT struct {
		Secret         string
		ExpireDuration time.Duration `mapstructure:"expire_duration"`
	}
	// language
	Language struct {
		Local string
	}

	// server
	Server struct {
		Port string
	}

	//common
	Common struct {
		BlackShortUrls   []string `mapstructure:"black_short_urls"`
		BlackShortUrlMap map[string]bool
		BaseString       string `mapstructure:"base_string"`
		DomainName       string `mapstructure:"domain_name"`
		Schema           string
		DomainLength     uint64 `mapstructure:"domain_length"`
	}
}

func init() {
	viper.SetConfigName(env.Active().Value() + "_config")
	viper.SetConfigType(viper.SupportedExts[1])
	viper.AddConfigPath("./config")
	// 读取配置
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("配置文件不存在")
			panic(err)
		} else {
			log.Printf("配置文件不存在，但其它错误")
			panic(err)
		}
	}
	//	解析
	if err := viper.Unmarshal(config); err != nil {
		log.Printf("配置文件解析错误")
		panic(err)
	}
	// 监听配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}
	})
	//初始化值
	config.Common.BlackShortUrlMap = make(map[string]bool)
	for _, url := range config.Common.BlackShortUrls {
		config.Common.BlackShortUrlMap[url] = true
	}
}

// Get 获取配置文件
func Get() *Config {
	return config
}
