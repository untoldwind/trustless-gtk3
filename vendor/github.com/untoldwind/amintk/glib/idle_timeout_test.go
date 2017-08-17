package glib_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/untoldwind/amintk/glib"
	"github.com/untoldwind/amintk/gtk"
)

func TestIdleTimeout(t *testing.T) {
	require := require.New(t)

	runtime.LockOSThread()

	require.False(gtk.EventsPending())

	_, err := glib.IdleAdd(func() {
		fmt.Println("bla")
	})
	require.Nil(err)
	require.Equal(1, glib.RegisteredClosures())

	require.True(gtk.EventsPending())

	gtk.MainIteration()

	require.False(gtk.EventsPending())
	require.Equal(0, glib.RegisteredClosures())
}
