package rtmsaw

import (
	"fmt"
	"testing"
)


func TestInsertmonitor(t *testing.T) {
	proker := "Pendaftaran"
	status := "Tidak sesuai target"
	about := "Membatasi waktu pendaftaran"
	karyawan := "Kelompok A"
	hsl := Insertmonitor(proker, status, about, karyawan)
	fmt.Println(hsl)
}