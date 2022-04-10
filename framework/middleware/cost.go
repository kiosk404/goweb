/**
* @Author: kiosk
* @Mail: weijiaxiang007@foxmail.com
* @Date: 2022/4/10
**/

package middleware

import (
	"fmt"
	"goweb/framework"
	"time"
)

func Cost() framework.ControllerHandler {
	// 使用函数回调
	return func(c *framework.Context) error {
		// 记录开始时间
		start := time.Now()

		// 使用next执行具体的业务逻辑
		c.Next()

		// 记录结束时间
		end := time.Now()
		cost := end.Sub(start)
		fmt.Printf("api uri: %v, cost: %v \n", c.GetRequest().RequestURI, cost.Seconds())

		return nil
	}
}