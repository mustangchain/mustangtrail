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
	// parts
	Year, Month, Day int
	// formats
	Extended, Basic string
}{
	{20180605, 2018, 6, 5, "2018-06-05", "20180605"},
}

func TestGoldenDates(t *testing.T) {
	for _, gold := range goldenDates {
		if got := gold.Date.String(); got != gold.Extended {
			t.Errorf("%d: got string %q, want %q", gold.Date, got, gold.Extended)
		}
		if got := ParseDate(gold.Extended); got != gold.Date {
			t.Errorf("%d: got %d for %q", gold.Date, got, gold.Extended)
		}
		if got := ParseDate(gold.Basic); got != gold.Date {
			t.Errorf("%d: got %d for %q", gold.Date, got, gold.Basic)
		}
		if y, m, d := gold.Date.Split(); y != gold.Year || m != gold.Month || d != gold.Day {
			t.Errorf("%d: got (%d, %d, %d), want (%d, %d, %d)",
				gold.Date, y, m, d, gold.Year, gold.Month, gold.Day)
		}
	}
}
