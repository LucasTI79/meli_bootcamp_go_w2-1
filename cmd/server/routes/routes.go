package routes

import (
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/docs"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/warehouse"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	docs.SwaggerInfo.Host = "localhost:8080/"
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

func (r *router) buildWarehouseRoutes() {

	warehouseRepo := warehouse.NewWarehouseRepository(r.db)
	warehouseService := warehouse.NewWarehouseService(warehouseRepo)
	warehouseHandler := handler.NewWarehouse(warehouseService)

	r.rg.GET("/warehouses", warehouseHandler.GetAll())
	r.rg.GET("/warehouses/:id", warehouseHandler.GetByID())
	r.rg.POST("/warehouses", warehouseHandler.Create())
	r.rg.PUT("/warehouses", warehouseHandler.Update())
	r.rg.PATCH("/warehouses/:id", warehouseHandler.UpdateByID())
	r.rg.DELETE("/warehouses/:id", warehouseHandler.Delete())
}

func (r *router) buildEmployeeRoutes() {}

func (r *router) buildBuyerRoutes() {}
