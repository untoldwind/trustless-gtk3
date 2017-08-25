package glib_test

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/untoldwind/amintk/fixtures"
	"github.com/untoldwind/amintk/glib"
)

func TestObject(t *testing.T) {
	require := require.New(t)

	createViewerFile(t)

	var nilObject *glib.Object = nil

	require.True(nilObject.Native() == nil)

	runtime.GC()
}

func createViewerFile(t *testing.T) {
	require := require.New(t)

	viewerFile := glib.NewObject(fixtures.ViewerFileGetType())

	require.Equal(glib.TYPE_STRING, viewerFile.GetPropertyType("filename"))
	require.Equal(glib.TYPE_UINT, viewerFile.GetPropertyType("zoom_level"))

	require.Equal("", viewerFile.GetProperty("filename").GetString())
	zoom, ok := viewerFile.GetProperty("zoom_level").GetUInt()
	require.Equal(uint(0), zoom)
	require.True(ok)

	viewerFile.SetProperty("filename", "blabla")
	viewerFile.SetProperty("zoom_level", uint(8))

	require.Equal("blabla", viewerFile.GetProperty("filename").GetString())
	zoom, ok = viewerFile.GetProperty("zoom_level").GetUInt()
	require.Equal(uint(8), zoom)
	require.True(ok)
}
