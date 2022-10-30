# golang-mini-project
This is a simple project for understanding gin framework with database connectivity.

## Configuration

### Developement Language
go version go1.17.1

### Database
MySQL


## API Reference

#### Get person

```http
  GET /person/:person_id/info
```

| Parameter | Type     | Description                         |
| :-------- | :------- | :--------------------------------   |
| `id`      | `int`    | **Required**. Id of person to fetch |


#### Create person

```http
  GET /person/create
```

| Parameter        | Type     | Description                         |
| :--------------  | :------- | :--------------------------------   |
| `Person` Object  | `int`    | **Required**. Id of person to fetch |


`Person` = {
        "name": "",
        "phone_number": "",
        "city": "",
        "state": "",
        "street1": "",
        "street2": "",
        "zip_code": ""
    }

  
## 🚀 About Me
I'm a full stack developer...