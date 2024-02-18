package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldNew(t *testing.T) {
	result := HelloWorld("Eko")
	if result != "Halo Eko" {
		// Unit Test gagal
		// t.Fail()
		t.Error("Harusnya Halo Eko")
	}
	fmt.Println("Tes aja hehe")
}

func TestHelloWorldSecond(t *testing.T) {
	result := HelloWorld("Kedua")
	if result != "Halo Kedua" {
		// Unit Test gagal
		// t.FailNow()
		t.Fatal("Hasilnya Harus Halo Kedua")
	}

	fmt.Println("Tes aja hehe second")
}
