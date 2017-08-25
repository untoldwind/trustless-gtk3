# Amin(o)TK

A minimalistic GNOME binding for GO

This project is inspired by the much more complete [gotk3 project](https://github.com/gotk3/gotk3)
and actually has a lot of code in common. There are some major differences though:

* Error handling (probably the main reason why I started this): In gotk3 every GTK error is mapped.
  In most cases this is not feasible at all. Most simple example: In gotk3 `gtk.LabelNew` returns
  an errors if the underlying `g_gtk_label_new` returns `NULL`. There is no reasonable way to handle
  these kind of errors (except logging maybe, but GTK is doing that by itself already). So instead of
  returning and error, we just return a go `nil` which you may or may not handle (all functions should be `nil` safe).
  This removes a ton of boilerplate code in your application.
* Linking: gotk3 made the honorable effort create GTK marshallers for all types.
  Even though this comes in handy, the drawback is that the `gtk` package has an `init()` function
  registering all marshallers, which in turn creates a dependency to literally everything.
  Therefore all the widget bindings are linked even if you do not need them, which creates
  rather large executables.
* Signals: To avoid the necessity of a global marshaller list, all signals should be bound
  via the `On<CamelCase>` rather the `Connect("<snake-case>")` functions. This removed a
  lot of magic strings and makes signal binding a bit more type safe. The drawback is
  that signal callbacks/closures do not support userdata any more - but with go one
  usually do not need those anyway.
   
