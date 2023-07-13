package carrier_test

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/carrier"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

var (
	mockedCarrierTemplate = domain.Carrier{
		ID:          1,
		CID:         "CID",
		CompanyName: "company",
		Address:     "address",
		Telephone:   "+554312343212",
		LocalityID:  1,
	}
)

func TestRepositoryExists(t *testing.T) {
	t.Run("Should return true", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"cid"}
		rows := sqlmock.NewRows(columns)
		cid := "CID"
		rows.AddRow(cid)

		mock.ExpectQuery(regexp.QuoteMeta(carrier.ExistsQuery)).
			WithArgs(cid).
			WillReturnRows(rows)

		repository := carrier.NewRepository(db)

		result := repository.Exists(cid)

		assert.True(t, result)
	})

	t.Run("Should return false when there are no query results", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		cid := "CID"

		mock.ExpectQuery(carrier.ExistsQuery).WithArgs(cid).WillReturnError(sql.ErrNoRows)

		repository := carrier.NewRepository(db)

		result := repository.Exists(cid)

		assert.False(t, result)
	})

	t.Run("Should return false when database has internal error", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		cid := "CID"

		mock.ExpectQuery(carrier.ExistsQuery).WithArgs(cid).WillReturnError(sql.ErrConnDone)

		repository := carrier.NewRepository(db)

		result := repository.Exists(cid)

		assert.False(t, result)
	})
}

func TestRepositoryGet(t *testing.T) {
	t.Run("Should return a carrier by specified id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id", "cid", "company_name", "address", "telephone", "locality_id"}
		rows := sqlmock.NewRows(columns)
		carrierId := 1
		rows.AddRow(carrierId, "1", "company", "address", "+554312343212", 1)

		mock.ExpectQuery(regexp.QuoteMeta(carrier.GetQuery)).
			WithArgs(carrierId).
			WillReturnRows(rows)

		repository := carrier.NewRepository(db)

		result := repository.Get(carrierId)

		assert.NotNil(t, result)
	})

	t.Run("Should not return a carrier", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		carrierId := 1

		mock.ExpectQuery(regexp.QuoteMeta(carrier.GetQuery)).
			WithArgs(carrierId).
			WillReturnError(sql.ErrNoRows)

		repository := carrier.NewRepository(db)

		result := repository.Get(carrierId)

		assert.Nil(t, result)
	})

	t.Run("Should throw panic", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		carrierId := 1

		mock.ExpectQuery(regexp.QuoteMeta(carrier.GetQuery)).
			WithArgs(carrierId).
			WillReturnError(sql.ErrConnDone)

		repository := carrier.NewRepository(db)

		assert.Panics(t, func() { repository.Get(carrierId) })
	})
}

func TestRepositorySave(t *testing.T) {
	t.Run("Should insert the carrier and return the carrier id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"cid", "company_name", "address", "telephone", "locality_id"}
		rows := sqlmock.NewRows(columns)
		rows.AddRow(mockedCarrierTemplate.CID, mockedCarrierTemplate.CompanyName, mockedCarrierTemplate.Address, mockedCarrierTemplate.Telephone, mockedCarrierTemplate.LocalityID)

		lastInsertId := 1
		mockedCarrier := mockedCarrierTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(carrier.InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(carrier.InsertQuery)).
			WithArgs(mockedCarrier.CID, mockedCarrier.CompanyName, mockedCarrier.Address, mockedCarrier.Telephone, mockedCarrier.LocalityID).
			WillReturnResult(sqlmock.NewResult(int64(lastInsertId), 1))

		repository := carrier.NewRepository(db)

		result := repository.Save(mockedCarrier)

		assert.Equal(t, lastInsertId, result)
	})

	t.Run("Should throw panic when expected prepare fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedCarrier := mockedCarrierTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(carrier.InsertQuery)).WillReturnError(sql.ErrConnDone)

		repository := carrier.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedCarrier) })
	})

	t.Run("Should throw panic when expected exec fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedCarrier := mockedCarrierTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(carrier.InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(carrier.InsertQuery)).
			WithArgs(mockedCarrier.CID, mockedCarrier.CompanyName, mockedCarrier.Address, mockedCarrier.Telephone, mockedCarrier.LocalityID).
			WillReturnError(sql.ErrConnDone)

		repository := carrier.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedCarrier) })
	})

	t.Run("Should throw panic when sql has error", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedCarrier := mockedCarrierTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(carrier.InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(carrier.InsertQuery)).
			WithArgs(mockedCarrier.CID, mockedCarrier.CompanyName, mockedCarrier.Address, mockedCarrier.Telephone, mockedCarrier.LocalityID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrConnDone))

		repository := carrier.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedCarrier) })
	})
}

func SetupMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	return db, mock
}
