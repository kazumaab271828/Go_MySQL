package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func gormConnect() *gorm.DB {
    DBMS    := "mysql"
    USER    := "root"
    PASS    := "#####"
    PROTOCOL:= "tcp(localhost)"
    DBNAME  := "testdb"
    CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
    db, err := gorm.Open(DBMS, CONNECT + "?parseTime=true")

    if err != nil {
        panic(err.Error())
    }
    fmt.Println("db connected: ", &db)
    return db
}

func main() {
 db := gormConnect()

 defer db.Close()
 db.LogMode(true)

type User struct {
	Id   int    `gorm:"primary_key"`
	Name string `gorm:"primary_key"`
	Col  string `gorm:"primary_key"`
}

var inuser = User{Id: 7, Name: "Kaneko", Col: "Maebashi"}
db.Create(inuser)

var user User
db.Find(&user)
db.First(&user)

db.Delete(&user, 7)

db.Save(&user)
}
