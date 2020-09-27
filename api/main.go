package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
)

func main() {
    router := gin.Default()
    cityEx := City{}

    db, err := connectDb()
    defer db.Close()

    // Permit singular name table
    db.SingularTable(true)

    if err != nil {
        router.GET("/", func(context *gin.Context) {
            context.JSON(http.StatusOK, gin.H{"data": err})
        })
    } else {
        // select id
        cityEx.ID = 12

        // get first record of id=12
        result := db.First(&cityEx)
        router.GET("/", func(context *gin.Context) {
            context.JSON(http.StatusOK, gin.H{"data": result.Value})
        })
    }

    router.Run()
}

func connectDb()(database *gorm.DB, err error) {
    DBMS := "mysql"
    USER := "root"
    PASS := "root"
    PROTOCOL := "tcp(db:3306)" // "db" is docker container name.
    DBNAME := "world"

    CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
    return gorm.Open(DBMS, CONNECT)
}


type City struct {
    // `gorm:"column:{Custom column name}"`
    ID int `gorm:"column:ID"`
    Name string `gorm:"column:Name"`
    CountryCode string `gorm:"column:CountryCode"`
    District string `gorm:"column:District"`
    Population int `gorm:"column:Population"`
}