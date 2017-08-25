package glib_test

import (
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

	called := false
	_, err := glib.IdleAdd(func() {
		called = true
	})
	require.Nil(err)
	require.Equal(1, glib.RegisteredClosures())
	require.False(called)

	require.True(gtk.EventsPending())

	gtk.MainIteration()

	require.False(gtk.EventsPending())
	require.True(called)
	require.Equal(0, glib.RegisteredClosures())
}
