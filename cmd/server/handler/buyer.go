package handler

import (
	"net/http"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/gin-gonic/gin"
)

type CreateBuyerRequest struct {
	CardNumberID *string `json:"card_number_id" binding:"required"`
	FirstName    *string `json:"first_name" binding:"required"`
	LastName     *string `json:"last_name" binding:"required"`
}

func (r CreateBuyerRequest) ToBuyer() domain.Buyer {
	return domain.Buyer{
		CardNumberID: *r.CardNumberID,
		FirstName:    *r.FirstName,
		LastName:     *r.LastName,
	}
}

type UpdateBuyerRequest struct {
	CardNumberID *string `json:"card_number_id"`
	FirstName    *string `json:"first_name"`
	LastName     *string `json:"last_name"`
}

func (r UpdateBuyerRequest) ToUpdateBuyer() domain.UpdateBuyer {
	return domain.UpdateBuyer{
		CardNumberID: r.CardNumberID,
		FirstName:    r.FirstName,
		LastName:     r.LastName,
	}
}

type Buyer struct {
	buyerService buyer.IService
}

func NewBuyer(b buyer.IService) *Buyer {
	return &Buyer{
		buyerService: b,
	}
}

// Get godoc
// @Summary Get a buyer by id
// @Description Get a buyer based on the provided id. Returns a not found error if the buyer does not exist.
// @Tags Buyers
// @Accept json
// @Produce json
// @Param id path int true "Buyer id"
// @Success 200 {object} domain.Buyer "Obtained buyer"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "Resource not found error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /buyers/{id} [get]
func (b *Buyer) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetInt("Id")

		buyer, err := b.buyerService.Get(c.Request.Context(), id)
		if err != nil {
			if apperr.Is[*apperr.ResourceNotFound](err) {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			}
		}

		web.Success(c, http.StatusOK, buyer)
	}
}

// GetAll godoc
// @Summary List all buyers
// @Description Returns a collection of existing buyers.
// @Tags Buyers
// @Accept json
// @Produce json
// @Success 200 {object} domain.Buyer "List of all buyers"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /buyers [get]
func (b *Buyer) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		buyers := b.buyerService.GetAll(c.Request.Context())
		web.Success(c, http.StatusOK, buyers)
	}
}

// Create godoc
// @Summary Create a buyer
// @Description Create a new buyer based on the provided JSON payload.
// @Tags Buyers
// @Accept json
// @Produce json
// @Param buyer body domain.Buyer true "Buyer to be created"
// @Success 201 {object} domain.Buyer "Created buyer"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 422 {object} web.ErrorResponse "Validation error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /buyers [post]
func (b *Buyer) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := c.MustGet(RequestParamContext).(CreateBuyerRequest)

		buyer, err := b.buyerService.Create(c.Request.Context(), request.ToBuyer())

		if err != nil {
			if apperr.Is[*apperr.ResourceAlreadyExists](err) {
				web.Error(c, http.StatusConflict, err.Error())
				return
			}
		}

		web.Success(c, http.StatusCreated, buyer)
	}
}

// Update godoc
// @Summary Update a buyer
// @Description Update an existent buyer based on the provided id and JSON payload.
// @Tags Buyers
// @Accept json
// @Produce json
// @Param buyer body domain.UpdateBuyer true "Buyer data to be updated"
// @Param id path int true "Buyer id"
// @Success 200 {object} domain.Buyer "Buyer with updated information"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "Resource not found error"
// @Failure 422 {object} web.ErrorResponse "Validation error"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /buyers/{id} [patch]
func (b *Buyer) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetInt("Id")
		request := c.MustGet(RequestParamContext).(UpdateBuyerRequest)

		updated, err := b.buyerService.Update(c.Request.Context(), id, request.ToUpdateBuyer())

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
// @Summary Delete a buyer
// @Description Delete a buyer based on the provided id.
// @Tags Buyers
// @Param id path int true "Buyer id"
// @Success 204 "No content"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "Resource not found error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /buyers/{id} [delete]
func (b *Buyer) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetInt("Id")

		err := b.buyerService.Delete(c.Request.Context(), id)

		if err != nil {
			if apperr.Is[*apperr.ResourceNotFound](err) {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			}
		}

		web.Success(c, http.StatusNoContent, nil)
	}
}
