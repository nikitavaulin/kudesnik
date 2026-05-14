package admin_transport_http

import "github.com/nikitavaulin/kudesnik/internal/core/domain"

type AdminResponseDTO struct {
	ID        string      `json:"id"`
	Email     string      `json:"email"`
	FullName  string      `json:"full_name"`
	AdminType domain.Role `json:"admin_type"`
}

type AdminsRepsonseDTO []AdminResponseDTO

func toAdminResponseDTO(admin domain.Admin) AdminResponseDTO {
	return AdminResponseDTO{
		ID:        admin.ID.String(),
		Email:     admin.Email,
		FullName:  admin.FullName,
		AdminType: admin.AdminType,
	}
}

func toAdminsResponseDTO(admins []domain.Admin) []AdminResponseDTO {
	dtos := make([]AdminResponseDTO, len(admins))
	for i, admin := range admins {
		dtos[i] = toAdminResponseDTO(admin)
	}
	return dtos
}
