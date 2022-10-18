package entity

type Mahasiswa struct {
	ID   uint8  `gorm:"type:int(10);not null;unique;autoIncrement:true" json:"id"`
	NPM  string `gorm:"type:varchar(9);primaryKey;unique;not null" json:"npm"`
	Nama string `gorm:"type:varchar(50);not null" json:"nama"`
}

func (Mahasiswa) TableName() string {
	return "mahasiswa"
}
