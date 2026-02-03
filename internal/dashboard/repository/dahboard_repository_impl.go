package repository

import (
	"github.com/Wenell09/MyStock/internal/models"
	"gorm.io/gorm"
)

type DashboardRepositoryImpl struct {
	DB *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) DashboardRepository {
	return &DashboardRepositoryImpl{
		DB: db,
	}
}

// CountItems implements [DashboardRepository].
func (d *DashboardRepositoryImpl) CountItems() (int64, error) {
	var totalItems int64
	if err := d.DB.Model(&models.Item{}).Count(&totalItems).Error; err != nil {
		return totalItems, err
	}
	return totalItems, nil
}

// CountSuppliers implements [DashboardRepository].
func (d *DashboardRepositoryImpl) CountSuppliers() (int64, error) {
	var totalSuppliers int64
	if err := d.DB.Model(&models.Supplier{}).Count(&totalSuppliers).Error; err != nil {
		return totalSuppliers, err
	}
	return totalSuppliers, nil
}

// CountTransactions implements [DashboardRepository].
func (d *DashboardRepositoryImpl) CountTransactions() (int64, error) {
	var totalTransactions int64
	if err := d.DB.Model(&models.StockTransaction{}).Count(&totalTransactions).Error; err != nil {
		return totalTransactions, err
	}
	return totalTransactions, nil
}

// CountWarehouses implements [DashboardRepository].
func (d *DashboardRepositoryImpl) CountWarehouses() (int64, error) {
	var totalWarehouses int64
	if err := d.DB.Model(&models.Warehouse{}).Count(&totalWarehouses).Error; err != nil {
		return totalWarehouses, err
	}
	return totalWarehouses, nil
}
