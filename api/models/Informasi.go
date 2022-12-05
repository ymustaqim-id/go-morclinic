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
	ID         int32  `gorm:"primary_key;auto_increment" json:"id"`
	Nama       string `gorm:"size:255;null;unique" json:"nama"`
	Keterangan string `gorm:"size:255;null;" json:"keterangan"`
	Foto       string `gorm:"size:255;null;" json:"foto"`
	Tipe       string `gorm:"size:255;null;" json:"tipe"`
	Id_klinik  int32  `gorm:"size:255;null;" json:"id_klinik"`
	Created_at string `gorm:"size:255;null;" json:"created_at"`
}

func (u *Informasi) FindAllFasilitas(db *gorm.DB) (*[]Informasi, error) {
	var err error
	fasilitas := []Informasi{}
	err = db.Debug().Model(&Informasi{}).Limit(100).Where(&Informasi{Tipe: "fasilitas"}).Find(&fasilitas).Error
	if err != nil {
		return &[]Informasi{}, err
	}
	return &fasilitas, err
}

func (p *Informasi) FindFasilitasByIdKlinik(db *gorm.DB, id int32) (*[]Informasi, error) {
	var err error
	fasilitas := []Informasi{}
	err = db.Debug().Model(&Informasi{}).Where(&Informasi{Tipe: "fasilitas"}).Where(&Informasi{Id_klinik: id}).Limit(100).Find(&fasilitas).Error
	if err != nil {
		return &[]Informasi{}, err
	}
	return &fasilitas, err
}

func (u *Informasi) FindAllNews(db *gorm.DB) (*[]Informasi, error) {
	var err error
	news := []Informasi{}
	err = db.Debug().Model(&Informasi{}).Limit(100).Where(&Informasi{Tipe: "news"}).Find(&news).Error
	if err != nil {
		return &[]Informasi{}, err
	}
	return &news, err
}

func (p *Informasi) FindNewsByIdKlinik(db *gorm.DB, id int32) (*[]Informasi, error) {
	var err error
	news := []Informasi{}
	err = db.Debug().Model(&Informasi{}).Where(&Informasi{Tipe: "news"}).Where(&Informasi{Id_klinik: id}).Limit(100).Find(&news).Error
	if err != nil {
		return &[]Informasi{}, err
	}
	return &news, err
}
