package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	// "net/http"
)

type User struct{
	FirstName string
	LastName string
	Email string
}

func listUsers(c *gin.Context){
	var users = []User{
		User{FirstName:"John",LastName:"Doe",Email:"john.doe@mail.com"},
		User{FirstName:"Jane",LastName:"Doe",Email:"jane.doe@mail.com"},
	}
	c.HTML(200, "users.html", gin.H{
		"users":users,
	})
}

func getCreateForm(c *gin.Context){	
	c.HTML(200, "create.html", gin.H{})
}

func createUser(c *gin.Context){	
	FirstName := c.PostForm("FirstName")
	LastName := c.PostForm("LastName")
	Email := c.PostForm("Email")	
	fmt.Println(FirstName)
	fmt.Println(LastName)
	fmt.Println(Email)		
	c.Redirect(302,c.FullPath())
}

func main(){
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/users",listUsers)
	router.GET("/create",getCreateForm)
	router.POST("/create",createUser)
	router.Run()
}