//  * @package	GO
//  * @author	Yulianto Mustaqim
//  * @copyleft (ɔ) 2022 - now , Medika Digital Nusantara
//  * @license	https://opensource.org/licenses/MIT	MIT License
//  * @github link https://github.com/ymustaqim-id/go-morclinic
//  * @version	Version 1.0.0

package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"go-morclinic/api/auth"
	"go-morclinic/api/models"
	"go-morclinic/api/responses"

	"golang.org/x/crypto/bcrypt"
)

func (server *Server) LoginWithUsername(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	pengguna := models.Pengguna_aplikasi{}
	err = json.Unmarshal(body, &pengguna)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}

	dataPengguna, err := pengguna.GetByDataUname(server.DB, pengguna.Username)
	if len(*dataPengguna) == 0 {
		formattedError := errors.New("Data pengguna tidak di temukan.")
		responses.ERROR(w, http.StatusNotFound, formattedError)
		return
	}

	token, err := server.SignInUname(pengguna.Username, pengguna.Password)
	if len(token) == 0 {
		formattedError := errors.New("Token tidak berhasil di generate.")
		responses.ERROR(w, http.StatusNotFound, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) LoginWithNorm(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	pengguna := models.Pengguna_aplikasi{}
	err = json.Unmarshal(body, &pengguna)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	fmt.Println("pengguna norm", pengguna)
	fmt.Println("err", err)

	dataPengguna, err := pengguna.GetByDataNorm(server.DB, pengguna.No_rm)
	if len(*dataPengguna) == 0 {
		formattedError := errors.New("Data pengguna tidak di temukan.")
		responses.ERROR(w, http.StatusNotFound, formattedError)
		return
	}
	token, err := server.SignInNorm(pengguna.No_rm, pengguna.Password)
	if len(token) == 0 {
		formattedError := errors.New("Token tidak berhasil di generate.")
		responses.ERROR(w, http.StatusNotFound, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignInUname(username, password string) (string, error) {

	var err error
	pengguna := models.Pengguna_aplikasi{}
	Pengguna, err := pengguna.GetByDataUname(server.DB, username)
	for _, result := range *Pengguna {
		err = models.VerifyPassword(result.Password, password)

		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
			formattedError := errors.New("Password tidak sesuai.")
			return "", formattedError
		}
	}
	return auth.CreateToken(pengguna.ID)
}

func (server *Server) SignInNorm(norm, password string) (string, error) {

	var err error
	pengguna := models.Pengguna_aplikasi{}
	Pengguna, err := pengguna.GetByDataNorm(server.DB, norm)
	for _, result := range *Pengguna {
		err = models.VerifyPassword(result.Password, password)

		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
			formattedError := errors.New("Password tidak sesuai.")
			return "", formattedError
		}
	}

	return auth.CreateToken(pengguna.ID)
}
