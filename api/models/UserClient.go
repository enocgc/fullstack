package models

import (

	"time"

)

type UserParkinClient struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name  string    `gorm:"size:255;not null;unique" json:"name"`
	LastName     string    `gorm:"size:100;not null;unique" json:"lastname"`
	Email     string    `gorm:"size:100;not null;" json:"email"`
	Phone     string    `gorm:"size:100;not null;" json:"phone"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	TipoRegistro string    `gorm:"size:100;not null;" json:"tiporegistro"`
	Token string    `"json:"token"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// func HashClient(password string) ([]byte, error) {
// 	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// }

// func VerifyPasswordClient(hashedPassword, password string) error {
// 	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
// }

// func (u *UserParkinClient) BeforeSave() error {
// 	hashedPassword, err := HashClient(u.Password)
// 	if err != nil {
// 		return err
// 	}
// 	u.Password = string(hashedPassword)
// 	return nil
// }

// func (u *UserParkinClient) Prepare() {
// 	u.ID = 0
// 	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
// 	u.LastName = html.EscapeString(strings.TrimSpace(u.LastName))
// 	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
// 	u.Phone = html.EscapeString(strings.TrimSpace(u.Email))
// 	u.TipoRegistro = html.EscapeString(strings.TrimSpace(u.TipoRegistro))
// 	u.Token = html.EscapeString(strings.TrimSpace(u.Token))
// 	u.CreatedAt = time.Now()
// 	u.UpdatedAt = time.Now()
// }

// func (u *UserParkinClient) Validate(action string) error {
// 	switch strings.ToLower(action) {
// 	case "update":
// 		if u.Name == "" {
// 			return errors.New("RequiredName")
// 		}
// 		if u.Password == "" {
// 			return errors.New("Required Password")
// 		}
// 		if u.Phone == "" {
// 			return errors.New("Required Phone")
// 		}
// 		if u.Email == "" {
// 			return errors.New("Required Email")
// 		}
// 		if err := checkmail.ValidateFormat(u.Email); err != nil {
// 			return errors.New("Invalid Email")
// 		}

// 		return nil
// 	case "login":
// 		if u.Password == "" {
// 			return errors.New("Required Password")
// 		}
// 		if u.Email == "" {
// 			return errors.New("Required Email")
// 		}
// 		if err := checkmail.ValidateFormat(u.Email); err != nil {
// 			return errors.New("Invalid Email")
// 		}
// 		return nil

// 	default:
// 		if u.Name == "" {
// 			return errors.New("Required Name")
// 		}
// 		if u.Password == "" {
// 			return errors.New("Required Password")
// 		}
// 		if u.Phone == "" {
// 			return errors.New("Required Phone")
// 		}
// 		if u.Email == "" {
// 			return errors.New("Required Email")
// 		}
// 		if err := checkmail.ValidateFormat(u.Email); err != nil {
// 			return errors.New("Invalid Email")
// 		}
// 		return nil
// 	}
// }

// func (u *UserParkinClient) SaveUser(db *gorm.DB) (*UserParkinClient, error) {

// 	var err error
// 	err = db.Debug().Create(&u).Error
// 	if err != nil {
// 		return &UserParkinClient{}, err
// 	}
// 	return u, nil
// }

// func (u *UserParkinClient) FindAllUsers(db *gorm.DB) (*[]UserParkinClient, error) {
// 	var err error
// 	users := []UserParkinClient{}
// 	err = db.Debug().Model(&UserParkinClient{}).Limit(100).Find(&users).Error
// 	if err != nil {
// 		return &[]UserParkinClient{}, err
// 	}
// 	return &users, err
// }

// func (u *UserParkinClient) FindUserByID(db *gorm.DB, uid uint32) (*UserParkinClient, error) {
// 	var err error
// 	err = db.Debug().Model(UserParkinClient{}).Where("id = ?", uid).Take(&u).Error
// 	if err != nil {
// 		return &UserParkinClient{}, err
// 	}
// 	if gorm.IsRecordNotFoundError(err) {
// 		return &UserParkinClient{}, errors.New("User Not Found")
// 	}
// 	return u, err
// }

// func (u *UserParkinClient) UpdateAUser(db *gorm.DB, uid uint32) (*UserParkinClient, error) {

// 	// To hash the password
// 	err := u.BeforeSave()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	db = db.Debug().Model(&UserParkinClient{}).Where("id = ?", uid).Take(&UserParkinClient{}).UpdateColumns(
// 		map[string]interface{}{
// 			"password":  u.Password,
// 			"name":  u.Name,
// 			"email":     u.Email,
// 			"phone":     u.Phone,
// 			"update_at": time.Now(),
// 		},
// 	)
// 	if db.Error != nil {
// 		return &UserParkinClient{}, db.Error
// 	}
// 	// This is the display the updated user
// 	err = db.Debug().Model(&UserParkinClient{}).Where("id = ?", uid).Take(&u).Error
// 	if err != nil {
// 		return &UserParkinClient{}, err
// 	}
// 	return u, nil
// }

// func (u *UserParkinClient) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

// 	db = db.Debug().Model(&UserParkinClient{}).Where("id = ?", uid).Take(&UserParkinClient{}).Delete(&UserParkinClient{})

// 	if db.Error != nil {
// 		return 0, db.Error
// 	}
// 	return db.RowsAffected, nil
// }
