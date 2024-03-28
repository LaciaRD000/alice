package database

type Shop struct {
	ID                string `gorm:"primaryKey"`
	WelcomeMention    bool
	AlmostTicket      int
	WelcomeMessage    string
	SupportMemberRole string
	Category          string
}

func (v *Shop) Create() error {
	return db.Create(&v).Error
}

func (v *Shop) Find(query string, args ...string) (err error) {
	return db.Where(query, args).Find(&v).Error
}
