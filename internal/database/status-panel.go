package database

type StatusPanel struct {
	ID     string `gorm:"primaryKey"`
	Status bool   `gorm:"true"`
}

func (v *StatusPanel) Create() error {
	return db.Create(&v).Error
}

func (v *StatusPanel) Find(query string, args ...string) (err error) {
	return db.Where(query, args).Find(&v).Error
}

func (v *StatusPanel) Update() error {
	return db.Save(&v).Error
}
