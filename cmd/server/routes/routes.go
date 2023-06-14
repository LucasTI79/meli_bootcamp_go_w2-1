package routes

import (
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/handler"
	// "github.com/extmatperez/meli_bootcamp_go_w2-1/docs"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/employee"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
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

func (r *router) buildSellerRoutes() {}

func (r *router) buildProductRoutes() {}

func (r *router) buildSectionRoutes() {}

func (r *router) buildWarehouseRoutes() {}

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

func (r *router) buildBuyerRoutes() {}
