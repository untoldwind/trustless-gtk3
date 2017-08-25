#ifndef __GTK_TREE_VIEW_COLUMN_GO_H__
#define __GTK_TREE_VIEW_COLUMN_GO_H__

#include <stdlib.h>
#include <gtk/gtk.h>

static GtkTreeViewColumn *
_gtk_tree_view_column_new_with_attributes_one(const gchar *title,
    GtkCellRenderer *renderer, const gchar *attribute, gint column)
{
	GtkTreeViewColumn	*tvc;

	tvc = gtk_tree_view_column_new_with_attributes(title, renderer,
	    attribute, column, NULL);
	return (tvc);
}

#endif
