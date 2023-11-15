package format

import (
	"github.com/therecipe/qt/core"
)

func FormatDayName(dayDate *core.QDate) string {
	dayID := int(dayDate.DayOfWeek())
	dayName := ""
	switch dayID {
		case 1:
			dayName = "L"
		case 2:
			dayName = "M"
		case 3:
			dayName = "M"
		case 4:
			dayName = "J"
		case 5:
			dayName = "V"
		case 6:
			dayName = "S"
		case 7:
			dayName = "D"
	}

	return (dayName)
}

func FormatDayName2(dayofweek int) string {
	dayName := ""
	switch dayofweek {
		case 1:
			dayName = "Lundi"
		case 2:
			dayName = "Mardi"
		case 3:
			dayName = "Mercredi"
		case 4:
			dayName = "Jeudi"
		case 5:
			dayName = "Vendredi"
		case 6:
			dayName = "Samedi"
		case 7:
			dayName = "Dimanche"
	}

	return (dayName)
}

var frenchMonths = []string{"Inconnu", "Janvier", "Février", "Mars", "Avril", "Mai", "Juin", "Juillet", "Août", "Septembre", "Octobre", "Novembre", "Décembre"}
func TranslateMonthNameToFrench(monthID int) string {
	name := frenchMonths[0]
	if (monthID >= 1 || monthID <= 12) {
		name = frenchMonths[monthID]
	}
	return (name)
}
