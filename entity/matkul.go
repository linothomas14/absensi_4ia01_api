package entity

type Matkul struct {
	ID     uint8  `gorm:"type:int(10);not null;unique;autoIncrement:true" json:"id"`
	Matkul string `gorm:"type:varchar(50);not null" json:"matkul"`
}
