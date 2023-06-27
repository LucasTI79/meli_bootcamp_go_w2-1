package handler

import (
	"net/http"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/seller"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/gin-gonic/gin"
)

type Seller struct {
	service seller.Service
}

type CreateSellerRequest struct {
	CID         int    `json:"cid" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Telephone   string `json:"telephone" binding:"required,e164"`
}

func (r CreateSellerRequest) ToSeller() domain.Seller {
	return domain.Seller{
		ID:          0,
		CID:         r.CID,
		CompanyName: r.CompanyName,
		Address:     r.Address,
		Telephone:   r.Telephone,
	}
}

type UpdateSellerRequest struct {
	CID         *int    `json:"cid"`
	CompanyName *string `json:"company_name"`
	Address     *string `json:"address"`
	Telephone   *string `json:"telephone" binding:"omitempty,e164"`
}

func (r UpdateSellerRequest) ToUpdateSeller() domain.UpdateSeller {
	return domain.UpdateSeller{
		CID:         r.CID,
		CompanyName: r.CompanyName,
		Address:     r.Address,
		Telephone:   r.Telephone,
	}
}

func NewSeller(service seller.Service) *Seller {
	return &Seller{service}
}

// Create godoc
// @Summary List sellers
// @Description List all sellers
// @Tags Sellers
// @Accept json
// @Produce json
// @Success 200 {object} []domain.Seller "List of sellers"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /sellers [get]
func (p *Seller) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		sellers := p.service.GetAll(c.Request.Context())
		web.Success(c, http.StatusOK, sellers)
	}
}

// Get godoc
// @Summary Get a seller by id
// @Description Get a seller based on the provided id
// @Tags Sellers
// @Accept json
// @Produce json
// @Param id path int true "Seller Id"
// @Success 200 {object} []domain.Seller "Created seller"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "Resource not found error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /sellers/{id} [get]
func (p *Seller) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetInt("Id")

		seller, err := p.service.Get(c.Request.Context(), id)

		if err != nil {
			if apperr.Is[*apperr.ResourceNotFound](err) {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			}
		}

		web.Success(c, http.StatusOK, seller)
	}
}

// Create godoc
// @Summary Create a new seller
// @Description Create a new seller based on the provided JSON payload
// @Tags Sellers
// @Accept json
// @Produce json
// @Param request body CreateSellerRequest true "Seller data"
// @Success 201 {object} domain.Seller "Created seller"
// @Failure 422 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "Not found error"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /sellers [post]
func (p *Seller) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := c.MustGet(RequestParamContext).(CreateSellerRequest)

		created, err := p.service.Create(c.Request.Context(), request.ToSeller())

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
// @Summary Update a seller
// @Description Update an existent seller based on the provided JSON payload
// @Tags Sellers
// @Accept json
// @Produce json
// @Param id path int true "Seller id"
// @Param request body UpdateSellerRequest true "Seller data"
// @Success 200 {object} domain.Seller "Updated seller"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 422 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "Resource not found error"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /sellers/{id} [patch]
func (p *Seller) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetInt("Id")
		request := c.MustGet(RequestParamContext).(UpdateSellerRequest)

		response, err := p.service.Update(c.Request.Context(), id, request.ToUpdateSeller())

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

		web.Success(c, http.StatusOK, response)
	}
}

// Delete godoc
// @Summary Delete a seller
// @Description Delete a seller based on the provided JSON payload
// @Tags Sellers
// @Accept json
// @Produce json
// @Param id path int true "Seller id"
// @Success 204
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "Resource not found error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /sellers/{id} [delete]
func (p *Seller) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetInt("Id")

		err := p.service.Delete(c.Request.Context(), id)

		if err != nil {
			if apperr.Is[*apperr.ResourceNotFound](err) {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			}
		}

		web.Success(c, http.StatusNoContent, nil)
	}
}
