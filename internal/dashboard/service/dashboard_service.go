package service

type DashboardService interface {
	CountData() (map[string]int64, error)
}
