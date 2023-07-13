package buyer_test

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/assert"
)

var (
	mockedBuyerTemplate = domain.Buyer{
		ID:           1,
		CardNumberID: "CARD1234",
		FirstName:    "Nome",
		LastName:     "Sobrenome",
	}
)

func TestRepositoryGetAll(t *testing.T) {
	t.Run("Should return all buyers", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id", "card_number_id", "first_name", "last_name"}
		rows := sqlmock.NewRows(columns)
		buyerID := 1
		rows.AddRow(buyerID, "", "", "")

		mock.ExpectQuery(regexp.QuoteMeta(buyer.GetAllQuery)).WillReturnRows(rows)

		repository := buyer.NewRepository(db)

		result := repository.GetAll()

		assert.NotNil(t, result)
	})
	t.Run("Should throw panic when query execution fail", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(buyer.GetAllQuery)).WillReturnError(sql.ErrConnDone)

		repository := buyer.NewRepository(db)

		assert.Panics(t, func() { repository.GetAll() })
	})
}

func TestRepositoryGet(t *testing.T) {
	t.Run("Should return a buyer by specified ID", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id", "card_number_id", "first_name", "last_name"}
		rows := sqlmock.NewRows(columns)
		buyerID := 1
		rows.AddRow(buyerID, "", "", "")

		mock.ExpectQuery(regexp.QuoteMeta(buyer.GetQuery)).WithArgs(buyerID).WillReturnRows(rows)

		repository := buyer.NewRepository(db)

		result := repository.Get(buyerID)

		assert.NotNil(t, result)
	})

	t.Run("Should not return a buyer", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		buyerID := 1

		mock.ExpectQuery(regexp.QuoteMeta(buyer.GetQuery)).WithArgs(buyerID).WillReturnError(sql.ErrNoRows)

		repository := buyer.NewRepository(db)

		result := repository.Get(buyerID)

		assert.Nil(t, result)
	})

	t.Run("Should throw panic", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		buyerID := 1

		mock.ExpectQuery(regexp.QuoteMeta(buyer.GetQuery)).WithArgs(buyerID).WillReturnError(sql.ErrConnDone)

		repository := buyer.NewRepository(db)

		assert.Panics(t, func() { repository.Get(buyerID) })
	})
}

func TestRepositoryExists(t *testing.T) {
	t.Run("Should return true", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"card_number_id"}
		rows := sqlmock.NewRows(columns)
		cardNumberID := "CARD1234"
		rows.AddRow(cardNumberID)

		mock.ExpectQuery(regexp.QuoteMeta(buyer.ExistsQuery)).
			WithArgs(cardNumberID).
			WillReturnRows(rows)

		repository := buyer.NewRepository(db)

		result := repository.Exists(cardNumberID)

		assert.True(t, result)
	})

	t.Run("Should return false when there are no query results", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		cardNumberID := "CARD1234"

		mock.ExpectQuery(buyer.ExistsQuery).WithArgs(cardNumberID).WillReturnError(sql.ErrNoRows)

		repository := buyer.NewRepository(db)

		result := repository.Exists(cardNumberID)

		assert.False(t, result)
	})

	t.Run("Should return false when database has internal error", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		cardNumberID := "CARD1234"

		mock.ExpectQuery(buyer.ExistsQuery).WithArgs(cardNumberID).WillReturnError(sql.ErrConnDone)

		repository := buyer.NewRepository(db)

		result := repository.Exists(cardNumberID)

		assert.False(t, result)
	})
}

func TestRepositorySave(t *testing.T) {
	t.Run("Should create the buyer and return buyer`s id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		lastInsertId := 1
		mockedBuyer:= mockedBuyerTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(buyer.InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(buyer.InsertQuery)).
			WithArgs(
				mockedBuyer.CardNumberID,
				mockedBuyer.FirstName,
				mockedBuyer.LastName,
			).WillReturnResult(sqlmock.NewResult(int64(lastInsertId), 1))

		repository := buyer.NewRepository(db)

		result := repository.Save(mockedBuyer)

		assert.Equal(t, lastInsertId, result)
	})

	t.Run("Should throw a panic when ExpectPrepare fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedBuyer := mockedBuyerTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(buyer.InsertQuery)).WillReturnError(sql.ErrConnDone)

		repository := buyer.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedBuyer) })
	})

	t.Run("Should throw a panic when ExpectExec fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedBuyer := mockedBuyerTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(buyer.InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(buyer.InsertQuery)).
			WithArgs(
				mockedBuyer.CardNumberID,
				mockedBuyer.FirstName,
				mockedBuyer.LastName,
			).WillReturnError(sql.ErrConnDone)

		repository := buyer.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedBuyer) })
	})

	t.Run("Should throw panic when sql has error", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedBuyer := mockedBuyerTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(buyer.InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(buyer.InsertQuery)).
			WithArgs(
				mockedBuyer.CardNumberID,
				mockedBuyer.FirstName,
				mockedBuyer.LastName,
			).WillReturnResult(sqlmock.NewErrorResult(sql.ErrConnDone))

		repository := buyer.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedBuyer) })
	})
}

func TestRepositoryUpdate(t *testing.T) {
	t.Run("Should update buyer by specific ID", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedBuyer := mockedBuyerTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(buyer.UpdateQuery))
		mock.ExpectExec(regexp.QuoteMeta(buyer.UpdateQuery)).
			WithArgs(
				mockedBuyer.CardNumberID,
				mockedBuyer.FirstName,
				mockedBuyer.LastName,
				mockedBuyer.ID,
			).WillReturnResult(sqlmock.NewResult(int64(mockedBuyer.ID), 1))

		repository := buyer.NewRepository(db)

		assert.NotPanics(t, func() { repository.Update(mockedBuyer) })
	})

	t.Run("Should throw a panic when ExpectPrepare fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedBuyer := mockedBuyerTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(buyer.UpdateQuery)).WillReturnError(sql.ErrConnDone)

		repository := buyer.NewRepository(db)

		assert.Panics(t, func() { repository.Update(mockedBuyer) })
	})

	t.Run("Should throw a panic when ExpectExec fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedBuyer := mockedBuyerTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(buyer.UpdateQuery))
		mock.ExpectExec(regexp.QuoteMeta(buyer.UpdateQuery)).
			WithArgs(
				mockedBuyer.CardNumberID,
				mockedBuyer.FirstName,
				mockedBuyer.LastName,
			).WillReturnError(sql.ErrConnDone)

		repository := buyer.NewRepository(db)

		assert.Panics(t, func() { repository.Update(mockedBuyer) })
	})
}

func TestRepositoryDelete(t *testing.T) {
	t.Run("Should delete the buyer by specific ID", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedBuyer := mockedBuyerTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(buyer.DeleteQuery))
		mock.ExpectExec(regexp.QuoteMeta(buyer.DeleteQuery)).
			WithArgs(mockedBuyer.ID).
			WillReturnResult(sqlmock.NewResult(int64(mockedBuyer.ID), 1))

		repository := buyer.NewRepository(db)

		assert.NotPanics(t, func() { repository.Delete(mockedBuyer.ID) })
	})

	t.Run("Should throw a panic when ExpectPrepare fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedBuyer := mockedBuyerTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(buyer.DeleteQuery)).WillReturnError(sql.ErrConnDone)

		repository := buyer.NewRepository(db)

		assert.Panics(t, func() { repository.Delete(mockedBuyer.ID) })
	})

	t.Run("Should throw a panic when ExpecExec fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedBuyer := mockedBuyerTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(buyer.DeleteQuery))
		mock.ExpectExec(regexp.QuoteMeta(buyer.DeleteQuery)).
			WithArgs(mockedBuyer.ID).
			WillReturnError(sql.ErrConnDone)

		repository := buyer.NewRepository(db)

		assert.Panics(t, func() { repository.Delete(mockedBuyer.ID) })
	})
}

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
