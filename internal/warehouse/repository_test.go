package warehouse_test

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/warehouse"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryGetAll(t *testing.T) {
	t.Run("should return all warehouses", func(t *testing.T) {
		db, mock := SteupMock(t)
		defer db.Close()

		colums := []string{"id", "address", "telephone", "warehouse_code", "minimum,_capacity", "minimum_temperature"}
		rows := sqlmock.NewRows(colums)
		warehouseId := 1
		rows.AddRow(warehouseId, "address", "telephone", "warehouse_code", 1, 1)

		mock.ExpectQuery(warehouse.G).
		WillReturnRows(rows)