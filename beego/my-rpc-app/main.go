package main

import (
	"fmt"
	"reflect"

	"my-rpc-app/models"
	"github.com/hprose/hprose-golang/rpc"

	"github.com/astaxie/beego"
)

func logInvokeHandler(
	name string,
	args []reflect.Value,
	context rpc.Context,
	next rpc.NextInvokeHandler) (results []reflect.Value, err error) {
	fmt.Printf("%s(%v) = ", name, args)
	results, err = next(name, args, context)
	fmt.Printf("%v %v\r\n", results, err)
	return
}

func main() {
	// Create WebSocketServer
	// service := rpc.NewWebSocketService()

	// Create Http Server
	service := rpc.NewHTTPService()

	// Use Logger Middleware
	service.AddInvokeHandler(logInvokeHandler)

	// Publish Functions
	service.AddFunction("AddOne", models.AddOne)
	service.AddFunction("GetOne", models.GetOne)

	// Start Service
	beego.Handler("/", service)
	beego.Run()
}
