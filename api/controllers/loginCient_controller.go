package controllers

import (
	"encoding/json"
	"fmt"
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

func (server *Server) LoginClientSocial(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user := models.UserParkinClient{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	fmt.Println(user)
	passwordInicial := user.Password
	fmt.Println("inicial 2 Password" + user.Password)
	err = server.DB.Debug().Model(models.UserParkinClient{}).Where("email = ?", user.Email).Take(&user).Error
	if err != nil {
		passwordInicial := user.Password
		user.Prepare()
		err = user.Validate("")
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		fmt.Println("save Password" + user.Password)

		token, err := auth.CreateTokenClient(user.Email, user.Name, user.LastName, user.Phone)
		user.Token = token
		userCreated, err := user.SaveUserClient(server.DB)

		if err != nil {

			formattedError := formaterror.FormatErrorClient(err.Error())

			responses.ERROR(w, http.StatusInternalServerError, formattedError)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))

		if err != nil {
			formattedError := formaterror.FormatErrorClient(err.Error())
			responses.ERROR(w, http.StatusInternalServerError, formattedError)
			return
		}
		err = user.Validate("login")
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		fmt.Println("email: " + user.Email + " Password: " + passwordInicial)
		token, err = server.SignInClient(user.Email, passwordInicial)

		if err != nil {
			formattedError := formaterror.FormatErrorClient(err.Error())
			responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
			return
		}
		// obtener los datos del usuario que se loguea
		err = server.DB.Debug().Model(models.UserParkinClient{}).Where("email = ?", user.Email).Take(&user).Error
		// fmt.Println(err)
		responses.JSON(w, http.StatusOK, struct {
			Message string                 `json:"menssage"`
			Status  int                    `json:"status"`
			Error   bool                   `json:"error"`
			Data    map[string]interface{} `json:"data"`
		}{
			Message: "Login  Exitoso",
			Status:  http.StatusUnprocessableEntity,
			Error:   false,
			Data:    map[string]interface{}{"token": token, "Name": userCreated.Name, "Last Name": userCreated.LastName, "email": userCreated.Email, "tipoRegistro": userCreated.TipoRegistro},
		})
	} else {

		user.Prepare()
		err = user.Validate("login")
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		fmt.Println("email: " + user.Email + " Password: " + passwordInicial)
		token, err := server.SignInClient(user.Email, passwordInicial)
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
	// user.Prepare()

}

func (server *Server) SignInClient(email, password string) (string, error) {

	var err error

	user := models.UserParkinClient{}
	fmt.Println("sigin Password" + password)
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
