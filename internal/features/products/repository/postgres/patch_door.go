// patch_doors.go
package products_repository_postges

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_postgres_pool "github.com/nikitavaulin/kudesnik/internal/core/repository/postgres/pool"
)

// Общая функция для обновления базовых полей двери
func (r *ProductsRepositoryPostgres) updateDoorBaseInTx(
	ctx context.Context,
	tx core_postgres_pool.Tx,
	id uuid.UUID,
	door domain.DoorBase,
) error {
	query := `
		UPDATE kudesnik.doors
		SET 
			collection = $2,
			width = $3,
			height = $4,
			outside_material = $5,
			outside_color = $6,
			outside_picture = $7,
			inside_material = $8,
			inside_color = $9,
			inside_picture = $10
		WHERE door_id = $1;
	`

	_, err := tx.Exec(
		ctx, query,
		id,
		door.Collection,
		door.Width,
		door.Height,
		door.OutsideMaterial,
		door.OutsideColor,
		door.OutsidePicture,
		door.InsideMaterial,
		door.InsideColor,
		door.InsidePicture,
	)

	return err
}

// Функция для обновления входной двери
func (r *ProductsRepositoryPostgres) PatchEntranceDoor(
	ctx context.Context,
	id uuid.UUID,
	door domain.EntranceDoor,
) (domain.EntranceDoor, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return domain.EntranceDoor{}, fmt.Errorf("PatchEntranceDoor (begin tx): %w", err)
	}
	defer tx.Rollback(ctx)

	updatedProduct, err := r.updateProductBaseInTx(ctx, tx, id, door.ProductBase, door.Version)
	if err != nil {
		return domain.EntranceDoor{}, fmt.Errorf("PatchEntranceDoor (update product): %w", err)
	}

	if err := r.updateDoorBaseInTx(ctx, tx, id, door.DoorBase); err != nil {
		return domain.EntranceDoor{}, fmt.Errorf("PatchEntranceDoor (update door base): %w", err)
	}

	updateEntranceDoorQuery := `
		UPDATE kudesnik.entrance_doors
		SET 
			strength_class = $2,
			sound_insulation = $3,
			metal_thickness = $4,
			box_thickness = $5,
			leaf_thickness = $6,
			leaf_description = $7,
			filling_description = $8,
			main_lock = $9,
			additional_lock = $10,
			insulation_description = $11,
			hinges = $12
		WHERE door_id = $1
		RETURNING 
			strength_class,
			sound_insulation,
			metal_thickness,
			box_thickness,
			leaf_thickness,
			leaf_description,
			filling_description,
			main_lock,
			additional_lock,
			insulation_description,
			hinges;
	`

	var entranceSpecific struct {
		StrengthClass         string
		SoundInsulation       string
		MetalThickness        string
		BoxThickness          string
		LeafThickness         string
		LeafDescription       string
		FillingDescription    string
		MainLock              string
		AdditionalLock        string
		InsulationDescription string
		Hinges                string
	}

	err = tx.QueryRow(
		ctx, updateEntranceDoorQuery,
		id,
		door.StrengthClass,
		door.SoundInsulation,
		door.MetalThickness,
		door.BoxThickness,
		door.LeafThickness,
		door.LeafDescription,
		door.FillingDescription,
		door.MainLock,
		door.AdditionalLock,
		door.InsulationDescription,
		door.Hinges,
	).Scan(
		&entranceSpecific.StrengthClass,
		&entranceSpecific.SoundInsulation,
		&entranceSpecific.MetalThickness,
		&entranceSpecific.BoxThickness,
		&entranceSpecific.LeafThickness,
		&entranceSpecific.LeafDescription,
		&entranceSpecific.FillingDescription,
		&entranceSpecific.MainLock,
		&entranceSpecific.AdditionalLock,
		&entranceSpecific.InsulationDescription,
		&entranceSpecific.Hinges,
	)

	if err != nil {
		return domain.EntranceDoor{}, fmt.Errorf("PatchEntranceDoor (update entrance_doors): %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return domain.EntranceDoor{}, fmt.Errorf("PatchEntranceDoor (commit): %w", err)
	}

	// Собираем результат
	patchedDoor := domain.EntranceDoor{
		ProductBase:           updatedProduct,
		DoorBase:              door.DoorBase,
		StrengthClass:         entranceSpecific.StrengthClass,
		SoundInsulation:       entranceSpecific.SoundInsulation,
		MetalThickness:        entranceSpecific.MetalThickness,
		BoxThickness:          entranceSpecific.BoxThickness,
		LeafThickness:         entranceSpecific.LeafThickness,
		LeafDescription:       entranceSpecific.LeafDescription,
		FillingDescription:    entranceSpecific.FillingDescription,
		MainLock:              entranceSpecific.MainLock,
		AdditionalLock:        entranceSpecific.AdditionalLock,
		InsulationDescription: entranceSpecific.InsulationDescription,
		Hinges:                entranceSpecific.Hinges,
	}

	return patchedDoor, nil
}

// Функция для обновления межкомнатной двери
func (r *ProductsRepositoryPostgres) PatchInteriorDoor(
	ctx context.Context,
	id uuid.UUID,
	door domain.InteriorDoor,
) (domain.InteriorDoor, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return domain.InteriorDoor{}, fmt.Errorf("PatchInteriorDoor (begin tx): %w", err)
	}
	defer tx.Rollback(ctx)

	updatedProduct, err := r.updateProductBaseInTx(ctx, tx, id, door.ProductBase, door.Version)
	if err != nil {
		return domain.InteriorDoor{}, fmt.Errorf("PatchInteriorDoor (update product): %w", err)
	}

	if err := r.updateDoorBaseInTx(ctx, tx, id, door.DoorBase); err != nil {
		return domain.InteriorDoor{}, fmt.Errorf("PatchInteriorDoor (update door base): %w", err)
	}

	updateInteriorDoorQuery := `
		UPDATE kudesnik.interior_doors
		SET 
			opening_system = $2,
			leaf_coating = $3,
			handle = $4
		WHERE door_id = $1
		RETURNING 
			opening_system,
			leaf_coating,
			handle;
	`

	var interiorSpecific struct {
		OpeningSystem string
		LeafCoating   string
		Handle        string
	}

	err = tx.QueryRow(
		ctx, updateInteriorDoorQuery,
		id,
		door.OpeningSystem,
		door.LeafCoating,
		door.Handle,
	).Scan(
		&interiorSpecific.OpeningSystem,
		&interiorSpecific.LeafCoating,
		&interiorSpecific.Handle,
	)

	if err != nil {
		return domain.InteriorDoor{}, fmt.Errorf("PatchInteriorDoor (update interior_doors): %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return domain.InteriorDoor{}, fmt.Errorf("PatchInteriorDoor (commit): %w", err)
	}

	patchedDoor := domain.InteriorDoor{
		ProductBase:   updatedProduct,
		DoorBase:      door.DoorBase,
		OpeningSystem: interiorSpecific.OpeningSystem,
		LeafCoating:   interiorSpecific.LeafCoating,
		Handle:        interiorSpecific.Handle,
	}

	return patchedDoor, nil
}
