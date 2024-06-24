package artist

type ArtistService struct {
	APIKey    string
	APISecret string
}

type Artist struct {
	Name       string  `json:"name"`
	Mbid       string  `json:"mbid"`
	URL        string  `json:"url"`
	Image      []Image `json:"image"`
	Streamable string  `json:"streamable"`
	Stats      Stats   `json:"stats"`
	Similar    Similar `json:"similar"`
	Tags       Tags    `json:"tags"`
	Bio        Bio     `json:"bio"`
}

type Image struct {
	Text string `json:"#text"`
	Size string `json:"size"`
}

type Stats struct {
	Listeners     string `json:"listeners"`
	Playcount     string `json:"playcount"`
	UserPlaycount string `json:"userplaycount"`
}

type Similar struct {
	Artist []SimilarArtist `json:"artist"`
}

type SimilarArtist struct {
	Name  string  `json:"name"`
	Mbid  string  `json:"mbid"`
	Match string  `json:"match"`
	URL   string  `json:"url"`
	Image []Image `json:"image"`
}

type Tags struct {
	Tag []Tag `json:"tag"`
}

type Tag struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Bio struct {
	Published string `json:"published"`
	Summary   string `json:"summary"`
	Content   string `json:"content"`
}

type Tracks struct {
	Track []Track `json:"track"`
}

type Track struct {
	Name      string  `json:"name"`
	Mbid      string  `json:"mbid"`
	Playcount string  `json:"playcount"`
	Listeners string  `json:"listeners"`
	URL       string  `json:"url"`
	Image     []Image `json:"image"`
}

type Albums struct {
	Album []Album `json:"album"`
}

type Album struct {
	Name      string  `json:"name"`
	Mbid      string  `json:"mbid"`
	Listeners string  `json:"listeners"`
	URL       string  `json:"url"`
	Image     []Image `json:"image"`
}
