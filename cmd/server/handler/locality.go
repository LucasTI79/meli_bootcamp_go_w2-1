package handler

import (
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/locality"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/helpers"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/gin-gonic/gin"
)

const (
	InvalidId = "o id '%s' é inválido"
)

type Locality struct {
	service locality.Service
}

type CreateLocalityRequest struct {
	LocalityName *string `json:"locality_name" binding:"required"`
	ProvinceID   *int    `json:"province_id" binding:"required"`
}

func (s CreateLocalityRequest) ToLocality() domain.Locality {
	return domain.Locality{
		ID:           0,
		LocalityName: helpers.ToFormattedAddress(*s.LocalityName),
		ProvinceID:   *s.ProvinceID,
	}
}

func NewLocality(service locality.Service) *Locality {
	return &Locality{service}
}

// Create godoc
// @Summary Create a locality
// @Description Create a new locality based on the provided JSON payload.
// @Tags Localities
// @Accept json
// @Produce json
// @Param request body CreateLocalityRequest true "Locality to be created"
// @Success 201 {object} domain.Locality "Created locality"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 422 {object} web.ErrorResponse "Validation error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /localities [post]
func (l *Locality) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := c.MustGet(RequestParamContext).(CreateLocalityRequest)

		created, err := l.service.Create(request.ToLocality())

		if err != nil {
			if apperr.Is[*apperr.ResourceAlreadyExists](err) {
				web.Error(c, http.StatusConflict, err.Error())
				return
			}
			if apperr.Is[*apperr.DependentResourceNotFound](err) {
				web.Error(c, http.StatusConflict, err.Error())
				return
			}
		}

		web.Success(c, http.StatusCreated, created)
	}
}

// Create godoc
// @Summary Count sellers by locality
// @Description Seller count by locality.
// @Description If no query param is given, it brings the report to all localities.
// @Description If a location id is specified, it brings the number of sellers for this locality.
// @Tags Localities
// @Produce json
// @Param id query int false "Locality ID"
// @Success 200 {object} []domain.SellersByLocalityReport "Report of sellers by locality"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "Resource not found error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /localities/report-sellers [get]
func (l *Locality) ReportSellers() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Request.URL.Query().Get("id")

		if idParam == "" {
			result := l.service.CountSellersByAllLocalities()
			web.Success(c, http.StatusOK, result)
			return
		}

		id, err := strconv.Atoi(idParam)

		if err != nil {
			web.Error(c, http.StatusBadRequest, InvalidId, idParam)
			return
		}

		localities, err := l.service.CountSellersByLocality(id)

		if err != nil {
			if apperr.Is[*apperr.ResourceNotFound](err) {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			}
		}

		web.Success(c, http.StatusOK, localities)
	}
}

// @Summary Count carriers by locality
// @Description Carrier count by location.
// @Description If no query param is given, it brings the report to all localities.
// @Description If a location id is specified, it brings the number of carriers for this locality.
// @Tags Localities
// @Produce json
// @Param id query int false "Locality ID"
// @Success 200 {object} []domain.CarriersByLocalityReport "List of localities"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "Resource not found error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /localities/report-carriers [get]
func (l Locality) ReportCarriers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Request.URL.Query().Get("id")

		if idParam == "" {
			result := l.service.CountCarriersByAllLocalities()
			web.Success(ctx, http.StatusOK, result)
			return
		}

		id, err := strconv.Atoi(idParam)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, InvalidId, idParam)
			return
		}

		reportCarriers, err := l.service.CountCarriersByLocality(id)
		if err != nil {
			if apperr.Is[*apperr.ResourceNotFound](err) {
				web.Error(ctx, http.StatusNotFound, err.Error())
				return
			}
		}

		web.Success(ctx, http.StatusOK, reportCarriers)

	}
}
