package palette_gui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

const (
	PNone = iota
	PClear
	PWork
	PClosing
	PHoliday
	PVisio
	PRevision
	PExam
	PWeekEnd
	PIntegration
)

const (
	CNone = "#FFFFFF"
	CClear = "#F2F2F2"
	CWork = "#bd95ba"
	CClosing = "#c99b95"
	CHoliday = "#9ea7a8"
	CVisio = "#a9c995"
	CRevision = "#7a7ede"
	CExam = "#6b4770"
	CWeekEnd = "#cae1e4"
	CIntegration = "#f2c1ef"
)

var SelectedPalette int = PNone

func CreatePaletteWidget(name string, tPalette int, hexColor string) *widgets.QWidget {
	hLayout := widgets.NewQHBoxLayout2(nil)
	hWidget := widgets.NewQWidget(nil, core.Qt__Widget)

	headerPalette := widgets.NewQLabel(nil, core.Qt__Widget)
	headerPalette.SetAlignment(core.Qt__AlignHCenter)
	headerPalette.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	headerPalette.SetText(name)
	headerPalette.SetMaximumHeight(50)
	headerPalette.SetStyleSheet(`
		font-family: Arial, sans-serif;
		font-size: 16px;
		border: 1;
	`)
	hLayout.AddWidget(headerPalette, 0, 0)

	colorPicker := widgets.NewQPushButton(nil)
	colorPicker.ConnectClicked(func(_ bool) {
		SelectedPalette = tPalette
	})
	colorPicker.SetStyleSheet(`
		background-color: ` + hexColor + `;
	`)
	hLayout.AddWidget(colorPicker, 0, 0)

	hWidget.SetLayout(hLayout)
	return (hWidget)
}

func CreatePalettesWidget() *widgets.QWidget {
	vLayout := widgets.NewQVBoxLayout2(nil)
	widgetVLayout := widgets.NewQWidget(nil, core.Qt__Widget)

	headerPalette := widgets.NewQLabel(nil, core.Qt__Widget)
	headerPalette.SetAlignment(core.Qt__AlignCenter)
	headerPalette.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	headerPalette.SetText("Palette de Couleur")
	headerPalette.SetMaximumHeight(50)
	headerPalette.SetStyleSheet(`
		font-family: Arial, sans-serif;
		font-size: 16px;
		border: 1;
	`)
	vLayout.AddWidget(headerPalette, 0, 0)

	vLayout.AddWidget(CreatePaletteWidget("Gomme", PClear, CClear), 0, 0)
	vLayout.AddWidget(CreatePaletteWidget("Formation", PWork, CWork), 0, 0)
	vLayout.AddWidget(CreatePaletteWidget("Férié", PHoliday, CHoliday), 0, 0)
	vLayout.AddWidget(CreatePaletteWidget("Fermeture", PClosing, CClosing), 0, 0)
	vLayout.AddWidget(CreatePaletteWidget("Visio", PVisio, CVisio), 0, 0)
	vLayout.AddWidget(CreatePaletteWidget("Révision", PRevision, CRevision), 0, 0)
	vLayout.AddWidget(CreatePaletteWidget("Epreuve", PExam, CExam), 0, 0)
	vLayout.AddWidget(CreatePaletteWidget("Intégration", PIntegration, CIntegration), 0, 0)

	vLayout.AddStretch(1)
	widgetVLayout.SetLayout(vLayout)
	widgetVLayout.SetStyleSheet(`
		background-color: #f2f2f2;
		color: #333;
		border: 1px solid #bbb;
		border-radius: 5px;
		padding: 0;
	`)
	widgetVLayout.SetMinimumSize2(150, 200)
	return (widgetVLayout)
}
