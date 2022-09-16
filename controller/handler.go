package controller

import (
	"database/sql"
	"fmt"
	"gin_project/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	db *sql.DB
}

func (h Handler) GetUsers(c *gin.Context) {
	rows, err := h.db.Query("select * from products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var products []model.Product
	var p model.Product

	for rows.Next() {
		err := rows.Scan(&p.Id, &p.Model, &p.Company, &p.Price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}

	c.IndentedJSON(http.StatusOK, products)
}

func (h Handler) PostUser(c *gin.Context) {
	var newUser model.Product
	var id int
	if err := c.BindJSON(&newUser); err != nil {
		panic(err)
	}
	h.db.QueryRow("insert into products (model, company, price) values ($1, $2, $3) returning id",
		newUser.Model, newUser.Company, newUser.Price).Scan(&id)
	newUser.Id = id
	c.IndentedJSON(http.StatusCreated, newUser)
}

func (h Handler) GetByID(c *gin.Context) {
	id := c.Param("id")
	rows, err := h.db.Query("select * from Products where id = $1", id)
	if err != nil {
		panic(err)
	}
	prod := model.Product{}
	rows.Next()
	err = rows.Scan(&prod.Id, &prod.Model, &prod.Company, &prod.Price)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"massage": "user not found"})
	} else {
		c.IndentedJSON(http.StatusOK, prod)
	}
}

func (h Handler) DeleteById(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) UpdateUser(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

type UserHandler interface {
	GetUsers(c *gin.Context)
	PostUser(c *gin.Context)
	GetByID(c *gin.Context)
	DeleteById(c *gin.Context)
	UpdateUser(c *gin.Context)
}

func NewHandler(db *sql.DB) UserHandler {
	return &Handler{
		db: db,
	}
}
