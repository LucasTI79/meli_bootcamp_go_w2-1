package routes

import (
	"database/sql"
	"os"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/docs"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/employee"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product"
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

	docs.SwaggerInfo.BasePath = "/api/v1"
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
	repo := seller.NewRepository(r.db)
	service := seller.NewService(repo)
	handler := handler.NewSeller(service)
	sellerRoutes := r.rg.Group("/sellers")

	sellerRoutes.GET("/", handler.GetAll())
	sellerRoutes.GET("/:id", handler.Get())
	sellerRoutes.POST("/", handler.Create())
	sellerRoutes.PATCH("/:id", handler.Update())
	sellerRoutes.DELETE("/:id", handler.Delete())
}

func (r *router) buildProductRoutes() {
	repo := product.NewRepository(r.db)
	service := product.NewService(repo)
	handler := handler.NewProduct(service)
	productRoutes := r.rg.Group("/products")

	productRoutes.GET("/", handler.GetAll())
	productRoutes.GET("/:id", handler.Get())
	productRoutes.POST("/", handler.Create())
	productRoutes.PATCH("/:id", handler.Update())
	productRoutes.DELETE("/:id", handler.Delete())
}

func (r *router) buildSectionRoutes() {
	repository := section.NewRepository(r.db)
	service := section.NewService(repository)
	handler := handler.NewSection(service)
	sectionRoutes := r.rg.Group("/sections")

	sectionRoutes.GET("/", handler.GetAll())
	sectionRoutes.POST("/", handler.Save())
	sectionRoutes.PATCH("/:id", handler.Update())
	sectionRoutes.GET("/:id", handler.Get())
	sectionRoutes.GET("/sectionNumber", handler.Exists())
	sectionRoutes.DELETE("/:id", handler.Delete())
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
	employeeRoutes := r.rg.Group("/employees")

	employeeRoutes.GET("/", handler.GetAll())
	employeeRoutes.POST("/", handler.Save())
	employeeRoutes.PATCH("/:id", handler.Update())
	employeeRoutes.GET("/:id", handler.Get())
	employeeRoutes.GET("/cardNumber", handler.Exists())
	employeeRoutes.DELETE("/:id", handler.Delete())
}

func (r *router) buildBuyerRoutes() {
	repo := buyer.NewRepository(r.db)
	service := buyer.NewService(repo)
	handler := handler.NewBuyer(service)
	buyerRoutes := r.rg.Group("/buyers")

	buyerRoutes.GET("/", handler.GetAll())
	buyerRoutes.GET("/:id", handler.Get())
	buyerRoutes.POST("/", handler.Create())
	buyerRoutes.PATCH("/:id", handler.Update())
	buyerRoutes.DELETE("/:id", handler.Delete())
}
