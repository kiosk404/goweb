/**
* @Author: kiosk
* @Mail: weijiaxiang007@foxmail.com
* @Date: 2022/4/11
**/
package middleware

import (
	"context"
	"fmt"
	"goweb/framework"
	"log"
	"net/http"
	"time"
)

func Timeout(d time.Duration) framework.ControllerHandler {
	// 使用函数回调
	return func(c *framework.Context) error {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)
		// 执行业务逻辑前预操作：初始化超时context
		durationCtx, cancel := context.WithTimeout(c.BaseContext(), d)
		defer cancel()

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()
			// 使用next执行具体的业务逻辑
			c.Next()

			finish <- struct{}{}
		}()
		// 执行业务逻辑后操作
		select {
		case p := <-panicChan:
			c.Json(500, "server internal error")
			log.Println(p)
		case <-finish:
			fmt.Println("finish")
		case <-durationCtx.Done():
			fmt.Println("超时了")
			c.SetHasTimeout()
			c.Json(http.StatusBadGateway, "time out")
		}
		return nil
	}
}
