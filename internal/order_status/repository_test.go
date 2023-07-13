package order_status_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/order_status"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryGet(t *testing.T) {
	t.Run("Should return an order status by specified id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id"}
		rows := sqlmock.NewRows(columns)
		orderStatusID := 1
		rows.AddRow(orderStatusID)

		mock.ExpectQuery(order_status.GetQuery).WithArgs(orderStatusID).WillReturnRows(rows)

		repository := order_status.NewRepository(db)

		result := repository.Get(orderStatusID)

		assert.NotNil(t, result)
	})

	t.Run("Should not return an order status", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		orderStatusID := 1

		mock.ExpectQuery(order_status.GetQuery).WithArgs(orderStatusID).WillReturnError(sql.ErrNoRows)

		repository := order_status.NewRepository(db)

		result := repository.Get(orderStatusID)

		assert.Nil(t, result)
	})

	t.Run("Should throw a panic", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		orderStatusID := 1

		mock.ExpectQuery(order_status.GetQuery).WithArgs(orderStatusID).WillReturnError(sql.ErrConnDone)

		repository := order_status.NewRepository(db)

		assert.Panics(t, func() { repository.Get(orderStatusID) })
	})
}

func SetupMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	return db, mock
}