/* ****************************************************************************
Filename    : dbAccess.go
File-type   : golang source code.
Compiler    : go version go1.17.1 windows/amd64
Author      : Vaibhav Jaiswal
Date        : 30-Oct-2022

Description :
- Uses as db helper functions.

Version History
Version     : 1.0
Author      : Vaibhav Jaiswal
Description : Initial version
**************************************************************************** */

package person

import (
    "github.com/jmoiron/sqlx"
	"net/http"
	"fmt"
	"strconv"
    "github.com/gin-gonic/gin"
)

var dbEngine *sqlx.DB

func Init(driver, connectionString string) *sqlx.DB {
	dbEngine = sqlx.MustConnect(driver, connectionString)
	return dbEngine
}

func GetPerson(c *gin.Context) {

    param_id := c.Param("person_id")
	id, err := strconv.Atoi(param_id)
    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid Person Id"})
		return
    }
	if id < 1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Person Id Should be Greater Than 0"})
		return
	}
    recList := []Person{}

    err = dbEngine.Select(&recList, "SELECT p.name, pn.number as phone_number, ad.city, ad.state, ad.street1, ad.street2, ad.zip_code from person p, phone pn, address ad, address_join aj where p.id = pn.person_id and ad.id = aj.address_id and aj.person_id = p.id and p.id = ?;", id)
    if err != nil {
        fmt.Println("Error in db...!")
        fmt.Println(err)
		return
    }

    if len(recList) < 1 {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Person Not Found"})
		return
    }

    c.JSON(http.StatusOK, recList)
	return
}


func CreatePerson(c *gin.Context) {

    var person Person

    // Call BindJSON to bind the received JSON to person
    if err := c.BindJSON(&person); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Client Data Binding Error..."})
		return
    }

    _, err := dbEngine.NamedExec("Insert into person (name) values (:name)", person)
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Database transaction Failed"})
		return
    }
    
    var person_id = make([]int, 0)

    err = dbEngine.Select(&person_id, "SELECT MAX(id) FROM person")
    if err != nil {
        fmt.Println(err)
        c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Did Not Get Person Id"})
        return
    }

    var phone Phone
    phone.PersonId = person_id[0]
    phone.Number = person.Number

    _, err = dbEngine.NamedExec("Insert into phone (number, person_id) values (:phone_number, :person_id)", phone)
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Database transaction Failed"})
		return
    }

    _, err = dbEngine.NamedExec("Insert into address (city, state, street1, street2, zip_code) values (:city, :state, :street1, :street2, :zip_code)", person)
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Database transaction Failed"})
		return
    }

    var address_id = make([]int, 0)

    err = dbEngine.Select(&address_id, "SELECT MAX(id) FROM address")
    if err != nil {
        fmt.Println(err)
        c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Did Not Get Address Id"})
		return
    }

    var address_join Address_Join
    address_join.PersonId = person_id[0]
    address_join.AddressId = address_id[0]

    _, err = dbEngine.NamedExec("Insert into address_join (person_id, address_id) values (:person_id, :address_id)", address_join)
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Database transaction Failed"})
		return
    }

	return
}
