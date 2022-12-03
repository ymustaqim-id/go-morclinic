//  * @package	GO
//  * @author	Yulianto Mustaqim
//  * @copyleft (ɔ) 2022 - now , Medika Digital Nusantara
//  * @license	https://opensource.org/licenses/MIT	MIT License
//  * @github link https://github.com/ymustaqim-id/go-morclinic
//  * @version	Version 1.0.0

package controllers

import (
	"net/http"
	"strconv"

	"go-morclinic/api/models"
	"go-morclinic/api/responses"

	"github.com/gorilla/mux"
)

func (server *Server) GetFasilitas(w http.ResponseWriter, r *http.Request) {

	fasilitas := models.Informasi{}
	apps, err := fasilitas.FindAllFasilitas(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusInternalServerError, apps)

}

func (server *Server) GetFasilitasById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	news := models.Informasi{}
	newsData, err := news.FindFasilitasByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, newsData)
}
