package external_package_demo

import (
	"fmt"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PersonTab struct {
	Id    uint64  `gorm:"column:pId;type:bigint unsigned;AUTO_INCREMENT;primary_key"`
	Name  *string `gorm:"column:name;type:varchar(20)"`
	Age   int     `gorm:"column:age;type:int(11)"`
	Grade *int    `gorm:"column:grade;type:int(11)"`
	Ctime uint64  `gorm:"column:ctime;type:bigint unsigned"`
	Mtime uint64  `gorm:"column:mtime;type:int(11）"` // 和数据库 bigint unsigned 不符, 但是 type 没有强制校验
}

func (p *PersonTab) TableName() string {
	return "person"
}

func (p *PersonTab) BeforeCreate(tx *gorm.DB) error {
	now := uint64(time.Now().Unix())
	p.Ctime = now
	p.Mtime = now
	return nil
}

func (p *PersonTab) BeforeUpdate(tx *gorm.DB) error {
	now := uint64(time.Now().Unix())
	p.Mtime = now
	return nil
}

func TestA(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:123456@(127.0.0.1:3306)/testdb"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	person := &PersonTab{}
	res := db.Create(&person) //INSERT INTO `person` (`name`,`age`,`grade`,`ctime`,`mtime`) VALUES (NULL,0,NULL,1661245911,1661245911)
	fmt.Println(res.Error, res.RowsAffected)
}
