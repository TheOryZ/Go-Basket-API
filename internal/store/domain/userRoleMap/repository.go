package userrolemap

import "gorm.io/gorm"

type UserRoleMapRepository struct {
	db *gorm.DB
}

func NewUserRoleMapRepository(db *gorm.DB) *UserRoleMapRepository {
	return &UserRoleMapRepository{db: db}
}
func (r *UserRoleMapRepository) Migration() {
	r.db.AutoMigrate(&UserRoleMap{})
}

//Create a new userrolemap
func (r *UserRoleMapRepository) Create(userrolemap *UserRoleMap) error {
	return r.db.Create(userrolemap).Error
}

//Update a userrolemap
func (r *UserRoleMapRepository) Update(userrolemap *UserRoleMap) error {
	return r.db.Save(userrolemap).Error
}

//Delete a userrolemap
func (r *UserRoleMapRepository) Delete(userrolemap *UserRoleMap) error {
	return r.db.Delete(userrolemap).Error
}

//Delete a userrolemap by id
func (r *UserRoleMapRepository) DeleteByID(id uint) error {
	userrolemap := UserRoleMap{}
	userrolemap.ID = id
	return r.db.Delete(&userrolemap).Error
}

//Find all userrolemaps
func (r *UserRoleMapRepository) FindAll() ([]UserRoleMap, error) {
	var userrolemaps []UserRoleMap
	err := r.db.Find(&userrolemaps).Error
	return userrolemaps, err
}

//Find a userrolemap by id
func (r *UserRoleMapRepository) FindByID(id uint) (*UserRoleMap, error) {
	userrolemap := UserRoleMap{}
	err := r.db.First(&userrolemap, id).Error
	return &userrolemap, err
}

//Find userrolemap by userid
func (r *UserRoleMapRepository) FindByUserID(userid uint) (*[]UserRoleMap, error) {
	userrolemap := []UserRoleMap{}
	err := r.db.Where("user_id = ?", userid).Find(&userrolemap).Error
	return &userrolemap, err
}

//Find userrolemap by roleid
func (r *UserRoleMapRepository) FindByRoleID(roleid uint) (*[]UserRoleMap, error) {
	userrolemap := []UserRoleMap{}
	err := r.db.Where("role_id = ?", roleid).Find(&userrolemap).Error
	return &userrolemap, err
}

//Seed a userrolemap
func (r *UserRoleMapRepository) Seed() error {
	userrolemaps := []UserRoleMap{
		{
			UserID:   1,
			RoleID:   1,
			IsActive: true,
		},
		{
			UserID:   1,
			RoleID:   2,
			IsActive: true,
		},
	}
	for _, userrolemap := range userrolemaps {
		err := r.db.Create(&userrolemap).Error
		if err != nil {
			return err
		}
	}
	return nil
}
