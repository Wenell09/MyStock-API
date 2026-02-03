package service

import (
	"github.com/Wenell09/MyStock/internal/dashboard/repository"
)

type DashboardServiceImpl struct {
	DashboardRepository repository.DashboardRepository
}

func NewDashboardService(dashboardRepository repository.DashboardRepository) DashboardService {
	return &DashboardServiceImpl{
		DashboardRepository: dashboardRepository,
	}
}

// CountData implements [DashboardService].
func (d *DashboardServiceImpl) CountData() (map[string]int64, error) {
	responseItems, err := d.DashboardRepository.CountItems()
	if err != nil {
		return nil, err
	}
	responseSuppliers, err := d.DashboardRepository.CountSuppliers()
	if err != nil {
		return nil, err
	}
	responseWarehouses, err := d.DashboardRepository.CountWarehouses()
	if err != nil {
		return nil, err
	}
	responseTransactions, err := d.DashboardRepository.CountTransactions()
	if err != nil {
		return nil, err
	}
	data := map[string]int64{
		"total_items":        responseItems,
		"total_suppliers":    responseSuppliers,
		"total_warehouses":   responseWarehouses,
		"total_transactions": responseTransactions,
	}
	return data, nil
}
