package qris

import "testing"

func Test_crc(t *testing.T) {
	input := "00020101021126640017ID.CO.BANKBSI.WWW0118936004510000305481021000002039790303URE51440014ID.CO.QRIS.WWW0215ID10243141810570303URE5204866153033605802ID5918SABILILLAH YAYASAN6006MALANG6105651416304"
	expected := "522C"

	crc := crc(input)
	if crc != expected {
		t.Errorf("Expected %s, got %s", expected, crc)
	}
}
