package product_record_test

import (
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	record "github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product_record"
	"github.com/stretchr/testify/assert"
)

var (
	mockedProductRecordTemplate = domain.ProductRecord{
		ID:             1,
		ProductID:      1,
		LastUpdateDate: time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
		PurchasePrice:  1,
		SalePrice:      1,
	}
)

func TestRepositoryGet(t *testing.T) {
	t.Run("Should return a record by specified id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id", "last_update_date", "purchase_price", "sale_price", "product_id"}
		rows := sqlmock.NewRows(columns)
		recordId := 1
		rows.AddRow(recordId, "2023-01-01 00:00:00", 1, 1, 1)

		mock.ExpectQuery(regexp.QuoteMeta(record.GetQuery)).
			WithArgs(recordId).
			WillReturnRows(rows)

		repository := record.NewRepository(db)

		result := repository.Get(recordId)

		assert.NotNil(t, result)
	})

	t.Run("Should not return a record", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		recordId := 1

		mock.ExpectQuery(regexp.QuoteMeta(record.GetQuery)).
			WithArgs(recordId).
			WillReturnError(sql.ErrNoRows)

		repository := record.NewRepository(db)

		result := repository.Get(recordId)

		assert.Nil(t, result)
	})

	t.Run("Should throw panic", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		recordId := 1

		mock.ExpectQuery(regexp.QuoteMeta(record.GetQuery)).
			WithArgs(recordId).
			WillReturnError(sql.ErrConnDone)

		repository := record.NewRepository(db)

		assert.Panics(t, func() { repository.Get(recordId) })
	})
}

func SetupMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	return db, mock
}

func TestRepositoryExists(t *testing.T) {
	t.Run("Should return true", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"product_id", "last_update_date"}
		rows := sqlmock.NewRows(columns)
		productId := 1
		lastUpdateDate := time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC)
		rows.AddRow(productId, "2023-01-01 00:00:00")

		mock.ExpectQuery(regexp.QuoteMeta(record.ExistsQuery)).
			WithArgs(productId, lastUpdateDate).
			WillReturnRows(rows)

		repository := record.NewRepository(db)

		result := repository.Exists(productId, lastUpdateDate)

		assert.True(t, result)
	})

	t.Run("Should return false when there are no query results", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		productId := 1
		lastUpdateDate := time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC)

		mock.ExpectQuery(record.ExistsQuery).WithArgs(productId, lastUpdateDate).WillReturnError(sql.ErrNoRows)

		repository := record.NewRepository(db)

		result := repository.Exists(productId, lastUpdateDate)

		assert.False(t, result)
	})

	t.Run("Should return false when database has internal error", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		productId := 1
		lastUpdateDate := time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC)

		mock.ExpectQuery(record.ExistsQuery).WithArgs(productId, lastUpdateDate).WillReturnError(sql.ErrConnDone)

		repository := record.NewRepository(db)

		result := repository.Exists(productId, lastUpdateDate)

		assert.False(t, result)
	})
}

func TestRepositorySave(t *testing.T) {
	t.Run("Should insert the record and return the record id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		lastInsertId := 1
		mockedProductRecord := mockedProductRecordTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(record.InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(record.InsertQuery)).
			WithArgs(mockedProductRecord.LastUpdateDate, mockedProductRecord.PurchasePrice, mockedProductRecord.SalePrice, mockedProductRecord.ProductID).
			WillReturnResult(sqlmock.NewResult(int64(lastInsertId), 1))

		repository := record.NewRepository(db)

		result := repository.Save(mockedProductRecord)

		assert.Equal(t, lastInsertId, result)
	})

	t.Run("Should throw panic when expected prepare fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedProductRecord := mockedProductRecordTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(record.InsertQuery)).WillReturnError(sql.ErrConnDone)

		repository := record.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedProductRecord) })
	})

	t.Run("Should throw panic when expected exec fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedProductRecord := mockedProductRecordTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(record.InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(record.InsertQuery)).
			WithArgs(mockedProductRecord.LastUpdateDate, mockedProductRecord.PurchasePrice, mockedProductRecord.SalePrice, mockedProductRecord.ProductID).
			WillReturnError(sql.ErrConnDone)

		repository := record.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedProductRecord) })
	})

	t.Run("Should throw panic when sql has error", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedProductRecord := mockedProductRecordTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(record.InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(record.InsertQuery)).
			WithArgs(mockedProductRecord.LastUpdateDate, mockedProductRecord.PurchasePrice, mockedProductRecord.SalePrice, mockedProductRecord.ProductID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrConnDone))

		repository := record.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedProductRecord) })
	})
}

func TestRepositoryCountCountRecordsByAllProducts(t *testing.T) {
	t.Run("Should return records count report by all products", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"product_id", "description", "records_count"}
		rows := sqlmock.NewRows(columns)
		rows.AddRow(1, "", 1)

		mock.ExpectQuery(regexp.QuoteMeta(record.CountRecordsByAllProductsQuery)).WillReturnRows(rows)

		repository := record.NewRepository(db)

		result := repository.CountRecordsByAllProducts()

		assert.Equal(t, len(result), 1)
	})

	t.Run("Should throw panic when query execution fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(record.CountRecordsByAllProductsQuery)).WillReturnError(sql.ErrConnDone)

		repository := record.NewRepository(db)

		assert.Panics(t, func() { repository.CountRecordsByAllProducts() })
	})
}

func TestRepositoryCountCountRecordsByProduct(t *testing.T) {
	t.Run("Should return records count report by specified product id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"product_id", "description", "records_count"}
		rows := sqlmock.NewRows(columns)
		productId := 1
		rows.AddRow(1, "", 1)

		mock.ExpectQuery(regexp.QuoteMeta(record.CountRecordsByProductQuery)).
			WithArgs(productId).
			WillReturnRows(rows)

		repository := record.NewRepository(db)

		result := repository.CountRecordsByProduct(productId)

		assert.NotNil(t, result)
	})

	t.Run("Should throw panic when query execution fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		recordId := 1
		mock.ExpectQuery(regexp.QuoteMeta(record.CountRecordsByProductQuery)).WillReturnError(sql.ErrNoRows)

		repository := record.NewRepository(db)

		result := repository.CountRecordsByProduct(recordId)

		assert.Nil(t, result)
	})

	t.Run("Should throw panic when query execution fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		recordId := 1
		mock.ExpectQuery(regexp.QuoteMeta(record.CountRecordsByProductQuery)).WillReturnError(sql.ErrConnDone)

		repository := record.NewRepository(db)

		assert.Panics(t, func() { repository.CountRecordsByProduct(recordId) })
	})
}
