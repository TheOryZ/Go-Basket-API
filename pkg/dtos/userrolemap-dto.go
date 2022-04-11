package dtos

import "github.com/gofrs/uuid"

type UserRoleMapListDTO struct {
	ID   uuid.UUID   `json:"id"`
	User UserListDTO `json:"user"`
	Role RoleListDTO `json:"role"`
}

type UserRoleMapCreateDTO struct {
	UserID uuid.UUID `json:"userId" form:"userId" binding:"required"`
	RoleID uuid.UUID `json:"roleId" form:"roleId" binding:"required"`
}

type UserRoleMapUpdateDTO struct {
	ID     uuid.UUID `json:"id" form:"id" binding:"required"`
	UserID uuid.UUID `json:"userId" form:"userId" binding:"required"`
	RoleID uuid.UUID `json:"roleId" form:"roleId" binding:"required"`
}
