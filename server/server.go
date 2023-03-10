package server

import (
	"fmt"
	"ldapAdapter/internal/ldap"
	"net/http"
)
import "github.com/gin-gonic/gin"

func Start() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/get", func(c *gin.Context) {

		// filter
		//
	})

	r.POST("/login", func(c *gin.Context) {
		var body ldap.User
		err := c.BindJSON(&body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "params unavailable"})
			return
		}
		fmt.Printf("Username:%v\n", body.Username)
		fmt.Printf("Password:%v\n", body.Password)

		// 1.调用 bind 校验用户信
		ldap.Bind(body)
		// 2.返回给用户的信息信息
	})

	r.Run()
}
