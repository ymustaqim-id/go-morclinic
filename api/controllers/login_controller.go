//  * @package	GO
//  * @author	Yulianto Mustaqim
//  * @copyleft (ɔ) 2022 - now , Medika Digital Nusantara
//  * @license	https://opensource.org/licenses/MIT	MIT License
//  * @github link https://github.com/ymustaqim-id/go-morclinic
//  * @version	Version 1.0.0

package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"go-morclinic/api/auth"
	"go-morclinic/api/models"
	"go-morclinic/api/responses"

	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, nil, "")
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, nil, "")
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, nil, "")
		return
	}
	token, err := server.SignIn(user.Username, user.Password)
	fmt.Println("user", token)

	if err != nil {
		responses.ERROR(w, http.StatusNotFound, nil, "")
		return
	}
	responses.JSON(w, http.StatusOK, token, "Data di temukan.")
}

func (server *Server) SignIn(email, password string) (string, error) {

	var err error

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("username = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
}
