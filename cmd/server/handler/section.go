package handler

import (
	"net/http"
	"strconv"

	_ "github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/section"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/gin-gonic/gin"
)

type SectionRequest struct {
	Id                 int `json:"id"`
	SectionNumber      int `json:"section_number"`
	CurrentTemperature int `json:"current_temperature"`
	MinimumTemperature int `json:"minimum_temperature"`
	CurrentCapacity    int `json:"current_capacity"`
	MinimumCapacity    int `json:"minimum_capacity"`
	MaximumCapacity    int `json:"maximum_capacity"`
	WarehouseId        int `json:"warehouse_id"`
	ProductTypeId      int `json:"id_product_type"`
}

type Section struct {
	service section.Service
}

func NewSection(s section.Service) *Section {
	return &Section{
		service: s,
	}
}

// Get All Sections godoc
// @Summary Get all sections
// @Description Get sections based on the provided JSON payload
// @Tags Sections
// @Accept json
// @Produce json
// @Success 200 {object} []domain.Section "Section"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /sections [get]
func (s *Section) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sections, err := s.service.GetAll()
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "erro interno.")
			return
		}
		web.Success(ctx, http.StatusOK, sections)
	}
}

// Get godoc
// @Summary Get a section
// @Description Get a section based on the provided JSON payload
// @Tags Sections
// @Accept json
// @Produce json
// @Success 200 {object} domain.Section "Section"
// @Failure 400 {object} web.ErrorResponse"Validation error"
// @Failure 404 {object} web.ErrorResponse "NotFound error"
// @Router /sections/{id} [get]
func (s *Section) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "id inválido.")
			return
		}

		sectionID, err := s.service.Get(id)
		if err != nil {
			web.Error(ctx, http.StatusNotFound, "seção não encontrada.")
			return
		}
		web.Success(ctx, http.StatusOK, sectionID)
	}
}

// Create godoc
// @Summary Create a new section
// @Description Create a new section based on the provided JSON payload
// @Tags Sections
// @Accept json
// @Produce json
// @Param request body domain.Section true "Section data"
// @Success 201 {object} domain.Section "Created section"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 422 {object} web.ErrorResponse "Unprocessable error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /sections [post]
func (s *Section) Save() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req SectionRequest
		if err := ctx.Bind(&req); err != nil {
			web.Error(ctx, http.StatusNotFound, "existem erros na formatação do json e não foi possível realizar o parse.")
			return
		}
		if req.SectionNumber == 0 && req.CurrentTemperature == 0 && req.MinimumTemperature == 0 && req.CurrentCapacity == 0 && req.MinimumCapacity == 0 &&
			req.MaximumCapacity == 0 && req.WarehouseId == 0 && req.ProductTypeId == 0 {
			web.Error(ctx, http.StatusUnprocessableEntity, "necessário adicionar todas as informações.")
			return
		}

		sectionId, err := s.service.Save(req.SectionNumber, req.CurrentTemperature, req.MinimumTemperature, req.CurrentCapacity, req.MinimumCapacity, req.MaximumCapacity,
			req.WarehouseId, req.ProductTypeId)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "erro interno de servidor.")
			return
		}
		sectionCreated, err := s.service.Get(sectionId)
		web.Success(ctx, http.StatusCreated, sectionCreated)

	}
}

// Update godoc
// @Summary Update a section
// @Description Update section based on the provided JSON payload
// @Tags Sections
// @Accept json
// @Produce json
// @Param request body domain.Section true "Section data"
// @Success 200 {object} domain.Section "Updated section"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "NotFound error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /sections/{id} [patch]
func (s *Section) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "id inválido.")
			return
		}

		var req SectionRequest
		err = ctx.Bind(&req)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "existem erros na formatação do json e não foi possível realizar o parse.")
			return
		}

		if req.SectionNumber == 0 && req.CurrentTemperature == 0 && req.MinimumTemperature == 0 && req.CurrentCapacity == 0 && req.MinimumCapacity == 0 && req.MaximumCapacity == 0 &&
			req.WarehouseId == 0 && req.ProductTypeId == 0 {

			web.Error(ctx, http.StatusUnprocessableEntity, "informe pelo menos um campo para concluir a atualização.")
			return
		}

		section, err := s.service.Get(id)
		if err != nil {
			web.Error(ctx, http.StatusNotFound, "seção não encontrada.")
			return
		}
		if req.SectionNumber != 0 {
			sectionNumber, err := s.service.Exists(req.SectionNumber)
			if err != nil {
				web.Error(ctx, http.StatusBadRequest, "número de seção cadastrado.")
				return
			} else {
				section.SectionNumber = sectionNumber
			}
		}
		if req.CurrentTemperature != 0 {
			section.CurrentTemperature = req.CurrentTemperature
		}
		if req.MinimumTemperature != 0 {
			section.MinimumTemperature = req.MinimumTemperature
		}
		if req.CurrentCapacity != 0 {
			section.CurrentCapacity = req.CurrentCapacity
		}
		if req.MinimumCapacity != 0 {
			section.MinimumCapacity = req.MinimumCapacity
		}
		if req.MaximumCapacity != 0 {
			section.MaximumCapacity = req.MaximumCapacity
		}
		if req.WarehouseId != 0 {
			section.WarehouseID = req.WarehouseId
		}
		if req.ProductTypeId != 0 {
			section.ProductTypeID = req.ProductTypeId
		}

		err = s.service.Update(section)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		web.Success(ctx, http.StatusOK, section)
	}
}

// Exists godoc
// @Summary Exist section number
// @Description Validate section number
// @Tags Sections
// @Accept json
// @Produce json
// @Success 200 {object} string "Section number"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /sections/sectionNumber [get]
func (s *Section) Exists() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req SectionRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			web.Error(ctx, http.StatusBadRequest, "existem erros na formatação do json e não foi possível realizar o parse.")
			return
		}
		if req.SectionNumber == 0 {
			web.Error(ctx, http.StatusUnprocessableEntity, "necessário adicionar número de seção.")
			return
		}
		sectionNumber, err := s.service.Exists(req.SectionNumber)
		if err != nil {
			web.Error(ctx, http.StatusConflict, "seção já cadastrada.")
			return
		}
		web.Success(ctx, http.StatusOK, sectionNumber)
	}
}

// Delete godoc
// @Summary Delete section
// @Description Delete section based on the provided JSON payload
// @Tags Sections
// @Accept json
// @Produce json
// @Success 204
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "NotFound error"
// @Router /sections/{id} [delete]
func (s *Section) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "id inválido.")
			return
		}

		err = s.service.Delete(id)
		if err != nil {
			web.Error(ctx, http.StatusNotFound, "seção não encontrada.")
			return
		}

		web.Success(ctx, http.StatusNoContent, nil)
	}
}
