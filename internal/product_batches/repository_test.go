package product_batches_test

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product_batches"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryCreate(t *testing.T) {
	t.Run("should create a product batch", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id"}
		rows := sqlmock.NewRows(columns)
		productBacth := 1
		rows.AddRow(productBacth, "", 1)
		mock.ExpectQuery(regexp.QuoteMeta(product_batches.CreateQuery)).
			WithArgs(productBacth).
			WillReturnRows(rows)

		repository := product_batches.NewRepository(db)
		result := repository.Create(productBacth)

		assert.Equal(t, productBacth, result.ID)
	})

	t.Run("should return an error when the product batch does not exist", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id"}
		rows := sqlmock.NewRows(columns)
		productBacth := 1
		rows.AddRow(productBacth, "", 1)
		mock.ExpectQuery(regexp.QuoteMeta(product_batches.CreateQuery)).
			WithArgs(productBacth).
			WillReturnRows(rows)

		repository := product_batches.NewRepository(db)
		result := repository.Create(productBacth)

		assert.Equal(t, productBacth, result.ID)
	})

}

func SetupMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	return db, mock
}

func TestRepositoryGet(t *testing.T) {
	t.Run("should return a product batch", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id", "name", "current_quantity"}
		rows := sqlmock.NewRows(columns)
		productBacth := 1
		rows.AddRow(productBacth, "", 1)
		mock.ExpectQuery(regexp.QuoteMeta(product_batches.GetQuery)).
			WithArgs(productBacth).
			WillReturnRows(rows)

		repository := product_batches.NewRepository(db)
		result := repository.Get(productBacth)

		assert.Equal(t, productBacth, result.ID)
	})

	t.Run("should return an error when the product batch does not exist", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id"}
		rows := sqlmock.NewRows(columns)
		productBacth := 1
		rows.AddRow(productBacth, "", 1)
		mock.ExpectQuery(regexp.QuoteMeta(product_batches.GetQuery)).
			WithArgs(productBacth).
			WillReturnRows(rows)

		repository := product_batches.NewRepository(db)
		result := repository.Get(productBacth)

		assert.Equal(t, productBacth, result.ID)
	})

}

func TestRepositorySave(t *testing.T) {
	t.Run("should save a product batch", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id"}
		rows := sqlmock.NewRows(columns)
		productBacth := 1
		rows.AddRow(productBacth, "", 1)
		mock.ExpectQuery(regexp.QuoteMeta(product_batches.GetQuery)).
			WithArgs(productBacth).
			WillReturnRows(rows)

		repository := product_batches.NewRepository(db)
		result := repository.Save(domain.ProductBatches)

		assert.Equal(t, productBacth, result)
	})

	t.Run("should return an error when the product batch does not exist", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id"}
		rows := sqlmock.NewRows(columns)
		productBacth := 1
		rows.AddRow(productBacth, "", 1)
		mock.ExpectQuery(regexp.QuoteMeta(product_batches.GetQuery)).
			WithArgs(productBacth).
			WillReturnRows(rows)

		repository := product_batches.NewRepository(db)
		result := repository.Save(domain.ProductBatches)

		assert.Equal(t, productBacth, result)
	})

}

func TestRepositoryExists(t *testing.T) {
	t.Run("should return true when the product batch exists", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id"}
		rows := sqlmock.NewRows(columns)
		productBacth := 1
		rows.AddRow(productBacth, "", 1)
		mock.ExpectQuery(regexp.QuoteMeta(product_batches.GetQuery)).
			WithArgs(productBacth).
			WillReturnRows(rows)

		repository := product_batches.NewRepository(db)
		result := repository.Exists(productBacth)

		assert.Equal(t, true, result)
	})

	t.Run("should return false when the product batch does not exist", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id"}
		rows := sqlmock.NewRows(columns)
		productBacth := 1
		rows.AddRow(productBacth, "", 1)
		mock.ExpectQuery(regexp.QuoteMeta(product_batches.GetQuery)).
			WithArgs(productBacth).
			WillReturnRows(rows)

		repository := product_batches.NewRepository(db)
		result := repository.Exists(productBacth)

		assert.Equal(t, false, result)
	})
}
