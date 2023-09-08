package Table

import (
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
	"time"
)

type Like struct {
	ID        int    `gorm:"primary_key"`
	Ip        string `gorm:"type:varchar(20);not null;index:ip_idx"`
	Ua        string `gorm:"type:varchar(256);not null;"`
	Title     string `gorm:"type:varchar(128);not null;index:title_idx"`
	Hash      uint64 `gorm:"unique_index:hash_idx;"`
	CreatedAt time.Time
}

func init() {
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if !db.HasTable(&Like{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Like{}).Error; err != nil {
			panic(err)
		}
	}
	/*
		-- 创建的表如下
		CREATE TABLE `likes` (
			`id` int NOT NULL AUTO_INCREMENT,
			`ip` varchar(20) NOT NULL,
			`ua` varchar(256) NOT NULL,
			`title` varchar(128) NOT NULL,
			`hash` bigint unsigned DEFAULT NULL,
			`created_at` datetime DEFAULT NULL,
			PRIMARY KEY (`id`),
			UNIQUE KEY `hash_idx` (`hash`),
			KEY `ip_idx` (`ip`),
			KEY `title_idx` (`title`)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

	*/
}
