package employee_test

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/employee"
	"github.com/stretchr/testify/assert"
)

var (
	allDataQuery = regexp.QuoteMeta(employee.GetQuery)
	mockedEmployeeTemplate = domain.Employee {
		ID: 1,
		CardNumberID: "123456",
		FirstName: "PrimeiroNome",
		LastName: "Sobrenome",
		WarehouseID: 3,
	}
	mockedEmployeeUInboundOrders = domain.InboundOrdersByEmployee{
		ID: 1,
		CardNumberID: "123456",
		FirstName: "PrimeiroNome",
		LastName: "Sobrenome",
		WarehouseID: 3,
		InboundOrdersCount: 1,
	}
)

func TestRepositoryGetAll(t *testing.T) {
	t.Run("Should return all employees", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id", "card_number_id", "first_name", "last_name", "warehouse_id"}
		rows := sqlmock.NewRows(columns)
		employeeId := 1
		rows.AddRow(employeeId, "", "", "", 1)

		mock.ExpectQuery(employee.GetAllQuery).WillReturnRows(rows)

		repository := employee.NewRepository(db)

		result := repository.GetAll()

		assert.NotNil(t, result)
	})

	t.Run("Should throw panic when query execution fail", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mock.ExpectQuery(employee.GetAllQuery).WillReturnError(sql.ErrConnDone)

		repository := employee.NewRepository(db)

		assert.Panics(t, func() { repository.GetAll() })
	})
}

func TestRepositoryGet(t *testing.T) {
	t.Run("Should return an employee by specified id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id", "card_number_id", "first_name", "last_name", "warehouse_id"}
		rows := sqlmock.NewRows(columns)
		employeeId := 1
		rows.AddRow(employeeId, "", "", "", 1)

		mock.ExpectQuery(allDataQuery).WithArgs(employeeId).WillReturnRows(rows)

		repository := employee.NewRepository(db)

		result := repository.Get(employeeId)

		assert.NotNil(t, result)
	})

	t.Run("Should not return a employee", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		employeeId := 1

		mock.ExpectQuery(allDataQuery).WithArgs(employeeId).WillReturnError(sql.ErrNoRows)

		repository := employee.NewRepository(db)

		result := repository.Get(employeeId)

		assert.Nil(t, result)
	})

	t.Run("Should throw panic", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		employeeId := 1

		mock.ExpectQuery(allDataQuery).WithArgs(employeeId).WillReturnError(sql.ErrConnDone)

		repository := employee.NewRepository(db)

		assert.Panics(t, func() { repository.Get(employeeId) })
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

		mock.ExpectQuery(regexp.QuoteMeta(employee.ExistsQuery)).
			WithArgs(cardNumberId).
			WillReturnRows(rows)

		repository := employee.NewRepository(db)

		result := repository.Exists(cardNumberId)

		assert.True(t, result)
	})

	t.Run("Should return false when there are no query results", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		cardNumberId := "123"

		mock.ExpectQuery(employee.ExistsQuery).WithArgs(cardNumberId).WillReturnError(sql.ErrNoRows)

		repository := employee.NewRepository(db)

		result := repository.Exists(cardNumberId)

		assert.False(t, result)
	})

	t.Run("Should return false when database has internal error", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		cardNumberId := "123"

		mock.ExpectQuery(employee.ExistsQuery).WithArgs(cardNumberId).WillReturnError(sql.ErrConnDone)

		repository := employee.NewRepository(db)

		result := repository.Exists(cardNumberId)

		assert.False(t, result)
	})
}

func TestRepositorySave(t *testing.T) {
	t.Run("Should insert the employee and return the employee id", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"card_number_id", "first_name", "last_name", "warehouse_id"}
		rows := sqlmock.NewRows(columns)
		rows.AddRow(1, "", "", 1)

		lastInsertId := 1
		mockedEmployee := mockedEmployeeTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(employee.SaveQuery))
		mock.ExpectExec(regexp.QuoteMeta(employee.SaveQuery)).
			WithArgs(mockedEmployee.CardNumberID, mockedEmployee.FirstName, mockedEmployee.LastName, mockedEmployee.WarehouseID).
			WillReturnResult(sqlmock.NewResult(int64(lastInsertId), 1))

		repository := employee.NewRepository(db)

		result := repository.Save(mockedEmployee)

		assert.Equal(t, lastInsertId, result)
	})

	t.Run("Should throw panic when expect prepare fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedEmployee := mockedEmployeeTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(employee.SaveQuery)).WillReturnError(sql.ErrConnDone)

		repository := employee.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedEmployee) })
	})

	t.Run("Should throw panic when expect exec fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedEmployee := mockedEmployeeTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(employee.SaveQuery))
		mock.ExpectExec(regexp.QuoteMeta(employee.SaveQuery)).
			WithArgs(mockedEmployee.CardNumberID, mockedEmployee.FirstName, mockedEmployee.LastName, mockedEmployee.WarehouseID).
			WillReturnError(sql.ErrConnDone)

		repository := employee.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedEmployee) })
	})

	t.Run("Should throw panic when sql has error", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedEmployee := mockedEmployeeTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(employee.SaveQuery))
		mock.ExpectExec(regexp.QuoteMeta(employee.SaveQuery)).
			WithArgs(mockedEmployee.CardNumberID, mockedEmployee.FirstName, mockedEmployee.LastName, mockedEmployee.WarehouseID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrConnDone))

		repository := employee.NewRepository(db)

		assert.Panics(t, func() { repository.Save(mockedEmployee) })
	})
}

func TestRepositoryUpdate(t *testing.T) {
	t.Run("Should update the employee", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedEmployee := mockedEmployeeTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(employee.UpdateQuery))
		mock.ExpectExec(regexp.QuoteMeta(employee.UpdateQuery)).
			WithArgs(mockedEmployee.CardNumberID, mockedEmployee.FirstName, mockedEmployee.LastName, mockedEmployee.WarehouseID, mockedEmployee.ID).
			WillReturnResult(sqlmock.NewResult(int64(mockedEmployee.ID), 1))

		repository := employee.NewRepository(db)

		assert.NotPanics(t, func() { repository.Update(mockedEmployee) })
	})

	t.Run("Should throw panic when expected prepare fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedEmployee := mockedEmployeeTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(employee.UpdateQuery)).WillReturnError(sql.ErrConnDone)

		repository := employee.NewRepository(db)

		assert.Panics(t, func() { repository.Update(mockedEmployee) })
	})

	t.Run("Should throw panic when expected exec fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedEmployee := mockedEmployeeTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(employee.UpdateQuery))
		mock.ExpectExec(regexp.QuoteMeta(employee.UpdateQuery)).
			WithArgs(mockedEmployee.CardNumberID, mockedEmployee.FirstName, mockedEmployee.LastName, mockedEmployee.WarehouseID, mockedEmployee.ID).
			WillReturnError(sql.ErrConnDone)

		repository := employee.NewRepository(db)

		assert.Panics(t, func() { repository.Update(mockedEmployee) })
	})
}

func TestRepositoryDelete(t *testing.T) {
	t.Run("Should delete the employee", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedEmployee := mockedEmployeeTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(employee.DeleteQuery))
		mock.ExpectExec(regexp.QuoteMeta(employee.DeleteQuery)).
			WithArgs(mockedEmployee.ID).
			WillReturnResult(sqlmock.NewResult(int64(mockedEmployee.ID), 1))

		repository := employee.NewRepository(db)

		assert.NotPanics(t, func() { repository.Delete(mockedEmployee.ID) })
	})

	t.Run("Should throw panic when expected prepare fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedEmployee := mockedEmployeeTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(employee.DeleteQuery)).WillReturnError(sql.ErrConnDone)

		repository := employee.NewRepository(db)

		assert.Panics(t, func() { repository.Delete(mockedEmployee.ID) })
	})

	t.Run("Should throw panic when expected exec fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mockedEmployee := mockedEmployeeTemplate
		mock.ExpectPrepare(regexp.QuoteMeta(employee.DeleteQuery))
		mock.ExpectExec(regexp.QuoteMeta(employee.DeleteQuery)).
			WithArgs(mockedEmployee.ID).
			WillReturnError(sql.ErrConnDone)

		repository := employee.NewRepository(db)

		assert.Panics(t, func() { repository.Delete(mockedEmployee.ID) })
	})
}

func TestRepositoryCountInboundOrdersByAllEmployees(t *testing.T) {

	t.Run("Should return all respective employees inbound orders", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id", "card_number_id", "first_name", "last_name", "warehouse_id", "inbound_orders_count"}
		rows := sqlmock.NewRows(columns)
		employeeId := 1
		rows.AddRow(employeeId, "", "", "", 1, 1)

		mock.ExpectQuery(regexp.QuoteMeta(employee.CountInboundOrdersByAllEmployeesQuery)).
		WillReturnRows(rows)

		repository := employee.NewRepository(db)
		result := repository.CountInboundOrdersByAllEmployees()

		assert.NotNil(t, result)
	})

	t.Run("should panic when expected exec fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(employee.CountInboundOrdersByAllEmployeesQuery)).
			WillReturnError(sql.ErrConnDone)

		repository := employee.NewRepository(db)

		assert.Panics(t, func() { repository.CountInboundOrdersByAllEmployees() })
	})
}

func TestRepositoryCountInboundOrdersByEmployee(t *testing.T) {

	t.Run("should return employee inbound orders", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		columns := []string{"id", "card_number_id", "first_name", "last_name", "warehouse_id", "inbound_orders_count"}
		rows := sqlmock.NewRows(columns)
		employeeId := 1
		rows.AddRow(employeeId, "", "", "", 1, 1)

		mock.ExpectQuery(regexp.QuoteMeta(employee.CountInboundOrdersByEmployeeQuery)).
		WithArgs(employeeId).
		WillReturnRows(rows)

		repository := employee.NewRepository(db)
		result := repository.CountInboundOrdersByEmployee(employeeId)

		assert.NotNil(t, result)
	})

	t.Run("should not return employee inbound orders", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		employeeId := 1

		mock.ExpectQuery(regexp.QuoteMeta(employee.CountInboundOrdersByEmployeeQuery)).
		WithArgs(employeeId).
		WillReturnError(sql.ErrNoRows)

		repository := employee.NewRepository(db)
		result := repository.CountInboundOrdersByEmployee(employeeId)

		assert.Nil(t, result)
	})
	
	t.Run("should panic when expected exec fails", func(t *testing.T) {
		db, mock := SetupMock(t)
		defer db.Close()

		employeeId := 1

		mock.ExpectQuery(regexp.QuoteMeta(employee.CountInboundOrdersByEmployeeQuery)).
		WithArgs(employeeId).
		WillReturnError(sql.ErrConnDone)

		repository := employee.NewRepository(db)

		assert.Panics(t, func() { repository.CountInboundOrdersByEmployee(employeeId) })
	})
}

func SetupMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	return db, mock
}
