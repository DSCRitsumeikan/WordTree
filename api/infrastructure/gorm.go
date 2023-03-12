package infrastructure

import (
	"WordTree/config"
	"WordTree/model"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db := InitGorm()
	return db
}

func InitGorm() *gorm.DB {
	driver := config.Mysql{
		Host:     config.GetEnvWithDefault("DB_HOST", "db"),
		Port:     config.GetEnvWithDefault("DB_PORT", "3306"),
		UserName: config.GetEnvWithDefault("DB_USER", "wordtree_user"),
		Password: config.GetEnvWithDefault("DB_PASSWORD", "password"),
		DBName:   config.GetEnvWithDefault("DB_NAME", "wordtree_dev"),
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", driver.UserName, driver.Password, driver.Host, driver.Port, driver.DBName)
	db := connectDB(dsn)
	autoMigrate(db)
	return db
}

func connectDB(dsn string) *gorm.DB {
	db, err := openDBWithTimeLimit(dsn, 5)
	if err != nil {
		panic(err)
	}
	return db
}

func openDBWithTimeLimit(dsn string, count int) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		if count == 0 {
			return nil, fmt.Errorf("retry count over")
		}
		time.Sleep(time.Second)
		count--
		return openDBWithTimeLimit(dsn, count)
	}
	return db, nil
}

func autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		new(model.User),
		new(model.CacheImage),
		new(model.Session),
		new(model.WordDefinition),
		new(model.WordNode),
		new(model.CurrentNodeRelation),
		new(model.WordNodeChildren),
	)
	if err != nil {
		panic(err)
	}
}
