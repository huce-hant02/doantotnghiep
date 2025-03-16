package model

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/lib/pq"
	"github.com/lithammer/shortuuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Đoàn đánh giá bên ngoài CTĐT
type ExternalAssessmentTeam struct {
	ID     uint `json:"id"`
	UserID uint `json:"userId" gorm:"default:0"`

	Code        string `json:"code" gorm:"unique"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Email       string `json:"email" gorm:"unique"`
	Phone       string `json:"phone"`

	ListDepartment pq.StringArray `json:"listDepartment" gorm:"type:text"` // {pud007, pud014}
	IsActive       *bool          `json:"isActive" gorm:"default:true"`

	User *User `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`
}

func (team *ExternalAssessmentTeam) AfterCreate(tx *gorm.DB) (err error) {
	var newEmail string
	if team.Email != "" {
		newEmail = team.Email
	}
	roleID := 0
	if err := tx.Model(&Role{}).Select("id").Where("code = 'doan-danh-gia-ngoai'").Find(&roleID).Error; err != nil {
		return err
	}
	password := team.Code + "@" + strconv.Itoa(time.Now().Year()) // ex: PUE0001@2023
	hashedPW := HashAndSalt(password)

	userTeam := User{
		Username: newEmail,
		Password: hashedPW,
		Type:     "DOANDANHGIANGOAI",
		PuCode:   team.Code,
		Name:     team.Name,
		UserRoles: []UserRole{
			{
				UserId: 0,
				RoleId: uint(roleID),
				Active: &TrueValue,
			},
		},
	}

	err = tx.Model(&User{}).Create(&userTeam).Error
	if err != nil {
		return err
	}
	err = tx.Model(&User{}).Where("id = ?", userTeam.ID).Update("password", hashedPW).Error
	if err != nil {
		return err
	}
	err = tx.Model(&ExternalAssessmentTeam{}).Where("lower(email) = ?", strings.ToLower(newEmail)).Update("user_id", userTeam.ID).Error
	if err != nil {
		return err
	}
	// if err := tx.Model(&ECGPartner{}).Where("id = ?", partner.ID).Update("user_partner_id", userPartner.ID).Error; err != nil {
	// 	return err
	// }
	// // send email
	// if partner.SendMail != nil && *partner.SendMail == true {
	// 	var system *System
	// 	if err := tx.Model(&System{}).Where("id > 0").Where("deleted_at IS NULL").Limit(1).
	// 		Find(&system).Error; err != nil {
	// 		return err
	// 	}
	// 	mailConfig := system.EmailConfig

	// 	if mailConfig != nil && mailConfig.Email != "" {
	// 		mailForm := struct {
	// 			PartnerName string
	// 			Username    string
	// 			Password    string
	// 		}{
	// 			PartnerName: partner.Name,
	// 			Username:    newEmail,
	// 			Password:    password,
	// 		}
	// 		htmlForm, err := template.New("partner-email").Parse(PartnerRegistrationEmailForm)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		var contentTemplate bytes.Buffer
	// 		err = htmlForm.Execute(&contentTemplate, mailForm)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		// log.Println(newEmail)

	// 		err = SendEmail(*mailConfig, []string{newEmail}, "Thông tin tài khoản", contentTemplate.String())
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}
	// }

	return nil
}

func (team *ExternalAssessmentTeam) AfterUpdate(tx *gorm.DB) (err error) {
	tmpDB := tx.Model(&User{}).Where("id = ? AND type = 'DOANDANHGIANGOAI'", team.UserID)
	changes := 0
	if team.Email != "" {
		tmpDB = tmpDB.Update("username", team.Email)
		changes++
	}
	if team.IsActive != nil {
		tmpDB = tmpDB.Update("active", *team.IsActive)
		changes++
	}
	if team.Name != "" {
		tmpDB = tmpDB.Update("name", team.Name)
		changes++
	}
	if team.Code != "" {
		tmpDB = tmpDB.Update("pu_code", team.Code)
		changes++
	}
	if changes > 0 {
		err = tmpDB.Error
		if err != nil {
			return err
		}
	}

	return nil
}

func GenCode() string {
	id := shortuuid.New()
	return strings.ToUpper(id[0:10])
}

func HashAndSalt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
