package app

import "accumulation_producer/internal/domain"

type PointService interface {
	CreatePoint(point *domain.Point) error
}

type PointApp struct {
	PointService PointService
}

func NewPointApplication(pointService PointService) *PointApp {
	return &PointApp{
		PointService: pointService,
	}
}

func (a *PointApp) CreatePoint(point *domain.Point) error {
	return a.PointService.CreatePoint(point)
}
