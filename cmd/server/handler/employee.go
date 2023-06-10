package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/employee"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Id             int    `json:"id"`
	Card_number_id string `json:"card_number_id"`
	First_name     string `json:"first_name"`
	Last_name      string `json:"last_name"`
	Warehouse_id   int    `json:"warehouse_id"`
}

type Employee struct {
	service employee.Service
}

func NewEmployee(e employee.Service) *Employee {
	return &Employee{
		service: e,
	}
}

func (e *Employee) Get() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (e *Employee) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		employees, err := e.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"Error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, employees, ""))
	}
}

func (e *Employee) Exists() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "ID inválido"})
			return
		}
		employee, err := e.service.Get(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, employee, ""))
	}
}

func (e *Employee) Save() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"Error": err.Error(),
			})
			return
		}
		employee, err := e.service.Save(req.Card_number_id, req.First_name, req.Last_name, req.Warehouse_id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, employee)
	}
}

func (e *Employee) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "ID inválido"})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		if req.Card_number_id == "" || req.First_name == "" || req.Last_name == "" || req.Warehouse_id == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Todas as informações do usuário devem ser preenchidas."})
			return
		}

		employee, err := e.service.Update(int(id), req.Card_number_id, req.First_name, req.Last_name, req.Warehouse_id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, employee, ""))
	}
}

func (e *Employee) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "ID inválido"})
			return
		}

		err = e.service.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"Data": fmt.Sprintf("O funcionário %d foi removido", id)})
	}
}
