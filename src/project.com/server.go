/* ****************************************************************************
Filename    : server.go
File-type   : golang source code.
Compiler    : go version go1.17.1 windows/amd64
Author      : Vaibhav Jaiswal
Date        : 30-Oct-2022

Description :
- Use as core server for the application.

Version History
Version     : 1.0
Author      : Vaibhav Jaiswal
Description : Initial version
**************************************************************************** */

package main

import (
	"fmt"
    "github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
    person "./person"
)

var dbEngine *sqlx.DB

func main() {

    cfg := mysql.Config{
        User:   "root",
        Passwd: "root",
        Addr:   "127.0.0.1:3306",
        DBName: "cetec",
        AllowNativePasswords: true,
    }

    // Get a database handle.
    dbEngine := person.Init("mysql", cfg.FormatDSN())
    
    pingErr := dbEngine.Ping()
    if pingErr != nil {
        fmt.Println("Database Connection Failed...!")
    } else {
        fmt.Println("Database Connection Success...")
    }

    router := gin.Default()

    router.GET("/person/:person_id/info", person.GetPerson)
    router.POST("/person/create", person.CreatePerson)

    router.Run("localhost:8080")
}