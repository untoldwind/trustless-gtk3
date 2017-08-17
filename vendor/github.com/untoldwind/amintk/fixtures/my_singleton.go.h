#ifndef __FIXTURES_MY_SINGLETON_GO_H__
#define __FIXTURES_MY_SINGLETON_GO_H__

#include <stdlib.h>
#include <glib.h>
#include <glib-object.h>

typedef GObject MySingletonObject;
typedef GObjectClass MySingletonObjectClass;

GType my_singleton_object_get_type (void);

G_DEFINE_TYPE (MySingletonObject, my_singleton_object, G_TYPE_OBJECT)

static MySingletonObject *singleton;

static void
my_singleton_object_init (MySingletonObject *obj)
{
}

static GObject *
my_singleton_object_constructor (GType                  type,
                                 guint                  n_construct_properties,
                                 GObjectConstructParam *construct_params)
{
  GObject *object;

  if (singleton)
    return g_object_ref (singleton);

  object = G_OBJECT_CLASS (my_singleton_object_parent_class)->
    constructor (type, n_construct_properties, construct_params);
  singleton = (MySingletonObject *)object;

  return object;
}

static void
my_singleton_object_finalize (MySingletonObject *obj)
{
  singleton = NULL;

  G_OBJECT_CLASS (my_singleton_object_parent_class)->finalize (obj);
}

static void
my_singleton_object_class_init (MySingletonObjectClass *klass)
{
  GObjectClass *object_class = G_OBJECT_CLASS (klass);

  object_class->constructor = my_singleton_object_constructor;
  object_class->finalize = my_singleton_object_finalize;
}

#endif
