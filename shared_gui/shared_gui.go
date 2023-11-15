package shared_gui

import (
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
)

/*
	PushButtonAnimated
*/

type PushButtonAnimated struct {
	*widgets.QPushButton
	InitialSize *core.QSize
	TargetSize *core.QSize
	DefaultStyle string
	HoveredStyle string
}

func NewPushButtonAnimated(parent *widgets.QWidget, icon_path string, f func(bool)) *PushButtonAnimated {
	dStyle := `
		background-color: #f2f2f2;
		color: #333;
		border: 0;
		padding: 0;
	`
	hStyle := `
		QWidget {
			background-color: #f2f2f2;
			border-left: 1px solid lightgray;
			border-bottom: 1px solid lightgray;
		}
		background-color: #f2f2f2;
		color: #333;
		border: 0;
		padding: 0;
	`
	pButton := &PushButtonAnimated{widgets.NewQPushButton(nil), core.NewQSize2(48, 48), core.NewQSize2(52, 52), dStyle, hStyle}
	pixmap := gui.NewQPixmap3(icon_path, "", core.Qt__AutoColor)
	icon := gui.NewQIcon2(pixmap)

	pButton.SetMinimumSize2(64, 64)
	pButton.SetIconSize(pButton.InitialSize)
	pButton.SetIcon(icon)
	pButton.ConnectClicked(f)
	pButton.SetStyleSheet(pButton.DefaultStyle)

	animation := core.NewQPropertyAnimation2(pButton, core.NewQByteArray2("iconSize", len("iconSize")), pButton)
	pButton.ConnectEnterEvent(func(event *core.QEvent) {
		if (pButton.IsEnabled()) {
			cursor := gui.NewQCursor2(core.Qt__PointingHandCursor)
			parent.SetCursor(cursor)
			animation.SetEndValue(core.NewQVariant25(pButton.TargetSize))
			animation.Start(core.QAbstractAnimation__KeepWhenStopped)
			pButton.SetStyleSheet(pButton.HoveredStyle)
		} else {
			cursor := gui.NewQCursor2(core.Qt__ForbiddenCursor)
			parent.SetCursor(cursor)
		}
		event.Ignore()
	})
	pButton.ConnectLeaveEvent(func(event *core.QEvent) {
		cursor := gui.NewQCursor2(core.Qt__ArrowCursor)
		parent.SetCursor(cursor)
		pButton.SetIconSize(pButton.InitialSize)
		pButton.SetStyleSheet(pButton.DefaultStyle)
		animation.Stop()
		event.Ignore()
	})

	return (pButton)
}
