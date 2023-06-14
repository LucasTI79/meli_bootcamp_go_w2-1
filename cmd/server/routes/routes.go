package routes

import (
	"database/sql"
	"os"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/docs"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/seller"
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
	// Example
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
