package glib_test

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/untoldwind/amintk/fixtures"
	"github.com/untoldwind/amintk/glib"
)

func TestObject(t *testing.T) {
	createViewerFile(t)

	runtime.GC()
}

func createViewerFile(t *testing.T) {
	require := require.New(t)

	viewerFile := glib.NewObject(fixtures.ViewerFileGetType())

	require.Equal(glib.TYPE_STRING, viewerFile.GetPropertyType("filename"))
	require.Equal(glib.TYPE_UINT, viewerFile.GetPropertyType("zoom_level"))

	require.Equal("", viewerFile.GetProperty("filename"))
	require.Equal(uint(0), viewerFile.GetProperty("zoom_level"))

	viewerFile.SetProperty("filename", "blabla")
	viewerFile.SetProperty("zoom_level", uint(8))

	require.Equal("blabla", viewerFile.GetProperty("filename"))
	require.Equal(uint(8), viewerFile.GetProperty("zoom_level"))
}
