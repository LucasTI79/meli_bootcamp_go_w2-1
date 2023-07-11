package section

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

var (
	mockedSectionTemplate = domain.Section{
		ID:                 1,
		SectionNumber:      1,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseID:        1,
		ProductTypeID:      1,
	}
)

func TestRepositoryGet(t *testing.T) {
	t.Run("Should return a section by specified id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id", "section_number", "current_temperature", "minimum_temperature", "current_capacity", "minimum_capacity", "maximum_capacity", "warehouse_id", "product_type_id"}
		rows := sqlmock.NewRows(columns)
		sectionId := 1

		rows.AddRow(sectionId, 1, 1, 1, 1, 1, 1, 1, 1)

		mock.ExpectQuery(regexp.QuoteMeta(GetQuery)).
			WithArgs(sectionId).
			WillReturnRows(rows)

		repository := NewRepository(db)

		result := repository.Get(sectionId)

		assert.NotNil(t, result)
	})

	t.Run("Should not return a section", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		sectionId := 1

		mock.ExpectQuery(regexp.QuoteMeta(GetQuery)).
			WithArgs(sectionId).
			WillReturnError(sql.ErrNoRows)

		repository := NewRepository(db)

		result := repository.Get(sectionId)

		assert.Nil(t, result)
	})

	t.Run("Should throw panic", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		localityId := 1

		mock.ExpectQuery(regexp.QuoteMeta(GetQuery)).
			WithArgs(localityId).
			WillReturnError(sql.ErrConnDone)

		repository := NewRepository(db)

		assert.Panics(t, func() { repository.Get(localityId) })
	})
}

func TestRepositoryExists(t *testing.T) {
	t.Run("Should return true", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"section_number"}
		rows := sqlmock.NewRows(columns)
		sectionNumber := 1
		rows.AddRow(sectionNumber)

		mock.ExpectQuery(regexp.QuoteMeta(ExistsQuery)).
			WithArgs(sectionNumber).
			WillReturnRows(rows)

		repository := NewRepository(db)

		result := repository.Exists(sectionNumber)

		assert.True(t, result)
	})

	t.Run("Should return false when there are no query results", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		sectionNumber := 1

		mock.ExpectQuery(ExistsQuery).WithArgs(sectionNumber).WillReturnError(sql.ErrNoRows)

		repository := NewRepository(db)

		result := repository.Exists(sectionNumber)

		assert.False(t, result)
	})

	t.Run("Should return false when database has internal error", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		sectionNumber := 1

		mock.ExpectQuery(ExistsQuery).WithArgs(sectionNumber).WillReturnError(sql.ErrConnDone)

		repository := NewRepository(db)

		result := repository.Exists(sectionNumber)

		assert.False(t, result)
	})
}

func TestRepositorySave(t *testing.T) {
	t.Run("Should insert the section and return the section id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"section_number"}
		rows := sqlmock.NewRows(columns)

		rows.AddRow(1)

		lastInsertId := 1

		mockedSection := mockedSectionTemplate

		mock.ExpectPrepare(regexp.QuoteMeta(InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(InsertQuery)).
			WithArgs(mockedSection.SectionNumber, mockedSection.ProductTypeID, mockedSection.CurrentCapacity, mockedSection.CurrentCapacity, mockedSection.MinimumCapacity, mockedSection.MaximumCapacity, mockedSection.WarehouseID, mockedSection.ProductTypeID).
			WillReturnResult(sqlmock.NewResult(int64(lastInsertId), 1))

		repository := NewRepository(db)

		result := repository.Save(mockedSection)

		assert.Equal(t, lastInsertId, result)
	})

	t.Run("Should throw panic when expected prepare fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedSection := mockedSectionTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(InsertQuery)).WillReturnError(sql.ErrConnDone)

		repository := NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedSection) })
	})

	t.Run("Should throw panic when expected exec fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedSection := mockedSectionTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(InsertQuery)).
			WithArgs(mockedSection.SectionNumber, mockedSection.ProductTypeID, mockedSection.CurrentCapacity, mockedSection.CurrentCapacity, mockedSection.MinimumCapacity, mockedSection.MaximumCapacity, mockedSection.WarehouseID, mockedSection.ProductTypeID).
			WillReturnError(sql.ErrConnDone)

		repository := NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedSection) })
	})

	t.Run("Should throw panic when sql has error", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedSection := mockedSectionTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(InsertQuery)).
			WithArgs(mockedSection.SectionNumber, mockedSection.ProductTypeID, mockedSection.CurrentCapacity, mockedSection.CurrentCapacity, mockedSection.MinimumCapacity, mockedSection.MaximumCapacity, mockedSection.WarehouseID, mockedSection.ProductTypeID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrConnDone))

		repository := NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedSection) })
	})
}

func TestRepositoryDelete(t *testing.T) {
	t.Run("Should delete the section", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedSection := mockedSectionTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(DeleteQuery))
		mock.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
			WithArgs(mockedSection.ID).
			WillReturnResult(sqlmock.NewResult(int64(mockedSection.ID), 1))

		repository := NewRepository(db)

		assert.NotPanics(t, func() { repository.Delete(mockedSection.ID) })
	})

	t.Run("Should throw panic when expected prepare fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedSection := mockedSectionTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(DeleteQuery)).WillReturnError(sql.ErrConnDone)

		repository := NewRepository(db)

		assert.Panics(t, func() { repository.Delete(mockedSection.ID) })
	})

	t.Run("Should throw panic when expected exec fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedSection := mockedSectionTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(DeleteQuery))
		mock.ExpectExec(regexp.QuoteMeta(DeleteQuery)).
			WithArgs(mockedSection.ID).
			WillReturnError(sql.ErrConnDone)

		repository := NewRepository(db)

		assert.Panics(t, func() { repository.Delete(mockedSection.ID) })
	})
}

func TestRepositoryUpdate(t *testing.T) {
	t.Run("Should update the section", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedSection := mockedSectionTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(UpdateQuery))
		mock.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
			WithArgs(
				mockedSection.ID,
				mockedSection.SectionNumber,
				mockedSection.CurrentTemperature,
				mockedSection.MinimumTemperature,
				mockedSection.CurrentCapacity,
				mockedSection.MinimumCapacity,
				mockedSection.MaximumCapacity,
				mockedSection.WarehouseID,
				mockedSection.ProductTypeID,
			).
			WillReturnResult(sqlmock.NewResult(int64(mockedSection.ID), 1))

		repository := NewRepository(db)

		assert.NotPanics(t, func() { repository.Update(mockedSection) })
	})

	t.Run("Should throw panic when expected prepare fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedSection := mockedSectionTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(UpdateQuery)).WillReturnError(sql.ErrConnDone)

		repository := NewRepository(db)

		assert.Panics(t, func() { repository.Update(mockedSection) })
	})

	t.Run("Should throw panic when expected exec fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedSection := mockedSectionTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(UpdateQuery))
		mock.ExpectExec(regexp.QuoteMeta(UpdateQuery)).
			WithArgs(
				mockedSection.SectionNumber,
				mockedSection.CurrentTemperature,
				mockedSection.MinimumTemperature,
				mockedSection.CurrentCapacity,
				mockedSection.MinimumCapacity,
				mockedSection.MaximumCapacity,
				mockedSection.WarehouseID,
				mockedSection.ProductTypeID,
			).
			WillReturnError(sql.ErrConnDone)

		repository := NewRepository(db)

		assert.Panics(t, func() { repository.Update(mockedSection) })
	})
}

func TestRepositoryGetAll(t *testing.T) {
	t.Run("Should return all sections", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id", "section_number", "current_temperature", "minimum_temperature", "current_capacity", "minimum_capacity", "maximum_capacity", "warehouse_id", "product_type_id"}
		rows := sqlmock.NewRows(columns)
		sectionId := 1
		rows.AddRow(sectionId, 1, 1, 1, 1, 1, 1, 1, 1)

		mock.ExpectQuery(regexp.QuoteMeta(GetAllQuery)).WillReturnRows(rows)

		repository := NewRepository(db)

		result := repository.GetAll()

		assert.NotNil(t, result)
	})

	t.Run("Should throw panic when query execution fail", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(GetAllQuery)).WillReturnError(sql.ErrConnDone)

		repository := NewRepository(db)

		assert.Panics(t, func() { repository.GetAll() })
	})
}

func SetupMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	return db, mock
}
