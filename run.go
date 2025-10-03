package main

import "fmt"

func main() {
	// integer and float
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga koma lima = ", 3.5)

	//bool
	fmt.Println("Benar = ", true)
	fmt.Println("Salah = ", false)

	fmt.Println("Ini adalah tipe data string")

	//string
	fmt.Println(len("Zelkova"))
	fmt.Println("Zelkova"[0])
	fmt.Println("Zelkova"[1])

	//variabel dengan inisiasi tipe data
	var name string
	name = "Zelkova Prunus"
	fmt.Println(name)

	//variabel tanpa mendefinisikan tipe data
	var name2 = "Zelkova Prunusssss"
	fmt.Println(name2)

	//tidak perlu pake keyword var
	var name3 = "zelkova keyaki"
	fmt.Println(name3)

	// kalo cuman ngubah nilai variabel bisa aja
	// cuman ngga bisa kalo define ulang variabel yg sudah ada
	name = "dsadawasdas"
	fmt.Println(name)

	//membuat multiple variable
	var (
		firstname = "abc"
		lastname  = "xyz"
	)
	fmt.Println(firstname, lastname)

	//constant
	// variabel yang valuenya tidak bisa di ubah
	const FIRSTNAME = "blaziken"
	const LASTNAME = "torchic"

	fmt.Println(FIRSTNAME, LASTNAME)

	// error
	//FIRSTNAME = "mudkip"

	const (
		FIRSTNAME1 = "charmander"
		LASTNAME2  = "Charmeleon"
	)
	fmt.Println(FIRSTNAME1, LASTNAME2)

	// konversi tipe data
	var nilai32 int32 = 32768
	var nilai64 = int64(nilai32)
	var nilai16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var nameEko = "Eko Kurniawan"
	var e uint8 = nameEko[0] // ini belum ngeslice berdasarkan index
	// ini baru ngerubah menjadi byte 69 harus di cast lagi make string

	// konversi e ke tipe yang string
	var eString string = string(e)

	fmt.Println(e)
	fmt.Println(eString)

	// type declaration
	//digunakan untuk membuat ALIAS atau nama lain dari tipe data yg sudah ada
	type NoKTP string

	var ktpEko NoKTP = "111111111111111"

	fmt.Println(ktpEko)
	fmt.Println(NoKTP("222222222222222222222"))

	type datetime string

	var tanggal string = "12102002"

	var tanggal_dttime datetime = datetime(tanggal)

	fmt.Println(tanggal_dttime)

}
