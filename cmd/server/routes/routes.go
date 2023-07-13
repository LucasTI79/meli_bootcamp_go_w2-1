package routes

import (
	"database/sql"
	"os"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/middleware"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/docs"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/carrier"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/employee"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/inbound_orders"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/locality"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/order_status"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product_batch"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product_record"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product_type"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/province"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/purchase_orders"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/section"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/seller"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/warehouse"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	CreateCanBeBlank = true
	UpdateCanBeBlank = false
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
	r.buildLocalityRoutes()
	r.buildProductRecordRoutes()
}

func (r *router) setGroup() {
	r.rg = r.eng.Group("/api/v1")
}

func (r *router) defineGlobalMiddlewares() {
	r.rg.Use(middleware.InternalError())
	r.rg.Use(middleware.IdValidation())
}

func (r *router) buildDocumentationRoutes() {
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.rg.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (r *router) buildSellerRoutes() {
	repo := seller.NewRepository(r.db)
	localityRepo := locality.NewRepository(r.db)
	service := seller.NewService(repo, localityRepo)
	controller := handler.NewSeller(service)
	sellerRoutes := r.rg.Group("/sellers")

	sellerRoutes.GET("/", controller.GetAll())
	sellerRoutes.GET("/:id", controller.Get())
	sellerRoutes.POST("/", middleware.RequestValidation[handler.CreateSellerRequest](CreateCanBeBlank), controller.Create())
	sellerRoutes.PATCH("/:id", middleware.RequestValidation[handler.UpdateSellerRequest](UpdateCanBeBlank), controller.Update())
	sellerRoutes.DELETE("/:id", controller.Delete())
}

func (r *router) buildProductRoutes() {
	repo := product.NewRepository(r.db)
	productTypeRepo := product_type.NewRepository(r.db)
	sellerRepo := seller.NewRepository(r.db)
	service := product.NewService(repo, productTypeRepo, sellerRepo)
	controller := handler.NewProduct(service)
	productRoutes := r.rg.Group("/products")

	productRoutes.GET("/", controller.GetAll())
	productRoutes.GET("/:id", controller.Get())
	productRoutes.POST("/", middleware.RequestValidation[handler.CreateProductRequest](CreateCanBeBlank), controller.Create())
	productRoutes.PATCH("/:id", middleware.RequestValidation[handler.UpdateProductRequest](UpdateCanBeBlank), controller.Update())
	productRoutes.DELETE("/:id", controller.Delete())
	productRoutes.GET("/report-records", controller.ReportRecords())
}

func (r *router) buildSectionRoutes() {
	repository := section.NewRepository(r.db)
	service := section.NewService(repository)
	controller := handler.NewSection(service)
	sectionRoutes := r.rg.Group("/sections")

	sectionRoutes.GET("/", controller.GetAll())
	sectionRoutes.POST("/", middleware.RequestValidation[handler.CreateSectionRequest](CreateCanBeBlank), controller.Create())
	sectionRoutes.PATCH("/:id", middleware.RequestValidation[handler.UpdateSectionRequest](UpdateCanBeBlank), controller.Update())
	sectionRoutes.GET("/:id", controller.Get())
	sectionRoutes.DELETE("/:id", controller.Delete())
}

func (r *router) buildWarehouseRoutes() {
	repo := warehouse.NewRepository(r.db)
	service := warehouse.NewService(repo)
	controller := handler.NewWarehouse(service)
	warehouseRoutes := r.rg.Group("/warehouses")

	warehouseRoutes.GET("/", controller.GetAll())
	warehouseRoutes.GET("/:id", controller.Get())
	warehouseRoutes.POST("/", middleware.RequestValidation[handler.CreateWarehouseRequest](CreateCanBeBlank), controller.Create())
	warehouseRoutes.PATCH("/:id", middleware.RequestValidation[handler.UpdateWarehouseRequest](UpdateCanBeBlank), controller.Update())
	warehouseRoutes.DELETE("/:id", controller.Delete())
}

func (r *router) buildEmployeeRoutes() {
	repository := employee.NewRepository(r.db)
	service := employee.NewService(repository)
	controller := handler.NewEmployee(service)
	employeeRoutes := r.rg.Group("/employees")

	employeeRoutes.GET("/", controller.GetAll())
	employeeRoutes.GET("/:id", controller.Get())
	employeeRoutes.GET("/report-inbound-orders", controller.ReportInboundOrders())
	employeeRoutes.POST("/", middleware.RequestValidation[handler.CreateEmployeeRequest](CreateCanBeBlank), controller.Create())
	employeeRoutes.PATCH("/:id", middleware.RequestValidation[handler.UpdateEmployeeRequest](UpdateCanBeBlank), controller.Update())
	employeeRoutes.DELETE("/:id", controller.Delete())
}

func (r *router) buildBuyerRoutes() {
	repo := buyer.NewRepository(r.db)
	service := buyer.NewService(repo)
	controller := handler.NewBuyer(service)
	buyerRoutes := r.rg.Group("/buyers")

	buyerRoutes.GET("/", controller.GetAll())
	buyerRoutes.GET("/:id", controller.Get())
	buyerRoutes.POST("/", middleware.RequestValidation[handler.CreateBuyerRequest](CreateCanBeBlank), controller.Create())
	buyerRoutes.PATCH("/:id", middleware.RequestValidation[handler.UpdateBuyerRequest](UpdateCanBeBlank), controller.Update())
	buyerRoutes.DELETE("/:id", controller.Delete())
	buyerRoutes.GET("/report-purchase-orders", controller.ReportPurchases())
}

func (r *router) buildLocalityRoutes() {
	repo := locality.NewRepository(r.db)
	provinceRepo := province.NewRepository(r.db)
	service := locality.NewService(repo, provinceRepo)
	controller := handler.NewLocality(service)
	localityRoutes := r.rg.Group("/localities")

	localityRoutes.POST("/", middleware.RequestValidation[handler.CreateLocalityRequest](CreateCanBeBlank), controller.Create())
	localityRoutes.GET("/report-sellers", controller.ReportSellers())
}

func (r *router) buildProductRecordRoutes() {
	repo := product_record.NewRepository(r.db)
	productRepo := product.NewRepository(r.db)
	service := product_record.NewService(repo, productRepo)
	controller := handler.NewProductRecord(service)
	productRecordRoutes := r.rg.Group("/product-records")

	productRecordRoutes.POST("/", middleware.RequestValidation[handler.CreateProductRecordRequest](CreateCanBeBlank), controller.Create())
}

func (r *router) buildPurchaseOrdersRoutes() {
	repo := purchase_orders.NewRepository(r.db)
	buyerRepo := buyer.NewRepository(r.db)
	orderStatusRepo := order_status.NewRepository(r.db)
	warehouseRepo := warehouse.NewRepository(r.db)
	carrierRepo := carrier.NewRepository(r.db)
	productRecordRepo := product_record.NewRepository(r.db)
	service := purchase_orders.NewService(repo, buyerRepo, orderStatusRepo, warehouseRepo, carrierRepo, productRecordRepo)
	controller := handler.NewPurchaseOrder(service)
	purchaseOrdersRoutes := r.rg.Group("/purchase-orders")

	purchaseOrdersRoutes.POST("/", middleware.RequestValidation[handler.CreatePurchaseOrderRequest](CreateCanBeBlank), controller.Create())
}

func (r *router) buildInboundOrdersRoutes() {
	repo := inbound_orders.NewRepository(r.db)
	repoEmployee := employee.NewRepository(r.db)
	repoProductBatch := product_batch.NewRepository(r.db)
	repoWarehouse := warehouse.NewRepository(r.db)
	service := inbound_orders.NewService(repo, repoEmployee, repoProductBatch, repoWarehouse)
	controller := handler.NewInboundOrder(service)
	inboundOrdersRoutes := r.rg.Group("/inbound-orders")

	inboundOrdersRoutes.POST("/", middleware.RequestValidation[handler.CreateInboundOrderRequest](CreateCanBeBlank), controller.Create())
}
