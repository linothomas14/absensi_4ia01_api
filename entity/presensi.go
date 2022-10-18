package entity

// import "time"

type Presensi struct {
	ID        uint8     `gorm:"primaryKey;autoIncrement:true" json:"id"`
	NPM       string    `gorm:"type:varchar(9);not null" json:"npm"`
	Mahasiswa Mahasiswa `gorm:"foreignKey:NPM" json:"mahasiswa"`
	Matkul    string    `gorm:"type:varchar(50)" json:"matkul"`
	Tanggal   string    `gorm:"type:varchar(10)" json:"tanggal"`
}

func (Presensi) TableName() string {
	return "presensi"
}
