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

var goldenRFID15s = []struct {
	Feed string
	Want RFID15
	// parts
	Manufacturer, ID string
}{
	{"", 0, "", ""},
	{"012345678987654", 12345678987654, "012", "345678987654"},
	{"321098765432100", 321098765432100, "321", "098765432100"},
	{"1234567890-1234", 0, "", ""},
	{"1234567890X1234", 0, "", ""},
}

func TestGoldenRFID15s(t *testing.T) {
	for _, gold := range goldenRFID15s {
		got := ParseRFID15(gold.Feed)
		if got != gold.Want {
			t.Errorf("%s: parsed to %d, want %d", gold.Feed, got, gold.Want)
		}
		if got == 0 {
			continue
		}

		if s := got.String(); s != gold.Feed {
			t.Errorf("%s: got string %q", gold.Feed, s)
		}
		if s := got.Manufacturer(); s != gold.Manufacturer {
			t.Errorf("%s: got manufacturer %q, want %q", gold.Feed, s, gold.Manufacturer)
		}
		if s := got.ID(); s != gold.ID {
			t.Errorf("%s: got ID %q, want %q", gold.Feed, s, gold.ID)
		}
	}
}

var goldenUELNs = []struct {
	Feed string
	OK   bool
	// parts
	Country, Database, NationalID string
}{
	{"", false, "", "", ""},
	{"372414001234567", true, "372", "414", "001234567"},
	{"25000100155928M", true, "250", "001", "00155928M"},
	{"25000100155928m", true, "250", "001", "00155928M"},
	{"12a456012345678", false, "", "", ""},
	{"12345a012345678", false, "", "", ""},
	{"123456-12345678", false, "", "", ""},
	{"1234560123456789", false, "", "", ""},
}

func TestGoldenUELNs(t *testing.T) {
	for _, gold := range goldenUELNs {
		got, ok := ParseUELN(gold.Feed)
		if ok != gold.OK {
			t.Errorf("%s: parse OK %t, want %t", gold.Feed, ok, gold.OK)
			continue
		}
		if !gold.OK {
			var want UELN
			if got != want {
				t.Errorf("%s: parse got %#x, want %#x", gold.Feed, got, want)
			}

			continue
		}

		if s := got.Country(); s != gold.Country {
			t.Errorf("%s: got country %q, want %q", gold.Feed, s, gold.Country)
		}
		if s := got.Database(); s != gold.Database {
			t.Errorf("%s: got database %q, want %q", gold.Feed, s, gold.Database)
		}
		if s := got.NationalID(); s != gold.NationalID {
			t.Errorf("%s: got national ID %q, want %q", gold.Feed, s, gold.NationalID)
		}

		want := gold.Country + gold.Database + gold.NationalID
		if s := got.String(); s != want {
			t.Errorf("%s: got string %q, want %q", gold.Feed, s, want)
		}
	}
}
