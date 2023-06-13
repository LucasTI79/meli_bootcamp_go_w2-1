package handler

import (
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/employee"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Id             int    `json:"id"`
	Card_number_id string `json:"card_number_id"`
	First_name     string `json:"first_name"`
	Last_name      string `json:"last_name"`
	Warehouse_id   int    `json:"warehouse_id"`
}

type Employee struct {
	service employee.Service
}

func NewEmployee(e employee.Service) *Employee {
	return &Employee{
		service: e,
	}
}

func (e *Employee) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "Error: ID inválido.")
			return
		}

		employeeID, err := e.service.Get(int(id))
		if err != nil {
			web.Error(ctx, http.StatusNotFound, "Error: Funcionário não encontrado.")
			return
		}
		web.Success(ctx, http.StatusOK, employeeID)
	}
}

func (e *Employee) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		employees, err := e.service.GetAll()
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "Error: Funcionários não encontrados.")
			return
		}
		web.Success(ctx, http.StatusOK, employees)
	}
}

func (e *Employee) Exists() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			web.Error(ctx, http.StatusBadRequest, "Error")
			return
		}
		if req.Card_number_id == "" {
			web.Error(ctx, http.StatusUnprocessableEntity, "Error: Necessário adicionar número do cartão.")
			return
		}
		cardNumberId, err := e.service.Exists(req.Card_number_id)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "Error")
			return
		}
		web.Success(ctx, http.StatusOK, cardNumberId)
	}
}

func (e *Employee) Save() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		if err := ctx.Bind(&req); err != nil {
			web.Error(ctx, http.StatusNotFound, "Error")
			return
		}
		if req.Card_number_id == "" && req.First_name == "" && req.Last_name == "" && req.Warehouse_id == 0 {
			web.Error(ctx, http.StatusUnprocessableEntity, "Error: Necessário adicionar todas as informações.")
			return
		}
		var err error
		if req.Card_number_id != "" {
			req.Card_number_id, err = e.service.Exists(req.Card_number_id)
			if err != nil {
				web.Error(ctx, http.StatusBadRequest, "Error: Número de cartão já cadastrado.")
				return
			}
		}
		employeeSaved, err := e.service.Save(req.Card_number_id, req.First_name, req.Last_name, req.Warehouse_id)
		if err != nil {
			web.Error(ctx, http.StatusNotFound, "Error")
			return
		}
		web.Success(ctx, http.StatusCreated, employeeSaved)
	}
}

func (e *Employee) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			web.Error(ctx, http.StatusNotFound, "Error: ID inválido, funcionário não encontrado.")
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			web.Error(ctx, http.StatusBadRequest, "Error.")
			return
		}

		if req.Card_number_id != "" {
			req.Card_number_id, err = e.service.Exists(req.Card_number_id)
			if err != nil {
				web.Error(ctx, http.StatusBadRequest, "Error: Cartão já cadastrado.")
				return
			} else {
				err = e.service.Update(domain.Employee{
					ID:           (int(id)),
					CardNumberID: req.Card_number_id})
				if err != nil {
					web.Error(ctx, http.StatusNotFound, "Erro ao atualizar número de cartão.")
					return
				}
				web.Success(ctx, http.StatusOK, "Número de cartão atualizado com sucesso!")
			}
		}
		if req.First_name != "" {
			err = e.service.Update(domain.Employee{
				ID:        (int(id)),
				FirstName: req.First_name})
			if err != nil {
				web.Error(ctx, http.StatusNotFound, "Erro ao atualizar nome de funcionáerio.")
				return
			}
			web.Success(ctx, http.StatusOK, "Nome atualizado com sucesso!")
		}
		if req.Last_name != "" {
			err = e.service.Update(domain.Employee{
				ID:       (int(id)),
				LastName: req.Last_name})
			if err != nil {
				web.Error(ctx, http.StatusNotFound, "Erro ao atualizar sobrenome de funcionário.")
				return
			}
			web.Success(ctx, http.StatusOK, "Sobrenome atualizado com sucesso!")

		}
		if req.Warehouse_id != 0 {
			err = e.service.Update(domain.Employee{
				ID:          (int(id)),
				WarehouseID: req.Warehouse_id})
			if err != nil {
				web.Error(ctx, http.StatusNotFound, "Erro ao atualizar número de armazém.")
				return
			}
			web.Success(ctx, http.StatusOK, "Armazém atualizado com sucesso!")
		}
	}
}

func (e *Employee) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "Error: ID inválido.")
			return
		}

		err = e.service.Delete(int(id))
		if err != nil {
			web.Error(ctx, http.StatusNotFound, "Error")
			return
		}
		web.Success(ctx, http.StatusNoContent, "Funcionário deletado com sucesso.")
	}
}
