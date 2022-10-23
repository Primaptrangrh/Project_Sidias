package entities

type Permission struct {
	Id          int64
	NamaLengkap string `validate:"required" label:"Nama Lengkap"`
	Email       string `validate:"required" label:"Email"`
	Departemen  string `validate:"required" label:"Departemen"`
	Position    string `validate:"required" label:"Position"`
	Reason      string `validate:"required" label:"Konfirmasi Reason"`
}
