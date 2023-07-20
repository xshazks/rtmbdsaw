package rtmsaw

import (
	"fmt"
	"testing"
	"github.com/aiteung/atdb"
)

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "dbmonitor",
}

func TestInsertmonitor(t *testing.T) {
	proker := "Melakukan maintance website"
	status := "Dalam pengerjaan"
	about := "Memperbaiki website karena terdapat kendala"
	karyawan := "Acep Suhendar"
	hsl := Insertmonitor(proker, status, about, karyawan)
	fmt.Println(hsl)
}