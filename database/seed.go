package database

import (
	"go-resto/internal/models"
	"go-resto/internal/models/constant"

	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
	foodMenu := []models.MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     15000,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Nasi Goreng",
			OrderCode: "nasi_goreng",
			Price:     20000,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Fuyunghai",
			OrderCode: "fuyunghai",
			Price:     30000,
			Type:      constant.MenuTypeFood,
		},
	}

	drinksMenu := []models.MenuItem{
		{
			Name:      "Es Teh",
			OrderCode: "es_teh",
			Price:     5000,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Es Jeruk",
			OrderCode: "jeruk",
			Price:     7000,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Es Kopi",
			OrderCode: "es_kopi",
			Price:     10000,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Es Tawar",
			OrderCode: "es_tawar",
			Price:     5000,
			Type:      constant.MenuTypeDrink,
		},
	}

	if err := db.First(&models.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinksMenu)
	}
}
