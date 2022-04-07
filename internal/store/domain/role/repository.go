package role

import "gorm.io/gorm"

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db: db}
}
func (r *RoleRepository) Migration() {
	r.db.AutoMigrate(&Role{})
}

//Create a new role
func (r *RoleRepository) Create(role *Role) error {
	return r.db.Create(role).Error
}

//Update a role
func (r *RoleRepository) Update(role *Role) error {
	return r.db.Save(role).Error
}

//Delete a role
func (r *RoleRepository) Delete(role *Role) error {
	return r.db.Delete(role).Error
}

//Delete a role by id
func (r *RoleRepository) DeleteByID(id uint) error {
	role := Role{}
	role.ID = id
	return r.db.Delete(&role).Error
}

//Find all roles
func (r *RoleRepository) FindAll() ([]Role, error) {
	var roles []Role
	err := r.db.Find(&roles).Error
	return roles, err
}

//Find a role by id
func (r *RoleRepository) FindByID(id uint) (*Role, error) {
	role := Role{}
	err := r.db.First(&role, id).Error
	return &role, err
}

//Find a role by name
func (r *RoleRepository) FindByName(name string) (*Role, error) {
	role := Role{}
	err := r.db.Where("name = ?", name).First(&role).Error
	return &role, err
}

//Get with users relationship  //TODO: Check this
func (r *RoleRepository) GetWithUsers(id uint) (*Role, error) {
	role := Role{}
	err := r.db.Preload("Users").First(&role, id).Error
	return &role, err
}

//Seed a role
func (r *RoleRepository) Seed() error {
	roles := []Role{
		{Name: "admin"},
		{Name: "member"},
	}
	for _, role := range roles {
		if err := r.db.Create(&role).Error; err != nil {
			return err
		}
	}
	return nil
}
