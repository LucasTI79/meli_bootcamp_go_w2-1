package product_batch_test

import (
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product_batch"
	"github.com/stretchr/testify/assert"
)

var (
	pb = domain.ProductBatch{
		ID:                 1,
		BatchNumber:        1,
		CurrentQuantity:    1,
		CurrentTemperature: 2,
		DueDate:            time.Date(2021, 01, 01, 10, 10, 10, 10, time.UTC),
		InitialQuantity:    10,
		ManufacturingDate:  time.Date(2021, 01, 01, 10, 10, 10, 10, time.UTC),
		ManufacturingHour:  10,
		MinimumTemperature: 0,
		ProductID:          1,
		SectionID:          1,
	}
)

func TestRepositoryExists(t *testing.T) {
	t.Run("Should return true when product batches number exists", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"BatchNumber"}
		rows := sqlmock.NewRows(columns)
		rows.AddRow(pb.BatchNumber)

		mock.ExpectQuery(regexp.QuoteMeta(product_batch.ExistsQuery)).
			WithArgs(pb.BatchNumber).
			WillReturnRows(rows)

		repository := product_batch.NewRepository(db)
		result := repository.Exists(pb.BatchNumber)
		assert.True(t, result)
	})

	t.Run("Should return false when product batches number does not exist", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		batchNumber := 1

		mock.ExpectQuery(product_batch.ExistsQuery).WithArgs(batchNumber).WillReturnError(sql.ErrNoRows)

		repository := product_batch.NewRepository(db)
		result := repository.Exists(batchNumber)
		assert.False(t, result)
	})
}

func TestRepositoryGet(t *testing.T) {
	t.Run("Should return a product batches by ID", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id", "batch_number", "current_quantity", "current_temperature", "due_date", "initial_quantity", "manufacturing_date", "manufacturing_hour", "minimum_temperature", "product_id", "section_id"}
		rows := sqlmock.NewRows(columns)
		batchNumber := 1

		rows.AddRow(batchNumber, 1, 1, 2, "2021-01-01 10:00:00", 10, "2021-01-01 10:00:00", 10, 0, 1, 1)

		mock.ExpectQuery(regexp.QuoteMeta(product_batch.GetQuery)).
			WithArgs(batchNumber).
			WillReturnRows(rows)

		repository := product_batch.NewRepository(db)

		result := repository.Get(batchNumber)

		assert.NotNil(t, result)
	})
	t.Run("Should not return a product batches", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id", "batch_number", "current_quantity", "current_temperature", "due_date", "initial_quantity", "manufacturing_date", "manufacturing_hour", "minimum_temperature", "product_id", "section_id"}
		rows := sqlmock.NewRows(columns)
		batchNumber := 1
		rows.AddRow(batchNumber, 1, 1, 2, "2021-01-01 10:00:00", 10, "2021-01-01 10:00:00", 10, 0, 1, 1)

		mock.ExpectQuery(product_batch.GetQuery).WithArgs(batchNumber).WillReturnError(sql.ErrNoRows)

		repository := product_batch.NewRepository(db)
		result := repository.Get(batchNumber)
		assert.Nil(t, result)
	})

	t.Run("Should throw panic", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		batchNumber := 1

		mock.ExpectQuery(product_batch.GetQuery).
			WithArgs(batchNumber).WillReturnError(sql.ErrConnDone)

		repository := product_batch.NewRepository(db)

		assert.Panics(t, func() {
			repository.Get(batchNumber)
		})

	})
}

func TestRepositorySave(t *testing.T) {
	t.Run("Should insert the product batches", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"batch_number", "current_quantity", "current_temperature", "due_date", "initial_quantity", "manufacturing_date", "manufacturing_hour", "minimum_temperature", "product_id", "section_id"}
		rows := sqlmock.NewRows(columns)
		rows.AddRow(pb.BatchNumber, pb.CurrentQuantity, pb.CurrentTemperature, pb.DueDate, pb.InitialQuantity, pb.ManufacturingDate, pb.ManufacturingHour, pb.MinimumTemperature, pb.ProductID, pb.SectionID)

		LastInsertId := 1
		mockedProductBatches := pb
		mock.ExpectPrepare(regexp.QuoteMeta(product_batch.InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(product_batch.InsertQuery)).
			WithArgs(mockedProductBatches.BatchNumber, mockedProductBatches.CurrentQuantity, mockedProductBatches.CurrentTemperature, mockedProductBatches.DueDate, mockedProductBatches.InitialQuantity, mockedProductBatches.ManufacturingDate, mockedProductBatches.ManufacturingHour, mockedProductBatches.MinimumTemperature, mockedProductBatches.ProductID, mockedProductBatches.SectionID).
			WillReturnResult(sqlmock.NewResult(int64(LastInsertId), 1))

		repository := product_batch.NewRepository(db)
		result := repository.Save(mockedProductBatches)
		assert.Equal(t, LastInsertId, result)
	})
	t.Run("Should throw panic when expected prepare fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedProductBatches := pb
		mock.ExpectPrepare(regexp.QuoteMeta(product_batch.InsertQuery)).WillReturnError(sql.ErrNoRows)

		repository := product_batch.NewRepository(db)
		assert.Panics(t, func() {
			repository.Save(mockedProductBatches)
		})
	})
	t.Run("Should throw panic when expected exec fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedProductBatches := pb
		mock.ExpectPrepare(regexp.QuoteMeta(product_batch.InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(product_batch.InsertQuery)).WillReturnError(sql.ErrNoRows).
			WithArgs(mockedProductBatches.BatchNumber, mockedProductBatches.CurrentQuantity, mockedProductBatches.CurrentTemperature, mockedProductBatches.DueDate, mockedProductBatches.InitialQuantity, mockedProductBatches.ManufacturingDate, mockedProductBatches.ManufacturingHour, mockedProductBatches.MinimumTemperature, mockedProductBatches.ProductID, mockedProductBatches.SectionID).
			WillReturnError(sql.ErrConnDone)

		repository := product_batch.NewRepository(db)
		assert.Panics(t, func() {
			repository.Save(mockedProductBatches)
		})
	})
	t.Run("Should throw panic when sql has error", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedProductBatches := pb
		mock.ExpectPrepare(regexp.QuoteMeta(product_batch.InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(product_batch.InsertQuery)).
			WithArgs(mockedProductBatches.BatchNumber, mockedProductBatches.CurrentQuantity, mockedProductBatches.CurrentTemperature, mockedProductBatches.DueDate, mockedProductBatches.InitialQuantity, mockedProductBatches.ManufacturingDate, mockedProductBatches.ManufacturingHour, mockedProductBatches.MinimumTemperature, mockedProductBatches.ProductID, mockedProductBatches.SectionID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrConnDone))

		repository := product_batch.NewRepository(db)
		assert.Panics(t, func() {
			repository.Save(mockedProductBatches)
		})
	})
}

func SetupMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	return db, mock
}
