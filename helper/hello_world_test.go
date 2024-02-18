package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Eko")
	if result != "Halo Eko" {
		// Unit Test gagal
		panic("Hasilnya gagal")
	}
}

func TestHelloWorldSecond(t *testing.T) {
	result := HelloWorld("Kedua")
	if result != "Halo Kedua" {
		// Unit Test gagal
		panic("Hasilnya gagal")
	}
}
