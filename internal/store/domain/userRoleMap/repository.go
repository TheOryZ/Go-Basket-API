package userrolemap

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type userRoleMapRepository struct {
	db *gorm.DB
}

//interface
type IUserRoleMapRepository interface {
	Migration()
	Create(userrolemap *UserRoleMap) error
	Update(userrolemap *UserRoleMap) error
	Delete(userrolemap *UserRoleMap) error
	DeleteByID(id uuid.UUID) error
	FindAll() ([]UserRoleMap, error)
	FindByID(id uuid.UUID) (*UserRoleMap, error)
	FindByUserID(userid uuid.UUID) (*[]UserRoleMap, error)
	FindByRoleID(roleid uuid.UUID) (*[]UserRoleMap, error)
}

var UserRoleMapRepository IUserRoleMapRepository = &userRoleMapRepository{}

func NewUserRoleMapRepository(db *gorm.DB) *userRoleMapRepository {
	return &userRoleMapRepository{db: db}
}
func (r *userRoleMapRepository) Migration() {
	r.db.AutoMigrate(&UserRoleMap{})
}

//Create a new userrolemap
func (r *userRoleMapRepository) Create(userrolemap *UserRoleMap) error {
	return r.db.Create(userrolemap).Error
}

//Update a userrolemap
func (r *userRoleMapRepository) Update(userrolemap *UserRoleMap) error {
	return r.db.Save(userrolemap).Error
}

//Delete a userrolemap
func (r *userRoleMapRepository) Delete(userrolemap *UserRoleMap) error {
	return r.db.Delete(userrolemap).Error
}

//Delete a userrolemap by id
func (r *userRoleMapRepository) DeleteByID(id uuid.UUID) error {
	userrolemap := UserRoleMap{}
	userrolemap.ID = id
	return r.db.Delete(&userrolemap).Error
}

//Find all userrolemaps
func (r *userRoleMapRepository) FindAll() ([]UserRoleMap, error) {
	var userrolemaps []UserRoleMap
	err := r.db.Find(&userrolemaps).Error
	return userrolemaps, err
}

//Find a userrolemap by id
func (r *userRoleMapRepository) FindByID(id uuid.UUID) (*UserRoleMap, error) {
	userrolemap := UserRoleMap{}
	err := r.db.First(&userrolemap, id).Error
	return &userrolemap, err
}

//Find userrolemap by userid
func (r *userRoleMapRepository) FindByUserID(userid uuid.UUID) (*[]UserRoleMap, error) {
	userrolemap := []UserRoleMap{}
	err := r.db.Where("user_id = ? AND deleted_at = ?", userid, nil).Find(&userrolemap).Error
	return &userrolemap, err
}

//Find userrolemap by roleid
func (r *userRoleMapRepository) FindByRoleID(roleid uuid.UUID) (*[]UserRoleMap, error) {
	userrolemap := []UserRoleMap{}
	err := r.db.Where("role_id = ? AND deleted_at = ?", roleid, nil).Find(&userrolemap).Error
	return &userrolemap, err
}

//Seed a userrolemap
// func (r *UserRoleMapRepository) Seed() error {
// 	userrolemaps := []UserRoleMap{
// 		{
// 			UserID:   1,
// 			RoleID:   1,
// 			IsActive: true,
// 		},
// 		{
// 			UserID:   1,
// 			RoleID:   2,
// 			IsActive: true,
// 		},
// 	}
// 	for _, userrolemap := range userrolemaps {
// 		err := r.db.Create(&userrolemap).Error
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
