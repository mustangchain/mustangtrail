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

// Date specifies a Gregorian calendar day (with 23 bits).
type Date uint

// Split returns the corresponding year, month and day of the month.
// Months count from 1 [January] to 12 [December] conform time.Month.
func (d Date) Split() (year, month, day int) {
	return int(d >> 9), int(d >> 5 & 15), int(d & 31)
}

// String returns the ISO 8601 extended format.
func (d Date) String() string {
	year, month, day := d.Split()
	return fmt.Sprintf("%04d-%02d-%02d", year, month, day)
}

// ParseDate parses an ISO calendar date.
// The return is zero for malformed input.
func ParseDate(s string) Date {
	var yyyy, mm, dd string
	switch len(s) {
	case 8: // basic format YYYYMMDD
		yyyy, mm, dd = s[:4], s[4:6], s[6:]
	case 10: // extended format YYYY-MM-DD
		if s[4] != '-' || s[7] != '-' {
			return 0
		}
		yyyy, mm, dd = s[:4], s[5:7], s[8:]
	default:
		return 0
	}

	year, err := strconv.ParseUint(yyyy, 10, 14)
	if err != nil {
		return 0
	}
	month, err := strconv.ParseUint(mm, 10, 4)
	if err != nil {
		return 0
	}
	day, err := strconv.ParseUint(dd, 10, 5)
	if err != nil {
		return 0
	}
	return Date(year<<9 | month<<5 | day)
}
