package settings_gui

import (
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"calendar/config"
	"calendar/shared_gui"
	"strings"
	"calendar/format"
	"fmt"
	"strconv"
	"runtime"
)

/*
	ListElement
*/

type ListElement struct {
	*widgets.QWidget
	deleteButton *shared_gui.PushButtonAnimated
}

func NewListElement(parent *widgets.QWidget, label string) *ListElement {
	element := &ListElement{widgets.NewQWidget(nil, core.Qt__Widget), shared_gui.NewPushButtonAnimated(parent, "assets/delete.png", func(_ bool) {})}
	element.SetMinimumHeight(64)

	hLayout := widgets.NewQHBoxLayout()
	curLabel := widgets.NewQLabel(nil, core.Qt__Widget)
	curLabel.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	curLabel.SetAlignment(core.Qt__AlignVCenter)
	curLabel.SetMaximumHeight(25)
	curLabel.SetStyleSheet(`
		border: 0;
		font-family: Arial, sans-serif;
		font-size: 16px;
	`)
	curLabel.SetText(label)
	hLayout.AddWidget(curLabel, 0, 0)

	element.deleteButton.SetMinimumSize2(48, 48)
	element.deleteButton.SetMaximumSize2(48, 48)
	element.deleteButton.InitialSize = core.NewQSize2(44, 44)
	element.deleteButton.TargetSize = core.NewQSize2(48, 48)
	element.deleteButton.SetIconSize(element.deleteButton.InitialSize)
	element.deleteButton.HoveredStyle = `
		QWidget {
			background-color: #f2f2f2;
			border: 0;
		}
		background-color: #f2f2f2;
		color: #333;
		border: 0;
		padding: 0;
	`
	margin := core.NewQMargins2(5, 5, 5, 5)
	hLayout.SetContentsMargins2(margin)
	hLayout.AddWidget(element.deleteButton, 0, 0)

	element.SetLayout(hLayout)
	element.SetStyleSheet(`
		background-color: #f2f2f2;
		color: #333;
		border: 1px solid #333;
		border-radius: 15px;
		padding: 0;
	`)
	return (element)
}

/*
	Config
*/

type CategoryData struct {
	name string
	viewer func(parent *SettingsMenu, conf *config.AppConfig)
}

var categories []CategoryData = []CategoryData{
	{"JOURS FERIES", func(mSettings *SettingsMenu, conf *config.AppConfig) {
		scrollArea := widgets.NewQScrollArea(nil)
		scrollArea.SetMaximumWidth(625)
		scrollArea.SetMinimumWidth(625)
		scrollArea.SetWidgetResizable(true)

		vWidget := widgets.NewQWidget(nil, core.Qt__Widget)
		vWidget.SetMaximumWidth(600)
		vWidget.SetMinimumWidth(600)
		vLayout := widgets.NewQVBoxLayout()

		for i := 0; i < len(conf.HolidaysSort); i++ {
			date := conf.HolidaysSort[i]
			if (date == "") {
				continue
			}
			substrings := strings.Split(date, "/")
			dayID := substrings[0]
			num, err := strconv.Atoi(substrings[1])
			if (err != nil) {
				fmt.Println("Erreur lors de la lecture du jour férié '" + date + "'")
				fmt.Println(err)
				continue
			}
			monthID := format.TranslateMonthNameToFrench(num)
			new_element := NewListElement(&mSettings.QWidget, dayID + " " + monthID)
			vLayout.AddWidget(new_element, 0, 0)
		}

		vLayout.AddStretch(1)
		vWidget.SetLayout(vLayout)
		vWidget.SetStyleSheet(`
			border: 0;
		`)
		scrollArea.SetWidget(vWidget)
		scrollArea.SetStyleSheet(`
			border: 1px solid lightgray;
		`)
		mSettings.categoryViewerLayout.AddWidget(scrollArea, 0, core.Qt__AlignHCenter)

		v2Widget := widgets.NewQWidget(nil, core.Qt__Widget)
		v2Widget.SetMaximumHeight(350)
		v2Layout := widgets.NewQVBoxLayout()

		dayCalendar := widgets.NewQCalendarWidget(nil)
		dayCalendar.SetGridVisible(true)
		dayCalendar.SetStyleSheet(`
			color: black;
			QCalendarWidget QAbstractItemView:section {
				color: black;
			}
		`)
		dayCalendar.SetMinimumSize2(350, 250)
		v2Layout.AddWidget(dayCalendar, 0, 0)

		addButton := shared_gui.NewPushButtonAnimated(&mSettings.QWidget, "assets/add.png", func(_ bool) {})
		addButton.HoveredStyle = `
			QWidget {
				background-color: #f2f2f2;
				border: 1px solid lightgray;
			}
			background-color: #f2f2f2;
			color: #333;
			border: 1px solid lightgray;
			padding: 0;
		`
		addButton.ConnectClicked(func(_ bool) {
			date := dayCalendar.SelectedDate()
			strDate := strconv.Itoa(date.Day()) + "/" + strconv.Itoa(date.Month())
			if _, ok := conf.Holidays[strDate]; !ok {
				conf.Holidays[strDate] = true
				conf.InsertDateAndSort(conf.HolidaysSort, strDate)
				new_element := NewListElement(&mSettings.QWidget, strconv.Itoa(date.Day()) + " " + format.TranslateMonthNameToFrench(date.Month()))
				vLayout.AddWidget(new_element, 0, 0)
				conf.SaveHolidays()
			}
		})
		addButton.SetDisabled(true)
		addButton.SetMaximumWidth(600)
		v2Layout.AddWidget(addButton, 0, 0)

		dayCalendar.ConnectClicked(func(date *core.QDate) {
			strDate := strconv.Itoa(date.Day()) + "/" + strconv.Itoa(date.Month())
			if _, ok := conf.Holidays[strDate]; ok {
				addButton.SetDisabled(true)
			} else {
				addButton.SetEnabled(true)
			}
		})

		v2Layout.AddStretch(1)
		v2Widget.SetLayout(v2Layout)
		v2Widget.SetStyleSheet(`
			border: 1px solid lightgray;
		`)
		mSettings.categoryViewerLayout.AddWidget(v2Widget, 0, core.Qt__AlignHCenter)
	}},
	{"JOURS DE FERMETURE", func(mSettings *SettingsMenu, conf *config.AppConfig) {
		scrollArea := widgets.NewQScrollArea(nil)
		scrollArea.SetMaximumWidth(625)
		scrollArea.SetMinimumWidth(625)
		scrollArea.SetWidgetResizable(true)

		vWidget := widgets.NewQWidget(nil, core.Qt__Widget)
		vWidget.SetMaximumWidth(600)
		vWidget.SetMinimumWidth(600)
		vLayout := widgets.NewQVBoxLayout()

		for i := 0; i < len(conf.ClosingDaysSort); i++ {
			date := conf.ClosingDaysSort[i]
			if (date == "") {
				continue
			}
			substrings := strings.Split(date, "/")
			dayID := substrings[0]
			num, err := strconv.Atoi(substrings[1])
			if (err != nil) {
				fmt.Println("Erreur lors de la lecture du jour de fermeture '" + date + "'")
				fmt.Println(err)
				continue
			}
			monthID := format.TranslateMonthNameToFrench(num)
			new_element := NewListElement(&mSettings.QWidget, dayID + " " + monthID)
			vLayout.AddWidget(new_element, 0, 0)
		}

		vLayout.AddStretch(1)
		vWidget.SetLayout(vLayout)
		vWidget.SetStyleSheet(`
			border: 0;
		`)
		scrollArea.SetWidget(vWidget)
		scrollArea.SetStyleSheet(`
			border: 1px solid lightgray;
		`)
		mSettings.categoryViewerLayout.AddWidget(scrollArea, 0, core.Qt__AlignHCenter)

		v2Widget := widgets.NewQWidget(nil, core.Qt__Widget)
		v2Widget.SetMaximumHeight(350)
		v2Layout := widgets.NewQVBoxLayout()

		dayCalendar := widgets.NewQCalendarWidget(nil)
		dayCalendar.SetGridVisible(true)
		dayCalendar.SetStyleSheet(`
			color: black;
			QCalendarWidget QAbstractItemView:section {
				color: black;
			}
		`)
		dayCalendar.SetMinimumSize2(350, 250)
		v2Layout.AddWidget(dayCalendar, 0, 0)

		addButton := shared_gui.NewPushButtonAnimated(&mSettings.QWidget, "assets/add.png", func(_ bool) {})
		addButton.HoveredStyle = `
			QWidget {
				background-color: #f2f2f2;
				border: 1px solid lightgray;
			}
			background-color: #f2f2f2;
			color: #333;
			border: 1px solid lightgray;
			padding: 0;
		`
		addButton.ConnectClicked(func(_ bool) {
			date := dayCalendar.SelectedDate()
			strDate := strconv.Itoa(date.Day()) + "/" + strconv.Itoa(date.Month())
			if _, ok := conf.ClosingDays[strDate]; !ok {
				conf.ClosingDays[strDate] = true
				conf.InsertDateAndSort(conf.ClosingDaysSort, strDate)
				new_element := NewListElement(&mSettings.QWidget, strconv.Itoa(date.Day()) + " " + format.TranslateMonthNameToFrench(date.Month()))
				vLayout.AddWidget(new_element, 0, 0)
				conf.SaveClosingDays()
			}
		})
		addButton.SetDisabled(true)
		addButton.SetMaximumWidth(600)
		v2Layout.AddWidget(addButton, 0, 0)

		dayCalendar.ConnectClicked(func(date *core.QDate) {
			strDate := strconv.Itoa(date.Day()) + "/" + strconv.Itoa(date.Month())
			if _, ok := conf.ClosingDays[strDate]; ok {
				addButton.SetDisabled(true)
			} else {
				addButton.SetEnabled(true)
			}
		})

		v2Layout.AddStretch(1)
		v2Widget.SetLayout(v2Layout)
		v2Widget.SetStyleSheet(`
			border: 1px solid lightgray;
		`)
		mSettings.categoryViewerLayout.AddWidget(v2Widget, 0, core.Qt__AlignHCenter)
	}},
	{"FORMATIONS", func(mSettings *SettingsMenu, conf *config.AppConfig) {
		scrollArea := widgets.NewQScrollArea(nil)
		scrollArea.SetMaximumWidth(625)
		scrollArea.SetMinimumWidth(625)
		scrollArea.SetWidgetResizable(true)

		vWidget := widgets.NewQWidget(nil, core.Qt__Widget)
		vWidget.SetMaximumWidth(600)
		vWidget.SetMinimumWidth(600)
		vLayout := widgets.NewQVBoxLayout()

		for key, value := range conf.FormationTypes {
			formationAbr := key
			new_element := NewListElement(&mSettings.QWidget, key + " - " + value)
			new_element.deleteButton.ConnectClicked(func(_ bool) {
				delete(conf.FormationTypes, formationAbr)
				conf.SaveFormationTypes()
				new_element.DeleteLater()
				fUpdateFormationTypes()
			})
			vLayout.AddWidget(new_element, 0, 0)
		}

		vWidget.SetLayout(vLayout)
		vWidget.SetStyleSheet(`
			border: 1px solid lightgray;
		`)
		scrollArea.SetWidget(vWidget)
		mSettings.categoryViewerLayout.AddWidget(scrollArea, 0, core.Qt__AlignHCenter)

		v2Widget := widgets.NewQWidget(nil, core.Qt__Widget)
		v2Layout := widgets.NewQVBoxLayout()

		nameEdit := widgets.NewQLineEdit2("", nil)
		nameEdit.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
		nameEdit.SetPlaceholderText("Nom complet de la formation")
		nameEdit.SetMaximumHeight(50)
		nameEdit.SetMinimumHeight(50)
		nameEdit.SetMaximumWidth(600)
		nameEdit.SetMinimumWidth(300)
		nameEdit.SetStyleSheet(`
			font-family: Arial, sans-serif;
			font-size: 16px;
		`)
		v2Layout.AddWidget(nameEdit, 0, 0)

		abrEdit := widgets.NewQLineEdit2("", nil)
		abrEdit.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
		abrEdit.SetPlaceholderText("Abréviation du nom de la formation")
		abrEdit.SetMaximumHeight(50)
		abrEdit.SetMinimumHeight(50)
		abrEdit.SetMaximumWidth(600)
		abrEdit.SetMinimumWidth(300)
		abrEdit.SetStyleSheet(`
			font-family: Arial, sans-serif;
			font-size: 16px;
		`)
		v2Layout.AddWidget(abrEdit, 0, 0)

		addButton := shared_gui.NewPushButtonAnimated(&mSettings.QWidget, "assets/add.png", func(_ bool) {})
		addButton.HoveredStyle = `
			QWidget {
				background-color: #f2f2f2;
				border: 1px solid lightgray;
			}
			background-color: #f2f2f2;
			color: #333;
			border: 1px solid lightgray;
			padding: 0;
		`
		addButton.ConnectClicked(func(_ bool) {
			formationName := nameEdit.Text()
			formationAbr := abrEdit.Text()
			_, ok := conf.FormationTypes[formationAbr]
			if (!ok && formationName != "" && formationAbr != "") {
				conf.FormationTypes[formationAbr] = formationName
				new_element := NewListElement(&mSettings.QWidget, formationAbr + " - " + formationName)
				new_element.deleteButton.ConnectClicked(func(_ bool) {
					delete(conf.FormationTypes, formationAbr)
					conf.SaveFormationTypes()
					new_element.DeleteLater()
					fUpdateFormationTypes()
				})
				vLayout.AddWidget(new_element, 0, 0)
				conf.SaveFormationTypes()
				fUpdateFormationTypes()
			}
		})
		addButton.SetMaximumWidth(600)
		v2Layout.AddWidget(addButton, 0, 0)

		v2Layout.AddStretch(1)
		v2Widget.SetLayout(v2Layout)
		v2Widget.SetStyleSheet(`
			border: 1px solid lightgray;
		`)
		mSettings.categoryViewerLayout.AddWidget(v2Widget, 0, core.Qt__AlignHCenter)
	}},
	{"DOSSIER PDF", func(mSettings *SettingsMenu, conf *config.AppConfig) {
		vWidget := widgets.NewQWidget(nil, core.Qt__Widget)
		vLayout := widgets.NewQVBoxLayout()

		// Create a new QFileDialog instance
		fileDialog := widgets.NewQFileDialog(nil, core.Qt__Dialog)

		// Set the FileMode property to DirectoryOnly
		fileDialog.SetFileMode(widgets.QFileDialog__DirectoryOnly)

		hWidget := widgets.NewQWidget(nil, core.Qt__Widget)
		hWidget.SetMaximumSize2(800, 80)
		hLayout := widgets.NewQHBoxLayout()

		pathLabel := widgets.NewQLabel(nil, core.Qt__Widget)
		pathLabel.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
		pathLabel.SetAlignment(core.Qt__AlignVCenter)
		pathLabel.SetMinimumWidth(500)
		pathLabel.SetMaximumHeight(25)
		pathLabel.SetStyleSheet(`
			border: 0;
			font-family: Arial, sans-serif;
			font-size: 16px;
		`)

		outputpath := ""

		if conf.OutputDir != "." {
			outputpath = conf.OutputDir
		} else {
			// Get the executable path
			_, filename, _, ok := runtime.Caller(0)
			// exePath, err := os.Executable()
			if !ok {
				fmt.Println("Failed to get executable path:")
				return
			}

			endstr := len(filename) - 29
			outputpath = filename[:endstr]

			fmt.Println("Source path:", outputpath)
		}

		fileDialog.SetWindowFilePath(outputpath)
		fileDialog.SetDirectory(outputpath)
		pathLabel.SetText("Path: " + outputpath)
		hLayout.AddWidget(pathLabel, 0, 0)

		fileDialog.ConnectFileSelected(func(selectedDirectory string) {
			pathLabel.SetText("Path: " + selectedDirectory)
			conf.SaveOutputDirectory(selectedDirectory)
		})

		// fileExplorerButton := widgets.NewQPushButton(nil)
		fileExplorerButton := shared_gui.NewPushButtonAnimated(hWidget, "assets/search.png", func(_ bool) {
			fileDialog.Exec()
		})
		hLayout.AddWidget(fileExplorerButton, 0, 0)

		hWidget.SetLayout(hLayout)
		vLayout.AddWidget(hWidget, 0, core.Qt__AlignCenter)

		vWidget.SetLayout(vLayout)
		vWidget.SetStyleSheet(`
			border: 1px solid lightgray;
		`)
		mSettings.categoryViewerLayout.AddWidget(vWidget, 0, 0)
	}},
}

/*
	CategoryButton
*/

type CategoryButton struct {
	*widgets.QPushButton
}

func NewCategoryButton(parent *widgets.QWidget, name string) *CategoryButton {
	cButton := &CategoryButton{widgets.NewQPushButton(nil)}
	cButton.SetStyleSheet(`
		background-color: #f2f2f2;
		color: #bbb;
		border: 1px solid #f2f2f2;
		padding: 0;
		font-family: Arial, sans-serif;
		font-size: 20px;
	`)
	cButton.ConnectEnterEvent(func(event *core.QEvent) {
		cursor := gui.NewQCursor2(core.Qt__PointingHandCursor)
		parent.SetCursor(cursor)
		cButton.SetStyleSheet(`
			background-color: #f2f2f2;
			color: #333;
			border: 1px solid #f2f2f2;
			padding: 0;
			font-family: Arial, sans-serif;
			font-size: 20px;
		`)
		event.Ignore()
	})
	cButton.ConnectLeaveEvent(func(event *core.QEvent) {
		cursor := gui.NewQCursor2(core.Qt__ArrowCursor)
		parent.SetCursor(cursor)
		cButton.SetStyleSheet(`
			background-color: #f2f2f2;
			color: #bbb;
			border: 1px solid #f2f2f2;;
			padding: 0;
			font-family: Arial, sans-serif;
			font-size: 20px;
		`)
		event.Ignore()
	})
	cButton.SetText(name)
	return (cButton)
}

/*
	SettingsMenu
*/

type SettingsMenu struct {
	*widgets.QDialog
	conf *config.AppConfig
	hLayout *widgets.QHBoxLayout
	categoriesPanel *widgets.QWidget
	categoryViewerPanel *widgets.QWidget
	categoryViewerLayout *widgets.QHBoxLayout
	header *widgets.QLabel
}

func (mSettings *SettingsMenu) createCategoriesMenu() {
	vWidget := widgets.NewQWidget(nil, core.Qt__Widget)
	vWidget.SetMaximumWidth(300)
	vWidget.SetMinimumWidth(300)
	vLayout := widgets.NewQVBoxLayout()

	categoriesWidget := widgets.NewQWidget(vWidget, core.Qt__Widget)
	nb := len(categories)
	categoriesWidget.SetMinimumHeight(30 * nb + 50 * (nb - 1))
	categoriesWidget.SetMinimumWidth(300)
	categoriesWidgetLayout := widgets.NewQVBoxLayout()
	categoriesWidgetLayout.SetSpacing(50)

	// Fill categories list
	for i := 0; i < len(categories); i++ {
		catData := categories[i]
		curButton := NewCategoryButton(&mSettings.QWidget, catData.name)
		curButton.ConnectClicked(func(_ bool) {
			// Clear current viewer
			cLayout := mSettings.categoryViewerPanel.Layout()
			for cLayout.Count() > 0 {
				child := cLayout.TakeAt(0).Widget()
				child.DeleteLater()
			}
			mSettings.header.SetText(catData.name)
			catData.viewer(mSettings, mSettings.conf)
		})
		categoriesWidgetLayout.AddWidget(curButton, 0, 0)
	}

	categoriesWidgetLayout.AddStretch(1)
	categoriesWidget.SetLayout(categoriesWidgetLayout)
	// vLayout.AddWidget(categoriesWidget, 0, 0)
	vWidget.SetLayout(vLayout)
	mSettings.categoriesPanel = vWidget
	vWidget.SetStyleSheet(`
		border: 0;
	`)
	mSettings.hLayout.AddWidget(vWidget, 0, 0)
	categoriesWidget.Move(core.NewQPoint2(0, 425 - categoriesWidget.Height() / 2))
}

func (mSettings *SettingsMenu) addHoliday(date string) {
	// pLayout := mSettings.categoryViewerLayout.Layout()
	fmt.Println("Add Holiday", date)
}

func (mSettings *SettingsMenu) createCategoryViewer() {
	vWidget := widgets.NewQWidget(nil, core.Qt__Widget)
	vLayout := widgets.NewQVBoxLayout()

	mSettings.header.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	mSettings.header.SetAlignment(core.Qt__AlignCenter)
	mSettings.header.SetMaximumHeight(50)
	mSettings.header.SetStyleSheet(`
		font-family: Arial, sans-serif;
		font-size: 20px;
		border: 0;
	`)
	vLayout.AddWidget(mSettings.header, 0, 0)

	categoryElements := widgets.NewQWidget(vWidget, core.Qt__Widget)
	categoryElements.SetMinimumWidth(1200)
	categoryElementsLayout := widgets.NewQHBoxLayout()
	categoryElements.SetLayout(categoryElementsLayout)
	categoryElements.SetStyleSheet(`
		border: 0;
	`)
	vLayout.AddWidget(categoryElements, 0, 0)

	vWidget.SetLayout(vLayout)
	vWidget.SetStyleSheet(`
		border: 1px solid lightgray;
	`)
	mSettings.categoryViewerPanel = categoryElements
	mSettings.categoryViewerLayout = categoryElementsLayout
	mSettings.hLayout.AddWidget(vWidget, 0, 0)
}

var fUpdateFormationTypes func()

func NewSettingsMenu(conf *config.AppConfig, fUFT func()) *SettingsMenu {
	fUpdateFormationTypes = fUFT
	mSettings := &SettingsMenu{widgets.NewQDialog(nil, core.Qt__Widget), conf, widgets.NewQHBoxLayout(), widgets.NewQWidget(nil, core.Qt__Widget), widgets.NewQWidget(nil, core.Qt__Widget), widgets.NewQHBoxLayout(), widgets.NewQLabel(nil, core.Qt__Widget)}
	mSettings.SetWindowFlags(core.Qt__Dialog | core.Qt__WindowTitleHint | core.Qt__CustomizeWindowHint | core.Qt__WindowCloseButtonHint)
	mSettings.SetFixedSize2(1600, 900)
	mSettings.SetWindowTitle("Options")
	mSettings.createCategoriesMenu()
	mSettings.createCategoryViewer()
	mSettings.hLayout.AddStretch(1)
	mSettings.SetLayout(mSettings.hLayout)
	return (mSettings)
}
