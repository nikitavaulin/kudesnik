package products_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

func (s *ProductsService) PatchProduct(ctx context.Context, id uuid.UUID, patch domain.ProductPatch) (domain.Product, error) {
	switch p := patch.(type) {
	case *domain.WindowPatch:
		patched, err := s.patchWindow(ctx, id, *p)
		return &patched, err

	case *domain.ProductBasePatch:
		patched, err := s.patchProduct(ctx, id, *p)
		return &patched, err

	case *domain.EntranceDoorPatch:
		patched, err := s.patchEntranceDoor(ctx, id, *p)
		return &patched, err

	case *domain.InteriorDoorPatch:
		patched, err := s.patchInteriorDoor(ctx, id, *p)
		return &patched, err

	case *domain.BalconyPatch:
		patched, err := s.patchBalcony(ctx, id, *p)
		return &patched, err
	}

	return nil, fmt.Errorf("unknown product patch")
}

func (s *ProductsService) patchWindow(ctx context.Context, id uuid.UUID, patch domain.WindowPatch) (domain.Window, error) {
	window, err := s.productRepo.GetWindow(ctx, id)
	if err != nil {
		return domain.Window{}, fmt.Errorf("failed to get window from repo: %w", err)
	}

	if err := window.ApplyPatch(&patch); err != nil {
		return domain.Window{}, fmt.Errorf("failed to apply window patch: %w", err)
	}

	patchedProduct, err := s.productRepo.PatchWindow(ctx, id, window)
	if err != nil {
		return domain.Window{}, fmt.Errorf("failed to patch window in repo: %w", err)
	}

	return patchedProduct, nil
}

func (s *ProductsService) patchProduct(ctx context.Context, id uuid.UUID, patch domain.ProductBasePatch) (domain.ProductBase, error) {
	product, err := s.productRepo.GetProduct(ctx, id)
	if err != nil {
		return domain.ProductBase{}, fmt.Errorf("failed to get product from repo: %w", err)
	}

	if err := product.ApplyPatch(&patch); err != nil {
		return domain.ProductBase{}, fmt.Errorf("failed to apply product patch: %w", err)
	}

	patchedProduct, err := s.productRepo.PatchProduct(ctx, id, product)
	if err != nil {
		return domain.ProductBase{}, fmt.Errorf("failed to patch product in repo: %w", err)
	}

	return patchedProduct, nil
}

func (s *ProductsService) patchEntranceDoor(ctx context.Context, id uuid.UUID, patch domain.EntranceDoorPatch) (domain.EntranceDoor, error) {
	door, err := s.productRepo.GetEntranceDoor(ctx, id)
	if err != nil {
		return domain.EntranceDoor{}, fmt.Errorf("failed to get entrance door from repo: %w", err)
	}

	if err := door.ApplyPatch(&patch); err != nil {
		return domain.EntranceDoor{}, fmt.Errorf("failed to apply entrance door patch: %w", err)
	}

	patchedDoor, err := s.productRepo.PatchEntranceDoor(ctx, id, door)
	if err != nil {
		return domain.EntranceDoor{}, fmt.Errorf("failed to patch entrance door in repo: %w", err)
	}

	return patchedDoor, nil
}

func (s *ProductsService) patchInteriorDoor(ctx context.Context, id uuid.UUID, patch domain.InteriorDoorPatch) (domain.InteriorDoor, error) {
	door, err := s.productRepo.GetInteriorDoor(ctx, id)
	if err != nil {
		return domain.InteriorDoor{}, fmt.Errorf("failed to get interior door from repo: %w", err)
	}

	if err := door.ApplyPatch(&patch); err != nil {
		return domain.InteriorDoor{}, fmt.Errorf("failed to apply interior door patch: %w", err)
	}

	patchedDoor, err := s.productRepo.PatchInteriorDoor(ctx, id, door)
	if err != nil {
		return domain.InteriorDoor{}, fmt.Errorf("failed to patch interior door in repo: %w", err)
	}

	return patchedDoor, nil
}

func (s *ProductsService) patchBalcony(ctx context.Context, id uuid.UUID, patch domain.BalconyPatch) (domain.Balcony, error) {
	balcony, err := s.productRepo.GetBalcony(ctx, id)
	if err != nil {
		return domain.Balcony{}, fmt.Errorf("failed to get balcony from repo: %w", err)
	}

	if err := balcony.ApplyPatch(&patch); err != nil {
		return domain.Balcony{}, fmt.Errorf("failed to apply balcony patch: %w", err)
	}

	patchedBalcony, err := s.productRepo.PatchBalcony(ctx, id, balcony)
	if err != nil {
		return domain.Balcony{}, fmt.Errorf("failed to patch balcony in repo: %w", err)
	}

	return patchedBalcony, nil
}
