package warehouse_test

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/warehouse"
	"github.com/stretchr/testify/assert"
)

var (
	mockedWarehouseTemplate = domain.Warehouse{
		ID:                 1,
		Address:            "address",
		Telephone:          "telephone",
		WarehouseCode:      "warehouse_code",
		MinimumCapacity:    1,
		MinimumTemperature: 1,
	}
)

func TestRepositoryGetAll(t *testing.T) {
	t.Run("should return all warehouses", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		colums := []string{"id", "address", "telephone", "warehouse_code", "minimum,_capacity", "minimum_temperature"}
		rows := sqlmock.NewRows(colums)
		warehouseId := 1
		rows.AddRow(warehouseId, "address", "telephone", "warehouse_code", 1, 1)

		mock.ExpectQuery(warehouse.GetAllQuery).WillReturnRows(rows)

		repository := warehouse.NewRepository(db)
		result := repository.GetAll()
		assert.NotNil(t, result)
	})
	t.Run("Should throw panic when query execution fail", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mock.ExpectQuery(warehouse.GetAllQuery).WillReturnError(sql.ErrConnDone)

		repository := warehouse.NewRepository(db)
		assert.Panics(t, func() { repository.GetAll() })
	})
}

func TestRepositoryGet(t *testing.T) {
	t.Run("Should return a seller by specified id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id", "address", "telephone", "warehouse_code", "minimum_capacity", "minimum_temperature"}
		rows := sqlmock.NewRows(columns)
		warehouseId := 1
		rows.AddRow(warehouseId, "address", "telephone", "warehouse_code", 1, 1)

		mock.ExpectQuery(warehouse.GetQuery).WithArgs(warehouseId).WillReturnRows(rows)

		repository := warehouse.NewRepository(db)

		result := repository.Get(warehouseId)

		assert.NotNil(t, result)
	})

	t.Run("Should not return a warehouse", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		warehouseId := 1

		mock.ExpectQuery(warehouse.GetQuery).WithArgs(warehouseId).WillReturnError(sql.ErrNoRows)

		repository := warehouse.NewRepository(db)

		result := repository.Get(warehouseId)

		assert.Nil(t, result)
	})

	t.Run("Should throw panic", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		warehouseId := 1

		mock.ExpectQuery(warehouse.GetQuery).WithArgs(warehouseId).WillReturnError(sql.ErrConnDone)

		repository := warehouse.NewRepository(db)

		assert.Panics(t, func() { repository.Get(warehouseId) })
	})
}

func TestRepositoryExists(t *testing.T) {
	t.Run("Should return true", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"warehouseCode"}
		rows := sqlmock.NewRows(columns)
		warehouseCode := "123"
		rows.AddRow(warehouseCode)

		mock.ExpectQuery(regexp.QuoteMeta(warehouse.ExistsQuery)).
			WithArgs(warehouseCode).
			WillReturnRows(rows)

		repository := warehouse.NewRepository(db)

		result := repository.Exists(warehouseCode)

		assert.True(t, result)
	})

	t.Run("Should return false when there are no query results", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		warehouseCode := "123"

		mock.ExpectQuery(warehouse.ExistsQuery).WithArgs(warehouseCode).WillReturnError(sql.ErrNoRows)

		repository := warehouse.NewRepository(db)

		result := repository.Exists(warehouseCode)

		assert.False(t, result)
	})

	t.Run("Should return false when database has internal error", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		warehouseCode := "123"

		mock.ExpectQuery(warehouse.ExistsQuery).WithArgs(warehouseCode).WillReturnError(sql.ErrConnDone)

		repository := warehouse.NewRepository(db)

		result := repository.Exists(warehouseCode)

		assert.False(t, result)
	})
}

func TestRepositorySave(t *testing.T) {
	t.Run("Should insert the warehouse and return the warehouse id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"address", "telephone", "warehouse_code", "minimum_capacity", "minimum_temperature"}
		rows := sqlmock.NewRows(columns)
		rows.AddRow("address", "telephone", "warehouse_code", 1, 1)

		lastInsertId := 1
		mockedWarehouse := mockedWarehouseTemplate

		mock.ExpectPrepare(regexp.QuoteMeta(warehouse.InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(warehouse.InsertQuery)).
			WithArgs(mockedWarehouse.Address, mockedWarehouse.Telephone, mockedWarehouse.WarehouseCode, mockedWarehouse.MinimumCapacity, mockedWarehouse.MinimumTemperature).
			WillReturnResult(sqlmock.NewResult(int64(lastInsertId), 1))

		repository := warehouse.NewRepository(db)

		result := repository.Save(mockedWarehouse)

		assert.Equal(t, lastInsertId, result)

	})
	t.Run("Should throw panic when expect prepare fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedWarehouse := mockedWarehouseTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(warehouse.InsertQuery)).WillReturnError(sql.ErrConnDone)

		repository := warehouse.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedWarehouse) })
	})
	t.Run("Should throw panic when expect exec fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedWarehouse := mockedWarehouseTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(warehouse.InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(warehouse.InsertQuery)).
			WithArgs(mockedWarehouseTemplate.Address, mockedWarehouse.Telephone, mockedWarehouse.WarehouseCode, mockedWarehouse.MinimumCapacity, mockedWarehouse.MinimumTemperature).
			WillReturnError(sql.ErrConnDone)

		repository := warehouse.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedWarehouse) })
	})

	t.Run("Should throw panic when sql has error", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedWarehouse := mockedWarehouseTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(warehouse.InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(warehouse.InsertQuery)).
			WithArgs(mockedWarehouseTemplate.Address, mockedWarehouse.Telephone, mockedWarehouse.WarehouseCode, mockedWarehouse.MinimumCapacity, mockedWarehouse.MinimumTemperature).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrConnDone))

		repository := warehouse.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedWarehouse) })
	})
}

func TestRepositoryUpdate(t *testing.T) {
	t.Run("Should update the warehouse", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedWarehouse := mockedWarehouseTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(warehouse.UpdateQuery))
		mock.ExpectExec(regexp.QuoteMeta(warehouse.UpdateQuery)).
			WithArgs(mockedWarehouse.Address, mockedWarehouse.Telephone, mockedWarehouse.WarehouseCode, mockedWarehouse.MinimumCapacity, mockedWarehouse.MinimumTemperature, mockedWarehouse.ID).
			WillReturnResult(sqlmock.NewResult(int64(mockedWarehouse.ID), 1))

		repository := warehouse.NewRepository(db)

		assert.NotPanics(t, func() { repository.Update(mockedWarehouse) })
	})

	t.Run("Should throw panic when expected prepare fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedWarehouse := mockedWarehouseTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(warehouse.UpdateQuery)).WillReturnError(sql.ErrConnDone)

		repository := warehouse.NewRepository(db)

		assert.Panics(t, func() { repository.Update(mockedWarehouse) })
	})

	t.Run("Should throw panic when expected exec fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedWarehouse := mockedWarehouseTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(warehouse.UpdateQuery))
		mock.ExpectExec(regexp.QuoteMeta(warehouse.UpdateQuery)).
			WithArgs(mockedWarehouse.Address, mockedWarehouse.Telephone, mockedWarehouse.WarehouseCode, mockedWarehouse.MinimumCapacity, mockedWarehouse.MinimumTemperature, mockedWarehouse.ID).
			WillReturnError(sql.ErrConnDone)

		repository := warehouse.NewRepository(db)

		assert.Panics(t, func() { repository.Update(mockedWarehouse) })
	})
}

func TestRepositoryDelete(t *testing.T) {
	t.Run("Should delete the warehouse", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedWarehouse := mockedWarehouseTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(warehouse.DeleteQuery))
		mock.ExpectExec(regexp.QuoteMeta(warehouse.DeleteQuery)).
			WithArgs(mockedWarehouse.ID).
			WillReturnResult(sqlmock.NewResult(int64(mockedWarehouse.ID), 1))

		repository := warehouse.NewRepository(db)

		assert.NotPanics(t, func() { repository.Delete(mockedWarehouse.ID) })
	})

	t.Run("Should throw panic when expected prepare fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedWarehouse := mockedWarehouseTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(warehouse.DeleteQuery)).WillReturnError(sql.ErrConnDone)

		repository := warehouse.NewRepository(db)

		assert.Panics(t, func() { repository.Delete(mockedWarehouse.ID) })
	})

	t.Run("Should throw panic when expected exec fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedWarehouse := mockedWarehouseTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(warehouse.DeleteQuery))
		mock.ExpectExec(regexp.QuoteMeta(warehouse.DeleteQuery)).
			WithArgs(mockedWarehouse.ID).
			WillReturnError(sql.ErrConnDone)

		repository := warehouse.NewRepository(db)

		assert.Panics(t, func() { repository.Delete(mockedWarehouse.ID) })
	})
}

func SetupMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	return db, mock
}
