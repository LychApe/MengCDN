package model

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	//连接数据库
	db, err := gorm.Open(sqlite.Open("config/MengCDN.db"), &gorm.Config{})
	if err != nil {
		panic("[MengCDN][InitDB]|[Error]:未寻找到数据库MengCDN.db")
	}
	//数据库自动迁移
	err = db.AutoMigrate(&SystemConfig{}, &User{}, &WhiteList{})
	if err != nil {
		panic("[MengCDN][AutoMigrate]|[Error]:迁移数据库异常")
	}
	DB = db

	//siteInfo := []WhiteList{
	//	{
	//		Key:   "GithubWL",
	//		Value: `{"Lychape/DreamCat","LychApe/DreamCat","luoxin-spot/gitcdn","inkdust-dev/ink-oss","louie-senpai/Siren","hmsjy2017/cdn","Zisbusy/Akina-for-Typecho","P3TERX/script","cdnjs/cdnjs","solstice23/argon-theme","halo-dev/theme-earth","solstice23/argon-theme",}`,
	//	},
	//	{
	//		Key:   "NpmWL",
	//		Value: `{"hexo-theme-cards","mdui","alist-web","jquery","pjax","waline","fontawesome","fancybox"}`,
	//	},
	//	{
	//		Key:   "WPthWL",
	//		Value: `{}`,
	//	},
	//	{
	//		Key:   "WPplWL",
	//		Value: `{}`,
	//	},
	//}
	//db.Create(siteInfo)
}
