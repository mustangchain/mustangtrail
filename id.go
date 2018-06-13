// Copyright 2018 MustangChain Foundation. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package trail

import (
	"fmt"
	"unicode"
)

// RFID15 is a 15 decimal implant number based on ISO 11784 & 11785.
type RFID15 uint64

// ParseRFID15 returns the parsed value if, and only if, s is syntactically correct.
func ParseRFID15(s string) RFID15 {
	if len(s) != 15 {
		return 0
	}

	var n uint64
	for i := 0; i < 15; i++ {
		d := uint64(s[i]) - '0'
		if d > 9 {
			return 0
		}
		n = n*10 + d
	}

	return RFID15(n)
}

// String returns the 15 decimal code.
func (id RFID15) String() string {
	return fmt.Sprintf("%015d", id)
}

// Manufacturer returns the 3 decimal code from the 10-bit country code field.
// See https://www.service-icar.com/tables/Tabella3.php
func (id RFID15) Manufacturer() string {
	return fmt.Sprintf("%03d", id/1E12)
}

// ID returns the 12 decimal code from the 38-bit ID field.
func (id RFID15) ID() string {
	return fmt.Sprintf("%012d", id%1E12)
}

// UELN (Universal Equine Life Number) is a horse-specific identification number
// that can be used to identify each horse individual.
// See http://www.ueln.net/ueln-presentation/rules-of-attribution-of-the-ueln/
type UELN string

// ParseUELN returns the parsed value if, and only if, s is syntactically correct.
func ParseUELN(s string) UELN {
	if len(s) < 15 {
		return ""
	}

	for i := 0; i < 6; i++ {
		if c := s[i]; c < '0' || c > '9' {
			return ""
		}
	}

	var n int
	for _, r := range s[6:] {
		n++
		if !unicode.IsDigit(r) && !unicode.IsLetter(r) {
			return ""
		}
		// BUG(pascaldekloe): Are UELNs case sensitive?
	}
	if n != 9 {
		return ""
	}

	return UELN(s)
}

// Country returns the 3 decimal ISO-3166 country code
// of the database which registered the foal at birth.
func (n UELN) Country() string {
	return string(n[:3])
}

// Database returns the 3 decimal code of the database
// where the horse has been registered at birth.
func (n UELN) Database() string {
	return string(n[3:6])
}

// NationalID returns the 9 character horse national
// identification number given by the stud-book of birth.
func (n UELN) NationalID() string {
	return string(n[6:])
}