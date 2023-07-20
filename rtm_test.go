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

var MongoConn = atdb.MongoConnect(MongoInfo)

func TestInsertmonitor(t *testing.T) {
	proker := "Melakukan maintance website"
	status := "Dalam pengerjaan"
	about := "Memperbaiki website karena terdapat kendala"
	karyawan := "Acep Suhendar"
	hsl := Insertmonitor(proker, status, about, karyawan)
	fmt.Println(hsl)
}

func TestGetDatakaryawan(t *testing.T) {
	karyawan := "Acep Suhendar"
	dt := GetDatakaryawan(karyawan)
	fmt.Println(dt)
}

func TestGetDatamonitor(t *testing.T) {
	status := "Dalam pengerjaan"
	dt := GetDatamonitor(status)
	fmt.Println(dt)
}

func TestGetDataproker(t *testing.T) {
	about := "Melakukan maintance website"
	dt := GetDataproker(about)
	fmt.Println(dt)
}
