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

func NewSeller(s seller.Service) *Seller {
	return &Seller{
		sellerService: s,
	}

}

func (s *Seller) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		sellers, err := s.sellerService.GetAll(c.Request.Context())
		if err != nil {
			web.Error(c, http.StatusInternalServerError, "internal server error")
			return
		}
		if len(sellers) == 0 {
			web.Response(c, 204, sellers)
			return
		}
		web.Success(c, http.StatusOK, sellers)
	}
}

func (s *Seller) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")
		parsedId, err := strconv.Atoi(id)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "id received is invalid")
			return
		}
		foundSeller, err := s.sellerService.Get(c.Request.Context(), parsedId)
		if err != nil {
			if errors.Is(err, seller.ErrNotFound) {
				web.Error(c, http.StatusNotFound, "could not find id %v", parsedId)
				return
			} else {
				web.Error(c, http.StatusInternalServerError, "internal server error")
				return
			}
		}
		web.Response(c, http.StatusOK, foundSeller)
	}
}

func (s *Seller) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.CreateSeller
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, http.StatusBadRequest, "invalid request body")
			return
		}
		if req.Address == "" {
			web.Error(c, http.StatusUnprocessableEntity, "address is required")
			return
		}
		if req.CompanyName == "" {
			web.Error(c, http.StatusUnprocessableEntity, "company name is required")
			return
		}
		if req.Telephone == "" {
			web.Error(c, http.StatusUnprocessableEntity, "telephone is required")
			return
		}
		if req.CID == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "cid is required")
			return
		}
		id, err := s.sellerService.Save(c.Request.Context(), req)
		if err != nil {
			if errors.Is(err, seller.ErrAlreadyExists) {
				web.Error(c, http.StatusConflict, "there is already a seller with cid %v", req.CID)
				return
			} else {
				web.Error(c, http.StatusInternalServerError, "internal server error")
				return
			}
		}
		web.Response(c, http.StatusCreated, id)
	}
}

func (s *Seller) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")
		parsedId, err := strconv.Atoi(id)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "id received is invalid")
			return
		}
		var req domain.UpdateSeller
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, http.StatusBadRequest, "invalid request body")
			return
		}
		req.ID = parsedId
		updatedSeller, err := s.sellerService.Update(c.Request.Context(), req)
		if err != nil {
			if errors.Is(err, seller.ErrAlreadyExists) {
				web.Error(c, http.StatusConflict, "cid already registred")
				return
			} else if errors.Is(err, seller.ErrNotFound) {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			} else {
				web.Error(c, http.StatusInternalServerError, "internal server error")
				return
			}
		}
		web.Success(c, http.StatusOK, updatedSeller)
	}
}

func (s *Seller) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")
		parsedId, err := strconv.Atoi(id)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "id received is invalid")
			return
		}
		err = s.sellerService.Delete(c.Request.Context(), parsedId)
		if err != nil {
			if errors.Is(err, seller.ErrNotFound) {
				web.Error(c, http.StatusNotFound, "could not find seller with id %v", parsedId)
				return
			} else {
				web.Error(c, http.StatusInternalServerError, "internal server error")
				return
			}
		}
		web.Response(c, http.StatusNoContent, "")
	}
}
