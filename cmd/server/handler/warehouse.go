package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	service "github.com/extmatperez/meli_bootcamp_go_w2-1/internal/warehouse"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/gin-gonic/gin"
)

type Warehouse struct {
	warehouseService service.WarehouseService
}

type WarehouseData struct {
	ID                 int     `json:"id"`
	Address            string  `json:"address"`
	Telephone          string  `json:"telephone"`
	WarehouseCode      string  `json:"warehouse_code"`
	MinimumCapacity    int     `json:"minimum_capacity"`
	MinimumTemperature float64 `json:"minimum_temperature"`
}

func NewWarehouse(warehouseService service.WarehouseService) *Warehouse {
	return &Warehouse{
		warehouseService: warehouseService,
	}
}

type UpdateRequest struct {
	ID                 *int    `json:"id"`
	Address            *string `json:"address"`
	Telephone          *string `json:"telephone"`
	WarehouseCode      *string `json:"warehouse_code"`
	MinimumCapacity    *int    `json:"minimum_capacity"`
	MinimumTemperature *int    `json:"minimum_temperature"`
}

func (w UpdateRequest) IsBlank() bool {
	return w.ID == nil && w.Address == nil && w.Telephone == nil && w.WarehouseCode == nil && w.MinimumCapacity == nil && w.MinimumTemperature == nil
}

// GetByID godoc
// @Summary Get warehouse by ID
// @Tags Warehouses
// @Description This endpoint retrieves the information of a warehouse by its ID
// @Accept  json
// @Produce  json
// @Param id path string true "Warehouse ID"
// @Success 200 {object} domain.Warehouse
// @Failure 400 {object} web.ErrorResponse "Invalid warehouse ID"
// @Failure 404 {object} web.ErrorResponse "Warehouse not found"
// @Router /warehouses/{id} [get]
func (w *Warehouse) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		warehouseID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Invalid warehouse ID")
			return
		}
		warehouseData, err := w.warehouseService.GetWarehouse(context.TODO(), warehouseID)
		if err != nil {
			web.Error(c, http.StatusNotFound, "Warehouse not found")
			return
		}
		web.Response(c, http.StatusOK, warehouseData)
	}
}

// GetAll godoc
// @Summary Get all warehouses
// @Tags Warehouses
// @Description Get all warehouses
// @Accept  json
// @Produce  json
// @Success 200 {array} domain.Warehouse
// @Failure 500 {object} web.ErrorResponse "Failed to get the warehouses"
// @Router /warehouses [get]
func (w *Warehouse) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		warehouses, err := w.warehouseService.GetAllWarehouses(c)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, "Failed to get the warehouses")
			return
		}
		response := gin.H{
			"data": warehouses,
		}
		web.Success(c, http.StatusOK, response)
	}
}

// Create godoc
// @Summary Create a warehouse
// @Tags Warehouses
// @Description Create a new warehouse
// @Accept  json
// @Produce  json
// @Param token header string true "Authentication token"
// @Param warehouseData body WarehouseData true "Warehouse data to store"
// @Success 201 {object} domain.Warehouse
// @Failure 400 {object} web.ErrorResponse "Invalid request body"
// @Failure 409 {object} web.ErrorResponse "Codigo Warehouse incorreto, rever"
// @Router /warehouses [post]
func (w *Warehouse) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		usedCodes := make(map[string]bool)

		var warehouseData domain.Warehouse
		if err := c.ShouldBindJSON(&warehouseData); err != nil {
			web.Error(c, http.StatusBadRequest, "Invalid request body")
			return
		}
		if warehouseData.Address == "" || warehouseData.WarehouseCode == "" {
			web.Error(c, http.StatusBadRequest, "All fields must be provided and the phone number must be in the proper format")
			return
		}
		if usedCodes[warehouseData.WarehouseCode] {
			web.Error(c, http.StatusBadRequest, "warehouse_code already registered")
			return
		}
		usedCodes[warehouseData.WarehouseCode] = true

		createdWarehouse, err := w.warehouseService.CreateWarehouse(context.TODO(), warehouseData)
		if err != nil {
			web.Error(c, http.StatusConflict, "Codigo Warehouse duplicado, rever")
			return
		}
		web.Response(c, http.StatusCreated, createdWarehouse)
	}
}

// Update godoc
// @Summary Update a warehouse
// @Tags Warehouses
// @Description This endpoint allows updating the information of an existing warehouse
// @Accept  json
// @Produce  json
// @Param token header string true "Authentication token"
// @Success 200 {object} domain.Warehouse "Warehouse updated successfully"
// @Failure 400 {object} web.ErrorResponse "Invalid data"
// @Failure 500 {object} web.ErrorResponse "Failed to update warehouse"
// @Router /warehouses [put]
func (w *Warehouse) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

		var updatedWarehouse domain.Warehouse
		if err := c.ShouldBindJSON(&updatedWarehouse); err != nil {
			web.Error(c, http.StatusBadRequest, "dados inválidos fornecidos.")
			return
		}
		err := w.warehouseService.UpdateWarehouse(context.TODO(), updatedWarehouse)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, "falha ao atualizar o armazém.")
			return
		}
		web.Response(c, http.StatusOK, updatedWarehouse)
	}
}

// UpdateByID godoc
// @Summary Update a warehouse by ID
// @Tags Warehouses
// @Description This endpoint allows updating the information of an existing warehouse identified by its ID
// @Accept  json
// @Produce  json
// @Param token header string true "Authentication token"
// @Param id path string true "Warehouse ID"
// @Param warehouse body WarehouseData true "Updated warehouse data"
// @Success 200 {object} WarehouseData "Updated warehouse"
// @Failure 400 {object} web.ErrorResponse "Invalid data provided"
// @Failure 404 {object} web.ErrorResponse "Warehouse not found"
// @Failure 500 {object} web.ErrorResponse "Failed to update warehouse"
// @Router /warehouses/{id} [patch]
func (w *Warehouse) UpdateByID() gin.HandlerFunc {
	return func(c *gin.Context) {

		warehouseIDStr := c.Param("id")
		warehouseID, err := strconv.Atoi(warehouseIDStr)

		var request UpdateRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			web.Error(c, http.StatusBadRequest, "dados inválidos fornecidos.")
			return
		}
		existingWarehouse, err := w.warehouseService.GetWarehouse(c.Request.Context(), warehouseID)
		if err != nil {
			web.Error(c, http.StatusNotFound, "armazém não encontrado.")
			return
		}
		fmt.Println(request.IsBlank())
		if request.IsBlank() {
			web.Error(c, http.StatusBadRequest, "pelo menos um campo deve ser informado.")
			return
		}

		if request.Address != nil {
			existingWarehouse.Address = *request.Address
		}
		if request.Telephone != nil {
			existingWarehouse.Telephone = *request.Telephone
		}
		if request.WarehouseCode != nil {
			existingWarehouse.WarehouseCode = *request.WarehouseCode
		}
		if request.MinimumCapacity != nil {
			existingWarehouse.MinimumCapacity = *request.MinimumCapacity
		}
		if request.MinimumTemperature != nil {
			existingWarehouse.MinimumTemperature = *request.MinimumTemperature
		}

		err = w.warehouseService.UpdateWarehouse(c.Request.Context(), existingWarehouse)

		if err != nil {
			web.Error(c, http.StatusInternalServerError, "falha ao atualizar o armazém.")
			return
		}
		web.Success(c, http.StatusOK, existingWarehouse)
	}
}

// Delete godoc
// @Summary Delete a warehouse by ID
// @Tags Warehouses
// @Description This endpoint allows deleting a warehouse by its ID
// @Accept  json
// @Produce  json
// @Param id path string true "Warehouse ID"
// @Success 204 {object} domain.Warehouse "No Content"
// @Failure 400 {object} web.ErrorResponse "Invalid ID"
// @Failure 404 {object} web.ErrorResponse "Warehouse not found"
// @Failure 500 {object} web.ErrorResponse "Failed to delete warehouse"
// @Router /warehouses/{id} [delete]"Failed to delete warehouse"
func (w *Warehouse) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {

		warehouseID := c.Param("id")
		id, err := strconv.Atoi(warehouseID)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "id inválido.")
		}
		existingWarehouse, err := w.warehouseService.GetWarehouse(context.TODO(), id)
		if err != nil {
			web.Error(c, http.StatusNotFound, "armazém não encontrado.")
			return
		}
		err = w.warehouseService.DeleteWarehouse(context.TODO(), existingWarehouse.ID)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, "falha ao excluir armazém.")
			return
		}
		web.Success(c, http.StatusNoContent, nil)
	}
}
