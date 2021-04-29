package blog_m

import (
	"errors"
	"fmt"
	"github.com/99-66/go-gin-example-blog-api/libs"
	"github.com/99-66/go-gin-example-blog-api/repositories"
	"github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id int `json:"-" gorm:"column:id;primaryKey;autoIncrement"`
	Password string `json:"-" gorm:"column:password;type:varchar(32);not null"`
	RegisteredAt time.Time `json:"registeredAt" gorm:"column:registeredAt;type:datetime;not null"`
	LastLogin time.Time `json:"lastLogin" gorm:"column:lastLogin;type:datetime"`
	FirstName string `json:"firstName,required" gorm:"column:firstName;type:varchar(50);default:null"`
	LastName string `json:"lastName,required" gorm:"column:lastName;type:varchar(50);default:null"`
	Mobile string `json:"mobile,required" gorm:"column:mobile;type:varchar(15);uniqueIndex:uq_mobile,sort:asc;null"`
	Email string `json:"email,required" gorm:"column:email;type:varchar(50);uniqueIndex:uq_email,sort:asc;null"`
	Intro string `json:"intro" gorm:"column:intro;type:tinytext"`
	Profile string `json:"profile" gorm:"column:profile;type:text"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) PreUpdateCleaning() {
	if u.Password != "" {
		u.Password = ""
	}
}

func (u *User) Create() (newUser User, err error){
	var mysqlErr *mysql.MySQLError

	hashedPassword, err := libs.HashPassword(u.Password)
	if err != nil {
		return
	}

	newUser = User {
		FirstName: u.FirstName,
		LastName: u.LastName,
		Mobile: u.Mobile,
		Email: u.Email,
		Password: hashedPassword,
		RegisteredAt: time.Now(),
		Intro: u.Intro,
		Profile: u.Profile,
	}

	err = repositories.Connections.DB.Create(&newUser).Error
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return User{}, fmt.Errorf("already signed up")
	}

	if err != nil {
		return User{}, err
	}

	return newUser, nil
}

func (u *User) FindUserById(id int) (user User, err error) {
	err = repositories.Connections.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *User) UpdateUser(updateUser *User) error {
	var mysqlErr *mysql.MySQLError
	updateUser.PreUpdateCleaning()

	err := repositories.Connections.DB.Model(u).Updates(updateUser).Error
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return fmt.Errorf("this value is duplicated with another user")
	}
	if err != nil {
		return err
	}

	return nil
}

func (u *User) DeleteUserById(id int) error {
	err := repositories.Connections.DB.Where("id = ?", id).Delete(u).Error
	if err != nil {
		return err
	}

	return nil
}