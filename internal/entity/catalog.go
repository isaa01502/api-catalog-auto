package entity

type Catalog struct {
	ID     string `gorm:"type:uuid; primary_key;" json:"id"`
	RegNum string `gorm:"type:text; NOT NULL" json:"regNum"`
	Mark   string `gorm:"type:text; NOT NULL" json:"mark"`
	Model  string `gorm:"type:text; NOT NULL" json:"model"`
	Owner  string `gorm:"type:text; NOT NULL" json:"owner"`
	Year   int    `gorm:"NOT NULL" json:"year"`
}
