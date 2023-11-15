package config

import (
	"os"
	"bufio"
	"log"
	"strings"
	"fmt"
	"strconv"
)

/*
	AppConfig
*/

type AppConfig struct {
	FormationTypes map[string]string
	FormationDays []string
	ClosingDays map[string]bool
	ClosingDaysSort []string
	Holidays map[string]bool
	HolidaysSort []string
	OutputDir string
}

func (conf *AppConfig) fillFormationDays(path string) {
	file, err := os.Open(path)
	if err == nil {
		s := bufio.NewScanner(file)
		i := 1
		for s.Scan() {
			conf.FormationDays[i] = s.Text()
			i++
		}
		file.Close()
	} else {
		log.Fatal(err)
	}
}

func (conf *AppConfig) fillFormationTypes(path string) {
	file, err := os.Open(path)
	if err == nil {
		s := bufio.NewScanner(file)
		for s.Scan() {
			split := strings.Split(s.Text(), "#")
			if len(split) == 2 {
				conf.FormationTypes[split[0]] = split[1]
			} else {
				log.Fatal("Erreur dans la sauvegarde des types de formation!\n" + s.Text())
			}
		}
		file.Close()
	} else {
		log.Fatal(err)
	}
}

func (conf *AppConfig) fillGenericMap(cur_map map[string]bool, cur_array_sort []string, path string) {
	file, err := os.Open(path)
	if err == nil {
		s := bufio.NewScanner(file)
		i := 0
		for s.Scan() {
			cur_map[s.Text()] = true
			cur_array_sort[i] = s.Text()
			i++
		}
		file.Close()
	} else {
		log.Fatal(err)
	}
}

func (conf *AppConfig) SaveFormationTypes() {
	file, err := os.OpenFile("config/formations.txt", os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)
	if err == nil {
		defer file.Close()
		for key, value := range conf.FormationTypes {
			_, err := file.WriteString(key + "#" + value + "\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		log.Fatal(err)
	}
}

func (conf *AppConfig) genericSave(path string, cur_array []string) {
	file, err := os.OpenFile(path, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)
	if err == nil {
		defer file.Close()
		for _, value := range cur_array {
			if (value != "") {
				_, err := file.WriteString(value + "\n")
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	} else {
		log.Fatal(err)
	}
}

func (conf *AppConfig) SaveHolidays() {
	conf.genericSave("config/holidays.txt", conf.HolidaysSort)
}

func (conf *AppConfig) SaveClosingDays() {
	conf.genericSave("config/closingdays.txt", conf.ClosingDaysSort)
}

func (conf *AppConfig) InsertDateAndSort(cur_array []string, strDate string) {
	strTab := strings.Split(strDate, "/")
	dayID, _ := strconv.Atoi(strTab[0])
	monthID, _ := strconv.Atoi(strTab[1])

	// Find ID where inserts the new date
	for i := 0; i < len(cur_array); i++ {
		curStr := cur_array[i]
		nextStr := cur_array[i + 1]
		if curStr != "" && nextStr != "" {
			curStrTab := strings.Split(curStr, "/")
			nextStrTab := strings.Split(nextStr, "/")
			curDayID, _ := strconv.Atoi(curStrTab[0])
			curMonthID, _ := strconv.Atoi(curStrTab[1])
			nextDayID, _ := strconv.Atoi(nextStrTab[0])
			nextMonthID, _ := strconv.Atoi(nextStrTab[1])

			if (monthID >= curMonthID && dayID >= curDayID && ((monthID == nextMonthID && dayID < nextDayID) || (monthID < nextMonthID))) {
				fmt.Println(curStr, nextStr)
				i++
				cur_array[i] = strDate
				for i += 1; i < len(cur_array); i++ {
					tmp := cur_array[i]
					cur_array[i] = nextStr
					nextStr = tmp
				}
				break
			}
		} else {
			if (curStr == "") {
				cur_array[i] = strDate
			} else {
				cur_array[i + 1] = strDate
			}
			break
		}
	}
}

func (conf *AppConfig) FindEmptyIndex(cur_array []string) int {
	for i := 0; i < len(cur_array); i++ {
		str := cur_array[i]
		if (str == "") {
			return (i)
		}
	}
	return (-1)
}

func (conf *AppConfig) fillOutputDirectory() {
	file, err := os.Open("config/outputdirectory.txt")
	if err == nil {
		s := bufio.NewScanner(file)
		if s.Scan() {
			conf.OutputDir = s.Text()
		}
		file.Close()
	} else {
		log.Fatal(err)
	}
}

func (conf *AppConfig) SaveOutputDirectory(outputdirectory string) {
	conf.OutputDir = outputdirectory
	file, err := os.OpenFile("config/outputdirectory.txt", os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)
	if err == nil {
		defer file.Close()
		_, err := file.WriteString(outputdirectory)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}

func NewConfig() *AppConfig {
	appC := &AppConfig{make(map[string]string), make([]string, 6), make(map[string]bool), make([]string, 30), make(map[string]bool), make([]string, 30), "."}
	appC.fillFormationDays("config/days.txt")
	appC.fillFormationTypes("config/formations.txt")
	appC.fillGenericMap(appC.Holidays, appC.HolidaysSort, "config/holidays.txt")
	appC.fillGenericMap(appC.ClosingDays, appC.ClosingDaysSort, "config/closingdays.txt")
	appC.fillOutputDirectory()
	return (appC)
} 
