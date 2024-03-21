package database

type Welcome struct {
	GuildID   string `gorm:"primaryKey"`
	Enabled   bool   `gorm:"default:false"`
	ChannelID string `gorm:"default:-1"`
}

func (v *Welcome) Create() error {
	return db.Create(&v).Error
}

func (v *Welcome) Update() error {
	return db.Updates(&v).Error
}

func (v *Welcome) Find(query string, args ...string) error {
	return db.Where(query, args).Find(&v).Error
}
