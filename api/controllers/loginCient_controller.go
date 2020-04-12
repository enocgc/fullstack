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

func (server *Server) LoginClient(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.UserParkinClient{}
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
	token, err := server.SignInClient(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatErrorClient(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	// obtener los datos del usuario que se loguea
	err = server.DB.Debug().Model(models.UserParkinClient{}).Where("email = ?", user.Email).Take(&user).Error

	responses.JSON(w, http.StatusOK, struct {
		Message string                 `json:"menssage"`
		Status  int                    `json:"status"`
		Error   bool                   `json:"error"`
		Data    map[string]interface{} `json:"data"`
	}{
		Message: "Login Exitoso",
		Status:  http.StatusUnprocessableEntity,
		Error:   false,
		Data:    map[string]interface{}{"token": token, "Name": user.Name, "Last Name": user.LastName, "email": user.Email, "tipoRegistro": user.TipoRegistro},
	})
}

func (server *Server) SignInClient(email, password string) (string, error) {

	var err error

	user := models.UserParkinClient{}

	err = server.DB.Debug().Model(models.UserParkinClient{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	return auth.CreateTokenClientLogin(user.ID, user.Email, user.Name, user.LastName, user.Phone)
}
