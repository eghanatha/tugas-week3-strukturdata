package view

import (
	"arrayOfStruct/model"
	"arrayOfStruct/node"
	"bufio"
	"fmt"
	"os"
)


func Insert() {
	var kota, nama, notelp, email string
	var id, nomer int
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Masukkkan id pegawai : ")
	fmt.Scan(&id)

	fmt.Print("Masukkan nama pegawai : ")
	nama, _ = reader.ReadString('\n')
	nama = nama[:len(nama)-1]

	fmt.Print("Masukkan jalan : ")
	jalan, _ := reader.ReadString('\n')
	jalan = jalan[:len(jalan)-1]

	fmt.Print("Masukkan nomer rumah : ")
	fmt.Scan(&nomer)

	fmt.Print("Masukkan kota : ")
	fmt.Scan(&kota)

	fmt.Print("Masukkan nomer telepon : ")
	fmt.Scan(&notelp)

	fmt.Print("Masukkan email : ")
	fmt.Scan(&email)

	pegawai := node.Pegawai {
		ID: id,
		Nama: nama,
		Alamat: node.Address{jalan, kota, nomer},
		NoTelp: notelp,
		Email: email,
	}

	cek := model.UpdatePegawai(pegawai, id)
	if cek {
		fmt.Println("-- Pegawai berhasil diapdate --")
	} else {
		fmt.Println("pegawai gagal update")
	}
	fmt.Println()
}

func Views() {
	fmt.Println("Daftar Pegawai")
	for i, emp := range model.ReadPegawai() {
		fmt.Println("--- Pegawai ke -", i+1, "---")
		fmt.Println("ID Pegawai\t: ", emp.ID)
		fmt.Println("Nama Pegawai\t: ", emp.Nama)
		fmt.Println("Alamat\t\t: ", emp.Alamat.Jalan, emp.Alamat.Nomer, emp.Alamat.Kota)
		fmt.Println("No Telepon\t: ", emp.NoTelp)
		fmt.Println("Email\t: ", emp.Email)
		fmt.Println()
	}
}

func Delete() {
	var id int
	fmt.Print("Masukkan id pegawai yang akan dihapus : ")
	fmt.Scan(&id)

	cek := model.DeletePegawai(id)
	if cek {
		fmt.Println("-- Pegawai berhasil dihapus --")
	} else {
		fmt.Println("Pegawai gagal dihapus")
	}
	fmt.Println()
}