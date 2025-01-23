package models

type Role struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// Role permissions
	CanCreateRole bool `json:"can_create_role"`
	CanReadRole bool `json:"can_read_role"`
	CanUpdateRole bool `json:"can_update_role"`
	CanDeleteRole bool `json:"can_delete_role"`
	// User permissions
	CanCreateUser bool `json:"can_create_user"`
	CanReadUser bool `json:"can_read_user"`
	CanUpdateUser bool `json:"can_update_user"`
	CanDeleteUser bool `json:"can_delete_user"`
}
