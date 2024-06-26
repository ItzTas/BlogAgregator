package client

type RSS struct {
	XMLNSAtom string  `xml:"xmlns:atom,attr"`
	Version   string  `xml:"version,attr"`
	Channel   Channel `xml:"channel"`
}

type Channel struct {
	Title         string   `xml:"title"`
	Link          string   `xml:"link"`
	Description   string   `xml:"description"`
	Generator     string   `xml:"generator"`
	Language      string   `xml:"language"`
	LastBuildDate string   `xml:"lastBuildDate"`
	AtomLink      AtomLink `xml:"atom:link"`
	Items         []Item   `xml:"item"`
	Images        []Image  `xml:"image"`
}

type Image struct {
	Url    string `xml:"url"`
	Title  string `xml:"title"`
	Link   string `xml:"link"`
	Width  string `xml:"width"`
	Height string `xml:"height"`
}

type AtomLink struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
	Type string `xml:"type,attr"`
}

type Item struct {
	Title       string      `xml:"title"`
	Link        string      `xml:"link"`
	PubDate     string      `xml:"pubDate"`
	Guid        string      `xml:"guid"`
	Description string      `xml:"description"`
	Media       []Media     `xml:"media:content"`
	Thumbnail   []Media     `xml:"media:thumbnail"`
	Enclosure   []Enclosure `xml:"enclosure"`
}

type Media struct {
	URL  string `xml:"url,attr"`
	Type string `xml:"type,attr"`
}

type Enclosure struct {
	URL  string `xml:"url,attr"`
	Type string `xml:"type,attr"`
}
