package handler

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/warehouse"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/gin-gonic/gin"
)

type Warehouse struct {
	warehouseService warehouse.Service
}

type WarehouseData struct {
	ID                 int     `json:"id"`
	Address            string  `json:"address"`
	Telephone          string  `json:"telephone"`
	WarehouseCode      string  `json:"warehouse_code"`
	MinimumCapacity    int     `json:"minimum_capacity"`
	MinimumTemperature float64 `json:"minimum_temperature"`
}

func NewWarehouse(w warehouse.Service) *Warehouse {
	return &Warehouse{
		warehouseService: w,
	}
}

// GetByID godoc
// @Summary Get warehouse by ID
// @Tags Warehouses
// @Description This endpoint retrieves the information of a warehouse by its ID
// @Accept  json
// @Produce  json
// @Param id path string true "Warehouse ID"
// @Success 200 {object} domain.Warehouse
// @Failure 400 {object} web.ErrorResponse "ID de armazem invalido"
// @Failure 404 {object} web.ErrorResponse "Armazem nao encontrado"
// @Router /warehouses/{id} [get]
func (w *Warehouse) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		whouseCode, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, "ID de armazem invalido")
			return
		}
		warehouseData, err := w.warehouseService.Get(context.TODO(), whouseCode)
		if err != nil {
			web.Error(c, http.StatusNotFound, "Armazem nao encontrado")
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
// @Failure 500 {object} web.ErrorResponse "Falha ao obter os armazéns"
// @Router /warehouses [get]
func (w *Warehouse) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		warehouses, err := w.warehouseService.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, "Falha ao obter os armazéns")
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
// @Failure 400 {object} web.ErrorResponse "Invalid data"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 422 {object} web.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /warehouses [post]
func (w *Warehouse) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		var warehouseData domain.Warehouse
		if err := c.ShouldBindJSON(&warehouseData); err != nil {
			web.Error(c, http.StatusBadRequest, "Verifique os dados informados e tende novamente!")
			return
		}
		if err := WarehouseValidator(c, warehouseData); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, err.Error())
			return
		}
		wCode, err := w.warehouseService.Save(c.Request.Context(), warehouseData)
		if err != nil {
			if errors.Is(err, warehouse.ErrWarehouseExists) {
				web.Error(c, http.StatusConflict, "Ja existe um armazem com o codigo %v", warehouseData.WarehouseCode)
				return
			} else {
				web.Error(c, http.StatusInternalServerError, "Erro interno de servidor")
				return
			}
		}
		web.Response(c, http.StatusCreated, wCode)
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
// @Failure 400 {object} web.ErrorResponse "Corpo da requisição inválido"
// @Failure 404 {object} web.ErrorResponse "Not found"
// @Failure 409 {object} web.ErrorResponse "Codigo de armazem ja registrado!"
// @Failure 500 {object} web.ErrorResponse "Erro interno no servidor."
// @Router /warehouses [put]
func (w *Warehouse) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

		warehouseId := c.Param("id")
		warehouseData := domain.Warehouse{}

		if err := c.ShouldBindJSON(&warehouseData); err != nil {
			web.Error(c, http.StatusBadRequest, "Corpo da requisição inválido")
			return
		}

		if warehouseData.WarehouseCode != "" && warehouseData.WarehouseCode != warehouseId {
			web.Error(c, http.StatusBadRequest, "Não é permitido atualizar o campo WarehouseCode!")
			return
		}

		warehouseData.WarehouseCode = warehouseId

		updateWarehouse, err := w.warehouseService.Update(c.Request.Context(), warehouseData)
		if err != nil {
			if errors.Is(err, warehouse.ErrWarehouseExists) {
				web.Error(c, http.StatusConflict, "Codigo de armazem ja registrado!")
				return
			} else if errors.Is(err, warehouse.ErrWarehouseNotFound) {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			} else {
				web.Error(c, http.StatusInternalServerError, "Erro interno no servidor.")
				return
			}
		}
		web.Success(c, http.StatusOK, updateWarehouse)
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
// @Failure 400 {object} web.ErrorResponse "ID invalido"
// @Failure 404 {object} web.ErrorResponse "Armazém não encontrado!"
// @Failure 500 {object} web.ErrorResponse "Falha ao excluir o armazem"
// @Router /warehouses/{id} [delete] "Falha ao excluir o armazem"
func (w *Warehouse) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {

		warehouseID := c.Param("id")
		id, err := strconv.Atoi(warehouseID)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "ID invalido")
		}
		existingWarehouse, err := w.warehouseService.Get(context.TODO(), id)
		if err != nil {
			web.Error(c, http.StatusNotFound, "Armazém não encontrado!")
			return
		}
		err = w.warehouseService.Delete(context.TODO(), existingWarehouse.ID)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, "Falha ao excluir o armazem")
			return
		}
		web.Success(c, http.StatusNoContent, "Armazém excluído com sucesso")
	}
}

func WarehouseValidator(c *gin.Context, warehouseData domain.Warehouse) error {
	if warehouseData.Address == "" {
		web.Error(c, http.StatusBadRequest, "O endereco do armazem deve ser informado!")
	}
	if warehouseData.Telephone == "" {
		web.Error(c, http.StatusBadRequest, "O telefone do armazem deve ser informado!")
	}
	if warehouseData.MinimumTemperature == 0 {
		web.Error(c, http.StatusBadRequest, "A temperatura minima deve ser informada!")
	}
	if warehouseData.MinimumCapacity == 0 {
		web.Error(c, http.StatusBadRequest, "A capacidade minima deve ser informada!")
	}
	if warehouseData.WarehouseCode == "" {
		web.Error(c, http.StatusBadRequest, "O codigo do armazem deve ser informado!")
	}
	return nil
}
