package product_batches_test

import (
	"database/sql"
	//"regexp"
	"testing"

	//"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product_batches"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryCreate(t *testing.T) {
	t.Run("should create a product batch", func(t *testing.T) {
		db := new(sql.DB)
		defer db.Close()

		repository := product_batches.NewRepository(db)

		productBatch := domain.ProductBatches{
			BatchNumber:        1,
			CurrentQuantity:    1,
			CurrentTemperature: 2,
			DueDate:            "2021-01-01",
			InitialQuantity:    10,
			ManufacturingDate:  "2021-01-01",
			ManufacturingHour:  "10:00",
			MinimumTemperature: 0,
			ProductID:          1,
			SectionID:          1,
		}

		id, _ := repository.Save(productBatch)

		var request int
		row := db.QueryRow("SELECT batch_number FROM product_batches WHERE id = ?", id)
		row.Scan(&request)

		assert.Equal(t, productBatch.BatchNumber, request)
	})
}
