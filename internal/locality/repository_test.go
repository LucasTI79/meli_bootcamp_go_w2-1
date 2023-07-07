package locality_test

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/locality"
	"github.com/stretchr/testify/assert"
)

var (
	mockedLocalityTemplate = domain.Locality{
		ID:           1,
		LocalityName: "Locality",
		ProvinceID:   1,
	}
)

func TestRepositoryGet(t *testing.T) {
	t.Run("Should return a locality by specified id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id", "locality_name", "province_id"}
		rows := sqlmock.NewRows(columns)
		localityId := 1
		rows.AddRow(localityId, "", 1)

		mock.ExpectQuery(regexp.QuoteMeta(locality.GetQuery)).
			WithArgs(localityId).
			WillReturnRows(rows)

		repository := locality.NewRepository(db)

		result := repository.Get(localityId)

		assert.NotNil(t, result)
	})

	t.Run("Should not return a locality", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		localityId := 1

		mock.ExpectQuery(regexp.QuoteMeta(locality.GetQuery)).
			WithArgs(localityId).
			WillReturnError(sql.ErrNoRows)

		repository := locality.NewRepository(db)

		result := repository.Get(localityId)

		assert.Nil(t, result)
	})

	t.Run("Should throw panic", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		localityId := 1

		mock.ExpectQuery(regexp.QuoteMeta(locality.GetQuery)).
			WithArgs(localityId).
			WillReturnError(sql.ErrConnDone)

		repository := locality.NewRepository(db)

		assert.Panics(t, func() { repository.Get(localityId) })
	})
}

func SetupMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	return db, mock
}

func TestRepositoryExists(t *testing.T) {
	t.Run("Should return true", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"locality_name"}
		rows := sqlmock.NewRows(columns)
		localityName := "Locality"
		rows.AddRow(localityName)

		mock.ExpectQuery(regexp.QuoteMeta(locality.ExistsQuery)).
			WithArgs(localityName).
			WillReturnRows(rows)

		repository := locality.NewRepository(db)

		result := repository.Exists(localityName)

		assert.True(t, result)
	})

	t.Run("Should return false when there are no query results", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		localityName := "Locality"

		mock.ExpectQuery(locality.ExistsQuery).WithArgs(localityName).WillReturnError(sql.ErrNoRows)

		repository := locality.NewRepository(db)

		result := repository.Exists(localityName)

		assert.False(t, result)
	})

	t.Run("Should return false when database has internal error", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		localityName := "Locality"

		mock.ExpectQuery(locality.ExistsQuery).WithArgs(localityName).WillReturnError(sql.ErrConnDone)

		repository := locality.NewRepository(db)

		result := repository.Exists(localityName)

		assert.False(t, result)
	})
}

func TestRepositorySave(t *testing.T) {
	t.Run("Should insert the locality and return the locality id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"locality_name"}
		rows := sqlmock.NewRows(columns)
		localityName := "Locality"
		rows.AddRow(localityName)

		lastInsertId := 1
		mockedLocality := mockedLocalityTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(locality.InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(locality.InsertQuery)).
			WithArgs(mockedLocality.LocalityName, mockedLocality.ProvinceID).
			WillReturnResult(sqlmock.NewResult(int64(lastInsertId), 1))

		repository := locality.NewRepository(db)

		result := repository.Save(mockedLocality)

		assert.Equal(t, lastInsertId, result)
	})

	t.Run("Should throw panic when expected prepare fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedLocality := mockedLocalityTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(locality.InsertQuery)).WillReturnError(sql.ErrConnDone)

		repository := locality.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedLocality) })
	})

	t.Run("Should throw panic when expected exec fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedLocality := mockedLocalityTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(locality.InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(locality.InsertQuery)).
			WithArgs(mockedLocality.LocalityName, mockedLocality.ProvinceID).
			WillReturnError(sql.ErrConnDone)

		repository := locality.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedLocality) })
	})

	t.Run("Should throw panic when sql has error", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedLocality := mockedLocalityTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(locality.InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(locality.InsertQuery)).
			WithArgs(mockedLocality.LocalityName, mockedLocality.ProvinceID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrConnDone))

		repository := locality.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedLocality) })
	})
}

func TestRepositoryCountSellersByAllLocalities(t *testing.T) {
	t.Run("Should return sellers count report by all localities", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"locality_id", "locality_name", "sellers_count"}
		rows := sqlmock.NewRows(columns)
		localityId := 1
		rows.AddRow(localityId, "", 1)

		mock.ExpectQuery(regexp.QuoteMeta(locality.CountSellersByAllLocalitiesQuery)).WillReturnRows(rows)

		repository := locality.NewRepository(db)

		result := repository.CountSellersByAllLocalities()

		assert.Equal(t, len(result), 1)
	})

	t.Run("Should throw panic when query execution fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(locality.CountSellersByAllLocalitiesQuery)).WillReturnError(sql.ErrConnDone)

		repository := locality.NewRepository(db)

		assert.Panics(t, func() { repository.CountSellersByAllLocalities() })
	})
}

func TestRepositoryCountSellersByLocality(t *testing.T) {
	t.Run("Should return sellers count report by specified locality id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"locality_id", "locality_name", "sellers_count"}
		rows := sqlmock.NewRows(columns)
		localityId := 1
		rows.AddRow(localityId, "", 1)

		mock.ExpectQuery(regexp.QuoteMeta(locality.CountSellersByLocalityQuery)).
			WithArgs(localityId).
			WillReturnRows(rows)

		repository := locality.NewRepository(db)

		result := repository.CountSellersByLocality(localityId)

		assert.NotNil(t, result)
	})

	t.Run("Should throw panic when query execution fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		localityId := 1
		mock.ExpectQuery(regexp.QuoteMeta(locality.CountSellersByLocalityQuery)).WillReturnError(sql.ErrNoRows)

		repository := locality.NewRepository(db)

		result := repository.CountSellersByLocality(localityId)

		assert.Nil(t, result)
	})

	t.Run("Should throw panic when query execution fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		localityId := 1
		mock.ExpectQuery(regexp.QuoteMeta(locality.CountSellersByLocalityQuery)).WillReturnError(sql.ErrConnDone)

		repository := locality.NewRepository(db)

		assert.Panics(t, func() { repository.CountSellersByLocality(localityId) })
	})
}
