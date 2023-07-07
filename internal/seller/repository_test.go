package seller_test

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/seller"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryGetAll(t *testing.T) {
	t.Run("Should return all sellers", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id", "cid", "company_name", "address", "telephone", "locality_id"}
		rows := sqlmock.NewRows(columns)
		sellerId := 1
		rows.AddRow(sellerId, 1, "", "", "", 1)

		mock.ExpectQuery(seller.GetAllQuery).WillReturnRows(rows)

		repository := seller.NewRepository(db)

		result := repository.GetAll()

		assert.NotNil(t, result)
	})

	t.Run("Should throw panic when query execution fail", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mock.ExpectQuery(seller.GetAllQuery).WillReturnError(sql.ErrConnDone)

		repository := seller.NewRepository(db)

		assert.Panics(t, func() { repository.GetAll() })
	})
}

func TestRepositoryGet(t *testing.T) {
	t.Run("Should return a seller by specified id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id", "cid", "company_name", "address", "telephone", "locality_id"}
		rows := sqlmock.NewRows(columns)
		sellerId := 1
		rows.AddRow(sellerId, 1, "", "", "", 1)

		mock.ExpectQuery(seller.GetQuery).WithArgs(sellerId).WillReturnRows(rows)

		repository := seller.NewRepository(db)

		result := repository.Get(sellerId)

		assert.NotNil(t, result)
	})

	t.Run("Should not return a seller", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		sellerId := 1

		mock.ExpectQuery(seller.GetQuery).WithArgs(sellerId).WillReturnError(sql.ErrNoRows)

		repository := seller.NewRepository(db)

		result := repository.Get(sellerId)

		assert.Nil(t, result)
	})

	t.Run("Should throw panic", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		sellerId := 1

		mock.ExpectQuery(seller.GetQuery).WithArgs(sellerId).WillReturnError(sql.ErrConnDone)

		repository := seller.NewRepository(db)

		assert.Panics(t, func() { repository.Get(sellerId) })
	})
}

func TestRepositoryExists(t *testing.T) {
	t.Run("Should return true", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"cid"}
		rows := sqlmock.NewRows(columns)
		cid := 123
		rows.AddRow(cid)

		mock.ExpectQuery(regexp.QuoteMeta(seller.ExistsQuery)).
			WithArgs(cid).
			WillReturnRows(rows)

		repository := seller.NewRepository(db)

		result := repository.Exists(cid)

		assert.True(t, result)
	})

	t.Run("Should return false when there are no query results", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		cid := 123

		mock.ExpectQuery(seller.ExistsQuery).WithArgs(cid).WillReturnError(sql.ErrNoRows)

		repository := seller.NewRepository(db)

		result := repository.Exists(cid)

		assert.False(t, result)
	})

	t.Run("Should return false when database has internal error", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		cid := 123

		mock.ExpectQuery(seller.ExistsQuery).WithArgs(cid).WillReturnError(sql.ErrConnDone)

		repository := seller.NewRepository(db)

		result := repository.Exists(cid)

		assert.False(t, result)
	})
}

func TestRepositorySave(t *testing.T) {
	t.Run("Should insert the seller and return the seller id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"cid", "company_name", "address", "telephone", "locality_id"}
		rows := sqlmock.NewRows(columns)
		rows.AddRow(1, "", "", "", 1)

		lastInsertId := 1
		mockedSeller := mockedSellerTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(seller.InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(seller.InsertQuery)).
			WithArgs(mockedSeller.CID, mockedSeller.CompanyName, mockedSeller.Address, mockedSeller.Telephone, mockedSeller.LocalityID).
			WillReturnResult(sqlmock.NewResult(int64(lastInsertId), 1))

		repository := seller.NewRepository(db)

		result := repository.Save(mockedSeller)

		assert.Equal(t, lastInsertId, result)
	})

	t.Run("Should throw panic when expect prepare fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedSeller := mockedSellerTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(seller.InsertQuery)).WillReturnError(sql.ErrConnDone)

		repository := seller.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedSeller) })
	})

	t.Run("Should throw panic when expect exec fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedSeller := mockedSellerTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(seller.InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(seller.InsertQuery)).
			WithArgs(mockedSeller.CID, mockedSeller.CompanyName, mockedSeller.Address, mockedSeller.Telephone, mockedSeller.LocalityID).
			WillReturnError(sql.ErrConnDone)

		repository := seller.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedSeller) })
	})

	t.Run("Should throw panic when sql has error", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedSeller := mockedSellerTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(seller.InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(seller.InsertQuery)).
			WithArgs(mockedSeller.CID, mockedSeller.CompanyName, mockedSeller.Address, mockedSeller.Telephone, mockedSeller.LocalityID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrConnDone))

		repository := seller.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedSeller) })
	})
}

func TestRepositoryUpdate(t *testing.T) {
	t.Run("Should update the seller", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedSeller := mockedSellerTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(seller.UpdateQuery))
		mock.ExpectExec(regexp.QuoteMeta(seller.UpdateQuery)).
			WithArgs(mockedSeller.CID, mockedSeller.CompanyName, mockedSeller.Address, mockedSeller.Telephone, mockedSeller.LocalityID, mockedSeller.ID).
			WillReturnResult(sqlmock.NewResult(int64(mockedSeller.ID), 1))

		repository := seller.NewRepository(db)

		assert.NotPanics(t, func() { repository.Update(mockedSeller) })
	})

	t.Run("Should throw panic when expected prepare fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedSeller := mockedSellerTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(seller.UpdateQuery)).WillReturnError(sql.ErrConnDone)

		repository := seller.NewRepository(db)

		assert.Panics(t, func() { repository.Update(mockedSeller) })
	})

	t.Run("Should throw panic when expected exec fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedSeller := mockedSellerTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(seller.UpdateQuery))
		mock.ExpectExec(regexp.QuoteMeta(seller.UpdateQuery)).
			WithArgs(mockedSeller.CID, mockedSeller.CompanyName, mockedSeller.Address, mockedSeller.Telephone, mockedSeller.LocalityID, mockedSeller.ID).
			WillReturnError(sql.ErrConnDone)

		repository := seller.NewRepository(db)

		assert.Panics(t, func() { repository.Update(mockedSeller) })
	})
}

func TestRepositoryDelete(t *testing.T) {
	t.Run("Should delete the seller", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedSeller := mockedSellerTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(seller.DeleteQuery))
		mock.ExpectExec(regexp.QuoteMeta(seller.DeleteQuery)).
			WithArgs(mockedSeller.ID).
			WillReturnResult(sqlmock.NewResult(int64(mockedSeller.ID), 1))

		repository := seller.NewRepository(db)

		assert.NotPanics(t, func() { repository.Delete(mockedSeller.ID) })
	})

	t.Run("Should throw panic when expected prepare fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedSeller := mockedSellerTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(seller.DeleteQuery)).WillReturnError(sql.ErrConnDone)

		repository := seller.NewRepository(db)

		assert.Panics(t, func() { repository.Delete(mockedSeller.ID) })
	})

	t.Run("Should throw panic when expected exec fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedSeller := mockedSellerTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(seller.DeleteQuery))
		mock.ExpectExec(regexp.QuoteMeta(seller.DeleteQuery)).
			WithArgs(mockedSeller.ID).
			WillReturnError(sql.ErrConnDone)

		repository := seller.NewRepository(db)

		assert.Panics(t, func() { repository.Delete(mockedSeller.ID) })
	})
}

func SetupMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	return db, mock
}
