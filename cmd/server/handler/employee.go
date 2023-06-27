package handler

import (
	"net/http"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/employee"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/gin-gonic/gin"
)

type Employee struct {
	service employee.Service
}

type CreateEmployeeRequest struct {
	CardNumberID *string `json:"card_number_id" binding:"required"`
	FirstName    *string `json:"first_name" binding:"required"`
	LastName     *string `json:"last_name" binding:"required"`
	WarehouseID  *int    `json:"warehouse_id" binding:"required"`
}

func (r CreateEmployeeRequest) ToEmployee() domain.Employee {
	return domain.Employee{
		ID:           0,
		CardNumberID: *r.CardNumberID,
		FirstName:    *r.FirstName,
		LastName:     *r.LastName,
		WarehouseID:  *r.WarehouseID,
	}
}

type UpdateEmployeeRequest struct {
	CardNumberID *string `json:"card_number_id"`
	FirstName    *string `json:"first_name"`
	LastName     *string `json:"last_name"`
	WarehouseID  *int    `json:"warehouse_id"`
}

func (r UpdateEmployeeRequest) ToUpdateEmployee() domain.UpdateEmployee {
	return domain.UpdateEmployee{
		CardNumberID: r.CardNumberID,
		FirstName:    r.FirstName,
		LastName:     r.LastName,
		WarehouseID:  r.WarehouseID,
	}
}

func NewEmployee(e employee.Service) *Employee {
	return &Employee{
		service: e,
	}
}

// GetAll Employee godoc
// @Summary List all employees
// @Description List all employees
// @Tags Employees
// @Produce json
// @Success 200 {object} []domain.Employee "Employee"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /employees [get]
func (e *Employee) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		employees := e.service.GetAll(ctx.Request.Context())
		web.Success(ctx, http.StatusOK, employees)
	}
}

// Get godoc
// @Summary Get a employee by ID
// @Description Get a employee based on the ID parameter
// @Tags Employees
// @Produce json
// @Success 200 {object} domain.Employee "Employee"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "ID not found"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /employees/{id} [get]
func (e *Employee) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.GetInt("Id")

		employee, err := e.service.Get(ctx.Request.Context(), id)

		if err != nil {
			if apperr.Is[*apperr.ResourceNotFound](err) {
				web.Error(ctx, http.StatusNotFound, err.Error())
				return
			}
		}

		web.Success(ctx, http.StatusOK, employee)
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
// @Failure 422 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "ID not found"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /employees [post]
func (e *Employee) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := ctx.MustGet(RequestParamContext).(CreateEmployeeRequest)

		createdEmployee, err := e.service.Create(ctx.Request.Context(), request.ToEmployee())

		if err != nil {
			if apperr.Is[*apperr.ResourceAlreadyExists](err) {
				web.Error(ctx, http.StatusConflict, err.Error())
				return
			}
		}

		web.Success(ctx, http.StatusCreated, createdEmployee)
	}
}

// Update godoc
// @Summary Update a employee
// @Description Update employee based on the provided JSON payload
// @Tags Employees
// @Accept json
// @Produce json
// @Param request body domain.Employee true "Employee data to update"
// @Success 200 {object} domain.Employee "Updated employee"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 422 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "ID not found error"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /employees/{id} [patch]
func (e *Employee) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.GetInt("Id")
		request := ctx.MustGet(RequestParamContext).(UpdateEmployeeRequest)

		response, err := e.service.Update(ctx.Request.Context(), id, request.ToUpdateEmployee())

		if err != nil {
			if apperr.Is[*apperr.ResourceNotFound](err) {
				web.Error(ctx, http.StatusNotFound, err.Error())
				return
			}

			if apperr.Is[*apperr.ResourceAlreadyExists](err) {
				web.Error(ctx, http.StatusConflict, err.Error())
				return
			}
		}

		web.Success(ctx, http.StatusOK, response)
	}
}

// Delete godoc
// @Summary Delete employee
// @Description Delete employee based on ID
// @Tags Employees
// @Produce json
// @Success 204 "No content"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "ID not found error"
// @Router /employees/{id} [delete]
func (e *Employee) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.GetInt("Id")

		err := e.service.Delete(ctx.Request.Context(), id)

		if err != nil {
			if apperr.Is[*apperr.ResourceNotFound](err) {
				web.Error(ctx, http.StatusNotFound, err.Error())
				return
			}
		}

		web.Success(ctx, http.StatusNoContent, nil)
	}
}
