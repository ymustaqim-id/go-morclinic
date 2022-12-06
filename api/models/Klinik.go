//  * @package	GO
//  * @author	Yulianto Mustaqim
//  * @copyleft (ɔ) 2022 - now , Medika Digital Nusantara
//  * @license	https://opensource.org/licenses/MIT	MIT License
//  * @github link https://github.com/ymustaqim-id/go-morclinic
//  * @version	Version 1.0.0

package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Klinik struct {
	ID     int32  `gorm:"primary_key;auto_increment" json:"id"`
	Nama   string `gorm:"size:255;not null;" json:"nama"`
	Alamat string `gorm:"size:255;not null;" json:"alamat"`
	Status string `gorm:"size:255;not null;" json:"status"`
}

type jadwalKlinik struct {
	ID          int32       `gorm:"primary_key;auto_increment" json:"id"`
	Id_klinik   int32       `gorm:"size:255;null;" json:"Id_klinik"`
	Jam_praktek interface{} `gorm:"size:255;null;" json:"Jam_praktek"`
	Jam_selesai interface{} `gorm:"size:255;null;" json:"Jam_selesai"`
	Hari        string      `gorm:"size:255;null;" json:"Hari"`
	Shift       string      `gorm:"size:255;null;" json:"Shift"`
	Status      string      `gorm:"size:255;null;" json:"Status"`
	Interval    string      `gorm:"size:255;null;" json:"Interval"`
	Created_at  string      `gorm:"size:255;null;" json:"created_at"`
}

func (u *Klinik) FindKlinik(db *gorm.DB) (*[]Klinik, error) {
	var err error
	klinik := []Klinik{}
	err = db.Debug().Model(&Klinik{}).Where(&Klinik{Status: "Aktif"}).Limit(15).Find(&klinik).Error
	if err != nil {
		return &[]Klinik{}, err
	}
	return &klinik, err
}

func (u *Klinik) DetailKlinik(db *gorm.DB, id int32) (*[]Klinik, error) {
	var err error
	klinik := []Klinik{}
	err = db.Debug().Select("*").
		Model(&Klinik{}).
		Model(&jadwalKlinik{}).
		Joins("JOIN jadwal_klinik on jadwal_klinik.id_klinik=klinik.id").
		Where(&Klinik{Status: "Aktif"}).
		Where(&jadwalKlinik{Status: "Aktif"}).
		Where(&Klinik{ID: id}).
		Limit(15).
		Find(&klinik).Error
	fmt.Println("err", &klinik)
	if err != nil {
		return &[]Klinik{}, err
	}
	return &klinik, err
}
