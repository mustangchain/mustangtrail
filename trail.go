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

// Package trail provides the recording domain model.
package trail

import (
	"fmt"
	"strconv"
)

// Date specifies a Gregorian calendar day as the numeric value of
// the 8 decimal ISO 8601 basic format.
type Date uint32

// Split returns the corresponding year, month and day of the month.
// Months count from 1 [January] to 12 [December] conform time.Month.
func (d Date) Split() (year, month, day int) {
	i := int(d)
	return i / 10000, (i / 100) % 100, i % 100
}

// String returns the ISO 8601 extended format.
func (d Date) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", d/10000, (d/100)%100, d%100)
}

// ParseDate parses an ISO calendar date.
// The return is zero for malformed input.
func ParseDate(s string) Date {
	switch len(s) {
	case 8: // basic format YYYYMMDD
		i, _ := strconv.ParseInt(s, 10, 32)
		return Date(i)

	case 10: // extended format YYYY-MM-DD
		if s[4] != '-' || s[7] != '-' {
			return 0
		}
		i, _ := strconv.ParseInt(s[:4]+s[5:7]+s[8:], 10, 32)
		return Date(i)

	default:
		return 0
	}
}
