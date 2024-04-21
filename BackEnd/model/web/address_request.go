package web

type AddressCreateRequest struct {
	Kabupaten string `json:"kabupaten"`
	Kecamatan string `json:"kecamatan"`
	Kelurahan string `json:"kelurahan"`
	Alamat    string `json:"alamat"`
	Note      string `json:"note"`
}

type AddressUpdateRequest struct {
	Id        int    `json:"id"`
	Kabupaten string `json:"kabupaten"`
	Kecamatan string `json:"kecamatan"`
	Kelurahan string `json:"kelurahan"`
	Alamat    string `json:"alamat"`
	Note      string `json:"note"`
}
