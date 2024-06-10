package handlers

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/fgtago/fgweb/appsmodel"
	"github.com/fgtago/fgweb/main/apps"
)

func PageLoginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	app := apps.GetApplication()
	device := ctx.Value(appsmodel.DeviceKeyName).(appsmodel.Device)

	// TODO: implmentasikan tpl
	tpl, exists, err := app.Webservice.TplMgr.GetPage("login", device.Type)
	if err != nil {
		// error 500
		fmt.Println(err.Error())
	}

	if !exists {
		// error 404
		fmt.Println("404")
	}

	// render page
	buff := new(bytes.Buffer)
	err = tpl.Execute(buff, nil)
	if err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		return
	}

	// send bufer to browser
	_, err = buff.WriteTo(w)
	if err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		return
	}
}