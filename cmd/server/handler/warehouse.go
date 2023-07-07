package handler

import (
	"net/http"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/warehouse"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/gin-gonic/gin"
)

type CreateWarehouseRequest struct {
	Address            *string `json:"address" binding:"required"`
	Telephone          *string `json:"telephone" binding:"required,e164"`
	WarehouseCode      *string `json:"warehouse_code" binding:"required"`
	MinimumCapacity    *int    `json:"minimum_capacity" binding:"required"`
	MinimumTemperature *int    `json:"minimum_temperature" binding:"required"`
}

func (w CreateWarehouseRequest) ToWarehouse() domain.Warehouse {
	return domain.Warehouse{
		ID:                 0,
		Address:            *w.Address,
		Telephone:          *w.Telephone,
		WarehouseCode:      *w.WarehouseCode,
		MinimumCapacity:    *w.MinimumCapacity,
		MinimumTemperature: *w.MinimumTemperature,
	}
}

type UpdateWarehouseRequest struct {
	Address            *string `json:"address"`
	Telephone          *string `json:"telephone"`
	WarehouseCode      *string `json:"warehouse_code"`
	MinimumCapacity    *int    `json:"minimum_capacity"`
	MinimumTemperature *int    `json:"minimum_temperature"`
}

func (w UpdateWarehouseRequest) ToUpdateWarehouse() domain.UpdateWarehouse {
	return domain.UpdateWarehouse{
		Address:            w.Address,
		Telephone:          w.Telephone,
		WarehouseCode:      w.WarehouseCode,
		MinimumCapacity:    w.MinimumCapacity,
		MinimumTemperature: w.MinimumTemperature,
	}
}

type Warehouse struct {
	service warehouse.Service
}

func NewWarehouse(w warehouse.Service) *Warehouse {
	return &Warehouse{
		service: w,
	}
}

// Get godoc
// @Summary Get a warehouse by id
// @Tags Warehouses
// @Description Get a warehouse based on the provided id
// @Accept  json
// @Produce  json
// @Param id path string true "Warehouse id"
// @Success 200 {object} domain.Warehouse "Obtained warehouse"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "Resource not found error"
// @Router /warehouses/{id} [get]
func (w *Warehouse) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetInt("Id")

		warehouse, err := w.service.Get(c.Request.Context(), id)

		if err != nil {
			if apperr.Is[*apperr.ResourceNotFound](err) {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			}
		}

		web.Success(c, http.StatusOK, warehouse)
	}
}

// GetAll godoc
// @Summary List warehouses
// @Tags Warehouses
// @Description List all warehouses
// @Accept  json
// @Produce  json
// @Success 200 {array} domain.Warehouse "List of warehouses"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /warehouses [get]
func (w *Warehouse) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		warehouses := w.service.GetAll(c.Request.Context())
		web.Success(c, http.StatusOK, warehouses)
	}
}

// Create godoc
// @Summary Create a warehouse
// @Tags Warehouses
// @Description Create a new warehouse based on the provided JSON payload
// @Accept  json
// @Produce  json
// @Param request body CreateWarehouseRequest true "Warehouse data"
// @Success 201 {object} domain.Warehouse "Created warehouse"
// @Failure 400 {object} web.ErrorResponse "Invalid data"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 422 {object} web.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /warehouses [post]
func (w *Warehouse) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := c.MustGet(RequestParamContext).(CreateWarehouseRequest)

		created, err := w.service.Create(c, request.ToWarehouse())

		if err != nil {
			if apperr.Is[*apperr.ResourceAlreadyExists](err) {
				web.Error(c, http.StatusConflict, err.Error())
				return
			}
		}

		web.Success(c, http.StatusCreated, created)
	}
}

// Update godoc
// @Summary Update a warehouse
// @Tags Warehouses
// @Description Update an existent warehouse based on the provided id and JSON payload
// @Accept  json
// @Produce  json
// @Param id path string true "Warehouse id"
// @Param request body UpdateWarehouseRequest true "Warehouse data"
// @Success 200 {object} domain.Warehouse "Updated warehouse"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "Resource not found"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /warehouses [patch]
func (w *Warehouse) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetInt("Id")
		request := c.MustGet(RequestParamContext).(UpdateWarehouseRequest)

		updated, err := w.service.Update(c.Request.Context(), id, request.ToUpdateWarehouse())

		if err != nil {
			if apperr.Is[*apperr.ResourceNotFound](err) {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			}
			if apperr.Is[*apperr.ResourceAlreadyExists](err) {
				web.Error(c, http.StatusConflict, err.Error())
				return
			}
		}

		web.Success(c, http.StatusOK, updated)
	}
}

// Delete godoc
// @Summary Delete a warehouse
// @Tags Warehouses
// @Description Delete a warehouse based on the provided id
// @Accept  json
// @Produce  json
// @Param id path string true "Warehouse id"
// @Success 204
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "Resource not found error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /warehouses/{id} [delete]
func (w *Warehouse) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetInt("Id")

		err := w.service.Delete(c, id)

		if err != nil {
			if apperr.Is[*apperr.ResourceNotFound](err) {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			}
		}

		web.Success(c, http.StatusNoContent, nil)
	}
}
