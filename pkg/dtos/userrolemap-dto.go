package dtos

import "github.com/gofrs/uuid"

//UserRoleMapListDTO is a struct for listing userrolemaps
type UserRoleMapListDTO struct {
	ID   uuid.UUID   `json:"id"`
	User UserListDTO `json:"user"`
	Role RoleListDTO `json:"role"`
}

//UserRoleMapCreateDTO is a struct for creating a new userrolemap
type UserRoleMapCreateDTO struct {
	UserID uuid.UUID `json:"userId" form:"userId" binding:"required"`
	RoleID uuid.UUID `json:"roleId" form:"roleId" binding:"required"`
}

//UserRoleMapUpdateDTO is a struct for updating a userrolemap
type UserRoleMapUpdateDTO struct {
	ID     uuid.UUID `json:"id" form:"id" binding:"required"`
	UserID uuid.UUID `json:"userId" form:"userId" binding:"required"`
	RoleID uuid.UUID `json:"roleId" form:"roleId" binding:"required"`
}
