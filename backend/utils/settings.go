package utils

import (
	"github.com/spf13/viper"
	"time"
)

type App struct {
	TimeFormat   string `json:"time_format"`
	JwtSecret    string `json:"jwt_secret"`
	TokenTimeout int64  `json:"token_timeout"`
	PageSize     int    `json:"page_size"`
}

type Server struct {
	RunMode      string        `json:"run_mode"`
	Host         string        `json:"host"`
	Port         string        `json:"port"`
	ReadTimeout  time.Duration `json:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout"`
}

type DataBase struct {
	Mode     string `json:"mode"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
	Salt     string `json:"salt"`
}

var ServerInfo = &Server{}
var DBInfo = &DataBase{}
var AppInfo = &App{}

func init() {
	viper.AddConfigPath("conf")
	viper.SetConfigName("config")
	// viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	AppInfo.TimeFormat = viper.GetString("app.timeFormat")
	AppInfo.JwtSecret = viper.GetString("app.jwtSecret")
	AppInfo.TokenTimeout = viper.GetInt64("app.tokenTimeout")
	AppInfo.PageSize = viper.GetInt("app.pageSize")

	ServerInfo.RunMode = viper.GetString("server.runMode")
	ServerInfo.Host = viper.GetString("server.host")
	ServerInfo.Port = viper.GetString("server.port")
	ServerInfo.ReadTimeout = time.Duration(viper.GetInt("server.readTimeout")) * time.Second
	ServerInfo.WriteTimeout = time.Duration(viper.GetInt("server.writeTimeout")) * time.Second

	DBInfo.Mode = viper.GetString("database.mode")
	DBInfo.Host = viper.GetString("database.host")
	DBInfo.Port = viper.GetString("database.port")
	DBInfo.User = viper.GetString("database.user")
	DBInfo.Password = viper.GetString("database.password")
	DBInfo.DBName = viper.GetString("database.dbName")
	DBInfo.Salt = viper.GetString("database.salt")
}
