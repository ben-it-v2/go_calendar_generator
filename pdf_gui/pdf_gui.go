package pdf_gui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/gui"
	"calendar/palette_gui"
	"calendar/format"
	"calendar/shared_gui"
	"calendar/config"
)

/*
	PDFMenu
*/

type PDFMenu struct {
	*widgets.QDialog
}

func createLegendWidget() *widgets.QWidget {
	vLayout := widgets.NewQVBoxLayout2(nil)
	widgetVLayout := widgets.NewQWidget(nil, core.Qt__Widget)

	headerPalette := widgets.NewQLabel(nil, core.Qt__Widget)
	headerPalette.SetAlignment(core.Qt__AlignCenter)
	headerPalette.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	headerPalette.SetText("Légende")
	headerPalette.SetMaximumHeight(50)
	headerPalette.SetStyleSheet(`
		font-family: Arial, sans-serif;
		font-size: 16px;
		border: 1;
	`)
	vLayout.AddWidget(headerPalette, 0, 0)

	vLayout.AddWidget(palette_gui.CreatePaletteWidget("Formation", palette_gui.PWork, palette_gui.CWork), 0, 0)
	vLayout.AddWidget(palette_gui.CreatePaletteWidget("Férié", palette_gui.PHoliday, palette_gui.CHoliday), 0, 0)
	vLayout.AddWidget(palette_gui.CreatePaletteWidget("Fermeture", palette_gui.PClosing, palette_gui.CClosing), 0, 0)
	vLayout.AddWidget(palette_gui.CreatePaletteWidget("Visio", palette_gui.PVisio, palette_gui.CVisio), 0, 0)
	vLayout.AddWidget(palette_gui.CreatePaletteWidget("Révision", palette_gui.PRevision, palette_gui.CRevision), 0, 0)
	vLayout.AddWidget(palette_gui.CreatePaletteWidget("Epreuve", palette_gui.PExam, palette_gui.CExam), 0, 0)
	vLayout.AddWidget(palette_gui.CreatePaletteWidget("Intégration", palette_gui.PIntegration, palette_gui.CIntegration), 0, 0)

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

func createDataLabel(key string, value string) *widgets.QWidget {
	hLayout := widgets.NewQHBoxLayout()
	hWidget := widgets.NewQWidget(nil, core.Qt__Widget)

	keyLabel := widgets.NewQLabel(nil, core.Qt__Widget)
	keyLabel.SetAlignment(core.Qt__AlignVCenter)
	keyLabel.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	keyLabel.SetText(key + " :")
	keyLabel.SetStyleSheet(`
		font-family: Arial, sans-serif;
		font-size: 16px;
	`)
	hLayout.AddWidget(keyLabel, 0, 0)

	valueLabel := widgets.NewQLabel(nil, core.Qt__Widget)
	valueLabel.SetAlignment(core.Qt__AlignVCenter)
	valueLabel.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	valueLabel.SetText(value)
	valueLabel.SetStyleSheet(`
		font-family: Arial, sans-serif;
		font-size: 16px;
	`)
	hLayout.AddWidget(valueLabel, 0, 0)

	hLayout.AddStretch(1)
	hWidget.SetLayout(hLayout)
	hWidget.SetMaximumHeight(40)
	hWidget.SetStyleSheet(`
		border: 0;
	`)
	return (hWidget)
}

var conf *config.AppConfig

func NewPDFMenu(nameValue string, firstnameValue string, formationType string, years string, workDay int, beginDate *core.QDate, containerWidget *widgets.QWidget, config *config.AppConfig) *PDFMenu {
	conf = config
	pdfTitle := nameValue + "_" + firstnameValue + "_calendrier.pdf"
	mPdf := &PDFMenu{widgets.NewQDialog(nil, core.Qt__Widget)}
	mPdf.SetWindowFlags(core.Qt__Dialog | core.Qt__WindowTitleHint | core.Qt__CustomizeWindowHint | core.Qt__WindowCloseButtonHint)
	mPdf.SetFixedSize2(1600, 900)
	mPdf.SetWindowTitle(pdfTitle)
	mainLayout := widgets.NewQVBoxLayout()

	headerWidget := widgets.NewQWidget(nil, core.Qt__Widget)
	headerLayout := widgets.NewQHBoxLayout()

	logoLabel := widgets.NewQLabel(nil, core.Qt__Widget)
	logoImage := gui.NewQPixmap3("assets/idmn_logo.png", "", core.Qt__AutoColor)
	logoImageWidth := float64(logoImage.Width()) / 4
	logoImageHeight := float64(logoImage.Height()) / 4
	logoImage = logoImage.Scaled2(int(logoImageWidth), int(logoImageHeight), core.Qt__KeepAspectRatio, core.Qt__SmoothTransformation)
	logoLabel.SetPixmap(logoImage)
	headerLayout.AddWidget(logoLabel, 0, 0)

	titleLayout := widgets.NewQVBoxLayout()
	titleWidget := widgets.NewQWidget(nil, core.Qt__Widget)
	titleLayout.AddStretch(1)

	titleLabel := widgets.NewQLabel(nil, core.Qt__Widget)
	titleLabel.SetText("Planning de la formation " + years)
	titleFont := gui.NewQFont2("Helvetica", 24, 1, false)
	titleLabel.SetFont(titleFont)
	titleLabel.SetStyleSheet(`
		color: #218b9a
	`)
	titleLayout.AddWidget(titleLabel, 0, 0)

	subtitleLabel := widgets.NewQLabel(nil, core.Qt__Widget)
	subtitleLabel.SetText(formationType)
	subtitleFont := gui.NewQFont2("Helvetica", 16, 1, true)
	subtitleLabel.SetFont(subtitleFont)
	subtitleLabel.SetStyleSheet(`
		color: #9b4694
	`)
	titleLayout.AddWidget(subtitleLabel, 0, 0)

	titleWidget.SetLayout(titleLayout)
	titleWidget.SetMinimumWidth(770)
	headerLayout.AddWidget(titleWidget, 0, 0)

	stampCaseWidget := widgets.NewQWidget(nil, core.Qt__Widget)
	stampCaseWidget.SetFixedSize2(300, 150)
	stampCaseWidget.SetStyleSheet(`
		background-color: #f2f2f2;
		color: #333;
		border: 5px solid #ccc;
		border-radius: 5px;
		padding: 0;
	`)
	headerLayout.AddWidget(stampCaseWidget, 0, core.Qt__AlignCenter)

	stampCaseLabel := widgets.NewQLabel(stampCaseWidget, core.Qt__Widget)
	stampCaseLabel.SetAlignment(core.Qt__AlignHCenter)
	stampCaseLabel.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	stampCaseLabel.SetText("IDMN CFA")
	scLabelFont := gui.NewQFont2("Helvetica", 6, 1, true)
	stampCaseLabel.SetFont(scLabelFont)
	stampCaseLabel.SetStyleSheet(`
		background-color: #f2f2f2;
		color: #ccc;
		border: 0;
	`)
	stampCaseLabel.Move2(25, -3)
	stampCaseLabel.SetMinimumSize2(60, 8)

	headerLayout.AddStretch(1)
	headerWidget.SetLayout(headerLayout)
	mainLayout.AddWidget(headerWidget, 0, 0)

	bodyWidget := widgets.NewQWidget(nil, core.Qt__Widget)
	bodyLayout := widgets.NewQHBoxLayout()

	leftBodyWidget := widgets.NewQWidget(nil, core.Qt__Widget)
	leftBodyLayout := widgets.NewQVBoxLayout()

	alternantWidget := widgets.NewQWidget(nil, core.Qt__Widget)
	alternantLayout := widgets.NewQVBoxLayout()

	alternantLayout.AddWidget(createDataLabel("Nom", nameValue), 0, 0)
	alternantLayout.AddWidget(createDataLabel("Prénom", firstnameValue), 0, 0)
	alternantLayout.AddWidget(createDataLabel("Jour de formation", format.FormatDayName2(workDay)), 0, 0)
	alternantLayout.AddWidget(createDataLabel("Date de début", beginDate.ToString("dd-MM-yyyy")), 0, 0)

	alternantLayout.AddStretch(1)
	alternantWidget.SetLayout(alternantLayout)
	alternantWidget.SetStyleSheet(`
		background-color: #f2f2f2;
		color: #333;
		border: 1px solid #bbb;
		border-radius: 5px;
		padding: 0;
	`)

	leftBodyLayout.AddWidget(alternantWidget, 0, 0)
	leftBodyLayout.AddWidget(createLegendWidget(), 0, 0)

	leftBodyLayout.AddStretch(1)
	leftBodyWidget.SetLayout(leftBodyLayout)
	leftBodyWidget.SetMaximumWidth(300)
	bodyLayout.AddWidget(leftBodyWidget, 0, 0)

	// Display current calendar
	// Use pixmap instead of creating twice the calendar
	calendarLabel := widgets.NewQLabel(nil, core.Qt__Widget)
	pixmap := containerWidget.Grab(containerWidget.Rect())
	calendarWidth := float64(containerWidget.Width()) / 1.5
	calendarHeight := float64(containerWidget.Height()) / 1.5
	pixmap = pixmap.Scaled2(int(calendarWidth), int(calendarHeight), core.Qt__KeepAspectRatio, core.Qt__SmoothTransformation)
	calendarLabel.SetPixmap(pixmap)
	bodyLayout.AddWidget(calendarLabel, 0, core.Qt__AlignTop)

	var saveButton *shared_gui.PushButtonAnimated
	saveButton = shared_gui.NewPushButtonAnimated(&mPdf.QWidget, "assets/save_icone.png", func(_ bool) {
		// Setup PDF file
		pdfWriter := gui.NewQPdfWriter(conf.OutputDir + "/" + pdfTitle)
		pdfWriter.SetPageSize2(gui.QPagedPaintDevice__A4)
		pdfWriter.SetPageOrientation(gui.QPageLayout__Landscape)
		pdfWriter.SetResolution(160)

		// Create a QPainter to draw on the PDF
		painter := gui.NewQPainter2(pdfWriter)
		defer painter.DestroyQPainter()

		// Destroy the save button
		saveButton.Close()

		// Draw the Calendar
		pixmap := mPdf.Grab(mPdf.Rect())
		pixmapW := float64(mPdf.Width()) * 1.4
		pixmapH := float64(mPdf.Height()) * 1.4
		pixmap = pixmap.Scaled2(int(pixmapW), int(pixmapH), core.Qt__KeepAspectRatio, core.Qt__SmoothTransformation)
		painter.DrawPixmap8(core.NewQPoint2(0, 0), pixmap)

		// End the QPainter
		painter.End()

		// Close PDF Preview Window
		mPdf.Close()
	})
	bodyLayout.AddWidget(saveButton, 0, core.Qt__AlignCenter)

	bodyWidget.SetLayout(bodyLayout)
	mainLayout.AddWidget(bodyWidget, 0, 0)
	mainLayout.AddStretch(1)
	mPdf.SetLayout(mainLayout)
	return (mPdf)
}
