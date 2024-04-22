package main

import (
	"testing"
)

func TestPembayaranBarang(t *testing.T) {
	tests := []struct {
		hargaTotal         float64
		metodePembayaran   string
		dicicil            bool
		expectedError      string
		expectedSuccessful bool
	}{
		{hargaTotal: 0, metodePembayaran: "credit", dicicil: false, expectedError: "harga tidak bisa nol", expectedSuccessful: false},
		{hargaTotal: 500000, metodePembayaran: "cash", dicicil: false, expectedError: "Metode tidak dikenali", expectedSuccessful: false},
		{hargaTotal: 500000, metodePembayaran: "credit", dicicil: true, expectedError: "", expectedSuccessful: true},
		{hargaTotal: 100000, metodePembayaran: "debit", dicicil: true, expectedError: "credit harus dicicil", expectedSuccessful: false},
		{hargaTotal: 200000, metodePembayaran: "ovo", dicicil: false, expectedError: "Metode tidak dikenali", expectedSuccessful: false},
	}

	for _, test := range tests {
		err := PembayaranBarang(test.hargaTotal, test.metodePembayaran, test.dicicil)
		if test.expectedSuccessful && err != nil {
			t.Errorf("Expected successful payment, but got error: %v", err)
		}
		if !test.expectedSuccessful && err == nil {
			t.Error("Expected payment error, but got nil")
		}
		if err != nil && err.Error() != test.expectedError {
			t.Errorf("Expected error message '%s', but got '%s'", test.expectedError, err.Error())
		}
	}
}
