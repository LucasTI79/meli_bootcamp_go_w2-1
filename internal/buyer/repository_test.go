package buyer_test

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/buyer"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryCountPuchasesbyAllBuyers(t *testing.T) {
	t.Run("Should return purchases count report by all buyers", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id", "card_number_id", "first_name", "last_name", "purchase_orders_count"}
		rows := sqlmock.NewRows(columns)
		buyerID := 1
		rows.AddRow(buyerID, "", "", "", 1)

		mock.ExpectQuery(regexp.QuoteMeta(buyer.CountPuchasesbyAllBuyers)).WillReturnRows(rows)

		repository := buyer.NewRepository(db)

		result := repository.CountPuchasesbyAllBuyers()

		assert.Equal(t, len(result), 1)
	})

	t.Run("Should throw panic when query execution fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(buyer.CountPuchasesbyAllBuyers)).WillReturnError(sql.ErrConnDone)

		repository := buyer.NewRepository(db)

		assert.Panics(t, func() { repository.CountPuchasesbyAllBuyers() })
	})
}

func TestRepositoryCountPuchasesbyBuyer(t *testing.T) {
	t.Run("Should return purchases count report by specified buyer id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id", "card_number_id", "first_name", "last_name", "purchase_orders_count"}
		rows := sqlmock.NewRows(columns)
		buyerID := 1
		rows.AddRow(buyerID, "", "", "", 1)

		mock.ExpectQuery(regexp.QuoteMeta(buyer.CountPuchasesbyBuyer)).
			WithArgs(buyerID).
			WillReturnRows(rows)

		repository := buyer.NewRepository(db)

		result := repository.CountPuchasesbyBuyer(buyerID)

		assert.NotNil(t, result)
	})

	t.Run("Should return nil when query ID does not exist", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		buyerID := 1
		mock.ExpectQuery(regexp.QuoteMeta(buyer.CountPuchasesbyBuyer)).WillReturnError(sql.ErrNoRows)

		repository := buyer.NewRepository(db)

		result := repository.CountPuchasesbyBuyer(buyerID)

		assert.Nil(t, result)
	})

	t.Run("Should throw panic when query execution fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		localityId := 1
		mock.ExpectQuery(regexp.QuoteMeta(buyer.CountPuchasesbyBuyer)).WillReturnError(sql.ErrConnDone)

		repository := buyer.NewRepository(db)

		assert.Panics(t, func() { repository.CountPuchasesbyBuyer(localityId) })
	})
}

func SetupMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	return db, mock
}