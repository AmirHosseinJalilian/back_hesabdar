package convert_date

import (
	"fmt"
	"time"
)

func ConvertGregorianToSolar(gregorianDate time.Time) string {
	// Convert Gregorian date to Julian Day Number
	jdn := gregorianToJDN(gregorianDate)

	// Convert Julian Day Number to Solar Hijri date
	solarYear, solarMonth, solarDay := jdnToSolarHijri(jdn)

	// Format the Solar Hijri date
	solarDate := fmt.Sprintf("%04d/%02d/%02d", solarYear, solarMonth, solarDay)
	return solarDate
}

func gregorianToJDN(date time.Time) int {
	year, month, day := date.Date()
	a := (14 - int(month)) / 12
	y := year + 4800 - a
	m := int(month) + 12*a - 3
	jdn := day + (153*m+2)/5 + 365*y + y/4 - y/100 + y/400 - 32045
	return jdn
}

func jdnToSolarHijri(jdn int) (year, month, day int) {
	j := jdn - 1948320 + 10632
	n := (j - 1) / 10631
	j = j - 10631*n + 354
	j = (j + 59) / 354
	year = (int(j) * 33 / 32) + 4
	j = (j * 33 % 32) + 3
	month = (int(j) + 4) / 5
	day = int(j) - (month * 5) + 1
	return year, month, day
}
