package controllers
import (
	"github.com/jinzhu/gorm"
	"sample_revel/app/model"
	_ "github.com/go-sql-driver/mysql" // please import mysql driver (if not have please install)
	"github.com/revel/revel"
	"database/sql"
)

var db *gorm.DB

func initDB() {
	// connection string -> user:password@host/database
	dbm, err := gorm.Open("mysql", "root:@/test?charset=utf8&parseTime=True")
	if err != nil {
		panic(err)
	}

	db = &dbm

	// auto migrate model table
	db.AutoMigrate(&model.User{})
}

type GormController struct {
	*revel.Controller
	db *gorm.DB
}

func (c *GormController) Begin() revel.Result {
	dbm := db.Begin()
	if dbm.Error != nil {
		panic(dbm.Error)
	}
	c.db = dbm
	return nil
}

func (c *GormController) Commit() revel.Result {
	if c.db == nil {
		return nil
	}

	c.db.Commit()
	if err := c.db.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.db = nil
	return nil
}

func (c *GormController) Rollback() revel.Result {
	if c.db == nil {
		return nil
	}

	c.db.Rollback()
	if err := c.db.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.db = nil
	return nil
}