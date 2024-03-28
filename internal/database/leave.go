package database

type Leave struct {
	GuildID   string `gorm:"primaryKey"`
	Enabled   bool   `gorm:"default:false"`
	ChannelID string `gorm:"default:-1"`
}

func (v *Leave) Create() error {
	return db.Create(&v).Error
}

func (v *Leave) Update() error {
	return db.Save(&v).Error
}

func (v *Leave) Find(query string, args ...string) error {
	return db.Where(query, args).Find(&v).Error
}
