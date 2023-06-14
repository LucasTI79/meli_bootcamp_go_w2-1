package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/seller"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/gin-gonic/gin"
)

type Seller struct {
	sellerService seller.Service
}
type CreateResponse struct {
	Data int `json:"data"`
}

func NewSeller(s seller.Service) *Seller {
	return &Seller{
		sellerService: s,
	}

}

// GetAll godoc
// @Summary Get all sellers
// @Description Get all sellers from the system
// @Tags Sellers
// @Accept  json
// @Produce  json
// @Success 200 {array} domain.Seller
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /sellers [get]
func (s *Seller) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		sellers, err := s.sellerService.GetAll(c.Request.Context())
		if err != nil {
			web.Error(c, http.StatusInternalServerError, "erro interno de servidor")
			return
		}
		if len(sellers) == 0 {
			web.Response(c, http.StatusNoContent, sellers)
			return
		}
		web.Success(c, http.StatusOK, sellers)
	}
}

// Get godoc
// @Summary Get a seller
// @Description Get a seller by its ID
// @Tags Sellers
// @Accept  json
// @Produce  json
// @Param id path int true "Seller ID"
// @Success 200 {object} domain.Seller
// @Failure 400 {object} web.ErrorResponse "Invalid ID supplied"
// @Failure 404 {object} web.ErrorResponse "Seller not found"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /sellers/{id} [get]
func (s *Seller) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")
		parsedId, err := strconv.Atoi(id)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "id recebido é invalido")
			return
		}
		foundSeller, err := s.sellerService.Get(c.Request.Context(), parsedId)
		if err != nil {
			if errors.Is(err, seller.ErrNotFound) {
				web.Error(c, http.StatusNotFound, "não foi possível encontrar o id %v", parsedId)
				return
			} else {
				web.Error(c, http.StatusInternalServerError, "erro interno de servidor")
				return
			}
		}
		web.Success(c, http.StatusOK, foundSeller)
	}
}

// Create godoc
// @Summary Create a new seller
// @Description Create a new seller based on the provided JSON payload
// @Tags Sellers
// @Accept  json
// @Produce  json
// @Param request body domain.CreateSeller true "Seller data"
// @Success 201 {object} CreateResponse "Created seller"
// @Failure 400 {object} web.ErrorResponse "Invalid data"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /sellers [post]
func (s *Seller) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.CreateSeller
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, http.StatusBadRequest, "corpo da requisicao invalido")
			return
		}
		if req.Address == "" {
			web.Error(c, http.StatusUnprocessableEntity, "endereco é necessário")
			return
		}
		if req.CompanyName == "" {
			web.Error(c, http.StatusUnprocessableEntity, "nome da empresa é necessário")
			return
		}
		if req.Telephone == "" {
			web.Error(c, http.StatusUnprocessableEntity, "telephone é necessário")
			return
		}
		if req.CID == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "cid é necessário")
			return
		}
		id, err := s.sellerService.Save(c.Request.Context(), req)
		if err != nil {
			if errors.Is(err, seller.ErrAlreadyExists) {
				web.Error(c, http.StatusConflict, "já existe um vendedor com cid %v", req.CID)
				return
			} else {
				web.Error(c, http.StatusInternalServerError, "erro interno de servidor")
				return
			}
		}
		web.Response(c, http.StatusCreated, CreateResponse{id})
	}
}

// Update godoc
// @Summary Update a seller
// @Description Update a seller based on the provided JSON payload
// @Tags Sellers
// @Accept  json
// @Produce  json
// @Param id path int true "Seller ID"
// @Param request body domain.UpdateSeller true "Seller data"
// @Success 200 {object} domain.Seller
// @Failure 400 {object} web.ErrorResponse "Invalid ID supplied"
// @Failure 404 {object} web.ErrorResponse "Seller not found"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /sellers/{id} [patch]
func (s *Seller) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")
		parsedId, err := strconv.Atoi(id)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "id recebido é invalido")
			return
		}
		var req domain.UpdateSeller
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, http.StatusBadRequest, "corpo da requisicão invalido")
			return
		}
		req.ID = parsedId
		updatedSeller, err := s.sellerService.Update(c.Request.Context(), req)
		if err != nil {
			if errors.Is(err, seller.ErrAlreadyExists) {
				web.Error(c, http.StatusConflict, "cid já registrado")
				return
			} else if errors.Is(err, seller.ErrNotFound) {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			} else {
				web.Error(c, http.StatusInternalServerError, "erro interno no servidor")
				return
			}
		}
		web.Success(c, http.StatusOK, updatedSeller)
	}
}

// Delete godoc
// @Summary Delete a seller
// @Description Delete a seller by its ID
// @Tags Sellers
// @Accept  json
// @Produce  json
// @Param id path int true "Seller ID"
// @Success 204 "No Content"
// @Failure 400 {object} web.ErrorResponse "Invalid ID supplied"
// @Failure 404 {object} web.ErrorResponse "Seller not found"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /sellers/{id} [delete]
func (s *Seller) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")
		parsedId, err := strconv.Atoi(id)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "id recebido é invalido")
			return
		}
		err = s.sellerService.Delete(c.Request.Context(), parsedId)
		if err != nil {
			if errors.Is(err, seller.ErrNotFound) {
				web.Error(c, http.StatusNotFound, "não foi possível encontrar um vendedor com id %v", parsedId)
				return
			} else {
				web.Error(c, http.StatusInternalServerError, "erro interno no servidor")
				return
			}
		}
		web.Success(c, http.StatusNoContent, "")
	}
}
