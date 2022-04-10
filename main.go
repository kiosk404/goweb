/**
* @Author: kiosk
* @Mail: weijiaxiang007@foxmail.com
* @Date: 2022/4/9
**/
package main

import (
	"goweb/framework"
	"net/http"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)

	server := http.Server{
		Addr:              ":8080",
		Handler:           core,
	}
	// 监听8080端口
	server.ListenAndServe()
}

