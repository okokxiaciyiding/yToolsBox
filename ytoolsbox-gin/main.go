package main

import (
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type menus struct {
	Data []menus_data `json:"data"`
	Meta menus_meta   `json:"meta"`
}

type menus_data struct {
	Id       int    `json:"id"`
	AuthName string `json:"authname"`
	Path     string `json:"path"`
}

type menus_meta struct {
	Msg         string `json:"msg"`
	Status_code int    `json:"status_code"`
}

func main() {
	r := gin.Default()
	r.Use(Cors()) //开启中间件 允许使用跨域请求
	r.POST("/login", func(c *gin.Context) {

		name := c.PostForm("name")
		password := c.PostForm("password")
		log.Print(name, password)

		// 账户名密码错误
		if password != "admin" || name != "admin" {
			data := map[string]interface{}{
				"data": map[string]interface{}{
					"status_code": 401,
					"message":     "账号密码错误",
				},
			}
			c.JSON(200, data)
			return
		}

		data := map[string]interface{}{
			"data": map[string]interface{}{
				"status_code": 200,
				"message":     "登录成功",
				"token":       "123456",
			},
		}

		c.JSON(200, data)
	})

	r.GET("/menus", func(c *gin.Context) {

		// data := map[string]menus{
		// 	"0": menus{0, "能效总览", "dashboard"},
		// 	"1": menus{1, "工具盒", "toolbox"},
		// 	"2": menus{2, "全局配置", "globalconfig"},
		// 	"3": menus{3, "关于", "about"},
		// 	"meta": map[string]interface{}{
		// 		"msg":         "suc",
		// 		"status_code": 200,
		// 	},
		// }
		data := menus{
			Data: []menus_data{
				{0, "能效总览", "dashboard"},
				{1, "工具盒", "toolbox"},
				{2, "全局配置", "globalconfig"},
				{3, "关于", "about"},
			},
			Meta: menus_meta{
				Msg:         "suc",
				Status_code: 200,
			},
		}

		c.JSON(200, data)
	})
	r.Run(":80") // 监听并在 0.0.0.0:80 上启动服务
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func StructToMapDemo(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}

// func TestStructToMap(){
// 	student := Student{10, "jqw", 18}
// 	data := StructToMapDemo(student)
// 	fmt.Println(data)
