package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserParkinAdmin struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Phone     string    `gorm:"size:100;not null;" json:"phone"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	TipoLogin string    `gorm:"size:100;not null;" json:"tipologin"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *UserParkinAdmin) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *UserParkinAdmin) Prepare() {
	u.ID = 0
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *UserParkinAdmin) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Username == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Phone == "" {
			return errors.New("Required Phone")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.Username == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Phone == "" {
			return errors.New("Required Phone")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

func (u *UserParkinAdmin) SaveUser(db *gorm.DB) (*UserParkinAdmin, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &UserParkinAdmin{}, err
	}
	return u, nil
}

func (u *UserParkinAdmin) FindAllUsers(db *gorm.DB) (*[]UserParkinAdmin, error) {
	var err error
	users := []UserParkinAdmin{}
	err = db.Debug().Model(&UserParkinAdmin{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]UserParkinAdmin{}, err
	}
	return &users, err
}

func (u *UserParkinAdmin) FindUserByID(db *gorm.DB, uid uint32) (*UserParkinAdmin, error) {
	var err error
	err = db.Debug().Model(UserParkinAdmin{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &UserParkinAdmin{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &UserParkinAdmin{}, errors.New("User Not Found")
	}
	return u, err
}

func (u *UserParkinAdmin) UpdateAUser(db *gorm.DB, uid uint32) (*UserParkinAdmin, error) {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&UserParkinAdmin{}).Where("id = ?", uid).Take(&UserParkinAdmin{}).UpdateColumns(
		map[string]interface{}{
			"password":  u.Password,
			"nickname":  u.Username,
			"email":     u.Email,
			"phone":     u.Phone,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &UserParkinAdmin{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&UserParkinAdmin{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &UserParkinAdmin{}, err
	}
	return u, nil
}

func (u *UserParkinAdmin) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&UserParkinAdmin{}).Where("id = ?", uid).Take(&UserParkinAdmin{}).Delete(&UserParkinAdmin{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
