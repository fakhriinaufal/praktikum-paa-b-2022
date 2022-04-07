package main

import "praktikum-paa-b-2022/routes"

func main() {
	e := routes.New()

	e.Start(":8080")
}
