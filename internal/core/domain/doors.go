package domain

import (
	"fmt"

	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
	core_validation "github.com/nikitavaulin/kudesnik/internal/core/tools/validation"
)

const (
	MinDoorCollectionLength = 0
	MaxDoorCollectionLength = 60
)

type DoorBase struct {
	Collection      string `json:"collection"`
	Width           int    `json:"width"`
	Height          int    `json:"height"`
	OutsideMaterial string `json:"outside_material"`
	OutsideColor    string `json:"outside_color"`
	OutsidePicture  string `json:"outside_picture"`
	InsideMaterial  string `json:"inside_material"`
	InsideColor     string `json:"inside_color"`
	InsidePicture   string `json:"inside_picture"`
}

type EntranceDoor struct {
	ProductBase
	DoorBase
	StrengthClass         string `json:"strength_class"`
	SoundInsulation       string `json:"sound_insulation"`
	MetalThickness        string `json:"metal_thickness"`
	BoxThickness          string `json:"box_thickness"`
	LeafThickness         string `json:"leaf_thickness"`
	LeafDescription       string `json:"leaf_description"`
	FillingDescription    string `json:"filling_description"`
	MainLock              string `json:"main_lock"`
	AdditionalLock        string `json:"additional_lock"`
	InsulationDescription string `json:"insulation_description"`
	Hinges                string `json:"hinges"`
}

type InteriorDoor struct {
	ProductBase
	DoorBase
	OpeningSystem string `json:"opening_system"`
	LeafCoating   string `json:"leaf_coating"`
	Handle        string `json:"handle"`
}

func (d *DoorBase) Validate() error {
	if err := core_validation.ValidateIntInBounds(len(d.Collection), MinDoorCollectionLength, MaxDoorCollectionLength); err != nil {
		return fmt.Errorf("invalid length door collection: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	if d.Width < 0 {
		return fmt.Errorf("width cannot be negative: %w", core_errors.ErrInvalidArgument)
	}
	if d.Height < 0 {
		return fmt.Errorf("height cannot be negative: %w", core_errors.ErrInvalidArgument)
	}
	return nil
}

func (d *EntranceDoor) GetBase() *ProductBase                { return &d.ProductBase }
func (d *EntranceDoor) GetCategoryName() ProductCategoryCode { return EntranceDoorsCategory }

func (d *EntranceDoor) Validate() error {
	if err := d.DoorBase.Validate(); err != nil {
		return err
	}
	return nil
}

func (d *InteriorDoor) GetBase() *ProductBase                { return &d.ProductBase }
func (d *InteriorDoor) GetCategoryName() ProductCategoryCode { return InteriorDoorsCategory }

func (d *InteriorDoor) Validate() error {
	if err := d.DoorBase.Validate(); err != nil {
		return err
	}
	return nil
}

type DoorBasePatch struct {
	Collection      Nullable[string] `json:"collection" swaggertype:"string"`
	Width           Nullable[int]    `json:"width" swaggertype:"integer"`
	Height          Nullable[int]    `json:"height" swaggertype:"integer"`
	OutsideMaterial Nullable[string] `json:"outside_material" swaggertype:"string"`
	OutsideColor    Nullable[string] `json:"outside_color" swaggertype:"string"`
	OutsidePicture  Nullable[string] `json:"outside_picture" swaggertype:"string"`
	InsideMaterial  Nullable[string] `json:"inside_material" swaggertype:"string"`
	InsideColor     Nullable[string] `json:"inside_color" swaggertype:"string"`
	InsidePicture   Nullable[string] `json:"inside_picture" swaggertype:"string"`
}

// Patch структура для EntranceDoor
type EntranceDoorPatch struct {
	ProductBasePatch
	DoorBasePatch
	StrengthClass         Nullable[string] `json:"strength_class" swaggertype:"string"`
	SoundInsulation       Nullable[string] `json:"sound_insulation" swaggertype:"string"`
	MetalThickness        Nullable[string] `json:"metal_thickness" swaggertype:"string"`
	BoxThickness          Nullable[string] `json:"box_thickness" swaggertype:"string"`
	LeafThickness         Nullable[string] `json:"leaf_thickness" swaggertype:"string"`
	LeafDescription       Nullable[string] `json:"leaf_description" swaggertype:"string"`
	FillingDescription    Nullable[string] `json:"filling_description" swaggertype:"string"`
	MainLock              Nullable[string] `json:"main_lock" swaggertype:"string"`
	AdditionalLock        Nullable[string] `json:"additional_lock" swaggertype:"string"`
	InsulationDescription Nullable[string] `json:"insulation_description" swaggertype:"string"`
	Hinges                Nullable[string] `json:"hinges" swaggertype:"string"`
}

// Patch структура для InteriorDoor
type InteriorDoorPatch struct {
	ProductBasePatch
	DoorBasePatch
	OpeningSystem Nullable[string] `json:"opening_system" swaggertype:"string"`
	LeafCoating   Nullable[string] `json:"leaf_coating" swaggertype:"string"`
	Handle        Nullable[string] `json:"handle" swaggertype:"string"`
}

func (d *DoorBasePatch) Validate() error {
	if d.Collection.Set {
		if err := core_validation.ValidateIntInBounds(len(*d.Collection.Value), MinDoorCollectionLength, MaxDoorCollectionLength); err != nil {
			return fmt.Errorf("invalid length door collection: %v: %w", err, core_errors.ErrInvalidArgument)
		}
	}

	if d.Width.Set && *d.Width.Value < 0 {
		return fmt.Errorf("width cannot be negative: %w", core_errors.ErrInvalidArgument)
	}

	if d.Height.Set && *d.Height.Value < 0 {
		return fmt.Errorf("height cannot be negative: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}

func (d *DoorBase) ApplyPatch(patch *DoorBasePatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("invalid door base patch: %w", err)
	}

	tmp := *d

	if patch.Collection.Set {
		tmp.Collection = *patch.Collection.Value
	}

	if patch.Width.Set {
		tmp.Width = *patch.Width.Value
	}

	if patch.Height.Set {
		tmp.Height = *patch.Height.Value
	}

	if patch.OutsideMaterial.Set {
		tmp.OutsideMaterial = *patch.OutsideMaterial.Value
	}

	if patch.OutsideColor.Set {
		tmp.OutsideColor = *patch.OutsideColor.Value
	}

	if patch.OutsidePicture.Set {
		tmp.OutsidePicture = *patch.OutsidePicture.Value
	}

	if patch.InsideMaterial.Set {
		tmp.InsideMaterial = *patch.InsideMaterial.Value
	}

	if patch.InsideColor.Set {
		tmp.InsideColor = *patch.InsideColor.Value
	}

	if patch.InsidePicture.Set {
		tmp.InsidePicture = *patch.InsidePicture.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("invalid door base after patch: %w", err)
	}

	*d = tmp

	return nil
}

func (d *EntranceDoorPatch) Validate() error {
	if err := d.ProductBasePatch.Validate(); err != nil {
		return fmt.Errorf("product base validation failed: %w", err)
	}

	if err := d.DoorBasePatch.Validate(); err != nil {
		return fmt.Errorf("door base validation failed: %w", err)
	}

	return nil
}

func (d *EntranceDoor) ApplyPatch(patch *EntranceDoorPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("invalid entrance door patch: %w", err)
	}

	tmp := *d

	if err := tmp.ProductBase.ApplyPatch(&patch.ProductBasePatch); err != nil {
		return fmt.Errorf("failed to patch product base: %w", err)
	}

	if err := tmp.DoorBase.ApplyPatch(&patch.DoorBasePatch); err != nil {
		return fmt.Errorf("failed to patch door base: %w", err)
	}

	if patch.StrengthClass.Set {
		tmp.StrengthClass = *patch.StrengthClass.Value
	}

	if patch.SoundInsulation.Set {
		tmp.SoundInsulation = *patch.SoundInsulation.Value
	}

	if patch.MetalThickness.Set {
		tmp.MetalThickness = *patch.MetalThickness.Value
	}

	if patch.BoxThickness.Set {
		tmp.BoxThickness = *patch.BoxThickness.Value
	}

	if patch.LeafThickness.Set {
		tmp.LeafThickness = *patch.LeafThickness.Value
	}

	if patch.LeafDescription.Set {
		tmp.LeafDescription = *patch.LeafDescription.Value
	}

	if patch.FillingDescription.Set {
		tmp.FillingDescription = *patch.FillingDescription.Value
	}

	if patch.MainLock.Set {
		tmp.MainLock = *patch.MainLock.Value
	}

	if patch.AdditionalLock.Set {
		tmp.AdditionalLock = *patch.AdditionalLock.Value
	}

	if patch.InsulationDescription.Set {
		tmp.InsulationDescription = *patch.InsulationDescription.Value
	}

	if patch.Hinges.Set {
		tmp.Hinges = *patch.Hinges.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("invalid entrance door after patch: %w", err)
	}

	*d = tmp

	return nil
}

func (d *InteriorDoorPatch) Validate() error {
	if err := d.ProductBasePatch.Validate(); err != nil {
		return fmt.Errorf("product base validation failed: %w", err)
	}

	if err := d.DoorBasePatch.Validate(); err != nil {
		return fmt.Errorf("door base validation failed: %w", err)
	}

	return nil
}

func (d *InteriorDoor) ApplyPatch(patch *InteriorDoorPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("invalid interior door patch: %w", err)
	}

	tmp := *d

	if err := tmp.ProductBase.ApplyPatch(&patch.ProductBasePatch); err != nil {
		return fmt.Errorf("failed to patch product base: %w", err)
	}

	if err := tmp.DoorBase.ApplyPatch(&patch.DoorBasePatch); err != nil {
		return fmt.Errorf("failed to patch door base: %w", err)
	}

	if patch.OpeningSystem.Set {
		tmp.OpeningSystem = *patch.OpeningSystem.Value
	}

	if patch.LeafCoating.Set {
		tmp.LeafCoating = *patch.LeafCoating.Value
	}

	if patch.Handle.Set {
		tmp.Handle = *patch.Handle.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("invalid interior door after patch: %w", err)
	}

	*d = tmp

	return nil
}

func (p *EntranceDoorPatch) isProductPatch() {}
func (p *InteriorDoorPatch) isProductPatch() {}
