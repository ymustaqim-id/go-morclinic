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
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	fmt.Println("data", err)
	if err != nil {
		JSON(w, statusCode, struct {
			Status  string `json:"status"`
			Message string `json:"message"`
			Data    string `json:"data"`
		}{
			Status:  "false",
			Message: err.Error(),
			Data:    "",
		})
		return
	}
	JSON(w, http.StatusBadRequest, "kosong")
}
