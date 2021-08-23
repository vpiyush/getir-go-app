# getir-go-app
This application has Following Endpoints
1. records (POST) Filters and fetches data from mongoDB
2. pair (GET/POST) Creates and fetches data from an in memory DB

## set up and run
* git clone https://github.com/vpiyush/getir-go-app.git
* cd getir-go-app
* go run main.go

The app will start listining to ":9999"

## run test
* export DEV=1
* go test ./apis

### EndPoints

* https://salty-eyrie-76135.herokuapp.com/api/v1/pair
GET and POST methods are supported, examples are given below

* https://salty-eyrie-76135.herokuapp.com/api/v1/records
Only POST method is supported, example is given below


### 1. In Memory DB endpoints
### 1.1 POST
| Parameters | Description |
| ------ | ----------- |
| key    | string |
| value  | string |

#### Request
```jsx
 https://salty-eyrie-76135.herokuapp.com/api/v1/pair
```
#### Request Payload
```jsx
{
"key": "active-tabs",
"value": "getir"
}
```
#### Error Responses
| Status | Response |
| ------ | ----------- |
| 403 | `{ "message": "key already exists"}` |
| 400 | `{ "message": "{field} value is invalid"}` |

### 1.2 GET
#### Request
| Parameters | Description |
| ------ | ----------- |
| key    | string |

```jsx
 https://salty-eyrie-76135.herokuapp.com/api/v1/pair?key=active-tabs
```
#### Error Responses
| Status | Response |
| ------ | ----------- |
| 404 | `{ "message": "key not found"}` |

### 2. Mongo DB endpoints
#### Request Payload

| Parameters | Description |
| ------ | ----------- |
| startDate   | Date (YYYY-MM-DD) |
| endDate 	  | Date (YYYY-MM-DD) |
| minCount    | int |
| maxCount    | int |

```jsx
{
  "startDate": "2016-01-26",
  "endDate": "2018-02-02",
  "minCount": 2700,
  "maxCount": 3000
}
```

#### Response Payload
```jsx
{
  "code":0,
  "msg":"Success",
  "records":[
    {
    "key":"TAKwGc6Jr4i8Z487",
    "createdAt":"2017-01-28T01:22:14.398Z",
    "totalCount":2800
    },
    {
    "key":"NAeQ8eX7e5TEg7oH",
    "createdAt":"2017-01-27T08:19:14.135Z",
    "totalCount":2900
    }
  ]
}
```
#### Error Responses
| Status | Response |
| ------ | ----------- |
| 500 | `{ "message": "records not found"}` |
| 400 | `{ "message": "{field} value is invalid"}` |
| 400 | `{ "message": "parsing time {value} as \"2006-01-02\": cannot parse {value} as \"2006\"}` |

#### LICENSE
* MIT License
