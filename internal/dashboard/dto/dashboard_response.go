package dto

type DashboardResponse struct {
	TotalItems        int64 `json:"total_items"`
	TotalSuppliers    int64 `json:"total_suppliers"`
	TotalWarehouses   int64 `json:"total_warehouses"`
	TotalTransactions int64 `json:"total_transactions"`
}
