package purchase_order_test

import (
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/purchase_order"
	"github.com/stretchr/testify/assert"
)

var (
	mockedPurchaseOrderTemplate = domain.PurchaseOrder{
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

func TestRepositoryGet(t *testing.T) {
	t.Run("Should return a purchase order by specified id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id", "order_number", "order_date", "tracking_code", "buyer_id", "carrier_id", "product_record_id", "order_status_id", "warehouse_id"}
		rows := sqlmock.NewRows(columns)
		purchaseOrderID := 1
		rows.AddRow(purchaseOrderID, "order123", "2023-01-01 00:00:00", "tr123", 1, 1, 1, 1, 1)

		mock.ExpectQuery(regexp.QuoteMeta(purchase_order.GetQuery)).
			WithArgs(purchaseOrderID).
			WillReturnRows(rows)

		repository := purchase_order.NewRepository(db)

		result := repository.Get(purchaseOrderID)

		assert.NotNil(t, result)
	})

	t.Run("Should not return a purchase order", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		purchaseOrderID := 1

		mock.ExpectQuery(regexp.QuoteMeta(purchase_order.GetQuery)).
			WithArgs(purchaseOrderID).
			WillReturnError(sql.ErrNoRows)

		repository := purchase_order.NewRepository(db)

		result := repository.Get(purchaseOrderID)

		assert.Nil(t, result)
	})

	t.Run("Should throw a panic", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		purchaseOrderID := 1

		mock.ExpectQuery(regexp.QuoteMeta(purchase_order.GetQuery)).
			WithArgs(purchaseOrderID).
			WillReturnError(sql.ErrConnDone)

		repository := purchase_order.NewRepository(db)

		assert.Panics(t, func() { repository.Get(purchaseOrderID) })
	})
}

func TestRepositoryExists(t *testing.T) {
	t.Run("Should return true", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"order_number"}
		rows := sqlmock.NewRows(columns)
		orderNumber := "order123"
		rows.AddRow(orderNumber)

		mock.ExpectQuery(regexp.QuoteMeta(purchase_order.ExistsQuery)).
			WithArgs(orderNumber).
			WillReturnRows(rows)

		repository := purchase_order.NewRepository(db)

		result := repository.Exists(orderNumber)

		assert.True(t, result)
	})

	t.Run("Should return false when there are no query results", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		orderNumber := "order123"

		mock.ExpectQuery(purchase_order.ExistsQuery).WithArgs(orderNumber).WillReturnError(sql.ErrNoRows)

		repository := purchase_order.NewRepository(db)

		result := repository.Exists(orderNumber)

		assert.False(t, result)
	})

	t.Run("Should return false when database has internal error", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		orderNumber := "order123"

		mock.ExpectQuery(purchase_order.ExistsQuery).WithArgs(orderNumber).WillReturnError(sql.ErrConnDone)

		repository := purchase_order.NewRepository(db)

		result := repository.Exists(orderNumber)

		assert.False(t, result)
	})
}

func TestRepositorySave(t *testing.T) {
	t.Run("Should insert the purchase order and return the purchase order id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"order_number"}
		rows := sqlmock.NewRows(columns)
		orderNumber := "order123"
		rows.AddRow(orderNumber)

		lastInsertId := 1
		mockedPurchaseOrder := mockedPurchaseOrderTemplate

		mock.ExpectPrepare(regexp.QuoteMeta(purchase_order.InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(purchase_order.InsertQuery)).
			WithArgs(mockedPurchaseOrder.OrderNumber, mockedPurchaseOrder.OrderDate, mockedPurchaseOrder.TrackingCode, mockedPurchaseOrder.BuyerID, mockedPurchaseOrder.CarrierID, mockedPurchaseOrder.ProductRecordID, mockedPurchaseOrder.OrderStatusID, mockedPurchaseOrder.WarehouseID).
			WillReturnResult(sqlmock.NewResult(int64(lastInsertId), 1))

		repository := purchase_order.NewRepository(db)

		result := repository.Save(mockedPurchaseOrder)

		assert.Equal(t, lastInsertId, result)
	})

	t.Run("Should throw panic when expected prepare fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedPurchaseOrder := mockedPurchaseOrderTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(purchase_order.InsertQuery)).WillReturnError(sql.ErrConnDone)

		repository := purchase_order.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedPurchaseOrder) })
	})

	t.Run("Should throw panic when expected exec fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedPurchaseOrder := mockedPurchaseOrderTemplate

		mock.ExpectPrepare(regexp.QuoteMeta(purchase_order.InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(purchase_order.InsertQuery)).
			WithArgs(mockedPurchaseOrder.OrderNumber, mockedPurchaseOrder.OrderDate, mockedPurchaseOrder.TrackingCode, mockedPurchaseOrder.BuyerID, mockedPurchaseOrder.CarrierID, mockedPurchaseOrder.ProductRecordID, mockedPurchaseOrder.OrderStatusID, mockedPurchaseOrder.WarehouseID).
			WillReturnError(sql.ErrConnDone)

		repository := purchase_order.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedPurchaseOrder) })
	})

	t.Run("Should throw panic when sql has error", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedPurchaseOrder := mockedPurchaseOrderTemplate

		mock.ExpectPrepare(regexp.QuoteMeta(purchase_order.InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(purchase_order.InsertQuery)).
			WithArgs(mockedPurchaseOrder.OrderNumber, mockedPurchaseOrder.OrderDate, mockedPurchaseOrder.TrackingCode, mockedPurchaseOrder.BuyerID, mockedPurchaseOrder.CarrierID, mockedPurchaseOrder.ProductRecordID, mockedPurchaseOrder.OrderStatusID, mockedPurchaseOrder.WarehouseID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrConnDone))

		repository := purchase_order.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedPurchaseOrder) })
	})
}

func SetupMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	return db, mock
}
