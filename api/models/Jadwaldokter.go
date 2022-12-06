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

type JadwalDokter struct {
	ID               int32       `gorm:"primary_key;auto_increment" json:"id"`
	Id_dokter_klinik int32       `gorm:"size:255;null;" json:"Id_dokter_klinik"`
	Id_klinik        int32       `gorm:"size:255;null;" json:"Id_klinik"`
	Jam_praktek      interface{} `gorm:"size:255;null;" json:"Jam_praktek"`
	Jam_selesai      interface{} `gorm:"size:255;null;" json:"Jam_selesai"`
	Hari             string      `gorm:"size:255;null;" json:"Hari"`
	Shift            string      `gorm:"size:255;null;" json:"Shift"`
	Qty              int32       `gorm:"size:255;null;" json:"Qty"`
	Status           string      `gorm:"size:255;null;" json:"Status"`
	Interval         string      `gorm:"size:255;null;" json:"Interval"`
	Id_unit          int32       `gorm:"size:255;null;" json:"Id_unit"`
	Created_at       string      `gorm:"size:255;null;" json:"created_at"`
}

func (p *JadwalDokter) FindJadwalDokterByIdKlinik(db *gorm.DB, id int32) (*[]JadwalDokter, error) {
	var err error
	jadwaldokter := []JadwalDokter{}
	err = db.Debug().Model(&JadwalDokter{}).Where(&JadwalDokter{Status: "Aktif"}).Where(&JadwalDokter{Id_klinik: id}).Limit(100).Find(&jadwaldokter).Error
	if err != nil {
		return &[]JadwalDokter{}, err
	}
	return &jadwaldokter, err
}
