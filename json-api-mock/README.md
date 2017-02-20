# json-api-mock

## directory tree
- [handlers.go](https://github.com/yagi2/go-sandbox/tree/master/json-api-mock/handlers.go) : API Handlers Function.
- [main.go](https://github.com/yagi2/go-sandbox/tree/master/json-api-mock/main.go)     : golang main file.(define recipt port number)
- [response.go](https://github.com/yagi2/go-sandbox/tree/master/json-api-mock/response.go) : define response json struct.
- [routes.go](https://github.com/yagi2/go-sandbox/tree/master/json-api-mock/routes.go)   : define URI routing.

## command
* run local  
` $ go run *.go`  

* build and run for Android  
` $ GOOS=linux GOARCH=arm go build`  
` $ adb push json-api-mock /local/data/tmp`  
` $ adb shell`  
` [Android]$ cd /local/data/tmp`  
` [Android]$ chmod 755 json-api-mock `  
` [Android]$ ./json-api-mock `  