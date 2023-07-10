package buyer_test

import (
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/assert"
)

var (
	mockedPurchaseOrderTemplate = domain.PurchaseOrders{
		ID:              1,
		OrderNumber:     "order#123",
		OrderDate:       time.Date(2023, 07, 10, 0, 0, 0, 0, time.UTC),
		TrackingCode:    "TRACK007",
		BuyerID:         1,
		CarrierID:       1,
		ProductRecordID: 1,
		OrderStatusID:   1,
		WarehouseID:     1,
	}
)

func TestRepositoryCountPurchasesByAllBuyers(t *testing.T) {
	t.Run("Should return purchases count report by all buyers", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id", "card_number_id", "first_name", "last_name", "purchase_orders_count"}
		rows := sqlmock.NewRows(columns)
		buyerID := 1
		rows.AddRow(buyerID, "", "", "", 1)

		mock.ExpectQuery(regexp.QuoteMeta(buyer.CountPurchasesByAllBuyers)).WillReturnRows(rows)

		repository := buyer.NewRepository(db)

		result := repository.CountPurchasesByAllBuyers()

		assert.Equal(t, len(result), 1)
	})

	t.Run("Should throw panic when query execution fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(buyer.CountPurchasesByAllBuyers)).WillReturnError(sql.ErrConnDone)

		repository := buyer.NewRepository(db)

		assert.Panics(t, func() { repository.CountPurchasesByAllBuyers() })
	})
}

func TestRepositoryCountPurchasesByBuyer(t *testing.T) {
	t.Run("Should return purchases count report by specified buyer id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id", "card_number_id", "first_name", "last_name", "purchase_orders_count"}
		rows := sqlmock.NewRows(columns)
		buyerID := 1
		rows.AddRow(buyerID, "", "", "", 1)

		mock.ExpectQuery(regexp.QuoteMeta(buyer.CountPurchasesByBuyer)).
			WithArgs(buyerID).
			WillReturnRows(rows)

		repository := buyer.NewRepository(db)

		result := repository.CountPurchasesByBuyer(buyerID)

		assert.NotNil(t, result)
	})

	t.Run("Should return nil when query ID does not exist", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		buyerID := 1
		mock.ExpectQuery(regexp.QuoteMeta(buyer.CountPurchasesByBuyer)).WillReturnError(sql.ErrNoRows)

		repository := buyer.NewRepository(db)

		result := repository.CountPurchasesByBuyer(buyerID)

		assert.Nil(t, result)
	})

	t.Run("Should throw panic when query execution fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		localityId := 1
		mock.ExpectQuery(regexp.QuoteMeta(buyer.CountPurchasesByBuyer)).WillReturnError(sql.ErrConnDone)

		repository := buyer.NewRepository(db)

		assert.Panics(t, func() { repository.CountPurchasesByBuyer(localityId) })
	})
}

func SetupMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	return db, mock
}
