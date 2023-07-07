package province_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/province"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryGet(t *testing.T) {
	t.Run("Should return a province by specified id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id"}
		rows := sqlmock.NewRows(columns)
		provinceId := 1
		rows.AddRow(provinceId)

		mock.ExpectQuery(province.GetQuery).WithArgs(provinceId).WillReturnRows(rows)

		repository := province.NewRepository(db)

		result := repository.Get(provinceId)

		assert.NotNil(t, result)
	})

	t.Run("Should not return a province", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		provinceId := 1

		mock.ExpectQuery(province.GetQuery).WithArgs(provinceId).WillReturnError(sql.ErrNoRows)

		repository := province.NewRepository(db)

		result := repository.Get(provinceId)

		assert.Nil(t, result)
	})

	t.Run("Should throw panic", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		provinceId := 1

		mock.ExpectQuery(province.GetQuery).WithArgs(provinceId).WillReturnError(sql.ErrConnDone)

		repository := province.NewRepository(db)

		assert.Panics(t, func() { repository.Get(provinceId) })
	})
}

func SetupMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	return db, mock
}
