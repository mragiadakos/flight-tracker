# FlightTracker
FlightTracker is a simple microservice API to track the source and destination of a person from a list of paths.


## Table of Contents

- [Install](#install)
- [Examples](#examples)
- [REST API](#rest-api)

## Install

### Compile and run
```shell
go build

./flight-tracker
```

### Create docker image and run
```shell
docker build -t flight-tracker .

docker run -it --rm -p 1234:1234 flight-tracker
```

## Examples
The best way to guide you is through examples
### Example for curl
```shell
curl --location --request GET 'localhost:1234/sort' \
--header 'Content-Type: application/json' \
--data-raw '{
    "paths": [["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]
}'
```

### Examples from postman
Import file flight-tracker.postman_collection.json to the Postman client and find the folder 'flight-tracker'



## REST-API
**Sort paths**
----
  Returns JSON data with the sorted path based on a list of paths.

* **URL**

  /sort

* **Method:**

  `GET`

*  **Headers**

   Content-Type: application/json
  
*  **URL Params**

   None

* **Data Params**
  <br /> The 'paths' are a list of tuples. Each tuple is a path with two strings. The first string on the path is the source and the second string is the destination.

  **Required:**
 
   `{ "paths": [[][]string]}`

* **Success Response:**
  <br /> The 'path' is a list of two strings. The first string is the source and the second string is the destination.
  * **Code:** 200 <br />
    **Content:** `{ "status" : "success", "path" : [[]string] }`
 
* **Error Response:**

  * **Code:** 400 BAD REQUEST <br />
    **Content:** `{  "status" : "failure", "message" : [string] }`