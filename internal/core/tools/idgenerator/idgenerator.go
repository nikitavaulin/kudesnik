package core_tools_idgenerator

import (
	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

type IDGenerator struct{}

func (g IDGenerator) New() domain.ID {
	return uuid.New()
}
