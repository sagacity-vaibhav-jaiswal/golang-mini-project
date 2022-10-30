/* ****************************************************************************
Filename    : dbAccess.go
File-type   : golang source code.
Compiler    : go version go1.17.1 windows/amd64
Author      : Vaibhav Jaiswal
Date        : 30-Oct-2022

Description :
- Uses as request/response models for application.

Version History
Version     : 1.0
Author      : Vaibhav Jaiswal
Description : Initial version
**************************************************************************** */

package person

type Person struct {
    Name string `db:"name" json:"name"`
    Number string `db:"phone_number" json:"phone_number"`
    City string `db:"city" json:"city"`
    State string `db:"state" json:"state"`
    Street1 string `db:"street1" json:"street1"`
    Street2 string `db:"street2" json:"street2"`
    ZipCode string `db:"zip_code" json:"zip_code"`
}

type Phone struct {
    PersonId int `db:"person_id"`
    Number string `db:"phone_number" json:"phone_number"`
}

type Address_Join struct {
    PersonId int `db:"person_id"`
    AddressId int `db:"address_id"`
}