package products_repository_postges

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_postgres_pool "github.com/nikitavaulin/kudesnik/internal/core/repository/postgres/pool"
)

func (r *ProductsRepositoryPostgres) CreateEntranceDoor(ctx context.Context, door domain.EntranceDoor) (domain.EntranceDoor, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return domain.EntranceDoor{}, fmt.Errorf("CreateEntranceDoor (begin tx): %w", err)
	}
	defer tx.Rollback(ctx)

	productID, version, err := r.createProductInTx(ctx, tx, door.ProductBase)
	if err != nil {
		return domain.EntranceDoor{}, fmt.Errorf("CreateEntranceDoor (insert product): %w", err)
	}

	if err := r.createDoorInTx(ctx, tx, productID, door.DoorBase); err != nil {
		return domain.EntranceDoor{}, fmt.Errorf("CreateEntranceDoor (insert door): %w", err)
	}

	if err := r.createEntranceDoorDetailsInTx(ctx, tx, productID, door); err != nil {
		return domain.EntranceDoor{}, fmt.Errorf("CreateEntranceDoor (insert entrance details): %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return domain.EntranceDoor{}, fmt.Errorf("CreateEntranceDoor (commit): %w", err)
	}

	door.ID = productID
	door.Version = version

	return door, nil
}

func (r *ProductsRepositoryPostgres) CreateInteriorDoor(ctx context.Context, door domain.InteriorDoor) (domain.InteriorDoor, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return domain.InteriorDoor{}, fmt.Errorf("CreateInteriorDoor (begin tx): %w", err)
	}
	defer tx.Rollback(ctx)

	productID, version, err := r.createProductInTx(ctx, tx, door.ProductBase)
	if err != nil {
		return domain.InteriorDoor{}, fmt.Errorf("CreateInteriorDoor (insert product): %w", err)
	}

	if err := r.createDoorInTx(ctx, tx, productID, door.DoorBase); err != nil {
		return domain.InteriorDoor{}, fmt.Errorf("CreateInteriorDoor (insert door): %w", err)
	}

	if err := r.createInteriorDoorDetailsInTx(ctx, tx, productID, door); err != nil {
		return domain.InteriorDoor{}, fmt.Errorf("CreateInteriorDoor (insert interior details): %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return domain.InteriorDoor{}, fmt.Errorf("CreateInteriorDoor (commit): %w", err)
	}

	door.ID = productID
	door.Version = version

	return door, nil
}

func (r *ProductsRepositoryPostgres) createDoorInTx(ctx context.Context, tx core_postgres_pool.Tx, doorID uuid.UUID, door domain.DoorBase) error {
	doorQuery := `
		INSERT INTO kudesnik.doors (
			door_id, collection, width, height,
			outside_material, outside_color, outside_picture,
			inside_material, inside_color, inside_picture
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);
	`

	_, err := tx.Exec(
		ctx, doorQuery,
		doorID, door.Collection, door.Width, door.Height,
		door.OutsideMaterial, door.OutsideColor, door.OutsidePicture,
		door.InsideMaterial, door.InsideColor, door.InsidePicture,
	)
	return err
}

func (r *ProductsRepositoryPostgres) createEntranceDoorDetailsInTx(ctx context.Context, tx core_postgres_pool.Tx, doorID uuid.UUID, door domain.EntranceDoor) error {
	query := `
		INSERT INTO kudesnik.entrance_doors (
			door_id, strength_class, sound_insulation, metal_thickness, box_thickness,
			leaf_thickness, leaf_description, filling_description, main_lock,
			additional_lock, insulation_description, hinges
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);
	`

	_, err := tx.Exec(
		ctx, query,
		doorID, door.StrengthClass, door.SoundInsulation,
		door.MetalThickness, door.BoxThickness,
		door.LeafThickness, door.LeafDescription, door.FillingDescription,
		door.MainLock, door.AdditionalLock, door.InsulationDescription, door.Hinges,
	)
	return err
}

func (r *ProductsRepositoryPostgres) createInteriorDoorDetailsInTx(ctx context.Context, tx core_postgres_pool.Tx, doorID uuid.UUID, door domain.InteriorDoor) error {
	query := `
		INSERT INTO kudesnik.interior_doors (
			door_id, opening_system, leaf_coating, handle
		) VALUES ($1, $2, $3, $4);
	`

	_, err := tx.Exec(
		ctx, query,
		doorID, door.OpeningSystem, door.LeafCoating, door.Handle,
	)
	return err
}
