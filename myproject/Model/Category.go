package Model

type Category struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Name      string `json:"name" gorm:"type:varchar(50);not null;unique"`
	CreatedAt Time   `json:"createdAt" gorm:"type:tiemstamp"`
	UpdatedAt Time   `json:"updatedAt" gorm:"type:timestamp"`
}
