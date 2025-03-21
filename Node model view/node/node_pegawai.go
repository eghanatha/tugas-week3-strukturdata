package node

type Address struct {
	Jalan,Kota string
	Nomer string 
}

type Pegawai struct {
	ID int
	Nama string
	Alamat Address
	NoTelp string
	Email string
}