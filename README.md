# getir-go-app
This application has Following Endpoints
1. records (POST) Filters and fetches data from mongoDB
2. pair (GET/POST) Creates and fetches data from an in memory DB

## set up and run
* git clone https://github.com/vpiyush/getir-go-app.git
* cd getir-go-app
* go build
* go run main.go

## run test
* go test *.go

### EndPoints
TODO

### 1. In Memory DB endpoints
###1.1 POST
| Parameters | Description |
| ------ | ----------- |
| key    | string |
| value  | string |
#### Request Payload
```jsx
{
"key": "active-tabs",
"value": "getir"
}
```
#### Response Payload
```jsx
{
"key": "active-tabs",
"value": "getir"
}
```

###1.2 GET
#### Request Payload
| Parameters | Description |
| ------ | ----------- |
| key    | string |

```jsx
{
"key": "active-tabs",
}
```
#### Response Payload
```jsx
{
"key": "active-tabs",
"value": "getir"
}
```

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

#### Success Response Payload
