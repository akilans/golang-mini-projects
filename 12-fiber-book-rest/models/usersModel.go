package models

// User Type -> userss table
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name" validate:"required,min=1,max=100"`
	Email    string `json:"email" validate:"required,email,min=6,max=100" gorm:"unique"`
	Password string `json:"password" validate:"required,min=6,max=15"`
}

// add user
func AddUser(user User) (id int, err error) {
	result := db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	} else {
		return user.ID, nil
	}
}

// add user
func GetUserByEmail(email string) (User, error) {
	var user User
	result := db.Where("email = ?", email).Find(&user)
	if result.Error != nil {
		return User{}, result.Error
	} else {
		return user, nil
	}
}
