package main

import (
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	/// @title MELI Bootcamp GO Wave 2
	// @version 1.0
	// @description This API Handle MELI Products.
	// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

	// @contact.name API Support
	// @contact.url https://developers.mercadolibre.com.ar/support

	// @license.name Apache 2.0
	// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

	db, err := sql.Open("mysql", "meli_sprint_user:Meli_Sprint#123@/melisprint")
	if err != nil {
		panic(err)
	}

	eng := gin.Default()

	router := routes.NewRouter(eng, db)
	router.MapRoutes()

	if err := eng.Run(); err != nil {
		panic(err)
	}
}
