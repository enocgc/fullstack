package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/enocgc/fullstack/api/auth"
	"github.com/enocgc/fullstack/api/models"
	"github.com/enocgc/fullstack/api/responses"
	"github.com/enocgc/fullstack/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

var responseData = []models.Response{
	models.Response{
		Message: "",
		Status:  "",
		Error:   "",
		Data:    "",
	},
}

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.UserParkinAdmin{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	// obtener los datos del usuario que se loguea
	err = server.DB.Debug().Model(models.UserParkinAdmin{}).Where("email = ?", user.Email).Take(&user).Error

	responses.JSON(w, http.StatusOK, struct {
		Message string                 `json:"menssage"`
		Status  int                    `json:"status"`
		Error   bool                   `json:"error"`
		Data    map[string]interface{} `json:"data"`
	}{
		Message: "Login Exitoso",
		Status:  http.StatusUnprocessableEntity,
		Error:   false,
		Data:    map[string]interface{}{"token": token, "userName": user.Username, "email": user.Email, "phone": user.Phone},
	})
}

func (server *Server) SignIn(email, password string) (string, error) {

	var err error

	user := models.UserParkinAdmin{}

	err = server.DB.Debug().Model(models.UserParkinAdmin{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	// var token=auth.CreateToken(user.ID)
	// responseData[0].Message = "Login Exitoso"
	// responseData[0].Status = "200"
	// responseData[0].Error = "Null"
	// responseData[0].Data = token

	return auth.CreateToken(user.ID, user.Email, user.Username, user.Phone)
}
