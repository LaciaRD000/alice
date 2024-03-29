package database

type UserLevel struct {
	GuildChannelID string `gorm:"primaryKey"`
	Level          int
	MessagesCount  int
}

func (v *UserLevel) Create() error {
	return db.Create(&v).Error
}

func (v *UserLevel) Update() error {
	return db.Save(&v).Error
}

func (v *UserLevel) Find(query interface{}, args ...interface{}) error {
	return db.Where(query, args).Find(&v).Error
}
