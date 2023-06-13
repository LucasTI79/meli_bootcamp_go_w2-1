package routes

import (
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/section"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	eng *gin.Engine
	rg  *gin.RouterGroup
	db  *sql.DB
}

func NewRouter(eng *gin.Engine, db *sql.DB) Router {
	return &router{eng: eng, db: db}
}

func (r *router) MapRoutes() {
	r.setGroup()

	r.buildSellerRoutes()
	r.buildProductRoutes()
	r.buildSectionRoutes()
	r.buildWarehouseRoutes()
	r.buildEmployeeRoutes()
	r.buildBuyerRoutes()
}

func (r *router) setGroup() {
	r.rg = r.eng.Group("/api/v1")
}

func (r *router) buildSellerRoutes() {}

func (r *router) buildProductRoutes() {}

func (r *router) buildSectionRoutes() {
	repository := section.NewRepository(r.db)
	service := section.NewService(repository)
	handler := handler.NewSection(service)

	r.rg.GET("/sections", handler.GetAll())
	r.rg.POST("/section", handler.Save())
	r.rg.PATCH("/section/:id", handler.Update())
	r.rg.GET("/section/:id", handler.Get())
	r.rg.GET("/section/sectionNumber", handler.Exists())
	r.rg.DELETE("/section/:id", handler.Delete())
}

func (r *router) buildWarehouseRoutes() {}

func (r *router) buildEmployeeRoutes() {}

func (r *router) buildBuyerRoutes() {}
