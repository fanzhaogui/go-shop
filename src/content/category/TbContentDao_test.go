package category

import (
	"testing"
	"fmt"
)

func TestGetContentByIDDao(t *testing.T) {
	r := GetContentByIDDao(31)
	fmt.Println(r)
}
