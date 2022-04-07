package user

import "gorm.io/gorm"

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
func (r *UserRepository) Migration() {
	r.db.AutoMigrate(&User{})
}

//Create a new user
func (r *UserRepository) Create(user *User) error {
	return r.db.Create(user).Error
}

//Update a user
func (r *UserRepository) Update(user *User) error {
	return r.db.Save(user).Error
}

//Delete a user
func (r *UserRepository) Delete(user *User) error {
	return r.db.Delete(user).Error
}

//Delete a user by id
func (r *UserRepository) DeleteByID(id uint) error {
	user := User{}
	user.ID = id
	return r.db.Delete(&user).Error
}

//Find all users
func (r *UserRepository) FindAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

//Find a user by id
func (r *UserRepository) FindByID(id uint) (*User, error) {
	user := User{}
	err := r.db.First(&user, id).Error
	return &user, err
}

//Find a user by name
func (r *UserRepository) FindByName(name string) (*User, error) {
	user := User{}
	err := r.db.Where("name = ?", name).First(&user).Error
	return &user, err
}

//Search users by name or email
func (r *UserRepository) Search(s string) ([]User, error) {
	var users []User
	err := r.db.Where("name LIKE ? OR email LIKE ?", "%"+s+"%", "%"+s+"%").Find(&users).Error
	return users, err
}

//Get with roles relationship //TODO: Check this
func (r *UserRepository) GetWithRoles(id uint) (*User, error) {
	user := User{}
	err := r.db.Preload("Roles").First(&user, id).Error
	return &user, err
}

//Seed a user
func (r *UserRepository) Seed() {
	users := []User{
		{Name: "John Doe", Email: "johnDoe@gmail.com", Password: "123456", IsActive: true},
	}
	for _, user := range users {
		r.Create(&user)
	}
}
