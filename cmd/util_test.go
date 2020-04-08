/*
Copyright © 2020 Daniel J. Lauk <daniel.lauk@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package cmd

import (
	"testing"
	"time"
)

var startOfToday = startOfDay(time.Now())

func TestStartOfDay(t *testing.T) {
	val := startOfDay(time.Date(2020, 3, 24, 9, 10, 11, 12, time.Local))
	expected := time.Date(2020, 3, 24, 0, 0, 0, 0, time.Local)
	if !val.Equal(expected) {
		t.Errorf("val incorrect, expected %v, got %v", expected, val)
	}
}

func TestParseTimeUnderstandsNow(t *testing.T) {
	val := parseTime("now")
	expected := time.Now()
	if val.Sub(expected) > 0 {
		t.Errorf("Time math incorrect: val > expected")
	}
	if expected.Sub(val).Milliseconds() > 100 {
		t.Errorf("val incorrect, expected %v, got %v", expected, val)
	}
}

func TestParseTimeHHMM(t *testing.T) {
	val := parseTime("08:09")
	expected := startOfToday.Add(8 * time.Hour).Add(9 * time.Minute)
	if !val.Equal(expected) {
		t.Errorf("val incorrect, expected %v, got %v", expected, val)
	}
}

func TestParseTimeHHMMSS(t *testing.T) {
	val := parseTime("08:09:10")
	expected := startOfToday.Add(8 * time.Hour).Add(9 * time.Minute).Add(10 * time.Second)
	if !val.Equal(expected) {
		t.Errorf("val incorrect, expected %v, got %v", expected, val)
	}
}

func TestParseTimeRelativePlus5Seconds(t *testing.T) {
	val := parseTime("+5s")
	expected := time.Now().Local().Add(5 * time.Second)
	if val.Sub(expected) > 0 {
		t.Errorf("Time math incorrect: val > expected")
	}
	if expected.Sub(val).Milliseconds() > 100 {
		t.Errorf("val incorrect, expected %v, got %v", expected, val)
	}
}

func TestParseTimeRelativeMinus5Seconds(t *testing.T) {
	val := parseTime("-5s")
	expected := time.Now().Local().Add(-5 * time.Second)
	if val.Sub(expected) > 0 {
		t.Errorf("Time math incorrect: val > expected")
	}
	if expected.Sub(val).Milliseconds() > 100 {
		t.Errorf("val incorrect, expected %v, got %v", expected, val)
	}
}

func TestParseTimeRelativePlus5Minutes(t *testing.T) {
	val := parseTime("+5m")
	expected := time.Now().Local().Add(5 * time.Minute)
	if val.Sub(expected) > 0 {
		t.Errorf("Time math incorrect: val > expected")
	}
	if expected.Sub(val).Milliseconds() > 100 {
		t.Errorf("val incorrect, expected %v, got %v", expected, val)
	}
}

func TestParseTimeRelativeMinus5Minutes(t *testing.T) {
	val := parseTime("-5m")
	expected := time.Now().Local().Add(-5 * time.Minute)
	if val.Sub(expected) > 0 {
		t.Errorf("Time math incorrect: val > expected")
	}
	if expected.Sub(val).Milliseconds() > 100 {
		t.Errorf("val incorrect, expected %v, got %v", expected, val)
	}
}

func TestParseTimeRelativePlus5Hours(t *testing.T) {
	val := parseTime("+5h")
	expected := time.Now().Local().Add(5 * time.Hour)
	if val.Sub(expected) > 0 {
		t.Errorf("Time math incorrect: val > expected")
	}
	if expected.Sub(val).Milliseconds() > 100 {
		t.Errorf("val incorrect, expected %v, got %v", expected, val)
	}
}

func TestParseTimeRelativeMinus5Hours(t *testing.T) {
	val := parseTime("-5h")
	expected := time.Now().Local().Add(-5 * time.Hour)
	if val.Sub(expected) > 0 {
		t.Errorf("Time math incorrect: val > expected")
	}
	if expected.Sub(val).Milliseconds() > 100 {
		t.Errorf("val incorrect, expected %v, got %v", expected, val)
	}
}

func TestParseTimeRelativePlusComplex(t *testing.T) {
	val := parseTime("+1h2m3s")
	expected := time.Now().Local().Add(1 * time.Hour).Add(2 * time.Minute).Add(3 * time.Second)
	if val.Sub(expected) > 0 {
		t.Errorf("Time math incorrect: val > expected")
	}
	if expected.Sub(val).Milliseconds() > 100 {
		t.Errorf("val incorrect, expected %v, got %v", expected, val)
	}
}

func TestParseTimeRelativeMinusComplex(t *testing.T) {
	val := parseTime("-1h2m3s")
	expected := time.Now().Local().Add(-1 * time.Hour).Add(-2 * time.Minute).Add(-3 * time.Second)
	if val.Sub(expected) > 0 {
		t.Errorf("Time math incorrect: val > expected")
	}
	if expected.Sub(val).Milliseconds() > 100 {
		t.Errorf("val incorrect, expected %v, got %v", expected, val)
	}
}

func TestParseDate(t *testing.T) {
	val := parseDate("2020-04-07")
	expected := time.Date(2020, time.April, 7, 0, 0, 0, 0, time.Local)
	if val.Sub(expected) > 0 {
		t.Errorf("Time math incorrect: val > expected: %v > %v", val, expected)
	}
	if expected.Sub(val).Milliseconds() > 100 {
		t.Errorf("val incorrect, expected %v, got %v", expected, val)
	}
}

func TestParseDateUnderstandsNow(t *testing.T) {
	val := parseDate("now")
	expected := time.Now()
	if val.Sub(expected) > 0 {
		t.Errorf("Time math incorrect: val > expected")
	}
	if expected.Sub(val).Milliseconds() > 100 {
		t.Errorf("val incorrect, expected %v, got %v", expected, val)
	}
}

func TestParseDateUnderstandsToday(t *testing.T) {
	val := parseDate("today")
	expected := startOfToday
	if val.Sub(expected) > 0 {
		t.Errorf("Time math incorrect: val > expected")
	}
	if expected.Sub(val).Milliseconds() > 100 {
		t.Errorf("val incorrect, expected %v, got %v", expected, val)
	}
}

func TestParseDateUnderstandsTomorrow(t *testing.T) {
	val := parseDate("tomorrow")
	expected := startOfToday.Add(24 * time.Hour)
	if val.Sub(expected) > 0 {
		t.Errorf("Time math incorrect: val > expected")
	}
	if expected.Sub(val).Milliseconds() > 100 {
		t.Errorf("val incorrect, expected %v, got %v", expected, val)
	}
}

func TestParseDateUnderstandsYesterday(t *testing.T) {
	val := parseDate("yesterday")
	expected := startOfToday.Add(-24 * time.Hour)
	if val.Sub(expected) > 0 {
		t.Errorf("Time math incorrect: val > expected")
	}
	if expected.Sub(val).Milliseconds() > 100 {
		t.Errorf("val incorrect, expected %v, got %v", expected, val)
	}
}
