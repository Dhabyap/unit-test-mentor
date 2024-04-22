package main

import (
	"errors"
	"fmt"
)

func PembayaranBarang(hargaTotal float64, metodePembayaran string, dicicil bool) error {
	if hargaTotal <= 0 {
		return errors.New("harga tidak bisa nol")
	}

	if metodePembayaran == "cod" || metodePembayaran == "transfer" || metodePembayaran == "debit" || metodePembayaran == "credit" || metodePembayaran == "gerai" {

		if hargaTotal >= 500000 && !dicicil {
			return errors.New("cicilan tidak memenuhi syarat")
		}

		if metodePembayaran != "credit" && dicicil {
			return errors.New("credit harus dicicil")
		}

		return nil

	} else {
		return errors.New("Metode tidak dikenali")
	}
}

func main() {
	harga := 400000.0
	metode := "cod"
	cicil := false

	err := PembayaranBarang(harga, metode, cicil)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Pembayaran berhasil.")
	}
}
