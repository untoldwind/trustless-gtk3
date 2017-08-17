#ifndef __GTK_MAIN_GO_H__
#define __GTK_MAIN_GO_H__

#include <stdlib.h>
#include <gtk/gtk.h>

static inline gchar** make_strings(int count) {
	return (gchar**)malloc(sizeof(gchar*) * count);
}

static inline void destroy_strings(gchar** strings) {
	free(strings);
}

static inline gchar* get_string(gchar** strings, int n) {
	return strings[n];
}

static inline void set_string(gchar** strings, int n, gchar* str) {
	strings[n] = str;
}

#endif
