package qrpayment

import (
	"fmt"
	"github.com/sigurn/crc16"
	"strings"
)

// generateTLVFormat generates a TLV (Tag-Length-Value) formatted string from a map of fields.
// The fields can be strings, maps of strings, or nested maps of interfaces.
// It recursively processes nested maps to generate the TLV format.
//
// Parameters:
// - fields: A map where the key is a string and the value can be a string, map[string]string, or map[string]interface{}.
//
// Returns:
// - A string representing the TLV formatted data.
func generateTLVFormat(fields map[string]interface{}) string {
	var tlv []string
	for key, value := range fields {
		switch v := value.(type) {
		case string:
			tlv = append(tlv, fmt.Sprintf("%s%02d%s", key, len(v), v))
		case map[string]string:
			var subTlv []string
			for subKey, subValue := range v {
				subTlv = append(subTlv, fmt.Sprintf("%s%02d%s", subKey, len(subValue), subValue))
			}
			joinedSub := strings.Join(subTlv, EmptyString)
			tlv = append(tlv, fmt.Sprintf("%s%02d%s", key, len(joinedSub), joinedSub))
		case map[string]interface{}:
			var subTlv []string
			for subKey, subValue := range v {
				switch subVal := subValue.(type) {
				case string:
					subTlv = append(subTlv, fmt.Sprintf("%s%02d%s", subKey, len(subVal), subVal))
				case map[string]string:
					var innerTlv []string
					for innerKey, innerValue := range subVal {
						innerTlv = append(innerTlv, fmt.Sprintf("%s%02d%s", innerKey, len(innerValue), innerValue))
					}
					joinedInner := strings.Join(innerTlv, EmptyString)
					subTlv = append(subTlv, fmt.Sprintf("%s%02d%s", subKey, len(joinedInner), joinedInner))
				}
			}
			joinedSub := strings.Join(subTlv, EmptyString)
			tlv = append(tlv, fmt.Sprintf("%s%02d%s", key, len(joinedSub), joinedSub))
		}
	}

	return strings.Join(tlv, EmptyString)
}

// generateChecksumWithCRC16 generates a CRC16 checksum for the given data string.
// It appends a CRC16 tag to the data, calculates the checksum, and returns the data with the checksum.
//
// Parameters:
// - data: The input string for which the checksum is to be generated.
//
// Returns:
// - A string representing the input data appended with the CRC16 checksum.
func generateChecksumWithCRC16(data string) string {
	str := fmt.Sprintf("%s%s", data, CRC16Tag)
	checksum := crc16.Checksum([]byte(str), crc16.MakeTable(crc16.CRC16_CCITT_FALSE))
	return fmt.Sprintf("%s%04X", str, checksum)
}
