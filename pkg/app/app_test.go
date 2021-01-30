package app_test

import (
	"testing"

	"github.com/seandrucker/golang-starter-cmd/pkg/app"
)

func TestApp(t *testing.T) {
	app := &app.App{
		Verbose: true,
	}
	app.Run()
}
