package controller

//
//import (
//	"fmt"
//	user "gin_project/Users"
//	"gin_project/connection"
//	"gin_project/model"
//	"net/http"
//
//	_ "github.com/lib/pq"
//
//	"github.com/gin-gonic/gin"
//)
//
//func GetUsers(c *gin.Context) {
//	db := connection.Connection()
//	rows, err := db.Query("select * from Products")
//	if err != nil {
//		panic(err)
//	}
//	defer rows.Close()
//	products := []model.Product{}
//	var p model.Product
//
//	for rows.Next() {
//		err := rows.Scan(&p.Id, &p.Model, &p.Company, &p.Price)
//		if err != nil {
//			fmt.Println(err)
//			continue
//		}
//		products = append(products, p)
//	}
//	for _, p := range products {
//		fmt.Println(p.Id, p.Model, p.Company, p.Price)
//	}
//
//	c.IndentedJSON(http.StatusOK, user.List)
//}
//
//func PostUser(c *gin.Context) {
//	var newUser model.Model
//
//	if err := c.BindJSON(&newUser); err != nil {
//		return
//	}
//
//	user.List = append(user.List, newUser)
//	c.IndentedJSON(http.StatusCreated, newUser)
//}
//
//func GetByID(c *gin.Context) {
//	id := c.Param("id")
//	var found bool
//	for _, value := range user.List {
//		if value.ID == id {
//			c.IndentedJSON(http.StatusOK, value)
//			found = true
//		}
//	}
//	if !found {
//		c.IndentedJSON(http.StatusNotFound, gin.H{"massage": "user not found"})
//	}
//}
//
//func DeleteById(c *gin.Context) {
//	id := c.Param("id")
//
//	var filtered []model.Model
//
//	for _, value := range user.List {
//		if value.ID != id {
//			filtered = append(filtered, value)
//		}
//	}
//	user.List = filtered
//	c.IndentedJSON(http.StatusOK, user.List)
//}
//
//func UpdateUser(c *gin.Context) {
//	var newInfoUser model.Model
//	if err := c.BindJSON(&newInfoUser); err != nil {
//		return
//	}
//	id := newInfoUser.ID
//	for index, value := range user.List {
//		if value.ID == id {
//			user.List[index] = newInfoUser
//		}
//	}
//	c.IndentedJSON(http.StatusOK, newInfoUser)
//}
