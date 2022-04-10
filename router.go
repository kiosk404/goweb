/**
* @Author: kiosk
* @Mail: weijiaxiang007@foxmail.com
* @Date: 2022/4/9
**/
package main

import (
	"fmt"
	"goweb/framework"
	"net/http"
	"time"
)

// 注册路由规则
func registerRouter(core *framework.Core) {
	// 静态路由+HTTP方法匹配
	core.Get("/v1/hello", func(c *framework.Context) error {
		time.Sleep(20 * time.Second)
		return c.Text(http.StatusOK, "/v1/hello")
	})

	// 批量通用前缀
	subjectApi := core.Group("/subject")
	{
		// 动态路由
		subjectApi.Get("/:id", func(c *framework.Context) error {
			return c.Text(http.StatusOK, fmt.Sprintf("/subject/%d",1))
		})
		subjectApi.Get("/list/all", func(c *framework.Context) error {
			return c.Text(http.StatusOK, fmt.Sprintf("/subject/all"))
		})

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Get("/name", func (c *framework.Context) error {
				return c.Text(http.StatusOK, fmt.Sprintf("/subject/info/name"))
			})
		}
	}
}
