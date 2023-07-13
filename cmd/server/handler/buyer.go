package handler

import (
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/gin-gonic/gin"
)

const (
	invalidId = "o id '%s' é inválido"
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
	buyerService buyer.Service
}

func NewBuyer(b buyer.Service) *Buyer {
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

		buyer, err := b.buyerService.Get(id)
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
		buyers := b.buyerService.GetAll()
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

		buyer, err := b.buyerService.Create(request.ToBuyer())

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
// @Param id path int true "Buyer id"
// @Param buyer body domain.UpdateBuyer true "Buyer data to be updated"
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

		updated, err := b.buyerService.Update(id, request.ToUpdateBuyer())

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

		err := b.buyerService.Delete(id)

		if err != nil {
			if apperr.Is[*apperr.ResourceNotFound](err) {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			}
		}

		web.Success(c, http.StatusNoContent, nil)
	}
}

// Create godoc
// @Summary Count purchase orders by buyer
// @Description Purchase Orders count by buyer.
// @Description If no query param is given, bring the report to all purchase orders for all buyers.
// @Description If a buyer id is specified, bring the amount of purchase orders for this buyer.
// @Tags Buyers
// @Accept json
// @Produce json
// @Param id query int false "Buyer ID"
// @Success 200 {object} []domain.PurchasesByBuyerReport "List of purchase Orders"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "Resource not found error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /buyers/report-purchase-orders [get]
func (b *Buyer) ReportPurchases() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Request.URL.Query().Get("id")

		if idParam == "" {
			result := b.buyerService.CountPurchasesByAllBuyers()
			web.Success(c, http.StatusOK, result)
			return
		}

		id, err := strconv.Atoi(idParam)

		if err != nil {
			web.Error(c, http.StatusBadRequest, InvalidId, idParam)
			return
		}

		purchases, err := b.buyerService.CountPurchasesByBuyer(id)

		if err != nil {
			if apperr.Is[*apperr.ResourceNotFound](err) {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			}
		}

		web.Success(c, http.StatusOK, purchases)
	}
}
