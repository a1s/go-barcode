package qr

import (
	"fmt"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/utils"
)

func encodeAuto(content string, ecl ErrorCorrectionLevel, eci barcode.ECI) (*utils.BitList, *versionInfo, error) {
	bits, vi, _ := Numeric.getEncoder()(content, ecl, eci)
	if bits != nil && vi != nil {
		return bits, vi, nil
	}
	bits, vi, _ = AlphaNumeric.getEncoder()(content, ecl, eci)
	if bits != nil && vi != nil {
		return bits, vi, nil
	}
	bits, vi, _ = Unicode.getEncoder()(content, ecl, eci)
	if bits != nil && vi != nil {
		return bits, vi, nil
	}
	return nil, nil, fmt.Errorf("no encoding found to encode %q", content)
}
