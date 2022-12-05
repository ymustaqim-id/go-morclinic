//  * @package	GO
//  * @author	Yulianto Mustaqim
//  * @copyleft (ɔ) 2022 - now , Medika Digital Nusantara
//  * @license	https://opensource.org/licenses/MIT	MIT License
//  * @github link https://github.com/ymustaqim-id/go-morclinic
//  * @version	Version 1.0.0

package responses

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/gommon/log"
)

type Response struct {
	Status     bool        `json:"status"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	var respData Response
	var message string
	w.WriteHeader(statusCode)
	respData.Data = data
	respData.Status = true
	respData.StatusCode = statusCode
	message = "Data di temukan"

	if respData.StatusCode == 404 {
		respData.Status = false
		message = "Data tidak di temukan"
		respData.Data = nil
	} else if respData.StatusCode == 401 {
		respData.Status = false
		message = "Autentifikasi gagal di lakukan."
		respData.Data = nil
	}

	respData.Message = message

	err := json.NewEncoder(w).Encode(respData)
	if err != nil {
		log.Error("error if")

	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, err)
		return
	}
	JSON(w, http.StatusBadRequest, err)
}
