package entity

type Presensi struct {
	ID     uint8  `gorm:"primaryKey;autoIncrement:true;" json:"id"`
	NPM    string `gorm:"type:varchar(9);not null;" json:"-"`
	Matkul string `gorm:"type:varchar(50);" json:"matkul"`
	Minggu uint8  `gorm:"type:int(2);" json:"minggu"`
}

func (Presensi) TableName() string {
	return "presensi"
}
