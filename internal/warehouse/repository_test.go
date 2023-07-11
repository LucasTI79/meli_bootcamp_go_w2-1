package warehouse_test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/warehouse"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryGetAll(t *testing.T) {
	t.Run("should return all warehouses", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		colums := []string{"id", "address", "telephone", "warehouse_code", "minimum,_capacity", "minimum_temperature"}
		rows := sqlmock.NewRows(colums)
		warehouseId := 1
		rows.AddRow(warehouseId, "address", "telephone", "warehouse_code", 1, 1)

		mock.ExpectQuery(regexp.QuoteMeta(warehouse.GetAllQuery)).WillReturnRows(rows)

		repository := warehouse.NewRepository(db)
		result := repository.GetAll(context.Background())
		assert.NotNil(t, result)
	})
	t.Run("Should throw panic when query execution fail", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(warehouse.GetAllQuery)).WillReturnError(sql.ErrConnDone)

		repository := warehouse.NewRepository(db)
		assert.Panics(t, func() { repository.GetAll(context.Background()) })
	})
}

func SetupMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	return db, mock
}
