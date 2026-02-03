package repository

type DashboardRepository interface {
	CountItems() (int64, error)
	CountSuppliers() (int64, error)
	CountWarehouses() (int64, error)
	CountTransactions() (int64, error)
}
