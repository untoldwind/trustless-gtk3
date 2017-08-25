#ifndef __GTK_LIST_STORE_GO_H_
#define __GTK_LIST_STORE_GO_H_

#include <stdlib.h>
#include <gtk/gtk.h>

static GType *
alloc_types(int n) {
	return ((GType *)g_new0(GType, n));
}

static void
set_type(GType *types, int n, GType t)
{
	types[n] = t;
}

static void
_gtk_list_store_set(GtkListStore *list_store, GtkTreeIter *iter, gint column,
	void* value)
{
	gtk_list_store_set(list_store, iter, column, value, -1);
}

static void
_gtk_tree_store_set(GtkTreeStore *store, GtkTreeIter *iter, gint column,
	void* value)
{
	gtk_tree_store_set(store, iter, column, value, -1);
}

#endif
