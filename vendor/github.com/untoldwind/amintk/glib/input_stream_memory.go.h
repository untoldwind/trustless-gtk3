#ifndef __GLIB_INPUT_STREAM_MEMORY_GO_H__
#define __GLIB_INPUT_STREAM_MEMORY_GO_H__

#include <stdlib.h>
#include <glib.h>
#include <glib-object.h>
#include <gio/gio.h>

static
GInputStream *
_g_memory_input_stream_new_from_data(const void *data, gssize len) {
  return g_memory_input_stream_new_from_data(data, len, free);
}

#endif
