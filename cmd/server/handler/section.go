package handler

import (
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/section"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/gin-gonic/gin"
)

type Section struct {
	service section.Service
}

type CreateSectionRequest struct {
	SectionNumber      *int     `json:"section_number" binding:"required"`
	CurrentTemperature *float32 `json:"current_temperature" binding:"required"`
	MinimumTemperature *float32 `json:"minimum_temperature" binding:"required"`
	CurrentCapacity    *int     `json:"current_capacity" binding:"required"`
	MinimumCapacity    *int     `json:"minimum_capacity" binding:"required"`
	MaximumCapacity    *int     `json:"maximum_capacity" binding:"required"`
	WarehouseID        *int     `json:"warehouse_id" binding:"required"`
	ProductTypeID      *int     `json:"product_type_id" binding:"required"`
}

func (r CreateSectionRequest) ToSection() domain.Section {
	return domain.Section{
		ID:                 0,
		SectionNumber:      *r.SectionNumber,
		CurrentTemperature: *r.CurrentTemperature,
		MinimumTemperature: *r.MinimumTemperature,
		CurrentCapacity:    *r.CurrentCapacity,
		MinimumCapacity:    *r.MinimumCapacity,
		MaximumCapacity:    *r.MaximumCapacity,
		WarehouseID:        *r.WarehouseID,
		ProductTypeID:      *r.ProductTypeID,
	}
}

type UpdateSectionRequest struct {
	SectionNumber      *int     `json:"section_number"`
	CurrentTemperature *float32 `json:"current_temperature"`
	MinimumTemperature *float32 `json:"minimum_temperature"`
	CurrentCapacity    *int     `json:"current_capacity"`
	MinimumCapacity    *int     `json:"minimum_capacity"`
	MaximumCapacity    *int     `json:"maximum_capacity"`
	WarehouseID        *int     `json:"warehouse_id"`
	ProductTypeID      *int     `json:"product_type_id"`
}

func (r UpdateSectionRequest) ToUpdateSection() domain.UpdateSection {
	return domain.UpdateSection{
		SectionNumber:      r.SectionNumber,
		CurrentTemperature: r.CurrentTemperature,
		MinimumTemperature: r.MinimumTemperature,
		CurrentCapacity:    r.CurrentCapacity,
		MinimumCapacity:    r.MinimumCapacity,
		MaximumCapacity:    r.MaximumCapacity,
		WarehouseID:        r.WarehouseID,
		ProductTypeID:      r.ProductTypeID,
	}
}

func NewSection(s section.Service) *Section {
	return &Section{
		service: s,
	}
}

// Get All Sections godoc
// @Summary List all sections
// @Description Returns a collection of existing sections.
// @Tags Sections
// @Produce json
// @Success 200 {object} []domain.Section "Section"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /sections [get]
func (s *Section) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sections := s.service.GetAll()
		web.Success(ctx, http.StatusOK, sections)
	}
}

// Get godoc
// @Summary Get a section by id
// @Description Get a section based on the provided JSON payload
// @Tags Sections
// @Produce json
// @Param id path int true "Section ID"
// @Success 200 {object} domain.Section "Section"
// @Failure 400 {object} web.ErrorResponse"Validation error"
// @Failure 404 {object} web.ErrorResponse "NotFound error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /sections/{id} [get]
func (s *Section) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.GetInt("Id")

		section, err := s.service.Get(id)

		if err != nil {
			if apperr.Is[*apperr.ResourceNotFound](err) {
				web.Error(ctx, http.StatusNotFound, err.Error())
				return
			}
		}
		web.Success(ctx, http.StatusOK, section)
	}
}

// Create godoc
// @Summary Create a new section
// @Description Create a new section based on the provided JSON payload
// @Tags Sections
// @Accept json
// @Produce json
// @Param request body CreateSectionRequest true "Section data"
// @Success 201 {object} domain.Section "Created section"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 422 {object} web.ErrorResponse "Validation error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /sections [post]
func (s *Section) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := ctx.MustGet(RequestParamContext).(CreateSectionRequest)

		created, err := s.service.Create(request.ToSection())

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

		web.Success(ctx, http.StatusCreated, created)
	}
}

// Update godoc
// @Summary Update a section
// @Description Update section based on the provided JSON payload
// @Tags Sections
// @Accept json
// @Produce json
// @Param id path int true "Section ID"
// @Param request body UpdateSectionRequest true "Section data"
// @Success 200 {object} domain.Section "Updated section"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "Resource not found error"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 422 {object} web.ErrorResponse "Validation error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /sections/{id} [patch]
func (s *Section) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.GetInt("Id")
		request := ctx.MustGet(RequestParamContext).(UpdateSectionRequest)

		response, err := s.service.Update(id, request.ToUpdateSection())

		if err != nil {
			if apperr.Is[*apperr.ResourceNotFound](err) {
				web.Error(ctx, http.StatusNotFound, err.Error())
				return
			}
			if apperr.Is[*apperr.ResourceAlreadyExists](err) {
				web.Error(ctx, http.StatusConflict, err.Error())
				return
			}
			if apperr.Is[*apperr.DependentResourceNotFound](err) {
				web.Error(ctx, http.StatusConflict, err.Error())
				return
			}
		}

		web.Success(ctx, http.StatusOK, response)
	}
}

// Delete godoc
// @Summary Delete section
// @Description Delete section based on the provided JSON payload
// @Tags Sections
// @Success 204 "No content"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "Resource not found error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /sections/{id} [delete]
func (s *Section) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.GetInt("Id")

		err := s.service.Delete(id)

		if err != nil {
			if apperr.Is[*apperr.ResourceNotFound](err) {
				web.Error(ctx, http.StatusNotFound, err.Error())
				return
			}
		}

		web.Success(ctx, http.StatusNoContent, nil)
	}
}

// ReportProducts godoc
// @Summary Count products by section
// @Description Return the report of products by section
// @Description If no query param is given, it brings the report of all products by section
// @Description If a section id is specified, it brings the number of products for this section.
// @Tags Sections
// @Produce json
// @Param id query int false "Section ID"
// @Success 200 {object} []domain.ProductsBySectionReport "Report of products by section"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "Resource not found error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /sections/report-products [get]
func (s *Section) ReportProducts() gin.HandlerFunc {
	return func(c *gin.Context) {

		idParam := c.Request.URL.Query().Get("id")
		if idParam == "" {
			result := s.service.CountProductsByAllSections()
			web.Success(c, http.StatusOK, result)
			return
		}
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Error(c, http.StatusBadRequest, InvalidId, idParam)
			return
		}
		result, err := s.service.CountProductsBySection(id)
		if err != nil {
			if apperr.Is[*apperr.ResourceNotFound](err) {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			}
		}

		web.Success(c, http.StatusOK, result)
	}
}
