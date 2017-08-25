#ifndef __GDK_PIXBUF_GO_H__
#define __GDK_PIXBUF_GO_H__

#include <stdlib.h>
#include <gdk/gdk.h>

static gboolean
_gdk_pixbuf_save_png(GdkPixbuf *pixbuf,
const char *filename, GError ** err, const char *compression)
{
	return gdk_pixbuf_save(pixbuf, filename, "png", err, "compression", compression, NULL);
}

static gboolean
_gdk_pixbuf_save_jpeg(GdkPixbuf *pixbuf,
const char *filename, GError ** err, const char *quality)
{
	return gdk_pixbuf_save(pixbuf, filename, "jpeg", err, "quality", quality, NULL);
}

#endif
