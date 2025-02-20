package conf

import (
	"context"
	"fmt"
	"gin-chat-demo/model"
	logging "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/ini.v1"
	"strings"
)

var (
	MongoDBClient *mongo.Client
	AppMode       string
	HttpPort      string
	RedisAddr     string
	RedisPw       string
	RedisDbName   string
	Db            string
	DbHost        string
	DbPort        string
	DbUser        string
	DbPassWord    string
	DbName        string
	MongoDBName   string
	MongoDBAddr   string
	MongoDBPwd    string
	MongoDBPort   string
)

func Init() {
	//从本地读取环境
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("ini load failed", err)
	}
	LoadServer(file)
	LoadMySQL(file)
	LoadMongoDB(file)
	MongoDB()
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	model.Database(path)
}

func MongoDB() {
	clientOptions := options.Client().ApplyURI("mongodb://" + MongoDBAddr + ":" + MongoDBPort)
	var err error
	MongoDBClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logging.Info(err)
		panic(err)
	}
	logging.Info("MongoDB connect success")
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadMySQL(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func LoadMongoDB(file *ini.File) {
	MongoDBName = file.Section("MongoDB").Key("MongoDBName").String()
	MongoDBAddr = file.Section("MongoDB").Key("MongoDBAddr").String()
	MongoDBPwd = file.Section("MongoDB").Key("MongoDBPwd").String()
	MongoDBPort = file.Section("MongoDB").Key("MongoDBPort").String()
}
