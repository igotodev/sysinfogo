package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type ipAddress struct {
	IP      string `json:"ip"`
	Country string `json:"country"`
	CC      string `json:"cc"`
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getIP() ipAddress {
	resp, err := http.Get("https://api.myip.com/")
	checkErr(err)
	bodyJson, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	defer resp.Body.Close()
	var jsonAnsw ipAddress
	err = json.Unmarshal(bodyJson, &jsonAnsw)
	checkErr(err)
	//data := jsonAnsw.IP + " " + jsonAnsw.Country // + " " + jsonAnsw.CC
	return jsonAnsw
}

func main() {
	a := app.New()
	w := a.NewWindow("SysInfoGo")

	host, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
		return
	}
	hDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
		return
	}

	nameUserHomeDir := widget.NewLabel("Home dir")
	nameUserHost := widget.NewLabel("Hostname")
	numCPU := widget.NewLabel("Threads CPU")
	os := widget.NewLabel("OS")
	arch := widget.NewLabel("Architecture")
	ip := widget.NewLabel("IP Address")
	country := widget.NewLabel("Country")
	cc := widget.NewLabel("Country code")

	trueIP := getIP()

	w.SetContent(container.NewVBox(
		os,
		arch,
		numCPU,
		nameUserHomeDir,
		nameUserHost,
		ip,
		country,
		cc,
		container.NewHBox(
			widget.NewButton("Get System Info", func() {
				os.SetText("OS: " + runtime.GOOS)
				arch.SetText("Architecture: " + runtime.GOARCH)
				numCPU.SetText("Threads CPU: " + strconv.Itoa(runtime.NumCPU()))
				nameUserHomeDir.SetText("Home dir: " + hDir)
				nameUserHost.SetText("Hostname: " + host)
				ip.SetText("IP Address: " + trueIP.IP)
				country.SetText("Country: " + trueIP.Country)
				cc.SetText("Country code: " + trueIP.CC)
			}),
			widget.NewButton("Quit", func() {
				a.Quit()
			}),
		)))
	w.CenterOnScreen()
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(220, 380))
	w.ShowAndRun()
}
