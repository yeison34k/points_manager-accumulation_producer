package app

import (
	"accumulation_producer/internal/domain"
	"fmt"
)

type PointHandler struct {
	PointApp *PointApp
}

func NewPointHandler(pointApp *PointApp) *PointHandler {
	return &PointHandler{
		PointApp: pointApp,
	}
}

func (h *PointHandler) HandlePointCreation(point *domain.Point) error {
	err := h.PointApp.CreatePoint(point)
	if err != nil {
		return fmt.Errorf("failed to create point: %w", err)
	}

	return nil
}
