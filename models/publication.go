package models

import "time"

// Publication Main structure for a publication
type Publication struct {
	Context   []string `json:"@context,omitempty"`
	Metadata  Metadata `json:"metadata"`
	Links     []Link   `json:"links"`
	Spine     []Link   `json:"spine"`
	Resources []Link   `json:"resources,omitempty"` //Replaces the manifest but less redundant
	TOC       []Link   `json:"toc,omitempty"`
	PageList  []Link   `json:"page-list,omitempty"`
	Landmarks []Link   `json:"landmarks,omitempty"`
	LOI       []Link   `json:"loi,omitempty"` //List of illustrations
	LOA       []Link   `json:"loa,omitempty"` //List of audio files
	LOV       []Link   `json:"lov,omitempty"` //List of videos
	LOT       []Link   `json:"lot,omitempty"` //List of tables

	OtherLinks       []Link                  `json:"-"` //Extension point for links that shouldn't show up in the manifest
	OtherCollections []PublicationCollection `json:"-"` //Extension point for collections that shouldn't show up in the manifest
	Internal         []Internal              `json:"-"`
}

// Internal TODO
type Internal struct {
	Name  string
	Value interface{}
}

// Metadata for the default context
type Metadata struct {
	RDFType         string        `json:"@type,omitempty"` //Defaults to schema.org for EBook
	Title           string        `json:"title"`
	Identifier      string        `json:"identifier"`
	Author          []Contributor `json:"author,omitempty"`
	Translator      []Contributor `json:"translator,omitempty"`
	Editor          []Contributor `json:"editor,omitempty"`
	Artist          []Contributor `json:"artist,omitempty"`
	Illustrator     []Contributor `json:"illustrator,omitempty"`
	Letterer        []Contributor `json:"letterer,omitempty"`
	Penciler        []Contributor `json:"penciler,omitempty"`
	Colorist        []Contributor `json:"colorist,omitempty"`
	Inker           []Contributor `json:"inker,omitempty"`
	Narrator        []Contributor `json:"narrator,omitempty"`
	Contributor     []Contributor `json:"contributor,omitempty"`
	Publisher       []Contributor `json:"publisher,omitempty"`
	Imprint         []Contributor `json:"imprint,omitempty"`
	Language        []string      `json:"language,omitempty"`
	Modified        *time.Time    `json:"modified,omitempty"`
	PublicationDate *time.Time    `json:"published,omitempty"`
	Description     string        `json:"description,omitempty"`
	Direction       string        `json:"direction,omitempty"`
	Rendition       *Rendition    `json:"rendition,omitempty"`
	Source          string        `json:"source,omitempty"`
	EpubType        []string      `json:"epub-type,omitempty"`
	Right           string        `json:"right,omitempty"`
	Subject         []Subject     `json:"subject,omitempty"`

	OtherMetadata []Meta `json:"-"` //Extension point for other metadata
}

// Meta is a generic structure for other metadata
type Meta struct {
	property string
	value    string
	children []Meta
}

// Link object used in collections and links
type Link struct {
	Href       string         `json:"href"`
	TypeLink   string         `json:"type"`
	Rel        []string       `json:"rel,omitempty"`
	Height     int            `json:"height,omitempty"`
	Width      int            `json:"width,omitempty"`
	Title      string         `json:"title,omitempty"`
	Properties []string       `json:"properties,omitempty"`
	Duration   *time.Duration `json:"duration,omitempty"`
	Templated  bool           `json:"templated,omitempty"`
}

// Contributor construct used internally for all contributors
type Contributor struct {
	Name       string `json:"name"`
	SortAs     string `json:"sort_as,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Role       string `json:"role,omitempty"`
}

// Rendition object for reflow/FXL
type Rendition struct {
	Flow        string `json:"flow,omitempty"`
	Layout      string `json:"layout,omitempty"`
	Orientation string `json:"orientation,omitempty"`
	Spread      string `json:"spread,omitempty"`
}

// Subject as based on EPUB 3.1 and Webpub
type Subject struct {
	Name   string `json:"name"`
	SortAs string `json:"sort_as,omitempty"`
	Scheme string `json:"scheme,omitempty"`
	Code   string `json:"code,omitempty"`
}

// BelongsTo is a list of collections that a publication belongs to
type BelongsTo struct {
	Series     []Collection `json:"series,omitempty"`
	Collection []Collection `json:"collection,omitempty"`
}

// Collection construct used for collection/serie metadata
type Collection struct {
	Name       string  `json:"name"`
	SortAs     string  `json:"sort_as,omitempty"`
	Identifier string  `json:"identifier,omitempty"`
	Position   float32 `json:"position,omitempty"`
}

// PublicationCollection is used as an extension points for other collections in a Publication
type PublicationCollection struct {
	Role     string
	Metadata []Meta
	Links    []Link
	Children []PublicationCollection
}

// LinkCover return the link for the cover
func (publication *Publication) LinkCover() string {
	// returns the link object for the cover
	return ""
}

func (publication *Publication) linkNavDoc() {
	// returns the link object for the navigation doc (EPUB 3.x)
}
