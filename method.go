package main

import "fmt"

func main(){
	var x2 Pelajar
	
	x2.nama = "Jack"
	x2.kelas = 11
	x2.umur = 16
	
	x2.namasaya()
}

type Pelajar struct{
	nama string
	kelas int
	umur int
}

func (s Pelajar) namasaya() {
	fmt.Println("Nama saya adalah",s.nama)	
}