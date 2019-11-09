package param

import (
	"testing"
	"fmt"
)

func TestInsertParamService(t *testing.T) {
	re := InsertParamService(2, "sssssss")

	fmt.Println(re)
}

func TestCatidService(t *testing.T) {
	re := CatidService(2)
	fmt.Println(re)
}