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

import "testing"

var goldenDates = []*struct {
	Date Date
	// formats
	Extended, Basic string
	// parts
	Year, Month, Day int
}{
	{0x0fc4c5, "2018-06-05", "20180605", 2018, 6, 5},
	{0x00019f, "0000-12-31", "00001231", 0, 12, 31},
	{0x4e1e21, "9999-01-01", "99990101", 9999, 1, 1},
}

func TestGoldenDates(t *testing.T) {
	for _, gold := range goldenDates {
		if y, m, d := gold.Date.Split(); y != gold.Year || m != gold.Month || d != gold.Day {
			t.Errorf("split %s got (%d, %d, %d), want (%d, %d, %d)",
				gold.Extended, y, m, d, gold.Year, gold.Month, gold.Day)
		}
		if got := gold.Date.String(); got != gold.Extended {
			t.Errorf("got string %q, want %q", got, gold.Extended)
		}
		if got := ParseDate(gold.Extended); got != gold.Date {
			t.Errorf("parse %q got 0b%b, want 0b%b", gold.Extended, got, gold.Date)
		}
		if got := ParseDate(gold.Basic); got != gold.Date {
			t.Errorf("parse %q got 0b%b, want 0b%b", gold.Basic, got, gold.Date)
		}
	}
}
