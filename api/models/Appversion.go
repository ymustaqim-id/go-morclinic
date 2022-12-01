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

type App_version struct {
	ID          uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Version     string `gorm:"size:255;not null;unique" json:"version"`
	Version_ios string `gorm:"size:255;not null;" json:"version_ios"`
}

func (u *App_version) FindAllVersion(db *gorm.DB) (*[]App_version, error) {
	var err error
	appversion := []App_version{}
	err = db.Debug().Model(&App_version{}).Limit(100).Find(&appversion).Error
	if err != nil {
		return &[]App_version{}, err
	}
	return &appversion, err
}
