package database

type Verify struct {
	ID   string `gorm:"primaryKey"`
	Role string
	Type int
}

func (v *Verify) Create() error {
	return db.Create(&v).Error
}

func (v *Verify) Find(query string, args ...string) error {
	return db.Where(query, args).Find(&v).Error
}
