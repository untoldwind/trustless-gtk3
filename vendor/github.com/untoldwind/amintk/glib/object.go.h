#ifndef __GLIB_OBJECT_GO_H__
#define __GLIB_OBJECT_GO_H__

#include <stdlib.h>
#include <glib.h>
#include <glib-object.h>

static GObjectClass *
_g_object_get_class (GObject *object)
{
  return (G_OBJECT_GET_CLASS(object));
}

static GObject *
_g_object_new (GType gType)
{
  return g_object_new(gType, NULL);
}

#endif
