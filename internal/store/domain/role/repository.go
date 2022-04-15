package role

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

//interface
type IRoleRepository interface {
	Migration()
	Create(role *Role) error
	Update(role *Role) error
	Delete(role *Role) error
	DeleteByID(id uuid.UUID) error
	FindAll() ([]Role, error)
	FindByID(id uuid.UUID) (*Role, error)
	FindByUserID(id uuid.UUID) (*[]Role, error)
	FindByName(name string) (*Role, error)
	Seed() error
}

var RoleRepository IRoleRepository = &roleRepository{}

func NewRoleRepository(db *gorm.DB) *roleRepository {
	return &roleRepository{db: db}
}
func (r *roleRepository) Migration() {
	r.db.AutoMigrate(&Role{})
}

//Create a new role
func (r *roleRepository) Create(role *Role) error {
	return r.db.Create(role).Error
}

//Update a role
func (r *roleRepository) Update(role *Role) error {
	return r.db.Save(role).Error
}

//Delete a role
func (r *roleRepository) Delete(role *Role) error {
	return r.db.Delete(role).Error
}

//Delete a role by id
func (r *roleRepository) DeleteByID(id uuid.UUID) error {
	role := Role{}
	role.ID = id
	return r.db.Delete(&role).Error
}

//Find all roles
func (r *roleRepository) FindAll() ([]Role, error) {
	var roles []Role
	err := r.db.Find(&roles).Error
	return roles, err
}

//Find a role by id
func (r *roleRepository) FindByID(id uuid.UUID) (*Role, error) {
	role := Role{}
	err := r.db.First(&role, id).Error
	return &role, err
}

//Find a role by user id
func (r *roleRepository) FindByUserID(id uuid.UUID) (*[]Role, error) {
	role := []Role{}
	//err := r.db.Preload("Users").First(&role, id).Error
	err := r.db.Joins("INNER JOIN user_role_map map on map.role_id = roles.id").
		Where("map.user_id = ?", id).
		Table("roles").
		Select("roles.*").
		Find(&role).Error
	return &role, err
}

//Find a role by name
func (r *roleRepository) FindByName(name string) (*Role, error) {
	role := Role{}
	err := r.db.Where("name = ?", name).First(&role).Error
	return &role, err
}

//Seed a role
func (r *roleRepository) Seed() error {
	roles := []Role{
		{Name: "admin", CreatedAt: "2020-01-01 00:00:00", UpdatedAt: "2020-01-01 00:00:00"},
		{Name: "member", CreatedAt: "2020-01-01 00:00:00", UpdatedAt: "2020-01-01 00:00:00"},
	}
	for _, role := range roles {
		if err := r.db.Create(&role).Error; err != nil {
			return err
		}
	}
	return nil
}
