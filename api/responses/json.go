//  * @package	GO
//  * @author	Yulianto Mustaqim
//  * @copyleft (ɔ) 2022 - now , Medika Digital Nusantara
//  * @license	https://opensource.org/licenses/MIT	MIT License
//  * @github link https://github.com/ymustaqim-id/go-morclinic
//  * @version	Version 1.0.0

package responses

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/gommon/log"
)

type Response struct {
	Status     bool        `json:"status"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}, pesan string) {
	fmt.Println("mesage data", data)
	var respData Response
	w.WriteHeader(statusCode)
	respData.Data = data
	respData.Status = true
	respData.StatusCode = statusCode
	respData.Message = pesan

	if respData.StatusCode == 404 {
		respData.Status = false
		respData.Data = nil
	} else if respData.StatusCode == 401 {
		respData.Status = false
		respData.Data = nil
	}

	err := json.NewEncoder(w).Encode(respData)
	if err != nil {
		log.Error("error if")

	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error, pesan string) {
	if err != nil {
		JSON(w, statusCode, err, pesan)
		return
	}
	JSON(w, http.StatusBadRequest, err, pesan)
}
