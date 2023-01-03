//  * @package	GO
//  * @author	Yulianto Mustaqim
//  * @copyleft (ɔ) 2022 - now , Medika Digital Nusantara
//  * @license	https://opensource.org/licenses/MIT	MIT License
//  * @github link https://github.com/ymustaqim-id/go-morclinic
//  * @version	Version 1.0.0

package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"go-morclinic/api/auth"
	"go-morclinic/api/models"
	"go-morclinic/api/responses"

	"github.com/gorilla/mux"
)

func (server *Server) GetNews(w http.ResponseWriter, r *http.Request) {

	//CHeck if the auth token is valid and  get the user id from it
	errs := auth.ExtractToken(r)
	if errs == "" {
		responses.ERROR(w, http.StatusUnauthorized, nil, "Authentifikasi gagal di lakukan.")
		return
	}

	news := models.Informasi{}
	apps, err := news.FindAllNews(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, nil, "Data tidak di temukan.")
		return
	}

	responses.JSON(w, http.StatusOK, apps, "Data news di temukan.")

}

func (server *Server) GetNewsByIdKlinik(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id_klinik"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, nil, "Parameter tidak sesuai.")
		return
	}

	//CHeck if the auth token is valid and  get the user id from it
	errs := auth.ExtractToken(r)
	if errs == "" {
		responses.ERROR(w, http.StatusUnauthorized, nil, "Authentifikasi gagal di lakukan.")
		return
	}

	news := models.Informasi{}
	newsData, err := news.FindNewsByIdKlinik(server.DB, int32(uid))
	if err != nil || len(*newsData) == 0 {
		if len(*newsData) == 0 {
			err = errors.New("Empty Data")
		}
		responses.ERROR(w, http.StatusNotFound, nil, "Data news tidak di temukan.")
		return
	}
	responses.JSON(w, http.StatusOK, newsData, "Data news di temukan.")
}

func (server *Server) GetListThreeByIdKlinik(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id_klinik"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, nil, err.Error())
		return
	}

	//CHeck if the auth token is valid and  get the user id from it
	errs := auth.ExtractToken(r)
	if errs == "" {
		responses.ERROR(w, http.StatusUnauthorized, nil, err.Error())
		return
	}

	news := models.Informasi{}
	newsData, err := news.FindListThreeByIdKlinik(server.DB, int32(uid))
	if err != nil || len(*newsData) == 0 {
		if len(*newsData) == 0 {
			err = errors.New("Empty Data")
		}
		responses.ERROR(w, http.StatusNotFound, nil, err.Error())
		return
	}
	responses.JSON(w, http.StatusOK, newsData, "Data list three di temukan.")
}
