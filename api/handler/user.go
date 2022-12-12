package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/asadbekGo/golang_crud/models"
	"github.com/asadbekGo/golang_crud/storage"
)

func (h *HandlerV1) Create(c *gin.Context) {

	var (
		user models.User
	)

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Printf("error whiling create: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := storage.Create(h.db, user)
	if err != nil {
		log.Printf("error whiling create: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling create"))
		return
	}

	userId, err := storage.GetById(h.db, id)
	if err != nil {
		log.Printf("error whiling get by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id"))
		return
	}

	c.JSON(http.StatusOK, userId)
}

func (h *HandlerV1) GetById(c *gin.Context) {

	id := c.Param("id")

	user, err := storage.GetById(h.db, id)
	if err != nil {
		log.Printf("error whiling get by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id"))
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *HandlerV1) Update(c *gin.Context) {

	var (
		user1 models.User
	)

	err := c.ShouldBindJSON(&user1)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := storage.Update(h.db, user1)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update"))
		return
	}

	if rowsAffected == 0 {
		log.Fatalf("not found rows")
	}

	user, err := storage.GetById(h.db, user1.Id)
	if err != nil {
		log.Printf("error whiling get by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id"))
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *HandlerV1) Delete(c *gin.Context) {

	var (
		user1 models.User
	)

	err := c.ShouldBindJSON(&user1)
	if err != nil {
		log.Printf("error whiling delete: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = storage.Delete(h.db, user1.Id)
	if err != nil {
		log.Printf("error whiling get delete: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling delete id"))
		return
	}

}

func (h *HandlerV1) GetList(c *gin.Context) {

	users, err := storage.GetList(h.db)
	if err != nil {
		log.Printf("error whiling get list: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get list"))
		return
	}

	c.JSON(http.StatusOK, users)
}

// update
// patch
// delete

// github push
