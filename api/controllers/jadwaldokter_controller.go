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

func (server *Server) JadwalDokter(w http.ResponseWriter, r *http.Request) {

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

	jadwaldokter := models.JadwalDokter{}
	jadwalDokter, err := jadwaldokter.FindJadwalDokterByIdKlinik(server.DB, int32(uid))
	if err != nil || len(*jadwalDokter) == 0 {
		if len(*jadwalDokter) == 0 {
			err = errors.New("Empty Data")
		}
		responses.ERROR(w, http.StatusNotFound, nil, "Data jadwal tidak di temukan.")
		return
	}
	responses.JSON(w, http.StatusOK, jadwalDokter, "Data jadwal di temukan.")
}
