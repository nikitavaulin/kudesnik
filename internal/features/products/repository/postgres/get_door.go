package products_repository_postges

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
	core_postgres_pool "github.com/nikitavaulin/kudesnik/internal/core/repository/postgres/pool"
)

func (r *ProductsRepositoryPostgres) GetEntranceDoor(ctx context.Context, id uuid.UUID) (domain.EntranceDoor, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	query := `
        SELECT 
            d.door_id,
            d.collection,
            d.width,
            d.height,
            d.outside_material,
            d.outside_color,
            d.outside_picture,
            d.inside_material,
            d.inside_color,
            d.inside_picture,
            ed.strength_class,
            ed.sound_insulation,
            ed.metal_thickness,
            ed.box_thickness,
            ed.leaf_thickness,
            ed.leaf_description,
            ed.filling_description,
            ed.main_lock,
            ed.additional_lock,
            ed.insulation_description,
            ed.hinges,
            p.product_name,
            p.price,
            p.description,
            p.category_code,
            p.is_visible,
            p.version
        FROM kudesnik.doors d
        INNER JOIN kudesnik.products p ON d.door_id = p.product_id
        INNER JOIN kudesnik.entrance_doors ed ON d.door_id = ed.door_id
        WHERE d.door_id = $1
    `

	row := r.pool.QueryRow(ctx, query, id)

	var door domain.EntranceDoor

	err := row.Scan(
		&door.ID,
		&door.Collection,
		&door.Width,
		&door.Height,
		&door.OutsideMaterial,
		&door.OutsideColor,
		&door.OutsidePicture,
		&door.InsideMaterial,
		&door.InsideColor,
		&door.InsidePicture,
		&door.StrengthClass,
		&door.SoundInsulation,
		&door.MetalThickness,
		&door.BoxThickness,
		&door.LeafThickness,
		&door.LeafDescription,
		&door.FillingDescription,
		&door.MainLock,
		&door.AdditionalLock,
		&door.InsulationDescription,
		&door.Hinges,
		&door.ProductName,
		&door.Price,
		&door.Description,
		&door.CategoryCode,
		&door.IsVisible,
		&door.Version,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.EntranceDoor{}, fmt.Errorf("GetEntranceDoor from repo: %v: %w", err, core_errors.ErrNotFound)
		}
		return domain.EntranceDoor{}, fmt.Errorf("GetEntranceDoor from repo: %w", err)
	}

	return door, nil
}

func (r *ProductsRepositoryPostgres) GetInteriorDoor(ctx context.Context, id uuid.UUID) (domain.InteriorDoor, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	query := `
        SELECT 
            d.door_id,
            d.collection,
            d.width,
            d.height,
            d.outside_material,
            d.outside_color,
            d.outside_picture,
            d.inside_material,
            d.inside_color,
            d.inside_picture,
            idoor.opening_system,
            idoor.leaf_coating,
            idoor.handle,
            p.product_name,
            p.price,
            p.description,
            p.category_code,
            p.is_visible,
            p.version
        FROM kudesnik.doors d
        INNER JOIN kudesnik.products p ON d.door_id = p.product_id
        INNER JOIN kudesnik.interior_doors idoor ON d.door_id = idoor.door_id
        WHERE d.door_id = $1
    `

	row := r.pool.QueryRow(ctx, query, id)

	var door domain.InteriorDoor

	err := row.Scan(
		&door.ID,
		&door.Collection,
		&door.Width,
		&door.Height,
		&door.OutsideMaterial,
		&door.OutsideColor,
		&door.OutsidePicture,
		&door.InsideMaterial,
		&door.InsideColor,
		&door.InsidePicture,
		&door.OpeningSystem,
		&door.LeafCoating,
		&door.Handle,
		&door.ProductName,
		&door.Price,
		&door.Description,
		&door.CategoryCode,
		&door.IsVisible,
		&door.Version,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.InteriorDoor{}, fmt.Errorf("GetInteriorDoor from repo: %v: %w", err, core_errors.ErrNotFound)
		}
		return domain.InteriorDoor{}, fmt.Errorf("GetInteriorDoor from repo: %w", err)
	}

	return door, nil
}
