package routes

import (
	"database/sql"
	"os"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/middleware"
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

	r.buildDocumentationRoutes()
	r.defineGlobalMiddlewares()
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

func (r *router) defineGlobalMiddlewares() {
	r.rg.Use(middleware.InternalError())
}

func (r *router) buildDocumentationRoutes() {
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.rg.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (r *router) buildSellerRoutes() {
	repo := seller.NewRepository(r.db)
	service := seller.NewService(repo)
	controller := handler.NewSeller(service)
	sellerRoutes := r.rg.Group("/sellers")

	sellerRoutes.GET("/", controller.GetAll())
	sellerRoutes.GET("/:id", controller.Get())
	sellerRoutes.POST("/", middleware.Validation[handler.CreateSellerRequest](), controller.Create())
	sellerRoutes.PATCH("/:id", middleware.Validation[handler.UpdateSellerRequest](), controller.Update())
	sellerRoutes.DELETE("/:id", controller.Delete())
}

func (r *router) buildProductRoutes() {
	repo := product.NewRepository(r.db)
	service := product.NewService(repo)
	controller := handler.NewProduct(service)
	productRoutes := r.rg.Group("/products")

	productRoutes.GET("/", controller.GetAll())
	productRoutes.GET("/:id", controller.Get())
	productRoutes.POST("/", middleware.Validation[handler.CreateProductRequest](), controller.Create())
	productRoutes.PATCH("/:id", middleware.Validation[handler.UpdateProductRequest](), controller.Update())
	productRoutes.DELETE("/:id", controller.Delete())
}

func (r *router) buildSectionRoutes() {
	repository := section.NewRepository(r.db)
	service := section.NewService(repository)
	controller := handler.NewSection(service)
	sectionRoutes := r.rg.Group("/sections")
	
	sectionRoutes.GET("/", controller.GetAll())
	sectionRoutes.POST("/", middleware.Validation[handler.CreateSectionRequest](), controller.Create())
	sectionRoutes.PATCH("/:id", middleware.Validation[handler.UpdateSectionRequest](), controller.Update())
	sectionRoutes.GET("/:id", controller.Get())
	sectionRoutes.DELETE("/:id", controller.Delete())
}

func (r *router) buildWarehouseRoutes() {

	repo := warehouse.NewRepository(r.db)
	service := warehouse.NewService(repo)
	controller := handler.NewWarehouse(service)
	routes := r.rg.Group("/warehouses")

	routes.GET("/", controller.GetAll())
	routes.GET("/:id", controller.Get())
	routes.POST("/", middleware.Validation[handler.CreateWarehouseRequest](), controller.Create())
	routes.PATCH("/:id", middleware.Validation[handler.UpdateWarehouseRequest](), controller.Update())
	routes.DELETE("/:id", controller.Delete())
}

func (r *router) buildEmployeeRoutes() {
	repository := employee.NewRepository(r.db)
	service := employee.NewService(repository)
	controller := handler.NewEmployee(service)
	employeeRoutes := r.rg.Group("/employees")

	employeeRoutes.GET("/", controller.GetAll())
	employeeRoutes.GET("/:id", controller.Get())
	employeeRoutes.POST("/", middleware.Validation[handler.CreateEmployeeRequest](), controller.Create())
	employeeRoutes.PATCH("/:id", middleware.Validation[handler.UpdateEmployeeRequest](), controller.Update())
	employeeRoutes.DELETE("/:id", controller.Delete())
}

func (r *router) buildBuyerRoutes() {
	repo := buyer.NewRepository(r.db)
	service := buyer.NewService(repo)
	controller := handler.NewBuyer(service)
	buyerRoutes := r.rg.Group("/buyers")

	buyerRoutes.GET("/", controller.GetAll())
	buyerRoutes.GET("/:id", controller.Get())
	buyerRoutes.POST("/", middleware.Validation[handler.CreateBuyerRequest](), controller.Create())
	buyerRoutes.PATCH("/:id", middleware.Validation[handler.UpdateBuyerRequest](), controller.Update())
	buyerRoutes.DELETE("/:id", controller.Delete())
}
