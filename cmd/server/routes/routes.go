package routes

import (
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/docs"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/employee"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/section"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/seller"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/warehouse"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
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

	docs.SwaggerInfo.Host = os.Getenv("HOST")
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
	repo := seller.NewRepository(r.db)
	service := seller.NewService(repo)
	handler := handler.NewSeller(service)
	r.rg.GET("/sellers", handler.GetAll())
	r.rg.GET("/sellers/:id", handler.Get())
	r.rg.POST("/sellers", handler.Create())
	r.rg.PATCH("/sellers/:id", handler.Update())
	r.rg.DELETE("/sellers/:id", handler.Delete())
}

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

func (r *router) buildEmployeeRoutes() {
	repository := employee.NewRepository(r.db)
	service := employee.NewService(repository)
	handler := handler.NewEmployee(service)

	r.rg.GET("/employees", handler.GetAll())
	r.rg.POST("/employees", handler.Save())
	r.rg.PATCH("/employees/:id", handler.Update())
	r.rg.GET("/employees/:id", handler.Get())
	r.rg.GET("/employees/cardNumber", handler.Exists())
	r.rg.DELETE("/employees/:id", handler.Delete())
}

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
