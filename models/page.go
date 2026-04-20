package models

// HomePage is view data for the decoder UI template.
type HomePage struct {
	Title     string
	Tagline   string
	PagesBase string // URL prefix for tab routes, e.g. "/pages"
	ActiveTab string // "encode" or "decode"
}
