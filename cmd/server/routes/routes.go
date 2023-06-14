package routes

import (
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/handler"
	// "github.com/extmatperez/meli_bootcamp_go_w2-1/docs"
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
	// docs.SwaggerInfo.Host = "http://localhost:8080/"
	// r.rg.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
	r.rg.POST("/sections", handler.Save())
	r.rg.PATCH("/sections/:id", handler.Update())
	r.rg.GET("/sections/:id", handler.Get())
	r.rg.GET("/sections/sectionNumber", handler.Exists())
	r.rg.DELETE("/sections/:id", handler.Delete())
}

func (r *router) buildWarehouseRoutes() {}

func (r *router) buildEmployeeRoutes() {}

func (r *router) buildBuyerRoutes() {}
