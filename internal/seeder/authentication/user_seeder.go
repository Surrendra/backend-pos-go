package seeder

import (
	"BackendPOS/internal/constant"
	"BackendPOS/internal/helper"
	modelAuth "BackendPOS/internal/model/authentication"
	modelMaintenance "BackendPOS/internal/model/maintenance"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB) {
	seedPermission(db)
	seedRole(db)
	seedRolePermission(db)
	seedMerchant(db)
	seedUser(db)

}

func seedPermission(db *gorm.DB) {
	permissons := []modelAuth.Permission{
		{
			Name:        "authentication",
			Description: "Authentication",
			Active:      constant.IndicatorActive,
			Children: []modelAuth.Permission{
				{
					Name:        "user.index",
					Description: "View User Data",
					Active:      constant.IndicatorActive,
					Children: []modelAuth.Permission{
						{
							Name:        "user.create",
							Description: "Create User Data",
							Active:      constant.IndicatorActive,
						},
						{
							Name:        "user.edit",
							Description: "Edit User Data",
							Active:      constant.IndicatorActive,
						},
						{
							Name:        "user.delete",
							Description: "Delete User Data",
							Active:      constant.IndicatorActive,
						},
					},
				},
				{
					Name:        "role.index",
					Description: "View Role Data",
					Active:      constant.IndicatorActive,
					Children: []modelAuth.Permission{
						{
							Name:        "role.create",
							Description: "Create Role Data",
							Active:      constant.IndicatorActive,
						},
						{
							Name:        "role.edit",
							Description: "Edit Role Data",
							Active:      constant.IndicatorActive,
						},
						{
							Name:        "role.delete",
							Description: "Delete Role Data",
							Active:      constant.IndicatorActive,
						},
					},
				},
			},
		},
		{
			Name:        "maintenance",
			Description: "Maintenance",
			Active:      constant.IndicatorActive,
			Children: []modelAuth.Permission{
				{
					Name:        "merchant.index",
					Description: "View Merchant Data",
					Active:      constant.IndicatorActive,
					Children: []modelAuth.Permission{
						{
							Name:        "merchant.create",
							Description: "Create Merchant Data",
							Active:      constant.IndicatorActive,
						},
						{
							Name:        "merchant.edit",
							Description: "Edit Merchant Data",
							Active:      constant.IndicatorActive,
						},
						{
							Name:        "merchant.delete",
							Description: "Delete Merchant Data",
							Active:      constant.IndicatorActive,
						},
					},
				},
			},
		},
	}

	for _, p := range permissons {
		var parent modelAuth.Permission
		err := db.Where("name = ?", p.Name).First(&parent).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				parent = modelAuth.Permission{
					Code:        uuid.New().String(),
					Name:        p.Name,
					Description: p.Description,
					Active:      p.Active,
				}
				if err = db.Create(&parent).Error; err != nil {
					logrus.Error("seedPermission@createParent ", err)
					continue
				}
			} else {
				logrus.Error("seedPermission@findParent ", err)
				continue
			}
		}

		for _, c := range p.Children {
			var child modelAuth.Permission
			err = db.Where("name = ?", c.Name).First(&child).Error
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					newChild := modelAuth.Permission{
						Code:        uuid.New().String(),
						Name:        c.Name,
						Description: c.Description,
						Active:      c.Active,
						ParentID:    &parent.ID,
					}
					if err = db.Create(&newChild).Error; err != nil {
						logrus.Error("seedPermission@createChild ", err, c.Name)
					}
					child = newChild
				} else {
					logrus.Error("seedPermission@findChild ", err, c.Name)
				}
			}

			for _, subChild := range c.Children {
				var subChildPerm modelAuth.Permission
				err = db.Where("name = ?", subChild.Name).First(&subChildPerm).Error
				if err != nil {
					if err == gorm.ErrRecordNotFound {
						newSubChild := modelAuth.Permission{
							Code:        uuid.New().String(),
							Name:        subChild.Name,
							Description: subChild.Description,
							Active:      subChild.Active,
							ParentID:    &child.ID,
						}
						if err = db.Create(&newSubChild).Error; err != nil {
							logrus.Error("seedPermission@createSubChild ", err, subChild.Name)
						}
					} else {
						logrus.Error("seedPermission@findSubChild ", err, subChild.Name)
					}
				}
			}
		}
	}
}

func seedMerchant(db *gorm.DB) {
	merchants := []modelMaintenance.Merchant{
		{
			Name:            "Media Ceria",
			Address:         helper.NullableString("Jalan Raya Sayan, Banjar Gg. Pande No.77, Sayan, Kecamatan Ubud, Kabupaten Gianyar, Bali 80571"),
			PicName:         helper.NullableString("Angga Pamungkas"),
			Domain:          "console",
			Active:          constant.IndicatorActive,
			PicEmail:        "medcer@gmail.com",
			CreatedUserID:   1,
			CreatedUserName: "Surendra Made",
		},
		{
			Name:            "Bali-Zoo",
			Address:         helper.NullableString("Jl. Raya Singapadu, Singapadu, Kec. Sukawati, Kabupaten Gianyar, Bali 80582"),
			PicName:         helper.NullableString("Luke Nichols"),
			Domain:          "balizoo",
			Active:          constant.IndicatorActive,
			PicEmail:        "balizoo@gmail.com",
			CreatedUserID:   1,
			CreatedUserName: "Surendra Made",
		},
		{
			Code:            "sample",
			Name:            "Merchant Sample",
			Address:         helper.NullableString("1600 Amphitheatre Parkway, Mountain View, California."),
			PicName:         helper.NullableString("Luke Nichols"),
			Domain:          "sample",
			Active:          constant.IndicatorActive,
			PicEmail:        "merchantsample@gmail.com",
			CreatedUserID:   1,
			CreatedUserName: "Surendra Made",
		},
	}

	for _, merchant := range merchants {
		if merchant.Code == "" {
			merchant.Code = uuid.New().String()
		}
		err := db.Where("domain = ?", merchant.Domain).FirstOrCreate(&merchant).Error
		if err != nil {
			logrus.Error("SeedUser@merhcnat", err)
		}
	}
}

func seedRolePermission(db *gorm.DB) {
	permissons := []modelAuth.Permission{}
	err := db.Find(&permissons).Error
	if err != nil {
		logrus.Error("seedRolePermission@permission", err)
		return
	}

	var role modelAuth.Role
	err = db.Where("name = ?", "administrator").First(&role).Error
	if err != nil {
		logrus.Error("seedRolePermission@role", err)
		return
	}

	for _, p := range permissons {
		rolePermission := modelAuth.RolePermission{
			RoleId:       role.ID,
			PermissionId: p.ID,
			//CreatedUserId:   1,
			//CreatedUserName: "Surendra Made",
		}
		err = db.Where("role_id = ? AND permission_id = ?", rolePermission.RoleId, rolePermission.PermissionId).FirstOrCreate(&rolePermission).Error
		if err != nil {
			logrus.Error("seedRolePermission@create ", err, role.Name, p.Name)
		}
	}
}

func seedUser(db *gorm.DB) {
	var merchant modelMaintenance.Merchant
	var role modelAuth.Role
	_ = db.Where("domain = ?", "console").First(&merchant).Error
	_ = db.Where("name = ?", "administrator").First(&role).Error

	//err = db.Where("id is not null").Delete(&modelAuth.User{}).Error

	defaultPassword, _ := helper.HashPassword("11235811")
	users := []modelAuth.User{
		{
			RoleId:   role.ID,
			Name:     "administrator",
			Username: "administrator",
			Email:    "administrator@pos.com",
			Address:  helper.NullableString("Jl. Panjaitan No.7, Sumerta Kelod, Denpasar Sel, Sumerta Kelod, Denpasar Selatan, Kota Denpasar, Bali 80234"),
			Password: defaultPassword,
		},
		{
			RoleId:   role.ID,
			Name:     "admin",
			Username: "admin",
			Email:    "admin@pos.com",
			Address:  helper.NullableString("Jl. Panjaitan No.7, Sumerta Kelod, Denpasar Sel, Sumerta Kelod, Denpasar Selatan, Kota Denpasar, Bali 80234"),
			Password: defaultPassword,
		},
	}

	// insert using upsert
	for _, user := range users {
		user.Code = uuid.New().String()
		user.Active = constant.IndicatorActive
		user.Status = constant.StatusActive
		user.MerchantId = merchant.ID
		user.MerchantCode = &merchant.Code
		user.MerchantName = &merchant.Name
		// update or create
		err := db.Where("username = ?", user.Username).FirstOrCreate(&user).Error
		if err != nil {
			logrus.Error("SeedUser@user", err)
		}

		var userMerchant modelAuth.UserMerchant
		err = db.Where("user_id = ? AND merchant_id = ?", user.ID, merchant.ID).First(&userMerchant).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				userMerchant = modelAuth.UserMerchant{
					UserID:          user.ID,
					MerchantID:      merchant.ID,
					CreatedUserID:   1,
					CreatedUserName: "Surendra Made",
				}
				if err = db.Create(&userMerchant).Error; err != nil {
					logrus.Error("SeedUser@userMerchant", err)
				}
			} else {
				logrus.Error("SeedUser@findUserMerchant", err)
			}
		}

		var userRole modelAuth.UserRole
		err = db.Where("user_id = ? AND role_id = ?", user.ID, role.ID).First(&userRole).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				userRole = modelAuth.UserRole{
					UserID:          user.ID,
					RoleID:          role.ID,
					CreatedUserID:   1,
					CreatedUserName: "Surendra Made",
				}
				if err = db.Create(&userRole).Error; err != nil {
					logrus.Error("SeedUser@userRole", err)
				}
			} else {
				logrus.Error("SeedUser@findUserRole", err)
			}
		}
	}
}

func seedRole(db *gorm.DB) {
	roles := []modelAuth.Role{
		{
			Name:        "administrator",
			Description: helper.NullableString("Administrator"),
			Active:      constant.IndicatorActive,
		},
		{
			Name:        "operator",
			Description: helper.NullableString("Operator"),
			Active:      constant.IndicatorActive,
		},
	}

	for _, role := range roles {
		role.Code = uuid.New().String()

		err := db.Where("name = ?", role.Name).FirstOrCreate(&role).Error
		if err != nil {
			logrus.Error("SeedUser@role", err)
		}
	}
}
