package entity

type Presensi struct {
	ID     uint8  `gorm:"primaryKey;autoIncrement:true;" json:"id"`
	NPM    string `gorm:"type:varchar(9);not null;" json:"npm"`
	Matkul string `gorm:"type:varchar(50);" json:"matkul"`
	Minggu uint8  `gorm:"type:int;" json:"minggu"`
}

func (Presensi) TableName() string {
	return "presensi"
}
