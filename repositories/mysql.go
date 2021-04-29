package repositories

import (
	"errors"
	"fmt"
	"github.com/caarlos0/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConfig struct {
	Host string `env:"MYSQL_HOST"`
	Port int `env:"MYSQL_PORT"`
	Name string `env:"MYSQL_NAME"`
	User string `env:"MYSQL_USER"`
	Password string `env:"MYSQL_PASSWORD"`
}

func initDB() {
	config := DBConfig{}

	// env parsing
	if err := env.Parse(&config); err != nil {
		panic(errors.New("cloud not load MySQL Environment variables"))
	}

	// Make URL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name)

	// Make Database Connection
	ormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}
	db, err := gorm.Open(mysql.Open(dsn), ormConfig)
	if err != nil {
		panic(err)
	}

	//AutoMigrate
	//err = db.AutoMigrate(&blog_m.User{},
	//&blog_m.Post{},
	//&blog_m.PostMeta{},
	//&blog_m.PostComment{},
	//&blog_m.Category{},
	//&blog_m.PostCategory{},
	//&blog_m.Tag{},
	//&blog_m.PostTag{},
	//)
	//if err != nil {
	//	return nil, err
	//}

	Connections.DB = db
}
