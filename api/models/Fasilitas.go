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

type Informasi struct {
	ID         uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Nama       string `gorm:"size:255;null;unique" json:"nama"`
	Keterangan string `gorm:"size:255;null;" json:"keterangan"`
	Foto       string `gorm:"size:255;null;" json:"foto"`
	Tipe       string `gorm:"size:255;null;" json:"tipe"`
	Id_klinik  int    `gorm:"size:255;null;" json:"id_klinik"`
	Created_at string `gorm:"size:255;null;" json:"created_at"`
}

func (u *Informasi) FindAllFasilitas(db *gorm.DB) (*[]Informasi, error) {
	var err error
	fasilitas := []Informasi{}
	err = db.Debug().Model(&Informasi{}).Where(&Informasi{Tipe: "fasilitas"}).Find(&fasilitas).Error
	if err != nil {
		return &[]Informasi{}, err
	}
	return &fasilitas, err
}

func (p *Informasi) FindFasilitasByID(db *gorm.DB, id uint32) (*Informasi, error) {
	var err error
	err = db.Debug().Model(&Informasi{}).Where("id_klinik = ?", id).Take(&p).Error
	if err != nil {
		return &Informasi{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&Informasi{}).Where("id_klinik = ?", p.Id_klinik).Take(&p).Error
		if err != nil {
			return &Informasi{}, err
		}
	}
	return p, nil
}
