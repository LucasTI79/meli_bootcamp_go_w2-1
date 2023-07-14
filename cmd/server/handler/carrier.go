package handler

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/carrier"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/helpers"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateCarrierRequest struct {
	CID         *string `json:"cid" binding:"required"`
	CompanyName *string `json:"company_name" binding:"required"`
	Address     *string `json:"address" binding:"required"`
	Telephone   *string `json:"telephone" binding:"required,e164"`
	LocalityID  *int    `json:"locality_id" binding:"required"`
}

func (c CreateCarrierRequest) ToCarrier() domain.Carrier {
	return domain.Carrier{
		CID:         *c.CID,
		CompanyName: *c.CompanyName,
		Address:     helpers.ToFormattedAddress(*c.Address),
		Telephone:   *c.Telephone,
		LocalityID:  *c.LocalityID,
	}
}

type Carrier struct {
	carrierService carrier.Service
}

func NewCarrier(c carrier.Service) *Carrier {
	return &Carrier{
		carrierService: c,
	}
}

// Create godoc
// @Summary Create a carrier
// @Description Create a new carrier based on the provided JSON payload.
// @Tags Carriers
// @Accept json
// @Produce json
// @Param request body CreateCarrierRequest true "Carrier to be created"
// @Success 201 {object} domain.Carrier "Created carrier"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 422 {object} web.ErrorResponse "Validation error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /carriers [post]
func (c *Carrier) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := ctx.MustGet(RequestParamContext).(CreateCarrierRequest)

		ca, err := c.carrierService.Create(request.ToCarrier())

		if err != nil {
			if apperr.Is[*apperr.ResourceAlreadyExists](err) {
				web.Error(ctx, http.StatusConflict, err.Error())
				return
			}
			if apperr.Is[*apperr.DependentResourceNotFound](err) {
				web.Error(ctx, http.StatusConflict, err.Error())
				return
			}
		}

		web.Success(ctx, http.StatusCreated, ca)
	}
}
