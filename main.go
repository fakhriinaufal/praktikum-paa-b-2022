package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// deklarasi variabel
	// var namaVariabel tipeData = nilainya
	// var nama string = "Fakhrii"

	// deklarasi sekarang, diisi nanti
	// var namaAsdos string
	// namaAsdos = "Ovi"

	// deklarasi dengan tipe data otomatis dari golang
	// namaVariabel := nilai
	// namaPraktikan := "Dion"

	// tipe data => int, float32, bool, string

	// array | kumpulan data yang kapasitasnya ditentukan
	// tipedata[ukuran]
	// students := [2]string{"Dion", "Ovi"}
	// var students [2]string = [2]string{"Dion", "Ovi"}
	// var students = [...]{"Dion", "Ovi"}

	// slice | kumpulan data yang nilai nya dinamis / tidak ditentukan
	// students := []string{"Dion", "Ovi"}

	// students := []string{"Dion", "Ovi"}

	// menambahkan item pada slice
	// students = append(students, "Fakhrii")

	// foreach / iterasi tiap elemen dari array / struct
	// for _, student := range students {
	// 	fmt.Println("Bernama", student)
	// }

	// pointer | alamat suatu data dalam memori
	// var name string
	// var namePointer *string

	// name = "Ovi"
	// namePointer = &name

	// fmt.Println(name)
	// fmt.Println(*namePointer)

	// name = "Fakhrii"
	// fmt.Println(name)
	// fmt.Println(*namePointer)

	// map tipe data berbentuk key-value
	// var students map[string]string

	// students = map[string]string{}

	// students["1020"] = "Fakhrii"
	// students["3003"] = "Ovi"

	// students := map[string]string{
	// 	"3003": "Ovi",
	// }

	// fmt.Println(students["3003"])

	e := echo.New()

	handleHello := func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	}

	e.GET("/", handleHello)

	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, json{
			"status":  "OK",
			"message": "hello world",
		})
	})

	e.Logger.Fatal(e.Start(":3000"))
}

type json map[string]interface{}
