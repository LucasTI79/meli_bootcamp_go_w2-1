package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"strconv"
)

type Buyer struct {
	buyerService buyer.IService
}

func NewBuyer(b buyer.IService) *Buyer {
	return &Buyer{
		buyerService: b,
	}
}

// Get godoc
// @Summary List buyer based on ID
// @Description get buyer by ID
// @Tags Buyers
// @Accept json
// @Produce json
// @Param id path int true "ID do comprador"
// @Success 200 {object} domain.Buyer "List a specific Buyer according to ID"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "Buyer not found"
// @Router /api/v1/buyers/{id} [get]
func (b *Buyer) Get() gin.HandlerFunc {
	return func(c *gin.Context) {

		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "ID inválido")
			return
		}
		
		buyer, err := b.buyerService.Get(c.Request.Context(), id)
		if err != nil {
			web.Error(c, http.StatusNotFound, "Comprador não encontrado")
			return
		}	

		web.Success(c, http.StatusOK, buyer)

	}
}

// GetAll godoc
// @Summary List all buyers
// @Description getAll buyers
// @Tags Buyers
// @Accept json
// @Produce json
// @Success 200 {object} domain.Buyer "List of all Buyers"
// @Failure 204 {object} web.ErrorResponse "Buyer not found"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Router /api/v1/buyers [get]
func (b *Buyer) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		buyers, err := b.buyerService.GetAll(c.Request.Context())
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		if len(buyers) == 0 {
			web.Success(c, http.StatusNoContent, "Não há compradores cadastrados")
			return
		}

		web.Success(c, http.StatusOK, buyers)

	}
}

// Create godoc
// @Summary Create new buyer
// @Description Create a new buyer based on the provided JSON payload
// @Tags Buyers
// @Accept json
// @Produce json
// @Param buyer body domain.Request true "Buyer to be created"
// @Success 201 {object} domain.Buyer "Created buyer"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 422 {object} web.ErrorResponse "Json Parse error"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Router /api/v1/buyers [post]
func (b *Buyer) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		var req domain.Request
		err := c.Bind(&req)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Existem erros na formatação do Json. Não foi possível realizar o parse.")
			return	
		}

		if req.CardNumberID == "" {
			web.Error(c, http.StatusUnprocessableEntity, "O campo 'Card Number' do comprador é obrigatório")
			return
		}
		if req.FirstName == "" {
			web.Error(c, http.StatusUnprocessableEntity, "O campo 'Nome' do comprador é obrigatório")
			return
		}
		if req.LastName== "" {
			web.Error(c, http.StatusUnprocessableEntity, "O campo 'Sobrenome' do comprador é obrigatório")
			return
		}

		buyerId, err := b.buyerService.Save(c.Request.Context(), req)
		if err != nil {
			web.Error(c, http.StatusConflict, err.Error())
			return
		}

		buyer := domain.Buyer{}
		buyer.ID = buyerId
		buyer.CardNumberID = req.CardNumberID
		buyer.FirstName = req.FirstName
		buyer.LastName = req.LastName

		web.Success(c, http.StatusCreated, buyer)

	}
}

// Update godoc
// @Summary Update a buyer based on ID
// @Description Update a specific buyer based on the provided JSON payload
// @Tags Buyers
// @Accept json
// @Produce json
// @Param buyer body domain.Request true "Buyer to be updated"
// @Param id path int true "ID do comprador"
// @Success 200 {object} domain.Buyer "Buyer with updated information"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "Buyer not found"
// @Failure 422 {object} web.ErrorResponse "Json Parse error"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /api/v1/buyers/{id} [patch]
func (b *Buyer) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "ID inválido")
			return
		}

		var req domain.Request
		err = c.Bind(&req)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Existem erros na formatação do Json. Não foi possível realizar o parse.")
			return	
		}
		
		if req.CardNumberID == "" && req.FirstName == "" && req.LastName== "" {
			web.Error(c, http.StatusUnprocessableEntity, "Informe pelo menos um campo para atualização")
			return
		}

		buyer, err := b.buyerService.Get(c.Request.Context(), id)
		if err != nil {
			web.Error(c, http.StatusNotFound, err.Error())
			return
		}	
		
		if req.CardNumberID != "" {
			exists := b.buyerService.Exists(c.Request.Context(), req.CardNumberID)

			if !exists {
				buyer.CardNumberID = req.CardNumberID
			} else {
				web.Error(c, http.StatusConflict, "Não é possível atualizar um comprador com Card Number repetido.") 
				return
			}
		}

		if req.FirstName != "" {
			buyer.FirstName = req.FirstName
		}

		if req.LastName != "" {
			buyer.LastName = req.LastName
		}

		err = b.buyerService.Update(c.Request.Context(), buyer)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}	

		web.Success(c, http.StatusOK, buyer)

	}
}

// Delete godoc
// @Summary Delete a buyer based on ID
// @Description Delete a specific buyer based on ID
// @Tags Buyers
// @Param id path int true "ID do comprador"
// @Success 204 "No content"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "Buyer not found"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /api/v1/buyers/{id} [delete]
func (b *Buyer) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "ID inválido")
			return
		}
		
		err = b.buyerService.Delete(c.Request.Context(), id)
		if err != nil {
			web.Error(c, http.StatusNotFound, "O comprador com o ID correspondente não existe")
			return
		}	

		web.Success(c, http.StatusNoContent, "")
	}
}
