package domain

type Role string

const (
	AdminRole          Role = "superadmin"
	ManagerRole        Role = "manager"
	DismissedAdminRole Role = "dismissed"
)
