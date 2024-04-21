package domain

type Address struct {
	Id        int
	UserId    int
	Kabupaten string
	Kecamatan string
	Kelurahan string
	Alamat    string
	Note      string
	User      []User
}
