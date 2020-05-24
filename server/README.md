# Server

This API is developed with **Go** using **Gin framework** (https://github.com/gin-gonic/gin). 
It includes the necessary endpoints to handle the crud operations that the frontend requires.


## Requirements 
* GO >= 1.9
* MongoDB >= 4.0


## Installation 
* Create **.env.local** and set it up with environment variables. 
Use **.dist** files as examples.


## Parse and generate API documentation
* Run following command in the project's root folder:
```
$GOPATH/bin/swag init --generalInfo=cmd/items/main.go --output=cmd/items/docs
```


## Run application
* Run following command in the project's root folder:
```
go run cmd/items/main.go
```


## See API documentation
* Run application and browse to the following url:
```
http://{domain}/doc/index.html
```


## Check API Health & Metrics
* Run application and browse to the following urls:
```
http://{domain}/health
http://{domain}/metrics
```