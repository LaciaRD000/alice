package database

type LevelConfig struct {
	GuildID   string `gorm:"primaryKey"`
	Enabled   bool
	Option    int
	ChannelID string
	Format    string
}

func (v *LevelConfig) Create() error {
	return db.Create(&v).Error
}

func (v *LevelConfig) Update() error {
	return db.Save(&v).Error
}

func (v *LevelConfig) Find(query string, args ...string) error {
	return db.Where(query, args).Find(&v).Error
}
