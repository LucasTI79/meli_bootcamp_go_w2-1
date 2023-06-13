package handler

import (
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/section"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Id                  int `json:"id"`
	Section_number      int `json:"section_number"`
	Current_temperature int `json:"current_temperatur"`
	Minimum_temperature int `json:"minimum_temperature"`
	Current_capacity    int `json:"current_capacity"`
	Minimum_capacity    int `json:"minimum_capacity"`
	Maximum_capacity    int `json:"maximum_capacity"`
	Warehouse_id        int `json:"warehouse_id"`
	Id_product_type     int `json:"id_product_type"`
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
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /sections [get]
func (s *Section) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sections, err := s.service.GetAll()
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "Error: Seções não encontradas.")
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
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /sections/:id [get]
func (s *Section) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "Error: ID inválido.")
			return
		}

		sectionID, err := s.service.Get(int(id))
		if err != nil {
			web.Error(ctx, http.StatusNotFound, "Error: Seção não encontrada.")
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
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /sections [post]
func (s *Section) Save() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		if err := ctx.Bind(&req); err != nil {
			web.Error(ctx, http.StatusNotFound, "Error")
			return
		}
		if req.Section_number == 0 && req.Current_temperature == 0 && req.Minimum_temperature == 0 && req.Current_capacity == 0 && req.Minimum_capacity == 0 &&
			req.Maximum_capacity == 0 && req.Warehouse_id == 0 && req.Id_product_type == 0 {
			web.Error(ctx, http.StatusUnprocessableEntity, "Error: Necessário adicionar todas as informações.")
			return
		}
		sectionSaved, err := s.service.Save(req.Section_number, req.Current_temperature, req.Minimum_temperature, req.Current_capacity, req.Minimum_capacity, req.Maximum_capacity,
			req.Warehouse_id, req.Id_product_type)
		if err != nil {
			web.Error(ctx, http.StatusNotFound, "Error")
			return
		}
		web.Success(ctx, http.StatusCreated, sectionSaved)
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
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /sections/:id [patch]
func (s *Section) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			web.Error(ctx, http.StatusNotFound, "Error: ID inválido, seção não encontrada.")
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			web.Error(ctx, http.StatusBadRequest, "Error.")
			return
		}

		if req.Section_number != 0 {
			req.Section_number, err = s.service.Exists(req.Section_number)
			if err != nil {
				web.Error(ctx, http.StatusBadRequest, "Error: Número de seção já cadastrado.")
				return
			} else {
				err = s.service.Update(domain.Section{
					ID:            (int(id)),
					SectionNumber: req.Section_number})
				if err != nil {
					web.Error(ctx, http.StatusNotFound, "Erro ao atualizar número de seção.")
					return
				}
				web.Success(ctx, http.StatusOK, "Número de seção atualizado com sucesso!")
			}
		}
		if req.Current_temperature != 0 {
			err = s.service.Update(domain.Section{
				ID:                 (int(id)),
				CurrentTemperature: req.Current_temperature})
			if err != nil {
				web.Error(ctx, http.StatusNotFound, "Erro ao atualizar temperatura atual.")
				return
			}
			web.Success(ctx, http.StatusOK, "Temperatura atual atualizada com sucesso!")
		}
		if req.Minimum_temperature != 0 {
			err = s.service.Update(domain.Section{
				ID:                 (int(id)),
				MinimumTemperature: req.Minimum_temperature})
			if err != nil {
				web.Error(ctx, http.StatusNotFound, "Erro ao atualizar temperatura miníma.")
				return
			}
			web.Success(ctx, http.StatusOK, "Temperatura atualiza com sucesso!")

		}
		if req.Current_capacity != 0 {
			err = s.service.Update(domain.Section{
				ID:              (int(id)),
				CurrentCapacity: req.Current_capacity})
			if err != nil {
				web.Error(ctx, http.StatusNotFound, "Erro ao atualizar capacidade atual.")
				return
			}
			web.Success(ctx, http.StatusOK, "Capacidade atual atualizada com sucesso!")
		}
		if req.Minimum_capacity != 0 {
			err = s.service.Update(domain.Section{
				ID:              (int(id)),
				MinimumCapacity: req.Minimum_capacity})
			if err != nil {
				web.Error(ctx, http.StatusNotFound, "Erro ao atualizar capacidade miníma.")
				return
			}
			web.Success(ctx, http.StatusOK, "Capacidade atualizada com sucesso!")
		}
		if req.Maximum_capacity != 0 {
			err = s.service.Update(domain.Section{
				ID:              (int(id)),
				MaximumCapacity: req.Maximum_capacity})
			if err != nil {
				web.Error(ctx, http.StatusNotFound, "Erro ao atualizar capacidade máxima.")
				return
			}
			web.Success(ctx, http.StatusOK, "Capacidade atualizada com sucesso!")
		}
		if req.Warehouse_id != 0 {
			err = s.service.Update(domain.Section{
				ID:          (int(id)),
				WarehouseID: req.Warehouse_id})
			if err != nil {
				web.Error(ctx, http.StatusNotFound, "Erro ao atualizar número de armazém.")
				return
			}
			web.Success(ctx, http.StatusOK, "Número de armazém atualizado com sucesso!")
		}
		if req.Id_product_type != 0 {
			err = s.service.Update(domain.Section{
				ID:            (int(id)),
				ProductTypeID: req.Id_product_type})
			if err != nil {
				web.Error(ctx, http.StatusNotFound, "Erro ao atualizar tipo de produto.")
				return
			}
			web.Success(ctx, http.StatusOK, "Tipo de produto atualizado com sucesso!")
		}
	}
}

// Exists godoc
// @Summary Exist section number
// @Description Validate section number
// @Tags Sections
// @Accept json
// @Produce json
// @Success 204 {object} string "Section number"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /sections/sectionNumber [get]
func (s *Section) Exists() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			web.Error(ctx, http.StatusBadRequest, "Error")
			return
		}
		if req.Section_number == 0 {
			web.Error(ctx, http.StatusUnprocessableEntity, "Error: Necessário adicionar número de seção.")
			return
		}
		sectionNumber, err := s.service.Exists(req.Section_number)
		if err != nil {
			web.Error(ctx, http.StatusNoContent, "Seção não cadastrada.")
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
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /sections/:id [delete]
func (s *Section) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "Error: ID inválido.")
			return
		}

		err = s.service.Delete(int(id))
		if err != nil {
			web.Error(ctx, http.StatusNotFound, "Error")
			return
		}
		web.Success(ctx, http.StatusNoContent, "Seção deletada com sucesso.")
	}
}
