package routes

import (
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/docs"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/buyer"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type IRouter interface {
	MapRoutes()
}

type router struct {
	eng *gin.Engine
	rg  *gin.RouterGroup
	db  *sql.DB
}

func NewRouter(eng *gin.Engine, db *sql.DB) IRouter {
	return &router{eng: eng, db: db}
}

func (r *router) MapRoutes() {
	r.setGroup()

	docs.SwaggerInfo.Host = "http://localhost:8080/"
	r.rg.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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

func (r *router) buildSellerRoutes() {
	// Example
	// repo := seller.NewRepository(r.db)
	// service := seller.NewService(repo)
	// handler := handler.NewSeller(service)
	// r.r.GET("/seller", handler.GetAll)
}

func (r *router) buildProductRoutes() {}

func (r *router) buildSectionRoutes() {}

func (r *router) buildWarehouseRoutes() {}

func (r *router) buildEmployeeRoutes() {}

func (r *router) buildBuyerRoutes() {

	repo := buyer.NewRepository(r.db)
	service := buyer.NewService(repo)
	handler := handler.NewBuyer(service)

	r.rg.GET("/buyers", handler.GetAll())
	r.rg.GET("/buyers/:id", handler.Get())
	r.rg.POST("/buyers", handler.Create())
	r.rg.PATCH("/buyers/:id", handler.Update())
	r.rg.DELETE("/buyers/:id", handler.Delete())
}
