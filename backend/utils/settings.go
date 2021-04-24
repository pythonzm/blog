package utils

import (
	"github.com/spf13/viper"
	"time"
)

type App struct {
	TimeFormat     string `json:"time_format"`
	JwtSecret      string `json:"jwt_secret"`
	TokenTimeout   int64  `json:"token_timeout"`
	StaticBasePath string `json:"static_base_path"`
	UploadBasePath string `json:"upload_base_path"`
	ImageRelPath   string `json:"image_rel_path"`
	AvatarRelPath  string `json:"avatar_rel_path"`
	ApiBaseUrl     string `json:"api_base_url"`
}

type Server struct {
	RunMode      string        `json:"run_mode"`
	ServerAddr   string        `json:"server_addr"`
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
}

type Redis struct {
	Host      string `json:"host"`
	Port      string `json:"port"`
	Password  string `json:"password"`
	DB        int    `json:"db"`
	CacheTime int    `json:"cache_time"`
}

type ES struct {
	Enable bool   `json:"enable"`
	Host   string `json:"host"`
	Port   string `json:"port"`
	Index  string `json:"index"`
}

type Mail struct {
	Enable   bool   `json:"enable"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Pass     string `json:"pass"`
	Receiver string `json:"receiver"`
}

var ServerInfo = &Server{}
var DBInfo = &DataBase{}
var AppInfo = &App{}
var RedisInfo = &Redis{}
var ESInfo = &ES{}
var MailInfo = &Mail{}

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
	AppInfo.StaticBasePath = viper.GetString("app.staticBasePath")
	AppInfo.UploadBasePath = viper.GetString("app.uploadBasePath")
	AppInfo.ImageRelPath = viper.GetString("app.imageRelPath")
	AppInfo.AvatarRelPath = viper.GetString("app.avatarRelPath")
	AppInfo.ApiBaseUrl = viper.GetString("app.apiBaseUrl")

	ServerInfo.RunMode = viper.GetString("server.runMode")
	ServerInfo.ServerAddr = viper.GetString("server.serverAddr")
	ServerInfo.ReadTimeout = time.Duration(viper.GetInt("server.readTimeout")) * time.Second
	ServerInfo.WriteTimeout = time.Duration(viper.GetInt("server.writeTimeout")) * time.Second

	DBInfo.Mode = viper.GetString("database.mode")
	DBInfo.Host = viper.GetString("database.host")
	DBInfo.Port = viper.GetString("database.port")
	DBInfo.User = viper.GetString("database.user")
	DBInfo.Password = viper.GetString("database.password")
	DBInfo.DBName = viper.GetString("database.dbName")

	RedisInfo.Host = viper.GetString("redis.host")
	RedisInfo.Port = viper.GetString("redis.port")
	RedisInfo.Password = viper.GetString("redis.password")
	RedisInfo.DB = viper.GetInt("redis.db")
	RedisInfo.CacheTime = viper.GetInt("redis.cacheTime")

	ESInfo.Enable = viper.GetBool("es.enable")
	ESInfo.Host = viper.GetString("es.host")
	ESInfo.Port = viper.GetString("es.port")
	ESInfo.Index = viper.GetString("es.index")

	MailInfo.Enable = viper.GetBool("mail.enable")
	MailInfo.Host = viper.GetString("mail.host")
	MailInfo.Port = viper.GetInt("mail.port")
	MailInfo.User = viper.GetString("mail.user")
	MailInfo.Pass = viper.GetString("mail.pass")
	MailInfo.Receiver = viper.GetString("mail.receiver")
}
