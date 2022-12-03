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
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	var respData Response
	var message string

	w.WriteHeader(statusCode)
	respData.Data = data
	respData.Status = statusCode
	message = "Data di temukan"

	if respData.Status == 404 {
		message = "Data tidak di temukan"
		respData.Data = nil
	}

	respData.Message = message

	err := json.NewEncoder(w).Encode(respData)
	if err != nil {
		log.Error("error if")

	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	fmt.Println("Testtt 1 error")
	if err != nil {
		fmt.Println("Testtt if error")
		JSON(w, statusCode, err)
		return
	}
	fmt.Println("Testtt 2 error")
	JSON(w, http.StatusBadRequest, err)
}
