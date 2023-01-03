//  * @package	GO
//  * @author	Yulianto Mustaqim
//  * @copyleft (ɔ) 2022 - now , Medika Digital Nusantara
//  * @license	https://opensource.org/licenses/MIT	MIT License
//  * @github link https://github.com/ymustaqim-id/go-morclinic
//  * @version	Version 1.0.0

package models

import (
	"github.com/jinzhu/gorm"
)

type Pengguna_aplikasi struct {
	ID            uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Username      string `gorm:"size:255;not null;unique;" json:"username"`
	Password      string `gorm:"size:255;not null;" json:"password"`
	No_telp       string `gorm:"size:255;not null;unique;" json:"no_telp"`
	No_ktp        string `gorm:"size:255;not null;unique;" json:"no_ktp"`
	Tgl_lahir     string `gorm:"size:255;not null;" json:"tgl_lahir"`
	No_rm         string `gorm:"size:255;not null;unique;" json:"no_rm"`
	Nama          string `gorm:"size:255;not null;" json:"nama"`
	Jenis_kelamin string `gorm:"size:255;not null;" json:"jenis_kelamin"`
	Alamat        string `gorm:"size:255;not null;" json:"alamat"`
	Status_user   string `gorm:"size:255;not null;" json:"status_user"`
	Parent_user   string `gorm:"size:255;not null;" json:"parent_user"`
	Status_delete string `gorm:"size:255;not null;" json:"status_delete"`
}

func (u *Pengguna_aplikasi) GetByDataUname(db *gorm.DB, username string) (*[]Pengguna_aplikasi, error) {
	var err error
	getData := []Pengguna_aplikasi{}
	err = db.Debug().Select("*").
		Model(&Pengguna_aplikasi{}).
		Where(&Pengguna_aplikasi{Username: username}).
		// Or(&Pengguna_aplikasi{No_rm: norm}).
		Where(&Pengguna_aplikasi{Status_user: "Master"}).
		Where(&Pengguna_aplikasi{Status_delete: "N"}).
		Find(&getData).Error
	if err != nil {
		return &getData, err
	}

	return &getData, err
}

func (u *Pengguna_aplikasi) GetByDataNorm(db *gorm.DB, norm string) (*[]Pengguna_aplikasi, error) {
	var err error
	getData := []Pengguna_aplikasi{}
	err = db.Debug().Select("*").
		Model(&Pengguna_aplikasi{}).
		// Where(&Pengguna_aplikasi{Username: username}).
		Where(&Pengguna_aplikasi{No_rm: norm}).
		Where(&Pengguna_aplikasi{Status_user: "Master"}).
		Where(&Pengguna_aplikasi{Status_delete: "N"}).
		Find(&getData).Error
	if err != nil {
		return &getData, err
	}

	return &getData, err
}
