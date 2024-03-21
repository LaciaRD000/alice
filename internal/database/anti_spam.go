package database

type AntiSpam struct {
	ID              string `gorm:"primaryKey"`
	Enabled         bool   `gorm:"default:false"`
	AntiInvite      bool   `gorm:"default:false"`
	ResolvingLinks  bool   `gorm:"default:false"`
	MaximumMentions int    `gorm:"default:-1"`
	AntiDuplicate   int    `gorm:"default:-1"`
	AntiRaid        int    `gorm:"default:-1"`
	MaximumLines    int    `gorm:"default:-1"`
}

func (v *AntiSpam) Create() error {
	return db.Create(&v).Error
}

func (v *AntiSpam) Find(query string, args ...string) (err error) {
	return db.Where(query, args).Find(&v).Error
}

func (v *AntiSpam) Update() (err error) {
	return db.Save(&v).Error
}
