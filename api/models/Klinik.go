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

type Klinik struct {
	ID           int32  `gorm:"primary_key;auto_increment" json:"id"`
	Nama         string `gorm:"size:255;not null;" json:"nama"`
	Alamat       string `gorm:"size:255;not null;" json:"alamat"`
	Status       string `gorm:"size:255;not null;" json:"status"`
	Hari_mulai   string `gorm:"size:255;not null;" json:"hari_mulai"`
	Hari_selesai string `gorm:"size:255;not null;" json:"hari_selesai"`
	Jam_mulai    string `gorm:"size:255;not null;" json:"jam_mulai"`
	Jam_selesai  string `gorm:"size:255;not null;" json:"jam_selesai"`
}

type jadwalKlinik struct {
	Status string `gorm:"size:255;null;" json:"Status"`
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
	err = db.Debug().Select("klinik.id, klinik.nama, klinik.alamat, klinik.status,min(jadwal_klinik.hari) as hari_mulai, max(jadwal_klinik.hari) as hari_selesai, min(jadwal_klinik.jam_praktek) as jam_mulai, max(jadwal_klinik.jam_selesai) as jam_selesai").
		Model(&Klinik{}).
		Joins("JOIN jadwal_klinik on jadwal_klinik.id_klinik=klinik.id").
		Where(&Klinik{Status: "Aktif"}).
		Where(&jadwalKlinik{Status: "Aktif"}).
		Where(&Klinik{ID: id}).
		Limit(15).
		Find(&klinik).Error
	if err != nil {
		return &klinik, err
	}

	return &klinik, err
}
