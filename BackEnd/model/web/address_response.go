package web

type AddressResponse struct {
	Id        int            `json:"id"`
	UserId    int            `json:"user_id"`
	Kabupaten string         `json:"kabupaten"`
	Kecamatan string         `json:"kecamatan"`
	Kelurahan string         `json:"kelurahan"`
	Alamat    string         `json:"alamat"`
	Note      string         `json:"note"`
	User      []UserResponse `json:"user"`
}
