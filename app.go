package main

import (
	"github.com/guark/guark/app"
	"github.com/guark/guark/engine"
	"github.com/guark/guark/log"
	"github.com/username/appname/lib"
)

func main() {

	a := &app.App{
		Log:     log.New("app"),
		Hooks:   lib.Hooks,
		Funcs:   lib.Funcs,
		Embed:   lib.Embeds,
		Plugins: lib.Plugins,
		EngineConfig: map[string]interface{}{
			"useAutomationExtension": false,
			"goog:chromeOptions": map[string]interface{}{
				"excludeSwitches":        []string{"enable-automation"},
				"useAutomationExtension": false,
			},
		},
	}

	if err := a.Use(engine.New(a)); err != nil {
		a.Log.Fatal(err)
	}

	defer a.Quit()

	if err := a.Run(); err != nil {
		a.Log.Fatal(err)
	}
}
