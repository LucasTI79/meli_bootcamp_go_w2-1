package purchase_order_test

import (
	"testing"

	buyerMocks "github.com/extmatperez/meli_bootcamp_go_w2-1/internal/buyer/mocks"
	carrierMocks "github.com/extmatperez/meli_bootcamp_go_w2-1/internal/carrier/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	orderStatusMocks "github.com/extmatperez/meli_bootcamp_go_w2-1/internal/order_status/mocks"
	productRecordMocks "github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product_record/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/purchase_order"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/purchase_order/mocks"
	warehouseMocks "github.com/extmatperez/meli_bootcamp_go_w2-1/internal/warehouse/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/stretchr/testify/assert"
)

var (
	mockedBuyerTemplate = domain.Buyer{
		ID: 1,
	}
	mockedOrderStatusTemplate = domain.OrderStatus{
		ID: 1,
	}
	mockedWarehouseTemplate = domain.Warehouse{
		ID: 1,
	}
	mockedProductRecordTemplate = domain.ProductRecord{
		ID: 1,
	}
	mockedCarrierTemplate = domain.Carrier{
		ID: 1,
	}
)

func TestServiceCreate(t *testing.T) {
	t.Run("Should return a created purchase order", func(t *testing.T) {
		service, repository, buyerRepo, orderStatusRepo, warehouseRepo, carrierRepo, productRecordRepo := CreateService(t)

		mockedPurchaseOrder := mockedPurchaseOrderTemplate
		mockedBuyer := mockedBuyerTemplate
		mockedOrderStatus := mockedOrderStatusTemplate
		mockedWarehouse := mockedWarehouseTemplate
		mockedProductRecord := mockedProductRecordTemplate
		mockedCarrier := mockedCarrierTemplate

		purchaseOrderID := 1

		repository.On("Exists", mockedPurchaseOrder.OrderNumber).Return(false)
		buyerRepo.On("Get", mockedBuyer.ID).Return(&mockedBuyer)
		orderStatusRepo.On("Get", mockedOrderStatus.ID).Return(&mockedOrderStatus)
		warehouseRepo.On("Get", mockedWarehouse.ID).Return(&mockedWarehouse)
		productRecordRepo.On("Get", mockedProductRecord.ID).Return(&mockedProductRecord)
		carrierRepo.On("Get", mockedCarrier.ID).Return(&mockedCarrier)
		repository.On("Save", mockedPurchaseOrder).Return(purchaseOrderID)
		repository.On("Get", purchaseOrderID).Return(&mockedPurchaseOrder)

		result, err := service.Create(mockedPurchaseOrder)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, mockedPurchaseOrder, *result)
	})

	t.Run("Should return a conflict error when order number already exists", func(t *testing.T) {
		service, repository, _, _, _, _, _ := CreateService(t)

		mockedPurchaseOrder := mockedPurchaseOrderTemplate
		repository.On("Exists", mockedPurchaseOrder.OrderNumber).Return(true)

		result, err := service.Create(mockedPurchaseOrder)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceAlreadyExists](err))
	})

	t.Run("Should return a conflict error when buyer id does not exist", func(t *testing.T) {
		service, repository, buyerRepo, _, _, _, _ := CreateService(t)

		mockedPurchaseOrder := mockedPurchaseOrderTemplate
		mockedBuyer := mockedBuyerTemplate
		var buyerRepositoryGetResult *domain.Buyer

		repository.On("Exists", mockedPurchaseOrder.OrderNumber).Return(false)
		buyerRepo.On("Get", mockedBuyer.ID).Return(buyerRepositoryGetResult)

		result, err := service.Create(mockedPurchaseOrder)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.DependentResourceNotFound](err))
	})

	t.Run("Should return a conflict error when order status id does not exist", func(t *testing.T) {
		service, repository, buyerRepo, orderStatusRepo, _, _, _ := CreateService(t)

		mockedPurchaseOrder := mockedPurchaseOrderTemplate
		mockedBuyer := mockedBuyerTemplate
		mockedOrderStatus := mockedOrderStatusTemplate
		var orderStatusRepositoryGetResult *domain.OrderStatus

		repository.On("Exists", mockedPurchaseOrder.OrderNumber).Return(false)
		buyerRepo.On("Get", mockedBuyer.ID).Return(&mockedBuyer)
		orderStatusRepo.On("Get", mockedOrderStatus.ID).Return(orderStatusRepositoryGetResult)

		result, err := service.Create(mockedPurchaseOrder)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.DependentResourceNotFound](err))
	})

	t.Run("Should return a conflict error when warehouse id does not exist", func(t *testing.T) {
		service, repository, buyerRepo, orderStatusRepo, warehouseRepo, _, _ := CreateService(t)

		mockedPurchaseOrder := mockedPurchaseOrderTemplate
		mockedBuyer := mockedBuyerTemplate
		mockedOrderStatus := mockedOrderStatusTemplate
		mockedWarehouse := mockedWarehouseTemplate
		var warehouseRepositoryGetResult *domain.Warehouse

		repository.On("Exists", mockedPurchaseOrder.OrderNumber).Return(false)
		buyerRepo.On("Get", mockedBuyer.ID).Return(&mockedBuyer)
		orderStatusRepo.On("Get", mockedOrderStatus.ID).Return(&mockedOrderStatus)
		warehouseRepo.On("Get", mockedWarehouse.ID).Return(warehouseRepositoryGetResult)

		result, err := service.Create(mockedPurchaseOrder)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.DependentResourceNotFound](err))
	})
	t.Run("Should return a conflict error when product record id does not exist", func(t *testing.T) {
		service, repository, buyerRepo, orderStatusRepo, warehouseRepo, _, productRecordRepo := CreateService(t)

		mockedPurchaseOrder := mockedPurchaseOrderTemplate
		mockedBuyer := mockedBuyerTemplate
		mockedOrderStatus := mockedOrderStatusTemplate
		mockedWarehouse := mockedWarehouseTemplate
		mockedProductRecord := mockedProductRecordTemplate
		var productRecordRepositoryGetResult *domain.ProductRecord

		repository.On("Exists", mockedPurchaseOrder.OrderNumber).Return(false)
		buyerRepo.On("Get", mockedBuyer.ID).Return(&mockedBuyer)
		orderStatusRepo.On("Get", mockedOrderStatus.ID).Return(&mockedOrderStatus)
		warehouseRepo.On("Get", mockedWarehouse.ID).Return(&mockedWarehouse)
		productRecordRepo.On("Get", mockedProductRecord.ID).Return(productRecordRepositoryGetResult)

		result, err := service.Create(mockedPurchaseOrder)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.DependentResourceNotFound](err))
	})

	t.Run("Should return a conflict error when carrier id does not exist", func(t *testing.T) {
		service, repository, buyerRepo, orderStatusRepo, warehouseRepo, carrierRepo, productRecordRepo := CreateService(t)

		mockedPurchaseOrder := mockedPurchaseOrderTemplate
		mockedBuyer := mockedBuyerTemplate
		mockedOrderStatus := mockedOrderStatusTemplate
		mockedWarehouse := mockedWarehouseTemplate
		mockedProductRecord := mockedProductRecordTemplate
		mockedCarrier := mockedCarrierTemplate
		var carrierRepositoryGetResult *domain.Carrier

		repository.On("Exists", mockedPurchaseOrder.OrderNumber).Return(false)
		buyerRepo.On("Get", mockedBuyer.ID).Return(&mockedBuyer)
		orderStatusRepo.On("Get", mockedOrderStatus.ID).Return(&mockedOrderStatus)
		warehouseRepo.On("Get", mockedWarehouse.ID).Return(&mockedWarehouse)
		productRecordRepo.On("Get", mockedProductRecord.ID).Return(&mockedProductRecord)
		carrierRepo.On("Get", mockedCarrier.ID).Return(carrierRepositoryGetResult)

		result, err := service.Create(mockedPurchaseOrder)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.DependentResourceNotFound](err))
	})
}

func CreateService(t *testing.T) (purchase_order.Service, *mocks.Repository, *buyerMocks.Repository, *orderStatusMocks.Repository, *warehouseMocks.Repository, *carrierMocks.Repository, *productRecordMocks.Repository) {
	t.Helper()
	repository := new(mocks.Repository)
	buyerRepo := new(buyerMocks.Repository)
	orderStatusRepo := new(orderStatusMocks.Repository)
	warehouseRepo := new(warehouseMocks.Repository)
	carrierRepo := new(carrierMocks.Repository)
	productRecordRepo := new(productRecordMocks.Repository)
	service := purchase_order.NewService(repository, buyerRepo, orderStatusRepo, warehouseRepo, carrierRepo, productRecordRepo)

	return service, repository, buyerRepo, orderStatusRepo, warehouseRepo, carrierRepo, productRecordRepo
}
