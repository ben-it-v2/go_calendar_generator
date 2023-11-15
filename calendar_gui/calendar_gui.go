package calendar_gui

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/gui"
	"strconv"
	"sync"
	"calendar/palette_gui"
	"calendar/pdf_gui"
	"calendar/format"
	"calendar/shared_gui"
	"calendar/config"
	"calendar/settings_gui"
)

// var selectedPalette int = palette_gui.PNone
var mainWindow *widgets.QMainWindow
var isGenerating bool = false
var typeButton *widgets.QPushButton

/*
	DayWidget
*/

type DayWidget struct {
    *widgets.QWidget
	monthWidget *MonthWidget
	pType int
	layoutDay *widgets.QHBoxLayout
	dayNameLabel *widgets.QLabel
	dayIDLabel *widgets.QLabel
	tagLabel *widgets.QLabel
	hour int
}

func NewDayWidget(monthWidget *MonthWidget, dayName string, dayID string) *DayWidget {
	widgetDay := &DayWidget{widgets.NewQWidget(nil, core.Qt__Widget), monthWidget, palette_gui.PClear, widgets.NewQHBoxLayout2(nil), widgets.NewQLabel(nil, core.Qt__Widget), widgets.NewQLabel(nil, core.Qt__Widget), widgets.NewQLabel(nil, core.Qt__Widget), 0}

	button := widgets.NewQPushButton2("", widgetDay)
	button.SetStyleSheet(`
		background-color: rgba(0, 0, 0, 0);
		padding: 0;
	`)
	button.ConnectClicked(func(_ bool) {
		widgetDay.onClick(palette_gui.SelectedPalette)
	})

	layoutDayMargin := core.NewQMargins2(1, 1, 1, 1)
	widgetDay.layoutDay.SetContentsMargins2(layoutDayMargin)
	widgetDay.layoutDay.SetSpacing(0)

	widgetDay.dayNameLabel.SetAlignment(core.Qt__AlignVCenter)
	widgetDay.dayNameLabel.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	widgetDay.dayNameLabel.SetText(dayName)
	widgetDay.dayNameLabel.SetMinimumWidth(25)
	widgetDay.dayNameLabel.SetStyleSheet("border: 1;")
	widgetDay.layoutDay.AddWidget(widgetDay.dayNameLabel, 0, 0)

	widgetDay.dayIDLabel.SetAlignment(core.Qt__AlignVCenter)
	widgetDay.dayIDLabel.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	widgetDay.dayIDLabel.SetText(dayID)
	widgetDay.dayIDLabel.SetMinimumWidth(30)
	widgetDay.dayIDLabel.SetStyleSheet("border: 1;")
	widgetDay.layoutDay.AddWidget(widgetDay.dayIDLabel, 0, 0)

	widgetDay.tagLabel.SetAlignment(core.Qt__AlignVCenter)
	widgetDay.tagLabel.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	widgetDay.tagLabel.SetStyleSheet("border: 1;")
	widgetDay.layoutDay.AddWidget(widgetDay.tagLabel, 0, 0)

	widgetDay.layoutDay.AddStretch(1)
	widgetDay.SetLayout(widgetDay.layoutDay)
	if dayName == "S" || dayName == "D" {
		widgetDay.SetWeekEnd()
	} else {
		widgetDay.SetFreeDay()
	}

	button.SetFixedSize2(widgetDay.Width(), widgetDay.Height())
	button.Raise()
	widgetDay.Update()

	return (widgetDay)
}

func (wDay *DayWidget) setStyleSheet2(hexColor string) {
	wDay.SetStyleSheet(`
		background-color: ` + hexColor + `;
		color: #333;
		font-family: Arial, sans-serif;
		font-size: 14px;
		border: 1px solid #ccc;
		border-radius: 0px;
		padding: 3px;
	`)
}

func (wDay *DayWidget) SetWeekEnd() {
    wDay.setStyleSheet2(palette_gui.CWeekEnd)
	wDay.pType = palette_gui.PWeekEnd
	wDay.updateHour(0)
}

func (wDay *DayWidget) SetWorkDay() {
	wDay.setStyleSheet2(palette_gui.CWork)
	wDay.pType = palette_gui.PWork
	wDay.updateHour(7)
}

func (wDay *DayWidget) SetHoliday() {
	wDay.setStyleSheet2(palette_gui.CHoliday)
	wDay.pType = palette_gui.PHoliday
	wDay.tagLabel.SetText("FÉRIÉ")
	wDay.tagLabel.SetStyleSheet(`
		color: #FFFFFF;
		border: 1;
	`)
	wDay.dayNameLabel.SetStyleSheet(`
		color: #FFFFFF;
		border: 1;
	`)
	wDay.dayIDLabel.SetStyleSheet(`
		color: #FFFFFF;
		border: 1;
	`)
	wDay.updateHour(0)
}

func (wDay *DayWidget) SetClosingDay() {
	wDay.setStyleSheet2(palette_gui.CClosing)
	wDay.pType = palette_gui.PClosing
	wDay.updateHour(0)
}

func (wDay *DayWidget) SetRevisionDay() {
	wDay.setStyleSheet2(palette_gui.CRevision)
	wDay.pType = palette_gui.PRevision
	wDay.updateHour(7)
}

func (wDay *DayWidget) SetExamDay() {
	wDay.setStyleSheet2(palette_gui.CExam)
	wDay.pType = palette_gui.PExam
	wDay.tagLabel.SetStyleSheet(`
		color: #FFFFFF;
		border: 1;
	`)
	wDay.dayNameLabel.SetStyleSheet(`
		color: #FFFFFF;
		border: 1;
	`)
	wDay.dayIDLabel.SetStyleSheet(`
		color: #FFFFFF;
		border: 1;
	`)
	wDay.updateHour(7)
}

func (wDay *DayWidget) SetVisioDay() {
	wDay.setStyleSheet2(palette_gui.CVisio)
	wDay.pType = palette_gui.PVisio
	wDay.updateHour(4)
}

func (wDay *DayWidget) SetFreeDay() {
	wDay.setStyleSheet2(palette_gui.CClear)
	wDay.pType = palette_gui.PClear
	wDay.updateHour(0)
}

func (wDay *DayWidget) SetIntegrationDay() {
	wDay.setStyleSheet2(palette_gui.CIntegration)
	wDay.pType = palette_gui.PIntegration
	wDay.updateHour(7)
}

func (wDay *DayWidget) updateHour(hour int) {
	wDay.hour = hour
	wDay.monthWidget.UpdateHour()
}

func (wDay *DayWidget) Reset() {
	wDay.tagLabel.SetText("")
	wDay.tagLabel.SetStyleSheet(`
		color: black;
		border: 1;
	`)
	wDay.dayNameLabel.SetStyleSheet(`
		color: black;
		border: 1;
	`)
	wDay.dayIDLabel.SetStyleSheet(`
		color: black;
		border: 1;
	`)
}

func (wDay *DayWidget) onClick(sPalette int) {
	if wDay.pType != palette_gui.PWeekEnd {
		wDay.Reset()
		switch sPalette {
			case palette_gui.PClear:
				wDay.SetFreeDay()
			case palette_gui.PWork:
				wDay.SetWorkDay()
			case palette_gui.PClosing:
				wDay.SetClosingDay()
			case palette_gui.PHoliday:
				wDay.SetHoliday()
			case palette_gui.PVisio:
				wDay.SetVisioDay()
			case palette_gui.PRevision:
				wDay.SetRevisionDay()
			case palette_gui.PExam:
				wDay.SetExamDay()
			case palette_gui.PIntegration:
				wDay.SetIntegrationDay()
		}
		updateTotalHour()
	}
}

var formationType string = ""
var containerLayout *widgets.QHBoxLayout
var yearHeader *widgets.QLabel
var beginDate *core.QDate
var loadingWidget *LoadingWidget
var containerWidget *widgets.QWidget
var pdfButton *shared_gui.PushButtonAnimated
var workDay int = 0

/*
	ContainerLayout
*/

type ContainerLayout struct {
	*widgets.QHBoxLayout
	mutex sync.Mutex
}

func NewContainerLayout() *ContainerLayout {
	containerLayout := &ContainerLayout{}
	return (containerLayout)
}

/*
	MonthWidget
*/

type MonthWidget struct {
	*widgets.QWidget
	vLayout *widgets.QVBoxLayout
	header *widgets.QLabel
	hours int
	dayWidgets []*DayWidget
	totalHours *widgets.QLabel
}

func NewMonthWidget(monthID int) *MonthWidget {
	monthWidget := &MonthWidget{widgets.NewQWidget(nil, core.Qt__Widget), widgets.NewQVBoxLayout2(nil), widgets.NewQLabel(nil, core.Qt__Widget), 0, make([]*DayWidget, 31), widgets.NewQLabel(nil, core.Qt__Widget)}
	
	margin := core.NewQMargins2(1, 0, 0, 0)
	monthWidget.vLayout.SetContentsMargins2(margin)
	monthWidget.vLayout.SetSpacing(0)

	monthDate := core.NewQDate3(beginDate.Year(), beginDate.Month(), 1)
	monthDate = monthDate.AddMonths(monthID)
	curMonth := monthDate.Month()

	monthWidget.header.SetAlignment(core.Qt__AlignCenter)
	monthWidget.header.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	monthWidget.header.SetText(format.TranslateMonthNameToFrench(curMonth))
	monthWidget.header.SetMaximumHeight(50)
	monthWidget.header.SetStyleSheet(`
		background-color: #218b9a;
		color: white;
		font-family: Arial, sans-serif;
		font-size: 16px;
		border-radius: 0px;
		padding: 3px;
	`)
	monthWidget.vLayout.AddWidget(monthWidget.header, 0, 0)

	for i := 0; i < 31; i++ {
		dayDate := monthDate.AddDays(int64(i))
		if (dayDate.Month() != curMonth) {
			spacerWidget := widgets.NewQWidget(nil, core.Qt__Widget)
			// spacerWidget.SetStyleSheet("background-color: rgba(255, 0, 0, 255);")
			spacerWidget.SetMinimumHeight(24)
			monthWidget.vLayout.AddWidget(spacerWidget, 0, 0)
			continue
		}
		dayWidget := NewDayWidget(monthWidget, format.FormatDayName(dayDate), strconv.Itoa(dayDate.Day()))
		monthWidget.dayWidgets[i] = dayWidget
		if workDay == dayDate.DayOfWeek() &&
			((dayDate.Year() == beginDate.Year() && dayDate.Month() == beginDate.Month() && dayDate.Day() >= beginDate.Day()) || 
			(dayDate.Year() != beginDate.Year() || dayDate.Month() != beginDate.Month())) {
			if (integrationDays >= 2) {
				dayWidget.SetWorkDay()
			} else {
				integrationDays += 1
				dayWidget.SetIntegrationDay()
			}
		}
		dateStr := strconv.Itoa(dayDate.Day()) + "/" + strconv.Itoa(dayDate.Month())
		if _, ok := appConfig.ClosingDays[dateStr]; ok {
			dayWidget.SetClosingDay()
		}
		if _, ok := appConfig.Holidays[dateStr]; ok {
			dayWidget.SetHoliday()
		}
		monthWidget.vLayout.AddWidget(dayWidget, 0, 0)
	}

	spacerWidget := widgets.NewQWidget(nil, core.Qt__Widget)
	// spacerWidget.SetStyleSheet("background-color: rgba(255, 0, 0, 255);")
	spacerWidget.SetMinimumHeight(24)
	monthWidget.vLayout.AddWidget(spacerWidget, 0, 0)

	monthWidget.totalHours.SetAlignment(core.Qt__AlignCenter)
	monthWidget.totalHours.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	monthWidget.totalHours.SetStyleSheet(`
		background-color: #F2F2F2;
		color: black;
		font-family: Arial, sans-serif;
		font-size: 16px;
		border-radius: 0px;
		border: 1px solid black;
		padding: 3px;
	`)
	monthWidget.totalHours.SetMinimumHeight(25)
	monthWidget.vLayout.AddWidget(monthWidget.totalHours, 0, 0)

	monthWidget.vLayout.AddStretch(1)
	monthWidget.SetLayout(monthWidget.vLayout)
	monthWidget.SetMinimumWidth(110)
	monthWidget.SetStyleSheet("border: 1;")

	return (monthWidget)
}

func (mWidget *MonthWidget) addHour(hours int) {
	mWidget.hours += hours
}

func (mWidget *MonthWidget) UpdateHour() {
	mWidget.hours = 0
	for i := 0; i < 31; i++ {
		wDay := mWidget.dayWidgets[i]
		if (wDay != nil) {
			mWidget.addHour(wDay.hour)
		}
	}
	mWidget.totalHours.SetText(strconv.Itoa(mWidget.hours))
}

func updateTotalHour() {
	total := 0
	for i := 0; i < len(monthsCache); i++ {
		curMonth := monthsCache[i]
		if (curMonth != nil) {
			total += monthsCache[i].hours
		}
	}
	totalHoursCounter.SetText(strconv.Itoa(total))
}

/*
	MyThread
*/

type MyThread struct {
	*core.QThread
}

func NewMyThread() *MyThread {
	thread := &MyThread{core.NewQThread(nil)}
	thread.SetPriority(core.QThread__HighPriority)
	return (thread)
}

/*
	LoadingWidget
*/

type LoadingWidget struct {
	*widgets.QWidget
	vLayout *widgets.QVBoxLayout
	header *widgets.QLabel
	barContainer *widgets.QWidget
	bar *widgets.QWidget
	max int
	value int
}

func NewLoadingWidget(maxValue int) *LoadingWidget {
	wLoading := &LoadingWidget{widgets.NewQWidget(nil, core.Qt__Widget), widgets.NewQVBoxLayout(), widgets.NewQLabel2("Chargement du Calendrier...", nil, 0), widgets.NewQWidget(nil, core.Qt__Widget), widgets.NewQWidget(nil, core.Qt__Widget), maxValue, 0}

	wLoading.header.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	wLoading.header.SetMaximumHeight(50)
	wLoading.header.SetStyleSheet(`
		font-family: Arial, sans-serif;
		font-size: 20px;
		border: 1;
	`)
	wLoading.vLayout.AddWidget(wLoading.header, 0, core.Qt__AlignCenter)

	wLoading.barContainer.SetMaximumHeight(80)
	wLoading.barContainer.SetFixedSize2(wLoading.header.Width() * 80 / 100, 50)
	wLoading.barContainer.SetStyleSheet(`
		background-color: #F2F2F2;
		color: #333;
		border: 1px solid #bbb;
		border-radius: 5px;
	`)

	hLayout := widgets.NewQHBoxLayout()
	margin := core.NewQMargins2(5, 5, 5, 5)
	hLayout.SetContentsMargins2(margin)
	hLayout.SetSpacing(0)

	wLoading.bar.SetStyleSheet(`
		background-color: #218b9a;
		border: 1;
		padding: 0;
	`)
	hLayout.AddWidget(wLoading.bar, 0, core.Qt__AlignLeft)

	wLoading.barContainer.SetLayout(hLayout)

	wLoading.vLayout.AddWidget(wLoading.barContainer, 0, core.Qt__AlignCenter)

	wLoading.SetLayout(wLoading.vLayout)
	wLoading.SetStyleSheet(`
		background-color: #f2f2f2;
		color: #333;
		border: 1px solid #bbb;
		border-radius: 5px;
	`)
	return (wLoading)
}

func (wLoading *LoadingWidget) Increase(value int) {
	wLoading.value += value
	fValue := float64(wLoading.barContainer.Width()) * float64(wLoading.value) / float64(wLoading.max)
	wLoading.bar.SetFixedSize2(int(fValue), wLoading.bar.Height())
}

var monthsCache []*MonthWidget
var years string = ""
var integrationDays = 0;

func updateCalendarWidget() {
	isGenerating = true
	integrationDays = 0
	maxMonth := 12
	if beginDate.Day() >= 11 {
		maxMonth = 13
	}

	if beginDate.Month() + maxMonth > 13 {
		years = strconv.Itoa(beginDate.Year()) + "/" + strconv.Itoa(beginDate.Year() + 1)
		yearHeader.SetText("Calendrier " + years)
	} else {
		years = strconv.Itoa(beginDate.Year())
		yearHeader.SetText("Calendrier " + years)
	}

	// Clear current calendar
	for containerLayout.Count() > 0 {
        child := containerLayout.TakeAt(0).Widget()
        child.DeleteLater()
    }

	// Add Loading Widget
	loadingWidget = NewLoadingWidget(maxMonth)
	containerLayout.AddWidget(loadingWidget, 0, core.Qt__AlignCenter)

	monthsCache = make([]*MonthWidget, maxMonth)

	// Create calendar widgets for each month and add them to the grid layout
	thread := NewMyThread()
	thread.ConnectStarted(func() {
		fmt.Println("Thread started:", core.QThread_CurrentThread())
	})
	thread.ConnectFinished(func() {
		fmt.Println("Thread finished", core.QThread_CurrentThread())
		isGenerating = false
	})
	thread.ConnectRun(func() {
		fmt.Println("Thread running", core.QThread_CurrentThread())
		for i := 0; i < maxMonth; i++ {
			monthsCache[i] = NewMonthWidget(i)
			loadingWidget.Increase(1)
		}
		loadingWidget.DeleteLater()
		for i:= 0; i < maxMonth; i++ {
			containerLayout.AddWidget(monthsCache[i], 0, 0)
		}
		updateTotalHour()
		pdfButton.SetDisabled(false)
	})
	thread.Start()

	fmt.Println("Current Thread:", core.QThread_CurrentThread())
}

func createHeaderWidget() *widgets.QWidget {
	hLayout := widgets.NewQHBoxLayout2(nil)
	widgetHLayout := widgets.NewQWidget(nil, core.Qt__Widget)

	// Create a new label to display the header title
	titleLabel := widgets.NewQLabel(nil, core.Qt__Widget)
	titleLabel.SetAlignment(core.Qt__AlignLeft)
	titleLabel.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	titleLabel.SetText("Planning de formation")
	titleLabel.SetMaximumHeight(48)
	font := gui.NewQFont2("Arial", 24, 1, false)
	titleLabel.SetFont(font)
	hLayout.AddWidget(titleLabel, 0, 0)

	settingsButton := shared_gui.NewPushButtonAnimated(&mainWindow.QWidget, "assets/settings_icone.png", func(_ bool) {
		mSettings := settings_gui.NewSettingsMenu(appConfig, updateFormationTypes)
		mSettings.Exec()
	})
	hLayout.AddWidget(settingsButton, 0, 0)

	pdfButton = shared_gui.NewPushButtonAnimated(&mainWindow.QWidget, "assets/pdf_icone.png", func(_ bool) {
		mPdf := pdf_gui.NewPDFMenu(nameValue, firstnameValue, formationType, years, workDay, beginDate, containerWidget, appConfig)
		mPdf.Exec()
	})
	pdfButton.SetDisabled(true)
	hLayout.AddWidget(pdfButton, 0, 0)

	widgetHLayout.SetLayout(hLayout)
	widgetHLayout.SetMaximumHeight(100)
	// widgetHLayout.SetStyleSheet("border: 1px solid black;")
	return (widgetHLayout)
}

func updateFormationTypes() {
	typeMenu := widgets.NewQMenu(nil)
	typeMenu.SetTitle("Formations")

	for key, value := range appConfig.FormationTypes {
		fmt.Println(key + " - " + value)
		// Add the action to the menu
		fKey := key
		fValue := value
		action := typeMenu.AddAction(key + " - " + value)
		action.ConnectTriggered(func(_ bool) {
			typeButton.SetText(fKey)
			formationType = fValue
		})
	}

	typeButton.SetMenu(typeMenu)
}

func createFormationTypeWidget() *widgets.QWidget {
	// Type
	typeLayout := widgets.NewQHBoxLayout2(nil)
	widgetType := widgets.NewQWidget(nil, core.Qt__Widget)

	// Type label
	typeLabel := widgets.NewQLabel(nil, core.Qt__Widget)
	typeLabel.SetAlignment(core.Qt__AlignCenter)
	typeLabel.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	typeLabel.SetText("Type de formation")
	typeLabel.SetWordWrap(true)
	typeLayout.AddWidget(typeLabel, 0, 0)

	// Type select submenu
	typeButton = widgets.NewQPushButton2("", nil)
	typeButton.SetMinimumWidth(100)
	typeButton.SetMinimumHeight(30)
	typeButton.SetStyleSheet(`
		background-color: #f2f2f2;
		color: #333;
		border: 1px solid #bbb;
		border-radius: 5px;
		padding: 0;
	`)
	typeLayout.AddWidget(typeButton, 0, 0)

	updateFormationTypes()

	widgetType.SetLayout(typeLayout)
	widgetType.SetStyleSheet("border: 1;")
	return (widgetType)
}

func createFormationDayWidget() *widgets.QWidget {
	// Day
	dayLayout := widgets.NewQHBoxLayout2(nil)
	widgetDay := widgets.NewQWidget(nil, core.Qt__Widget)

	// Day label
	dayLabel := widgets.NewQLabel(nil, core.Qt__Widget)
	dayLabel.SetAlignment(core.Qt__AlignCenter)
	dayLabel.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	dayLabel.SetText("Jour de formation")
	dayLabel.SetWordWrap(true)
	dayLayout.AddWidget(dayLabel, 0, 0)

	// Day select submenu
	dayButton := widgets.NewQPushButton2("", nil)
	dayButton.SetMinimumWidth(100)
	dayButton.SetMinimumHeight(30)
	dayButton.SetStyleSheet(`
		background-color: #f2f2f2;
		color: #333;
		border: 1px solid #bbb;
		border-radius: 5px;
		padding: 0;
	`)
	dayLayout.AddWidget(dayButton, 0, 0)

	dayMenu := widgets.NewQMenu(nil)
	dayMenu.SetTitle("Jours")

	for i := 1; i < len(appConfig.FormationDays); i++ {
		// Add the action to the menu
		dayKey := i
		dayValue := appConfig.FormationDays[i]
		action := dayMenu.AddAction(dayValue)
		action.ConnectTriggered(func(_ bool) {
			dayButton.SetText(dayValue)
			workDay = dayKey
		})
	}

	// Set the menu for the submenu button
	dayButton.SetMenu(dayMenu)

	widgetDay.SetLayout(dayLayout)
	widgetDay.SetStyleSheet("border: 1;")
	return (widgetDay)
}

func createBeginDateWidget() *widgets.QWidget {
	// Begin Day
	vLayout := widgets.NewQVBoxLayout2(nil)
	vWidget := widgets.NewQWidget(nil, core.Qt__Widget)

	hLayout := widgets.NewQHBoxLayout2(nil)
	hWidget := widgets.NewQWidget(nil, core.Qt__Widget)

	dateHeader := widgets.NewQLabel(nil, core.Qt__Widget)
	dateHeader.SetAlignment(core.Qt__AlignHCenter)
	dateHeader.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	dateHeader.SetText("Date de début:")
	dateHeader.SetMaximumHeight(50)
	dateHeader.SetStyleSheet(`
		font-family: Arial, sans-serif;
		font-size: 16px;
	`)
	hLayout.AddWidget(dateHeader, 0, 0)

	dateSelected := widgets.NewQLabel(nil, core.Qt__Widget)
	dateSelected.SetAlignment(core.Qt__AlignHCenter)
	dateSelected.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	dateSelected.SetText("Aucune!")
	dateSelected.SetMaximumHeight(50)
	dateSelected.SetStyleSheet(`
		font-family: Arial, sans-serif;
		font-size: 16px;
	`)
	hLayout.AddWidget(dateSelected, 0, 0)

	hWidget.SetLayout(hLayout)
	hWidget.SetStyleSheet("border: 1;")
	vLayout.AddWidget(hWidget, 0, 0)

	dayCalendar := widgets.NewQCalendarWidget(nil)
	dayCalendar.SetGridVisible(true)
	dayCalendar.SetStyleSheet(`
		color: black;
		QCalendarWidget QAbstractItemView:section {
			color: black;
		}
	`)
	dayCalendar.ConnectClicked(func(date *core.QDate) {
		dateSelected.SetText(date.ToString("dd-MM-yyyy"))
		beginDate = core.NewQDate3(date.Year(), date.Month(), date.Day())
	})
	vLayout.AddWidget(dayCalendar, 0, 0)
	
	vWidget.SetLayout(vLayout)
	return (vWidget)
}

var nameValue string = ""
var firstnameValue string = ""

func createFormWidget() *widgets.QWidget {
	vLayout := widgets.NewQVBoxLayout2(nil)
	widgetVLayout := widgets.NewQWidget(nil, core.Qt__Widget)

	nameEdit := widgets.NewQLineEdit2("", nil)
	nameEdit.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	nameEdit.SetPlaceholderText("Nom")
	nameEdit.SetMaximumHeight(30)
	nameEdit.SetMinimumHeight(30)
	nameEdit.ConnectEditingFinished(func() {
		nameValue = nameEdit.Text()
	})
	vLayout.AddWidget(nameEdit, 0, 0)

	firstnameEdit := widgets.NewQLineEdit2("", nil)
	firstnameEdit.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	firstnameEdit.SetPlaceholderText("Prénom")
	firstnameEdit.SetMaximumHeight(30)
	firstnameEdit.SetMinimumHeight(30)
	firstnameEdit.ConnectEditingFinished(func() {
		firstnameValue = firstnameEdit.Text()
	})
	vLayout.AddWidget(firstnameEdit, 0, 0)

	vLayout.AddWidget(createFormationTypeWidget(), 0, 0)
	vLayout.AddWidget(createFormationDayWidget(), 0, 0)
	vLayout.AddWidget(createBeginDateWidget(), 0, 0)

	generateButton := widgets.NewQPushButton2("Générer le Calendrier", nil)
	generateButton.SetMinimumHeight(30)
	generateButton.SetMaximumHeight(30)
	generateButton.SetStyleSheet(`
		background-color: #f2f2f2;
		color: #333;
		border: 1px solid #bbb;
		border-radius: 5px;
		padding: 0;
	`)
	generateButton.ConnectClicked(func(_ bool) {
		if (!isGenerating && beginDate != nil && workDay != 0) {
			fmt.Println("Generating")
			updateCalendarWidget()
		}
    })
	vLayout.AddWidget(generateButton, 0, 0)

	vLayout.AddStretch(1)
	widgetVLayout.SetLayout(vLayout)
	widgetVLayout.SetStyleSheet(`
		background-color: #f2f2f2;
		color: #333;
		border: 1px solid #bbb;
		border-radius: 5px;
		padding: 0;
	`)
	widgetVLayout.SetMinimumSize2(150, 400)
	return (widgetVLayout)
}

var widgetPalettes *widgets.QWidget

func createCalendarOptionsWidget() *widgets.QWidget {
	vLayout := widgets.NewQVBoxLayout2(nil)
	widgetVLayout := widgets.NewQWidget(nil, core.Qt__Widget)

	vLayout.AddWidget(createFormWidget(), 0, 0)
	widgetPalettes = palette_gui.CreatePalettesWidget()
	vLayout.AddWidget(widgetPalettes, 0, 0)
	vLayout.AddStretch(1)

	widgetVLayout.SetLayout(vLayout)
	return (widgetVLayout)
}

var totalHoursCounter *widgets.QLabel

func createCalendarContainerWidget() *widgets.QWidget {
	vLayout := widgets.NewQVBoxLayout2(nil)
	containerWidget = widgets.NewQWidget(nil, core.Qt__Widget)

	vMargin := core.NewQMargins2(0, 0, 0, 0)
	vLayout.SetContentsMargins2(vMargin)
	vLayout.SetSpacing(2)

	yearHeader = widgets.NewQLabel(nil, core.Qt__Widget)
	yearHeader.SetAlignment(core.Qt__AlignCenter)
	yearHeader.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	yearHeader.SetText("En attente d'informations...")
	yearHeader.SetMaximumHeight(50)
	yearHeader.SetStyleSheet(`
		background-color: #218b9a;
		color: white;
		font-family: Arial, sans-serif;
		font-size: 24px;
		border-radius: 0px;
		padding: 3px;
	`)
	vLayout.AddWidget(yearHeader, 0, 0)

	containerLayout = widgets.NewQHBoxLayout()
	hWidget := widgets.NewQWidget(nil, core.Qt__Widget)

	margin := core.NewQMargins2(0, 0, 0, 0)
	containerLayout.SetContentsMargins2(margin)
	containerLayout.SetSpacing(2)

	hWidget.SetLayout(containerLayout)
	hWidget.SetFixedSize2(1405, 850)
	hWidget.SetStyleSheet("border: 1;")
	vLayout.AddWidget(hWidget, 0, core.Qt__AlignCenter)

	totalHoursLayout := widgets.NewQHBoxLayout()
	totalHoursWidget := widgets.NewQWidget(nil, core.Qt__Widget)

	totalHoursLabel := widgets.NewQLabel(nil, core.Qt__Widget)
	totalHoursLabel.SetAlignment(core.Qt__AlignCenter)
	totalHoursLabel.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	totalHoursLabel.SetText("Total (en heure) :")
	totalHoursLabel.SetStyleSheet(`
		font-family: Arial, sans-serif;
		font-size: 16px;
		border: 1;
	`)
	totalHoursLabel.SetMinimumHeight(25)
	totalHoursLayout.AddWidget(totalHoursLabel, 0, core.Qt__AlignVCenter)

	totalHoursCounter = widgets.NewQLabel(nil, core.Qt__Widget)
	totalHoursCounter.SetAlignment(core.Qt__AlignCenter)
	totalHoursCounter.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	totalHoursCounter.SetText("0")
	totalHoursCounter.SetStyleSheet(`
		font-family: Arial, sans-serif;
		font-size: 16px;
		border: 1;
	`)
	totalHoursCounter.SetMinimumHeight(25)
	totalHoursLayout.AddWidget(totalHoursCounter, 0, core.Qt__AlignVCenter)

	totalHoursWidget.SetLayout(totalHoursLayout)
	totalHoursWidget.SetStyleSheet(`
		background-color: #F2F2F2;
		color: black;
		font-family: Arial, sans-serif;
		font-size: 16px;
		border-radius: 0px;
		border: 1px solid black;
		padding: 3px;
	`)
	vLayout.AddWidget(totalHoursWidget, 0, core.Qt__AlignHCenter)
	
	vLayout.AddStretch(1)
	containerWidget.SetLayout(vLayout)
	containerWidget.SetFixedSize2(1410, 950)
	containerWidget.SetStyleSheet(`
		background-color: #f2f2f2;
		color: #333;
		border: 1px solid #bbb;
		border-radius: 5px;
		padding: 0;
	`)
	return (containerWidget)
}

func createBodyWidget() *widgets.QWidget {
	hLayout := widgets.NewQHBoxLayout2(nil)
	widgetHLayout := widgets.NewQWidget(nil, core.Qt__Widget)

	hLayout.AddWidget(createCalendarOptionsWidget(), 0, 0)
	hLayout.AddWidget(createCalendarContainerWidget(), 0, 0)
	hLayout.AddStretch(1)

	widgetHLayout.SetLayout(hLayout)
	widgetHLayout.SetMinimumWidth(1200)
	return (widgetHLayout)
}

func CreateMainWindow(argc int, args []string) {
	app := widgets.NewQApplication(argc, args)

	mainWindow = widgets.NewQMainWindow(nil, 0)
	mainWindow.SetWindowTitle("CFA Calendrier")
	windowIcon := gui.NewQIcon5("assets/idmn_logo.png")
	mainWindow.SetWindowIcon(windowIcon)

	vLayout := widgets.NewQVBoxLayout2(nil)
	widgetLayout := widgets.NewQWidget(nil, core.Qt__Widget)

	vLayout.AddWidget(createHeaderWidget(), 0, 0)
	vLayout.AddWidget(createBodyWidget(), 0, 0)
	vLayout.AddStretch(1)

	widgetLayout.SetLayout(vLayout)

	scrollArea := widgets.NewQScrollArea(nil)
	scrollArea.SetWidgetResizable(true)
	scrollArea.SetWidget(widgetLayout)
	mainWindow.SetCentralWidget(scrollArea)

	// Maximize the window to fill the screen
	mainWindow.ShowMaximized()

	// Show the main window
	mainWindow.Show()

	// Execute the application
	app.Exec()
}

var appConfig *config.AppConfig

func init() {
	// Create Application Configuration
	appConfig = config.NewConfig()

	fmt.Println("Package 'Calendar GUI': ✔️")
}
