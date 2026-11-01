package pdf417

import (
	"image"
	"reflect"
	"testing"

	"github.com/boombuler/barcode"
)

func sameImage(first, second image.Image) bool {
	if first.Bounds() != second.Bounds() {
		return false
	}
	bounds := first.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			if first.At(x, y) != second.At(x, y) {
				return false
			}
		}
	}
	return true
}

func Test_ECICodewords(t *testing.T) {
	tests := []struct {
		eci   barcode.ECI
		words []int
	}{
		{barcode.NoECI, nil},
		{0, []int{927, 0}},
		{barcode.ECIISO8859_1, []int{927, 3}},
		{barcode.ECIUTF8, []int{927, 26}},
		{899, []int{927, 899}},
		{900, []int{926, 0, 0}},
		{810899, []int{926, 899, 899}},
		{810900, []int{925, 0}},
		{811799, []int{925, 899}},
	}
	for _, test := range tests {
		got, err := eciCodewords(test.eci)
		if err != nil || !reflect.DeepEqual(got, test.words) {
			t.Errorf("eciCodewords(%d) = %v, %v, want %v",
				test.eci, got, err, test.words)
		}
	}

	for _, eci := range []barcode.ECI{-2, 811800, barcode.MaxECI} {
		if _, err := eciCodewords(eci); err == nil {
			t.Errorf("ECI %d should be refused", eci)
		}
	}
}

func Test_EncodeWithECI(t *testing.T) {
	data := "Caf\u00e9"
	plain, err := Encode(data, 2)
	if err != nil {
		t.Fatal(err)
	}
	none, err := EncodeWithECI(data, 2, barcode.NoECI)
	if err != nil {
		t.Fatal(err)
	}
	if !sameImage(plain, none) {
		t.Error("NoECI draws a different symbol from Encode")
	}
	withECI, err := EncodeWithECI(data, 2, barcode.ECIUTF8)
	if err != nil {
		t.Fatal(err)
	}
	if sameImage(plain, withECI) {
		t.Error("an ECI draws the same symbol as no ECI")
	}
	if _, err := EncodeWithECI(data, 2, barcode.MaxECI); err == nil {
		t.Errorf("ECI %d should be refused", barcode.MaxECI)
	}
}
