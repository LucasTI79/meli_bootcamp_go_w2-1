package product_test

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryGetAll(t *testing.T) {
	t.Run("Should return all products", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id", "description", "expiration_rate", "freezing_rate", "height", "lenght", "netweight", "product_code", "recommended_freezing_temperature", "width", "id_product_type", "id_seller"}
		rows := sqlmock.NewRows(columns)
		productId := 1
		rows.AddRow(productId, "", 1, 1, 1, 1, 1, "ABC", 1, 1, 1, 1)

		mock.ExpectQuery(regexp.QuoteMeta(product.GetAllQuery)).WillReturnRows(rows)

		repository := product.NewRepository(db)

		result := repository.GetAll()

		assert.NotNil(t, result)
	})

	t.Run("Should throw panic when query execution fail", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(product.GetAllQuery)).WillReturnError(sql.ErrConnDone)

		repository := product.NewRepository(db)

		assert.Panics(t, func() { repository.GetAll() })
	})
}

func TestRepositoryGet(t *testing.T) {
	t.Run("Should return a product by specified id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id", "description", "expiration_rate", "freezing_rate", "height", "lenght", "netweight", "product_code", "recommended_freezing_temperature", "width", "id_product_type", "id_seller"}
		rows := sqlmock.NewRows(columns)
		productId := 1
		rows.AddRow(productId, "", 1, 1, 1, 1, 1, "ABC", 1, 1, 1, 1)

		mock.ExpectQuery(regexp.QuoteMeta(product.GetQuery)).WithArgs(productId).WillReturnRows(rows)

		repository := product.NewRepository(db)

		result := repository.Get(productId)

		assert.NotNil(t, result)
	})

	t.Run("Should not return a product", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		productId := 1

		mock.ExpectQuery(regexp.QuoteMeta(product.GetQuery)).WithArgs(productId).WillReturnError(sql.ErrNoRows)

		repository := product.NewRepository(db)

		result := repository.Get(productId)

		assert.Nil(t, result)
	})

	t.Run("Should throw panic", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		productId := 1

		mock.ExpectQuery(regexp.QuoteMeta(product.GetQuery)).WithArgs(productId).WillReturnError(sql.ErrConnDone)

		repository := product.NewRepository(db)

		assert.Panics(t, func() { repository.Get(productId) })
	})
}

func TestRepositoryExists(t *testing.T) {
	t.Run("Should return true", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"cid"}
		rows := sqlmock.NewRows(columns)
		productCode := "ABC123"
		rows.AddRow(productCode)

		mock.ExpectQuery(regexp.QuoteMeta(product.ExistsQuery)).
			WithArgs(productCode).
			WillReturnRows(rows)

		repository := product.NewRepository(db)

		result := repository.Exists(productCode)

		assert.True(t, result)
	})

	t.Run("Should return false when there are no query results", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		productCode := "ABC123"

		mock.ExpectQuery(product.ExistsQuery).WithArgs(productCode).WillReturnError(sql.ErrNoRows)

		repository := product.NewRepository(db)

		result := repository.Exists(productCode)

		assert.False(t, result)
	})

	t.Run("Should return false when database has internal error", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		productCode := "ABC123"

		mock.ExpectQuery(product.ExistsQuery).WithArgs(productCode).WillReturnError(sql.ErrConnDone)

		repository := product.NewRepository(db)

		result := repository.Exists(productCode)

		assert.False(t, result)
	})
}

func TestRepositorySave(t *testing.T) {
	t.Run("Should insert the product and return the product id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		lastInsertId := 1
		mockedProduct := mockedProductTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(product.InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(product.InsertQuery)).
			WithArgs(
				mockedProduct.Description,
				mockedProduct.ExpirationRate,
				mockedProduct.FreezingRate,
				mockedProduct.Height,
				mockedProduct.Length,
				mockedProduct.Netweight,
				mockedProduct.ProductCode,
				mockedProduct.RecomFreezTemp,
				mockedProduct.Width,
				mockedProduct.ProductTypeID,
				mockedProduct.SellerID,
			).
			WillReturnResult(sqlmock.NewResult(int64(lastInsertId), 1))

		repository := product.NewRepository(db)

		result := repository.Save(mockedProduct)

		assert.Equal(t, lastInsertId, result)
	})

	t.Run("Should throw panic when expect prepare fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedProduct := mockedProductTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(product.InsertQuery)).WillReturnError(sql.ErrConnDone)

		repository := product.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedProduct) })
	})

	t.Run("Should throw panic when expect exec fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedProduct := mockedProductTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(product.InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(product.InsertQuery)).
			WithArgs(
				mockedProduct.Description,
				mockedProduct.ExpirationRate,
				mockedProduct.FreezingRate,
				mockedProduct.Height,
				mockedProduct.Length,
				mockedProduct.Netweight,
				mockedProduct.ProductCode,
				mockedProduct.RecomFreezTemp,
				mockedProduct.Width,
				mockedProduct.ProductTypeID,
				mockedProduct.SellerID,
			).
			WillReturnError(sql.ErrConnDone)

		repository := product.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedProduct) })
	})

	t.Run("Should throw panic when sql has error", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedProduct := mockedProductTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(product.InsertQuery))
		mock.ExpectExec(regexp.QuoteMeta(product.InsertQuery)).
			WithArgs(
				mockedProduct.Description,
				mockedProduct.ExpirationRate,
				mockedProduct.FreezingRate,
				mockedProduct.Height,
				mockedProduct.Length,
				mockedProduct.Netweight,
				mockedProduct.ProductCode,
				mockedProduct.RecomFreezTemp,
				mockedProduct.Width,
				mockedProduct.ProductTypeID,
				mockedProduct.SellerID,
			).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrConnDone))

		repository := product.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedProduct) })
	})
}

func TestRepositoryUpdate(t *testing.T) {
	t.Run("Should update the product", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedProduct := mockedProductTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(product.UpdateQuery))
		mock.ExpectExec(regexp.QuoteMeta(product.UpdateQuery)).
			WithArgs(
				mockedProduct.Description,
				mockedProduct.ExpirationRate,
				mockedProduct.FreezingRate,
				mockedProduct.Height,
				mockedProduct.Length,
				mockedProduct.Netweight,
				mockedProduct.ProductCode,
				mockedProduct.RecomFreezTemp,
				mockedProduct.Width,
				mockedProduct.ProductTypeID,
				mockedProduct.SellerID,
				mockedProduct.ID,
			).
			WillReturnResult(sqlmock.NewResult(int64(mockedProduct.ID), 1))

		repository := product.NewRepository(db)

		assert.NotPanics(t, func() { repository.Update(mockedProduct) })
	})

	t.Run("Should throw panic when expected prepare fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedProduct := mockedProductTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(product.UpdateQuery)).WillReturnError(sql.ErrConnDone)

		repository := product.NewRepository(db)

		assert.Panics(t, func() { repository.Update(mockedProduct) })
	})

	t.Run("Should throw panic when expected exec fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedProduct := mockedProductTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(product.UpdateQuery))
		mock.ExpectExec(regexp.QuoteMeta(product.UpdateQuery)).
			WithArgs(
				mockedProduct.Description,
				mockedProduct.ExpirationRate,
				mockedProduct.FreezingRate,
				mockedProduct.Height,
				mockedProduct.Length,
				mockedProduct.Netweight,
				mockedProduct.ProductCode,
				mockedProduct.RecomFreezTemp,
				mockedProduct.Width,
				mockedProduct.ProductTypeID,
				mockedProduct.SellerID,
				mockedProduct.ID,
			).
			WillReturnError(sql.ErrConnDone)

		repository := product.NewRepository(db)

		assert.Panics(t, func() { repository.Update(mockedProduct) })
	})
}

func TestRepositoryDelete(t *testing.T) {
	t.Run("Should delete the product", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedProduct := mockedProductTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(product.DeleteQuery))
		mock.ExpectExec(regexp.QuoteMeta(product.DeleteQuery)).
			WithArgs(mockedProduct.ID).
			WillReturnResult(sqlmock.NewResult(int64(mockedProduct.ID), 1))

		repository := product.NewRepository(db)

		assert.NotPanics(t, func() { repository.Delete(mockedProduct.ID) })
	})

	t.Run("Should throw panic when expected prepare fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedProduct := mockedProductTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(product.DeleteQuery)).WillReturnError(sql.ErrConnDone)

		repository := product.NewRepository(db)

		assert.Panics(t, func() { repository.Delete(mockedProduct.ID) })
	})

	t.Run("Should throw panic when expected exec fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedProduct := mockedProductTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(product.DeleteQuery))
		mock.ExpectExec(regexp.QuoteMeta(product.DeleteQuery)).
			WithArgs(mockedProduct.ID).
			WillReturnError(sql.ErrConnDone)

		repository := product.NewRepository(db)

		assert.Panics(t, func() { repository.Delete(mockedProduct.ID) })
	})
}

func TestRepositoryCountRecordsByAllProducts(t *testing.T) {
	t.Run("Should return records count report by all products", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"product_id", "description", "records_count"}
		rows := sqlmock.NewRows(columns)
		rows.AddRow(1, "", 1)

		mock.ExpectQuery(regexp.QuoteMeta(product.CountRecordsByAllProductsQuery)).WillReturnRows(rows)

		repository := product.NewRepository(db)

		result := repository.CountRecordsByAllProducts()

		assert.Equal(t, len(result), 1)
	})

	t.Run("Should throw panic when query execution fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(product.CountRecordsByAllProductsQuery)).WillReturnError(sql.ErrConnDone)

		repository := product.NewRepository(db)

		assert.Panics(t, func() { repository.CountRecordsByAllProducts() })
	})
}

func TestRepositoryCountRecordsByProduct(t *testing.T) {
	t.Run("Should return records count report by specified product id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"product_id", "description", "records_count"}
		rows := sqlmock.NewRows(columns)
		productId := 1
		rows.AddRow(1, "", 1)

		mock.ExpectQuery(regexp.QuoteMeta(product.CountRecordsByProductQuery)).
			WithArgs(productId).
			WillReturnRows(rows)

		repository := product.NewRepository(db)

		result := repository.CountRecordsByProduct(productId)

		assert.NotNil(t, result)
	})

	t.Run("Should throw panic when query execution fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		recordId := 1
		mock.ExpectQuery(regexp.QuoteMeta(product.CountRecordsByProductQuery)).WillReturnError(sql.ErrNoRows)

		repository := product.NewRepository(db)

		result := repository.CountRecordsByProduct(recordId)

		assert.Nil(t, result)
	})

	t.Run("Should throw panic when query execution fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		recordId := 1
		mock.ExpectQuery(regexp.QuoteMeta(product.CountRecordsByProductQuery)).WillReturnError(sql.ErrConnDone)

		repository := product.NewRepository(db)

		assert.Panics(t, func() { repository.CountRecordsByProduct(recordId) })
	})
}

func SetupMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	return db, mock
}
