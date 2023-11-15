# __PDF Calendar Generator__

## __Description__

<br/>

PDF Calendar Generator is a Windows desktop application created for IDMN CFA purpose. This project is based on Golang v1.20.2 using QT binding for Golang to create all User Interfaces.
<br/>
<br/>
Golang : https://go.dev
<br/>
Golang QT : https://github.com/therecipe/qt

<br/>

---

## __How it works__

<br/>

To generate a calendar, you have to fill student's name & firstname, formation type, formation day of the week and the beginning date. Once all data is correctly setup, you can create the calendar by clicking on "Generate" button. After the calendar loading, you will be able to modify the calendar with the colors picker on the left bottom. On adding exam day, visio day, revision day, etc... It automatically updates month formation hours and total hours of the calendar. After editing the calendar, you can create a PDF preview by clicking on the PDF button in right top. Finally, if the PDF preview suits you, you can save the PDF by clicking on the button with a disk icon.
<br/>

/!\ PDF file will be save at the path setting. By default, it sets to the root of the project but you can modify it in the application settings.

/!\ PDF file name is set as: `<NAME>_<FIRSTNAME>_calendar.pdf`
<br/>
`<NAME>`: Student's name
<br/>
`<FIRSTNAME>`: Students's firstname

<br/>

---

## __Installation__
<br/>

Installation is simplified! You have to run the powershell script named `./install.ps1` at root of the project. This script automatically installed and setup Golang v1.20.2 on your computer.

<br/>

---

## __Run Project__
<br/>

To execute the project, you have to run the following command at the root of the project: `go run main.go`
However, we highly recommand to generate an executable file and execute it because the previous command create a build dir and download Golang requirments at each call.

<br/>

---

## __Executable File__
<br/>

/!\ Golang v1.20.2 have to be installed, you can check it by running this command `go version`.

To generate an executable file, you have to run the command `go build` in the root of the projet. This command will install all Golang requirments. If build stage succeded, a file named `calendar.exe` will appear at the root.
