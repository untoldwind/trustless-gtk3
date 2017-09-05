#ifndef __GDK_PIXBUF_GO_H__
#define __GDK_PIXBUF_GO_H__

#include <stdlib.h>
#include <gdk/gdk.h>
#include <gio/gio.h>

static gboolean
_gdk_pixbuf_save_png(GdkPixbuf *pixbuf,
const char *filename, GError ** err, const char *compression)
{
	return gdk_pixbuf_save(pixbuf, filename, "png", err, "compression", compression, NULL);
}

static gboolean
_gdk_pixbuf_save_png_buffer(GdkPixbuf *pixbuf,
	                          gchar **buffer, gsize *buffer_size,
														GError ** err, const char *compression)
{
	return gdk_pixbuf_save_to_buffer(pixbuf, buffer, buffer_size, "png", err, "compression", compression, NULL);
}

static gboolean
_gdk_pixbuf_save_jpeg(GdkPixbuf *pixbuf,
const char *filename, GError ** err, const char *quality)
{
	return gdk_pixbuf_save(pixbuf, filename, "jpeg", err, "quality", quality, NULL);
}

static gboolean
_gdk_pixbuf_save_jpeg_buffer(GdkPixbuf *pixbuf,
	                          gchar **buffer, gsize *buffer_size,
														GError ** err, const char *quality)
{
	return gdk_pixbuf_save_to_buffer(pixbuf, buffer, buffer_size, "jpeg", err, "quality", quality, NULL);
}

#endif
