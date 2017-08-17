#ifndef __FIXTURES_VIEWER_FILE_GO_H__
#define __FIXTURES_VIEWER_FILE_GO_H__

#include <stdio.h>
#include <glib-object.h>

G_BEGIN_DECLS

#define VIEWER_TYPE_FILE viewer_file_get_type ()
G_DECLARE_FINAL_TYPE (ViewerFile, viewer_file, VIEWER, FILE, GObject)

ViewerFile *viewer_file_new (void);

G_END_DECLS

struct _ViewerFile
{
  GObject parent_instance;

  gchar *filename;
  guint zoom_level;
};

G_DEFINE_TYPE (ViewerFile, viewer_file, G_TYPE_OBJECT)

enum
{
  PROP_FILENAME = 1,
  PROP_ZOOM_LEVEL,
  N_PROPERTIES
};

static GParamSpec *obj_properties[N_PROPERTIES] = { NULL, };

enum
{
  CHANGED = 1,
  N_SIGNALS
};

static guint file_signals[N_SIGNALS] = { 0, };

static void
viewer_file_set_property (GObject      *object,
                          guint         property_id,
                          const GValue *value,
                          GParamSpec   *pspec)
{
  ViewerFile *self = VIEWER_FILE (object);

  switch (property_id)
    {
    case PROP_FILENAME:
      g_free (self->filename);
      self->filename = g_value_dup_string (value);
      g_print ("filename: %s\n", self->filename);
      break;

    case PROP_ZOOM_LEVEL:
      self->zoom_level = g_value_get_uint (value);
      g_print ("zoom level: %u\n", self->zoom_level);
      break;

    default:
      /* We don't have any other property... */
      G_OBJECT_WARN_INVALID_PROPERTY_ID (object, property_id, pspec);
      break;
    }
}

static void
viewer_file_get_property (GObject    *object,
                          guint       property_id,
                          GValue     *value,
                          GParamSpec *pspec)
{
  ViewerFile *self = VIEWER_FILE (object);

  switch (property_id)
    {
    case PROP_FILENAME:
      g_value_set_string (value, self->filename);
      break;

    case PROP_ZOOM_LEVEL:
      g_value_set_uint (value, self->zoom_level);
      break;

    default:
      /* We don't have any other property... */
      G_OBJECT_WARN_INVALID_PROPERTY_ID (object, property_id, pspec);
      break;
    }
}

static void
viewer_file_constructed (GObject *obj)
{
  fprintf(stderr, "ViewerFile constructed\n");
  G_OBJECT_CLASS (viewer_file_parent_class)->constructed (obj);
}

static void
viewer_file_dispose (GObject *gobject)
{
  fprintf(stderr, "ViewerFile disposed\n");
  G_OBJECT_CLASS (viewer_file_parent_class)->dispose (gobject);
}

static void
viewer_file_finalize (GObject *gobject)
{
  fprintf(stderr, "ViewerFile finalized\n");
  G_OBJECT_CLASS (viewer_file_parent_class)->finalize (gobject);
}

static void
viewer_file_class_init (ViewerFileClass *klass)
{
  GObjectClass *object_class = G_OBJECT_CLASS (klass);

  object_class->constructed = viewer_file_constructed;
  object_class->dispose = viewer_file_dispose;
  object_class->finalize = viewer_file_finalize;

  object_class->set_property = viewer_file_set_property;
  object_class->get_property = viewer_file_get_property;

  obj_properties[PROP_FILENAME] =
    g_param_spec_string ("filename",
                         "Filename",
                         "Name of the file to load and display from.",
                         NULL  /* default value */,
                         G_PARAM_READWRITE);

  obj_properties[PROP_ZOOM_LEVEL] =
    g_param_spec_uint ("zoom-level",
                       "Zoom level",
                       "Zoom level to view the file at.",
                       0  /* minimum value */,
                       10 /* maximum value */,
                       2  /* default value */,
                       G_PARAM_READWRITE);

  g_object_class_install_properties (object_class,
                                     N_PROPERTIES,
                                     obj_properties);

                                     file_signals[CHANGED] =
  g_signal_newv ("changed",
                G_TYPE_FROM_CLASS (object_class),
                G_SIGNAL_RUN_LAST | G_SIGNAL_NO_RECURSE | G_SIGNAL_NO_HOOKS,
                NULL /* closure */,
                NULL /* accumulator */,
                NULL /* accumulator data */,
                NULL /* C marshaller */,
                G_TYPE_NONE /* return_type */,
                0     /* n_params */,
                NULL  /* param_types */);
}

static void
viewer_file_init (ViewerFile *self)
{
}

static void
_viewer_file_emit_changed(gpointer self)
{
  g_signal_emit (self, file_signals[CHANGED], 0 /* details */);
}

#endif
