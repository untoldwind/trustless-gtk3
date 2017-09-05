package cairo

// MimeType is a representation of Cairo's CAIRO_MIME_TYPE_*
// preprocessor constants.
type MimeType string

const (
	MimeTypeJP2      MimeType = "image/jp2"
	MimeTypeJpeg     MimeType = "image/jpeg"
	MimeTypePng      MimeType = "image/png"
	MimeTypeUri      MimeType = "image/x-uri"
	MimeTypeUniqueId MimeType = "application/x-cairo.uuid"
)
