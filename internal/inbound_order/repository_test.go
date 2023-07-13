package inbound_order_test

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/inbound_order"
	"github.com/stretchr/testify/assert"
)

var (
	allDataQuery               = regexp.QuoteMeta(inbound_order.GetQuery)
	allDataInsertQuery         = regexp.QuoteMeta(inbound_order.InsertQuery)
	mockedInboundOrderTemplate = domain.InboundOrder{
		ID:             1,
		OrderDate:      dateString,
		OrderNumber:    "asdf",
		EmployeeId:     1,
		ProductBatchId: 1,
		WarehouseId:    1,
	}
)

func TestRepositoryGet(t *testing.T) {
	t.Run("Should return an inbound order by specified id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id", "order_date", "order_number", "inboundOrder_id", "product_batch_id", "warehouse_id"}
		rows := sqlmock.NewRows(columns)
		inboundOrderId := 1
		rows.AddRow(inboundOrderId, date, "", 1, 1, 1)

		mock.ExpectQuery(allDataQuery).WithArgs(inboundOrderId).WillReturnRows(rows)

		repository := inbound_order.NewRepository(db)

		result := repository.Get(inboundOrderId)

		assert.NotNil(t, result)
	})

	t.Run("Should not return a inboundOrder", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		inboundOrderId := 1

		mock.ExpectQuery(allDataQuery).WithArgs(inboundOrderId).WillReturnError(sql.ErrNoRows)

		repository := inbound_order.NewRepository(db)

		result := repository.Get(inboundOrderId)

		assert.Nil(t, result)
	})

	t.Run("Should throw panic", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		inboundOrderId := 1

		mock.ExpectQuery(allDataQuery).WithArgs(inboundOrderId).WillReturnError(sql.ErrConnDone)

		repository := inbound_order.NewRepository(db)

		assert.Panics(t, func() { repository.Get(inboundOrderId) })
	})
}

func TestRepositoryExists(t *testing.T) {
	t.Run("Should return true", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"card_number_id"}
		rows := sqlmock.NewRows(columns)
		cardNumberId := "123"
		rows.AddRow(cardNumberId)

		mock.ExpectQuery(regexp.QuoteMeta(inbound_order.ExistsQuery)).
			WithArgs(cardNumberId).
			WillReturnRows(rows)

		repository := inbound_order.NewRepository(db)

		result := repository.Exists(cardNumberId)

		assert.True(t, result)
	})

	t.Run("Should return false when there are no query results", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		cardNumberId := "123"

		mock.ExpectQuery(inbound_order.ExistsQuery).WithArgs(cardNumberId).WillReturnError(sql.ErrNoRows)

		repository := inbound_order.NewRepository(db)

		result := repository.Exists(cardNumberId)

		assert.False(t, result)
	})

	t.Run("Should return false when database has internal error", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		cardNumberId := "123"

		mock.ExpectQuery(inbound_order.ExistsQuery).WithArgs(cardNumberId).WillReturnError(sql.ErrConnDone)

		repository := inbound_order.NewRepository(db)

		result := repository.Exists(cardNumberId)

		assert.False(t, result)
	})
}

func TestRepositorySave(t *testing.T) {
	t.Run("Should insert the inbound_order and return the inbound_order id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"order_date", "order_number", "inboundOrder_id", "product_batch_id", "warehouse_id"}
		rows := sqlmock.NewRows(columns)
		rows.AddRow("", "", 1, 1, 1)

		lastInsertId := 1
		mockedInboundOrder := mockedInboundOrderTemplate
		mock.ExpectPrepare(allDataInsertQuery)
		mock.ExpectExec(allDataInsertQuery).
			WithArgs(mockedInboundOrder.OrderDate, mockedInboundOrder.OrderNumber, mockedInboundOrder.EmployeeId, mockedInboundOrder.ProductBatchId, mockedInboundOrder.WarehouseId).
			WillReturnResult(sqlmock.NewResult(int64(lastInsertId), 1))

		repository := inbound_order.NewRepository(db)

		result := repository.Save(mockedInboundOrder)

		assert.Equal(t, lastInsertId, result)
	})

	t.Run("Should throw panic when expect prepare fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedInboundOrder := mockedInboundOrderTemplate
		mock.ExpectPrepare(allDataInsertQuery).WillReturnError(sql.ErrConnDone)

		repository := inbound_order.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedInboundOrder) })
	})

	t.Run("Should throw panic when expect exec fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedInboundOrder := mockedInboundOrderTemplate
		mock.ExpectPrepare(allDataInsertQuery)
		mock.ExpectExec(allDataInsertQuery).
			WithArgs(mockedInboundOrder.OrderDate, mockedInboundOrder.OrderNumber, mockedInboundOrder.EmployeeId, mockedInboundOrder.ProductBatchId, mockedInboundOrder.WarehouseId).
			WillReturnError(sql.ErrConnDone)

		repository := inbound_order.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedInboundOrder) })
	})

	t.Run("Should throw panic when sql has error", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedInboundOrder := mockedInboundOrderTemplate
		mock.ExpectPrepare(allDataInsertQuery)
		mock.ExpectExec(allDataInsertQuery).
			WithArgs(mockedInboundOrder.OrderDate, mockedInboundOrder.OrderNumber, mockedInboundOrder.EmployeeId, mockedInboundOrder.ProductBatchId, mockedInboundOrder.WarehouseId).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrConnDone))

		repository := inbound_order.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedInboundOrder) })
	})
}

func SetupMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	return db, mock
}
