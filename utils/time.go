package utils

import (
	"fmt"
	"time"
)

var (
	bulan = map[int]string{
		1:  "Januari",
		2:  "Februari",
		3:  "Maret",
		4:  "April",
		5:  "Mei",
		6:  "Juni",
		7:  "Juli",
		8:  "Agustus",
		9:  "September",
		10: "Oktober",
		11: "November",
		12: "Desember",
	}

	hari = map[time.Weekday]string{
		time.Sunday:    "Minggu",
		time.Monday:    "Senin",
		time.Tuesday:   "Selasa",
		time.Wednesday: "Rabu",
		time.Thursday:  "Kamis",
		time.Friday:    "Jum'at",
		time.Saturday:  "Sabtu",
	}
)

func FormatIndonesiaTime(t time.Time) string {
	hour := fmt.Sprintf("%d", t.Hour())
	if t.Hour() < 10 {
		hour = fmt.Sprintf("0%d", t.Hour())
	}
	minute := fmt.Sprintf("%d", t.Minute())
	if t.Minute() < 10 {
		minute = fmt.Sprintf("0%d", t.Minute())
	}
	return fmt.Sprintf("%s, %d %s %d %s:%s WIB", hari[t.Weekday()], t.Day(), bulan[int(t.Month())], t.Year(), hour, minute)
}
