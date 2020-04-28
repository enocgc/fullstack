package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/enocgc/fullstack/api/auth"
	"github.com/enocgc/fullstack/api/models"
	"github.com/enocgc/fullstack/api/responses"
	"github.com/enocgc/fullstack/api/utils/formaterror"
	"github.com/gorilla/mux"
)

func (server *Server) CreateParkinDetail(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	detail := models.ParkInDetail{}
	err = json.Unmarshal(body, &detail)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	detail.PrepareDetail()
	err = detail.ValidateDetail()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	// log.Println(uid)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	if uid != detail.Fk_ParkInAdmin {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	detailPrkin, err := detail.SaveParkinDetail(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, detailPrkin.ID))
	responses.JSON(w, http.StatusCreated, struct {
		Message string                 `json:"menssage"`
		Status  int                    `json:"status"`
		Error   bool                   `json:"error"`
		Data    map[string]interface{} `json:"data"`
	}{
		Message: "Registro Exitoso",
		Status:  http.StatusUnprocessableEntity,
		Error:   false,
		Data: map[string]interface{}{"id": detailPrkin.ID, "nombreParqueo": detailPrkin.NombreParqueo,
			"horario": detailPrkin.Horario,
			"detalle": detailPrkin.Detalle},
	})
}

func (server *Server) GetParkinDetails(w http.ResponseWriter, r *http.Request) {

	detail := models.ParkInDetail{}
	details, err := detail.FindAllParkin(server.DB)
	log.Println(details)
	// if err != nil {
	// 	log.Fatal("Cannot encode to JSON ", err)
	// }
	// log.Fatal("succest encode to JSON ", pagesJson)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, details)
}

func (server *Server) GetParkinDetail(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	detail := models.ParkInDetail{}

	detailReceived, err := detail.FindDetailByID(server.DB, pid)
	// log.Println(detailReceived)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, struct {
		Message string                 `json:"menssage"`
		Status  int                    `json:"status"`
		Error   bool                   `json:"error"`
		Data    map[string]interface{} `json:"data"`
	}{
		Message: "Get Parkin detail Exitoso",
		Status:  http.StatusUnprocessableEntity,
		Error:   false,
		Data: map[string]interface{}{"id": detailReceived.ID, "nombreParqueo": detailReceived.NombreParqueo,
			"horario":  detailReceived.Horario,
			"detalle":  detailReceived.Detalle,
			"lat":      detailReceived.Lat,
			"long":     detailReceived.Long,
			"phone":    detailReceived.Phone,
			"sitioWeb": detailReceived.SitioWeb,
			"admin":    detailReceived.Fk_ParkInAdmin},
	})
}

func (server *Server) UpdateParkinDetail(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
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
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if tokenID != uint32(uid) {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	user.Prepare()
	err = user.Validate("update")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedUser, err := user.UpdateAUser(server.DB, uint32(uid))

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedUser)
}

func (server *Server) DeleteParkinDetail(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	user := models.UserParkinClient{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if tokenID != 0 && tokenID != uint32(uid) {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	_, err = user.DeleteAUser(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.JSON(w, http.StatusNoContent, "")
}
