//  * @package	GO
//  * @author	Yulianto Mustaqim
//  * @copyleft (ɔ) 2022 - now , Medika Digital Nusantara
//  * @license	https://opensource.org/licenses/MIT	MIT License
//  * @github link https://github.com/ymustaqim-id/go-morclinic
//  * @version	Version 1.0.0

package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"go-morclinic/api/auth"
	"go-morclinic/api/models"
	"go-morclinic/api/responses"

	"github.com/gorilla/mux"
)

func (server *Server) GetKlinik(w http.ResponseWriter, r *http.Request) {

	klinik := models.Klinik{}
	klinikData, err := klinik.FindKlinik(server.DB)
	if err != nil || len(*klinikData) == 0 {
		if len(*klinikData) == 0 {
			err = errors.New("Empty Data")
		}
		responses.ERROR(w, http.StatusNotFound, nil, "Data klinik tidak di temukan.")
		return
	}
	responses.JSON(w, http.StatusOK, klinikData, "Data klinik di temukan.")

}

func (server *Server) GetKlinikDetail(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id_klinik"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, nil, "Parameter tidak boleh kosong.")
		return
	}

	//CHeck if the auth token is valid and  get the user id from it
	errs := auth.ExtractToken(r)
	if errs == "" {
		responses.ERROR(w, http.StatusUnauthorized, nil, "Authentifikasi gagal di lakukan.")
		return
	}

	klinik := models.Klinik{}
	klinikData, err := klinik.DetailKlinik(server.DB, int32(uid))
	fmt.Println("klinikData", klinikData)
	if err != nil || len(*klinikData) == 0 {
		if len(*klinikData) == 0 {
			err = errors.New("Empty Data")
		}
		responses.ERROR(w, http.StatusNotFound, nil, "Data klinik tidak di temukan.")
		return
	}
	responses.JSON(w, http.StatusOK, klinikData, "Data klinik di temukan.")
}
