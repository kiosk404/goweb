/**
* @Author: kiosk
* @Mail: weijiaxiang007@foxmail.com
* @Date: 2022/4/9
**/
package main

import (
	"goweb/framework"
	"goweb/framework/middleware"
	"net/http"
	"time"
)

func main() {
	core := framework.NewCore()

	core.Use(middleware.Timeout(1 * time.Second))
	core.Use(middleware.Cost())

	registerRouter(core)
	server := http.Server{
		Addr:              ":8080",
		Handler:           core,
	}
	// 监听8080端口
	server.ListenAndServe()
}

