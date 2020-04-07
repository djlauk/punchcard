package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/djlauk/punchcard/data"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// DateTimeFormat is a ISO8601 conforming date time string
const DateTimeFormat = "2006-01-02T15:04:05"

// DateFormat is a ISO8601 conforming date string
const DateFormat = "2006-01-02"

// TimeFormat is a ISO8601 conforming time string
const TimeFormat = "15:04:05"

// TimeFmtHHMM is a time format with hours and minutes only
const TimeFmtHHMM = "15:04"

func readData() *data.PunchcardData {
	dataFile, err := homedir.Expand(viper.GetString("storage.data"))
	if err != nil {
		log.Fatal(err)
	}
	if verbose {
		fmt.Printf("Reading data from %s\n", dataFile)
	}
	pcd, err := data.ReadData(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			if verbose {
				fmt.Println("Data file does not exist yet")
			}
		} else {
			log.Fatal(err)
		}
	}
	return pcd
}

func writeData(pcd *data.PunchcardData) {
	dataFile, err := homedir.Expand(viper.GetString("storage.data"))
	if err != nil {
		log.Fatal(err)
	}
	if verbose {
		fmt.Printf("Writing data to %s\n", dataFile)
	}
	if err = data.WriteData(dataFile, pcd); err != nil {
		log.Fatal(err)
	}
}

func startOfDay(t time.Time) time.Time {
	y, m, d := t.Local().Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local)
}

func parseTime(s string) time.Time {
	if s == "now" {
		return time.Now().Local()
	}
	if len(s) == len(TimeFmtHHMM) {
		s += ":00"
	}
	if len(s) != len(TimeFormat) {
		log.Fatalf("Time format not supported: %s", s)
	}
	t, err := time.ParseInLocation(TimeFormat, s, time.Local)
	if err != nil {
		log.Fatal(err)
	}
	hour, minute, second := t.Clock()
	year, month, day := time.Now().Date()
	result := time.Date(year, month, day, hour, minute, second, 0, time.Local)
	return result
}

func parseDate(s string) time.Time {
	if s == "now" {
		return time.Now().Local()
	}
	t, err := time.Parse(DateFormat, s)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

func parseDateTime(s string) time.Time {
	if s == "now" {
		return time.Now().Local()
	}
	t, err := time.Parse(DateTimeFormat, s)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

func formatDate(t *time.Time) string {
	return t.Local().Format(DateFormat)
}

func formatDateTime(t *time.Time) string {
	return t.Local().Format(DateTimeFormat)
}

func zeroTime() time.Time {
	var t time.Time
	return t
}
