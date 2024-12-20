// Package qris contains functions to generate and process QRIS.
package qris

import (
	"fmt"

	"github.com/sigurn/crc16"
)

// crc calculate CRC16 from QRIS data
func crc(qrisString string) string {
	// Calculate CRC16
	table := crc16.MakeTable(crc16.CRC16_CCITT_FALSE)

	crc := crc16.Checksum([]byte(qrisString), table)
	return fmt.Sprintf("%04X", crc)
}
