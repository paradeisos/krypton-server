package models

import "github.com/astaxie/beego"

var (
	mongo *Model
)

func InitMongo() {
	mongoHost := beego.AppConfig.String("mongo_host")
	mongoUser := beego.AppConfig.String("mongo_user")
	mongoPassword := beego.AppConfig.String("mongo_password")
	mongoDatabase := beego.AppConfig.String("mongo_database")
	mongoMode := beego.AppConfig.String("mongo_mode")
	mongoPool := beego.AppConfig.DefaultInt("mongo_pool", 0)
	mongoTimeout := beego.AppConfig.DefaultInt("mongo_timeout", 0)

	option := &Option{
		Host:     mongoHost,
		User:     mongoUser,
		Password: mongoPassword,
		Database: mongoDatabase,
		Mode:     mongoMode,
		Pool:     mongoPool,
		Timeout:  mongoTimeout,
	}

	mongo = NewModel(option)

}

func Mongo() *Model {
	return mongo
}

func DropDatabase() {
	Mongo().Session().DB(mongo.Database()).DropDatabase()
}
