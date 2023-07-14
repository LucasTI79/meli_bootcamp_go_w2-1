package product_type_test

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product_type"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryGet(t *testing.T) {
	t.Run("Should return a product type by specified id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id"}
		rows := sqlmock.NewRows(columns)
		productId := 1
		rows.AddRow(productId)

		mock.ExpectQuery(regexp.QuoteMeta(product_type.GetQuery)).WithArgs(productId).WillReturnRows(rows)

		repository := product_type.NewRepository(db)

		result := repository.Get(productId)

		assert.NotNil(t, result)
	})

	t.Run("Should not return a product", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		productId := 1

		mock.ExpectQuery(regexp.QuoteMeta(product_type.GetQuery)).WithArgs(productId).WillReturnError(sql.ErrNoRows)

		repository := product_type.NewRepository(db)

		result := repository.Get(productId)

		assert.Nil(t, result)
	})

	t.Run("Should throw panic", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		productId := 1

		mock.ExpectQuery(regexp.QuoteMeta(product_type.GetQuery)).WithArgs(productId).WillReturnError(sql.ErrConnDone)

		repository := product_type.NewRepository(db)

		assert.Panics(t, func() { repository.Get(productId) })
	})
}

func SetupMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	return db, mock
}
