package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func (s *Service) initMySQL() {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local&charset=utf8mb4,utf8",
		s.Conf.MySQL.DBUsername,
		s.Conf.MySQL.DBPassword,
		s.Conf.MySQL.DBHost,
		s.Conf.MySQL.DBName,
	))
	if err != nil {
		log.Fatalln(err)
	}

	s.Mysql = db

	// 建表
	s.Mysql.AutoMigrate(
		&Manager{},
		&Challenge{},
		&Token{},
		&Team{},
		&GameBox{},
		&Bulletin{},
		&BulletinRead{},

		&AttackAction{},
		&DownAction{},
		&Score{},
		&Flag{},
		&Log{},
	)

	// 初始数据
	var managerCount int
	s.Mysql.Model(&Manager{}).Count(&managerCount)
	if managerCount == 0 {
		// 添加默认管理员
		s.Mysql.Create(&Manager{
			Name:     "e99",
			Password: s.addSalt("123456"),
		})
		s.NewLog(WARNING, "system", fmt.Sprintf("添加默认管理员账号成功，请注意修改密码！"))
		log.Println("添加默认管理员账号成功")
	}
}
