package mysql

import "testing"

func TestOpen(t *testing.T) {
	_, err := Open("root:CXZdsa12#$@tcp(192.168.50.198:3306)/aimodel?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		t.Fatal(err)
	}
}
