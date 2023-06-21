package handler

import (
	"errors"
	"net/http"
	"strconv"

	_ "github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/employee"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/gin-gonic/gin"
)

type EmployeeRequest struct {
	Id           int    `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseId  int    `json:"warehouse_id"`
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
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /employees/{id} [get]
func (e *Employee) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "id inválido.")
			return
		}

		employeeID, err := e.service.Get(int(id))
		if err != nil {
			web.Error(ctx, http.StatusNotFound, "funcionário não encontrado.")
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
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /employees [get]
func (e *Employee) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		employees, err := e.service.GetAll()
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "um erro interno ocorreu.")
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
		var req EmployeeRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			web.Error(ctx, http.StatusBadRequest, "existem erros na formatação do json e não foi possível realizar o parse.")
			return
		}
		if req.CardNumberId == "" {
			web.Error(ctx, http.StatusUnprocessableEntity, "necessário adicionar número do cartão.")
			return
		}
		cardNumberId, err := e.service.Exists(req.CardNumberId)
		if err != nil {
			web.Error(ctx, http.StatusNoContent, "não cadastrado.")
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
		var req EmployeeRequest
		if err := ctx.Bind(&req); err != nil {
			web.Error(ctx, http.StatusNotFound, "existem erros na formatação do json e não foi possível realizar o parse.")
			return
		}
		if req.CardNumberId == "" && req.FirstName == "" && req.LastName == "" && req.WarehouseId == 0 {
			web.Error(ctx, http.StatusUnprocessableEntity, "necessário adicionar todas as informações.")
			return
		}
		if req.CardNumberId != "" {
			e.service.Exists(req.CardNumberId)
		}
		employeeId, err := e.service.Save(req.CardNumberId, req.FirstName, req.LastName, req.WarehouseId)
		if err != nil {
			web.Error(ctx, http.StatusNotFound, "funcionário não encontrado.")
			return
		}
		employeeCreated, err := e.service.Get(employeeId)
		web.Success(ctx, http.StatusCreated, employeeCreated)
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
// @Router /employees/{id} [patch]
func (e *Employee) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")
		parsedId, err := strconv.Atoi(id)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "ID inválido.")
			return
		}
		var req employee.EmployeeRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, http.StatusBadRequest, "Corpo da requisição inválido.")
			return
		}
		req.Id = parsedId
		updatedEmployee, err := e.service.Update(req)
		if err != nil {
			if errors.Is(err, employee.ErrNotFound) {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			} else {
				web.Error(c, http.StatusInternalServerError, "Erro interno no servidor.")
				return
			}
		}
		web.Success(c, http.StatusOK, updatedEmployee)
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
// @Failure 404 {object} web.ErrorResponse "NotFound error"
// @Router /employees/{id} [delete]
func (e *Employee) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "id inválido.")
			return
		}

		err = e.service.Delete(int(id))
		if err != nil {
			web.Error(ctx, http.StatusNotFound, "funcionário não encontrado.")
			return
		}
		web.Success(ctx, http.StatusNoContent, nil)
	}
}
