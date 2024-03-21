package database

type Ticket struct {
	ID                string `gorm:"primaryKey"`
	UserID            string
	WelcomeMention    bool
	AlmostTicket      int
	WelcomeMessage    string
	SupportMemberRole string
}

func (v *Ticket) Create() error {
	return db.Create(&v).Error
}

func (v *Ticket) Find(query string, args ...string) (err error) {
	return db.Where(query, args).Find(&v).Error
}
