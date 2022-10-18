package dto

//MahasiswaUpdateDTO is a entity that client use when updating a Mahasiswa
// type MahasiswaUpdateDTO struct {
// 	ID          uint64 `json:"id" form:"id" binding:"required"`
// 	Title       string `json:"title" form:"title" binding:"required"`
// 	Description string `json:"description" form:"description" binding:"required"`
// 	UserID      uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
// }

//MahasiswaCreateDTO is is a entity that client use when create a new Mahasiswa
type MahasiswaCreateDTO struct {
	NPM  string `json:"npm" form:"npm" binding:"required"`
	Nama string `json:"nama" form:"nama" binding:"required"`
}
