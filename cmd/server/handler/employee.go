package handler

import (
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

// Get godoc
// @Summary Get a employee
// @Description Get a employee based on the provided JSON payload
// @Tags Employees
// @Accept json
// @Produce json
// @Success 200 {object} domain.Employee "Employee"
// @Failure 400 {object} web.ErrorResponse"Validation error"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /employees/:id [get]
func (e *Employee) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "Error: ID inválido.")
			return
		}

		employeeID, err := e.service.Get(int(id))
		if err != nil {
			web.Error(ctx, http.StatusNotFound, "Error: Funcionário não encontrado.")
			return
		}
		web.Success(ctx, http.StatusOK, employeeID)
	}
}

// Get All Employee godoc
// @Summary Get all employee
// @Description Get employee based on the provided JSON payload
// @Tags Employees
// @Accept json
// @Produce json
// @Success 200 {object} []domain.Employee "Employee"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /employees [get]
func (e *Employee) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		employees, err := e.service.GetAll()
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "Error: Funcionários não encontrados.")
			return
		}
		web.Success(ctx, http.StatusOK, employees)
	}
}

// Exists godoc
// @Summary Exist card number
// @Description Validate card number
// @Tags Employees
// @Accept json
// @Produce json
// @Success 204 {object} string "Card number"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /employees/cardNumber [get]
func (e *Employee) Exists() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			web.Error(ctx, http.StatusBadRequest, "Error")
			return
		}
		if req.Card_number_id == "" {
			web.Error(ctx, http.StatusUnprocessableEntity, "Error: Necessário adicionar número do cartão.")
			return
		}
		cardNumberId, err := e.service.Exists(req.Card_number_id)
		if err != nil {
			web.Error(ctx, http.StatusNoContent, "Não cadastrado.")
			return
		}
		web.Success(ctx, http.StatusOK, cardNumberId)
	}
}

// Create godoc
// @Summary Create a new employee
// @Description Create a new employee based on the provided JSON payload
// @Tags Employees
// @Accept json
// @Produce json
// @Param request body domain.Employee true "Employee data"
// @Success 201 {object} domain.Employee "Created employee"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /employees [post]
func (e *Employee) Save() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		if err := ctx.Bind(&req); err != nil {
			web.Error(ctx, http.StatusNotFound, "Error")
			return
		}
		if req.Card_number_id == "" && req.First_name == "" && req.Last_name == "" && req.Warehouse_id == 0 {
			web.Error(ctx, http.StatusUnprocessableEntity, "Error: Necessário adicionar todas as informações.")
			return
		}
		employeeSaved, err := e.service.Save(req.Card_number_id, req.First_name, req.Last_name, req.Warehouse_id)
		if err != nil {
			web.Error(ctx, http.StatusNotFound, "Error")
			return
		}
		web.Success(ctx, http.StatusCreated, employeeSaved)
	}
}

// Update godoc
// @Summary Update a employee
// @Description Update employee based on the provided JSON payload
// @Tags Employees
// @Accept json
// @Produce json
// @Param request body domain.Employee true "Employee data"
// @Success 200 {object} domain.Employee "Updated employee"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /employees/:id [patch]
func (e *Employee) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "ID inválido.")
			return
		}

		var req request
		err = c.Bind(&req)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Error.")
			return
		}

		if req.Card_number_id == "" && req.First_name == "" && req.Last_name == "" && req.Warehouse_id == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "Informe pelo menos um campo para concluir a atualização.")
			return
		}

		employee, err := e.service.Get(id)
		if err != nil {
			web.Error(c, http.StatusNotFound, "Funcionário não encontrado.")
			return
		}

		if req.First_name != "" {
			employee.FirstName = req.First_name
		}

		if req.Last_name != "" {
			employee.LastName = req.Last_name
		}

		if req.Warehouse_id != 0 {
			employee.WarehouseID = req.Warehouse_id
		}

		err = e.service.Update(employee)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}

		web.Success(c, http.StatusOK, employee)
	}
}

// Delete godoc
// @Summary Delete employee
// @Description Delete employee based on the provided JSON payload
// @Tags Employees
// @Accept json
// @Produce json
// @Success 204
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /employees/:id [delete]
func (e *Employee) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "Error: ID inválido.")
			return
		}

		err = e.service.Delete(int(id))
		if err != nil {
			web.Error(ctx, http.StatusNotFound, "Error")
			return
		}
		web.Success(ctx, http.StatusNoContent, "Funcionário deletado com sucesso.")
	}
}
