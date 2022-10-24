package entities

type User struct {
	Id          int64
	NamaLengkap string `validate:"required" label:"Nama Lengkap"`
	Email       string `validate:"required" label:"Email"`
	Username    string `validate:"required" label:"Username"`
	Password    string `validate:"required" label:"Password"`
	Cpassword   string `validate:"required" label:"Konfirmasi Password"`
	Role		string `validate:"required" label:"Role"`
}
