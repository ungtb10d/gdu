//go:build linux

package tui

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ungtb10d/gdu/v5/internal/testapp"
	"github.com/ungtb10d/gdu/v5/pkg/device"
)

func TestShowDevicesWithError(t *testing.T) {
	app, simScreen := testapp.CreateTestAppWithSimScreen(50, 50)
	defer simScreen.Fini()

	getter := device.LinuxDevicesInfoGetter{MountsPath: "/xyzxyz"}

	ui := CreateUI(app, simScreen, &bytes.Buffer{}, false, false, false, false, false)
	err := ui.ListDevices(getter)

	assert.Contains(t, err.Error(), "no such file")
}
