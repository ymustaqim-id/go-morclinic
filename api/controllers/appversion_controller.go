//  * @package	GO
//  * @author	Yulianto Mustaqim
//  * @copyleft (ɔ) 2022 - now , Medika Digital Nusantara
//  * @license	https://opensource.org/licenses/MIT	MIT License
//  * @github link https://github.com/ymustaqim-id/go-morclinic
//  * @version	Version 1.0.0

package controllers

import (
	"net/http"

	"go-morclinic/api/models"
	"go-morclinic/api/responses"
)

func (server *Server) GetAppversion(w http.ResponseWriter, r *http.Request) {

	appversion := models.App_version{}
	apps, err := appversion.FindAllVersion(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusInternalServerError, apps)

}
