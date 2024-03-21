package database

type Reaction struct {
	ID     string `gorm:"primaryKey"`
	Role1  string
	Role2  string
	Role3  string
	Role4  string
	Role5  string
	Role6  string
	Role7  string
	Role8  string
	Role9  string
	Role10 string
}

func (v *Reaction) Create() error {
	return db.Create(&v).Error
}

func (v *Reaction) Find(query string, args ...string) (err error) {
	return db.Where(query, args).Find(&v).Error
}
