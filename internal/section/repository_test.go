package section_test

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/section"
	"github.com/stretchr/testify/assert"
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

func TestRepositoryGetAll(t *testing.T) {
	t.Run("Should return all sections", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id", "section_number", "current_temperature", "minimum_temperature", "current_capacity", "minimum_capacity", "maximum_capacity", "warehouse_id", "product_type_id"}
		rows := sqlmock.NewRows(columns)
		sectionId := 1
		rows.AddRow(sectionId, 1, 1, 1, 1, 1, 1, 1, 1)

		mock.ExpectQuery(regexp.QuoteMeta(section.GetAllQuery)).
			WillReturnRows(rows)

		repository := section.NewRepository(db)
		result := repository.GetAll()

		assert.NotNil(t, result)
	})
	t.Run("Should throw panic when query execution fail", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mock.ExpectQuery(section.GetAllQuery).WillReturnError(sql.ErrConnDone)
		repository := section.NewRepository(db)

		assert.Panics(t, func() { repository.GetAll() })
	})
}

func TestRepositoryGet(t *testing.T) {
	t.Run("Should return a section by specified id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id", "section_number", "current_temperature", "minimum_temperature", "current_capacity", "minimum_capacity", "maximum_capacity", "warehouse_id", "product_type_id"}
		rows := sqlmock.NewRows(columns)
		sectionId := 1

		rows.AddRow(sectionId, 1, 1, 1, 1, 1, 1, 1, 1)

		mock.ExpectQuery(regexp.QuoteMeta(section.GetQuery)).
			WithArgs(sectionId).
			WillReturnRows(rows)

		repository := section.NewRepository(db)

		result := repository.Get(sectionId)

		assert.NotNil(t, result)
	})

	t.Run("Should not return a section", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		sectionId := 1

		mock.ExpectQuery(regexp.QuoteMeta(section.GetQuery)).WithArgs(sectionId).WillReturnError(sql.ErrNoRows)

		repository := section.NewRepository(db)

		result := repository.Get(sectionId)

		assert.Nil(t, result)
	})

	t.Run("Should throw panic", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		sectionId := 1

		mock.ExpectQuery(section.GetQuery).WithArgs(sectionId).WillReturnError(sql.ErrConnDone)

		repository := section.NewRepository(db)

		assert.Panics(t, func() { repository.Get(sectionId) })
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

		mock.ExpectQuery(regexp.QuoteMeta(section.ExistsQuery)).
			WithArgs(sectionNumber).
			WillReturnRows(rows)

		repository := section.NewRepository(db)

		result := repository.Exists(sectionNumber)

		assert.True(t, result)
	})

	t.Run("Should return false when there are no query results", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		sectionNumber := 123

		mock.ExpectQuery(section.ExistsQuery).WithArgs(sectionNumber).WillReturnError(sql.ErrNoRows)

		repository := section.NewRepository(db)

		result := repository.Exists(sectionNumber)

		assert.False(t, result)
	})

	t.Run("Should return false when database has internal error", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		sectionNumber := 123

		mock.ExpectQuery(section.ExistsQuery).WithArgs(sectionNumber).WillReturnError(sql.ErrConnDone)

		repository := section.NewRepository(db)

		result := repository.Exists(sectionNumber)

		assert.False(t, result)
	})
}

func TestRepositorySave(t *testing.T) {
	t.Run("Should insert the section and return the section id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		lastInsertId := 1
		mockedSection := mockedSectionTemplate

		mock.ExpectPrepare(regexp.QuoteMeta(section.InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(section.InsertQuery)).
			WithArgs(mockedSection.SectionNumber, mockedSection.CurrentTemperature, mockedSection.MinimumTemperature, mockedSection.CurrentCapacity, mockedSection.MinimumCapacity, mockedSection.MaximumCapacity, mockedSection.WarehouseID, mockedSection.ProductTypeID).
			WillReturnResult(sqlmock.NewResult(int64(lastInsertId), 1))

		repository := section.NewRepository(db)

		result := repository.Save(mockedSection)

		assert.Equal(t, lastInsertId, result)
	})

	t.Run("Should throw panic when expect prepare fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedSection1 := mockedSectionTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(section.InsertQuery)).WillReturnError(sql.ErrConnDone)

		repository := section.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedSection1) })
	})

	t.Run("Should throw panic when expect exec fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedSection := mockedSectionTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(section.InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(section.InsertQuery)).
			WithArgs(mockedSection.ID, mockedSection.SectionNumber, mockedSection.CurrentTemperature, mockedSection.MinimumTemperature, mockedSection.CurrentCapacity, mockedSection.MinimumCapacity, mockedSection.MaximumCapacity, mockedSection.WarehouseID, mockedSection.ProductTypeID).
			WillReturnError(sql.ErrConnDone)

		repository := section.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedSection) })
	})

	t.Run("Should throw panic when sql has error", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedSection := mockedSectionTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(section.InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(section.InsertQuery)).
			WithArgs(mockedSection.SectionNumber, mockedSection.CurrentTemperature, mockedSection.MinimumTemperature, mockedSection.CurrentCapacity, mockedSection.MinimumCapacity, mockedSection.MaximumCapacity, mockedSection.WarehouseID, mockedSection.ProductTypeID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrConnDone))

		repository := section.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedSection) })
	})
}

func TestRepositoryUpdate(t *testing.T) {
	t.Run("Should update the section", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedSection := mockedSectionTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(section.UpdateQuery))
		mock.ExpectExec(regexp.QuoteMeta(section.UpdateQuery)).
			WithArgs(mockedSection.SectionNumber, mockedSection.CurrentTemperature, mockedSection.MinimumTemperature, mockedSection.CurrentCapacity, mockedSection.MinimumCapacity, mockedSection.MaximumCapacity, mockedSection.WarehouseID, mockedSection.ProductTypeID, mockedSection.ID).
			WillReturnResult(sqlmock.NewResult(int64(mockedSection.ID), 1))

		repository := section.NewRepository(db)

		assert.NotPanics(t, func() { repository.Update(mockedSection) })
	})

	t.Run("Should throw panic when expected prepare fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedSection4 := mockedSectionTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(section.UpdateQuery)).WillReturnError(sql.ErrConnDone)

		repository := section.NewRepository(db)

		assert.Panics(t, func() { repository.Update(mockedSection4) })
	})

	t.Run("Should throw panic when expected exec fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedSection := mockedSectionTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(section.UpdateQuery))
		mock.ExpectExec(regexp.QuoteMeta(section.UpdateQuery)).
			WithArgs(mockedSection.ID, mockedSection.SectionNumber, mockedSection.CurrentTemperature, mockedSection.MinimumTemperature, mockedSection.CurrentCapacity, mockedSection.MinimumCapacity, mockedSection.MaximumCapacity, mockedSection.WarehouseID, mockedSection.ProductTypeID).
			WillReturnError(sql.ErrConnDone)

		repository := section.NewRepository(db)

		assert.Panics(t, func() { repository.Update(mockedSection) })
	})
}

func TestRepositoryDelete(t *testing.T) {
	t.Run("Should delete the section", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedSection := mockedSectionTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(section.DeleteQuery))
		mock.ExpectExec(regexp.QuoteMeta(section.DeleteQuery)).
			WithArgs(mockedSection.ID).
			WillReturnResult(sqlmock.NewResult(int64(mockedSection.ID), 1))

		repository := section.NewRepository(db)

		assert.NotPanics(t, func() { repository.Delete(mockedSection.ID) })
	})

	t.Run("Should throw panic when expected prepare fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedSection := mockedSectionTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(section.DeleteQuery)).WillReturnError(sql.ErrConnDone)

		repository := section.NewRepository(db)

		assert.Panics(t, func() { repository.Delete(mockedSection.ID) })
	})

	t.Run("Should throw panic when expected exec fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedSection := mockedSectionTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(section.DeleteQuery))
		mock.ExpectExec(regexp.QuoteMeta(section.DeleteQuery)).
			WithArgs(mockedSection.ID).
			WillReturnError(sql.ErrConnDone)

		repository := section.NewRepository(db)

		assert.Panics(t, func() { repository.Delete(mockedSection.ID) })
	})
}

func TestRepositoryCountProductsByAllSections(t *testing.T) {
	t.Run("Should return products count by all sections", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"section_id", "section_number", "product_count"}
		rows := sqlmock.NewRows(columns)
		id := 1
		rows.AddRow(id, 1, 1)

		mock.ExpectQuery(regexp.QuoteMeta(section.CountProductsByAllSectionsQuery)).
			WillReturnRows(rows)

		repository := section.NewRepository(db)

		result := repository.CountProductsByAllSections()

		assert.Equal(t, len(result), 1)
	})
	t.Run("Should throw panic when expected query fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(section.CountProductsByAllSectionsQuery)).
			WillReturnError(sql.ErrConnDone)

		repository := section.NewRepository(db)

		assert.Panics(t, func() { repository.CountProductsByAllSections() })
	})
}

func TestRepositoryCountProductsBySection(t *testing.T) {
	t.Run("Should return products count by specific section id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"section_id", "section_number", "product_count"}
		rows := sqlmock.NewRows(columns)
		id := 1
		rows.AddRow(id, 1, 1)

		mock.ExpectQuery(regexp.QuoteMeta(section.CountProductsBySectionQuery)).
			WithArgs(id).
			WillReturnRows(rows)

		repository := section.NewRepository(db)

		result := repository.CountProductsBySection(id)

		assert.NotNil(t, result)
	})
	t.Run("Should throw panic when expected execution fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		id := 1
		mock.ExpectQuery(regexp.QuoteMeta(section.CountProductsBySectionQuery)).
			WillReturnError(sql.ErrNoRows)

		repository := section.NewRepository(db)
		result := repository.CountProductsBySection(id)

		assert.Nil(t, result)
	})
	t.Run("Should throw panic when expected query fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		id := 1
		mock.ExpectQuery(regexp.QuoteMeta(section.CountProductsBySectionQuery)).
			WithArgs(id).
			WillReturnError(sql.ErrConnDone)

		repository := section.NewRepository(db)

		assert.Panics(t, func() { repository.CountProductsBySection(id) })
	})
}

func SetupMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	return db, mock
}
