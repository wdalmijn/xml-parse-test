package parser

// Image with all it's props within Discogs
type Image struct {
	Height      int    `json:"height" xml:"height,attr"`
	ResourceURL string `json:"resource_url"`
	Type        string `json:"type" xml:"type,attr"`
	URI         string `json:"uri" xml:"uri,attr"`
	URI150      string `json:"uri150" xml:"uri150,attr"`
	Width       int    `json:"width" xml:"width,attr"`
}

// ArtistReference is a compact reference to an artist
// containing it's ID and name
type ArtistReference struct {
	ID          int64  `json:"id" xml:"id,attr"`
	Name        string `json:"name" xml:",innerxml"`
	ResourceURL string `json:"resource_url"`
	MemberID    int64  `json:"-" xml:"id"`
}

// ArtistEntry is a compact version of an Artist within a release/master
type ArtistEntry struct {
	Anv         string `json:"anv" xml:"anv"`
	ID          int64  `json:"id" xml:"id"`
	Join        string `json:"join" xml:"join"`
	Name        string `json:"name" xml:"name"`
	ResourceURL string `json:"resource_url"`
	Role        string `json:"role" xml:"role"`
	Tracks      string `json:"tracks" xml:"tracks"`
	Index       int    `json:"-" xml:"-"`
}

// TrackEntry is a track part of a release in it's tracklist
type TrackEntry struct {
	TrackID     int64         `json:"-" xml:"-"`
	Duration    string        `json:"duration" xml:"duration"`
	Position    string        `json:"position" xml:"position"`
	Title       string        `json:"title" xml:"title"`
	Type        string        `json:"type_"`
	Artist      []ArtistEntry `json:"artists,omitempty" xml:"artists>artist"`
	ExtraArtist []ArtistEntry `json:"extraartists,omitempty" xml:"extraartists>artist"`
}

// Video is a video correlated to a release
type Video struct {
	VideoID     int64  `json:"-" xml:"-"`
	Description string `json:"description" xml:"description"`
	Duration    int    `json:"duration" xml:"duration,attr"`
	Embed       bool   `json:"embed" xml:"embed,attr"`
	Title       string `json:"title" xml:"title"`
	URI         string `json:"uri" xml:"src,attr"`
}

// LabelReference is a compact reference to a label
// containing it's ID and name
type LabelReference struct {
	ResourceURL string `json:"resource_url"`
	ID          int64  `json:"id" xml:"id,attr"`
	Name        string `json:"name" xml:",innerxml"`
}

// User is a user on the Discogs website
type User struct {
	Username    string `json:"username"`
	ResourceURL string `json:"resource_url"`
}

// Rating contains the average and the amount of ratings (count)
type Rating struct {
	Average float32 `json:"average"`
	Count   int     `json:"count"`
}

// Community property within a release which contains
// Discogs specific information
type Community struct {
	Contributors []User `json:"contributors"`
	DataQuality  string `json:"data_quality"`
	Have         int    `json:"have"`
	Rating       Rating `json:"rating"`
	Status       string `json:"status"`
	Submitter    User   `json:"submitter"`
	Want         int    `json:"want"`
}

// Company contains information about the company which pressed/produced the record
type Company struct {
	Catno          string `json:"catno" xml:"catno"`
	EntityType     string `json:"entity_type" xml:"entity_type"`
	EntityTypeName string `json:"entity_type_name" xml:"entity_type_name"`
	ID             int64  `json:"id" xml:"id"`
	Name           string `json:"name" xml:"name"`
	ResourceURL    string `json:"resource_url" xml:"resource_url"`
}

// Format is a presentation of a media format for a release (e.g. Vinyl/CD)
type Format struct {
	Descriptions []string `json:"descriptions" xml:"descriptions>description"`
	Name         string   `json:"name" xml:"name,attr"`
	Text         string   `json:"text" xml:"text"`
	Qty          string   `json:"qty" xml:"qty,attr"`
}

// Identifier for a release, e.g. type barcode
type Identifier struct {
	Type        string `json:"type" xml:"type,attr"`
	Value       string `json:"value" xml:"value,attr"`
	Description string `json:"description" xml:"description"`
}

// LabelEntry is how a release is part of a certain label
type LabelEntry struct {
	Catno          string `json:"catno" xml:"catno,attr"`
	EntityType     string `json:"entity_type"`
	EntityTypeName string `json:"entity_type_name" xml:"entity_type_name"`
	ID             int64  `json:"id" xml:"id,attr"`
	Name           string `json:"name" xml:"name,attr"`
	ResourceURL    string `json:"resource_url"`
	Index          int    `json:"-" xml:"-"`
}

// Artist as it's represented on Discogs
type Artist struct {
	Name           string            `json:"name" xml:"name"`
	Namevariations []string          `json:"namevariations" xml:"namevariations>name"`
	Profile        string            `json:"profile" xml:"profile"`
	ReleasesURL    string            `json:"releases_url"`
	Realname       string            `json:"realname" xml:"realname"`
	ResourceURL    string            `json:"resource_url"`
	URI            string            `json:"uri"`
	Urls           []string          `json:"urls" xml:"urls>url"`
	DataQuality    string            `json:"data_quality" xml:"data_quality"`
	Aliases        []ArtistReference `json:"aliases" xml:"aliases>name"`
	ID             int64             `json:"id" xml:"id"`
	Images         []Image           `json:"images" xml:"images>image"`
	Members        []ArtistReference `json:"members" xml:"members>name"`
	Groups         []ArtistReference `json:"groups" xml:"groups>name"`
}

// Label contains all the Discogs Label information
type Label struct {
	Profile     string           `json:"profile" xml:"profile"`
	ReleasesURL string           `json:"releases_url"`
	Name        string           `json:"name" xml:"name"`
	ContactInfo string           `json:"contact_info" xml:"contactinfo"`
	URI         string           `json:"uri"`
	ParentLabel LabelReference   `json:"parent_label,omitempty" xml:"parentLabel"`
	Sublabels   []LabelReference `json:"sublabels" xml:"sublabels>label"`
	Urls        []string         `json:"urls" xml:"urls>url"`
	Images      []Image          `json:"images"`
	ResourceURL string           `json:"resource_url"`
	ID          int64            `json:"id" xml:"id"`
	DataQuality string           `json:"data_quality" xml:"data_quality"`
}

// Master contains all the Discogs Master information
type Master struct {
	Title                string        `json:"title" xml:"title"`
	ID                   int64         `json:"id" xml:"id,attr"`
	Artists              []ArtistEntry `json:"artists" xml:"artists>artist"`
	DataQuality          string        `json:"data_quality" xml:"data_quality"`
	Genres               []string      `json:"genres" xml:"genres>genre"`
	Images               []Image       `json:"images" xml:"images>image"`
	LowestPrice          float32       `json:"lowest_price" xml:"lowest_price"`
	MainRelease          int64         `json:"main_release" xml:"main_release"`
	MainReleaseURL       string        `json:"main_release_url"`
	MostRecentRelease    int64         `json:"most_recent_release"`
	MostRecentReleaseURL string        `json:"most_recent_release_url"`
	NumForSale           int           `json:"num_for_sale"`
	Notes                string        `json:"notes" xml:"notes"`
	ResourceURL          string        `json:"resource_url"`
	Styles               []string      `json:"styles" xml:"styles>style"`
	Tracklist            []TrackEntry  `json:"tracklist"`
	URI                  string        `json:"uri"`
	VersionsURL          string        `json:"versions_url"`
	Videos               []Video       `json:"videos" xml:"videos>video"`
	Year                 int           `json:"year" xml:"year"`
}

// TODO: find a way to get is_main_release from XML

// Release contains all the Discogs release information
type Release struct {
	Title             string        `json:"title" xml:"title"`
	ID                int64         `json:"id" xml:"id,attr"`
	Artists           []ArtistEntry `json:"artists" xml:"artists>artist"`
	ArtistsSort       string        `json:"artists_sort,omitempty" xml:"artists_sort"`
	DataQuality       string        `json:"data_quality" xml:"data_quality"`
	Thumb             string        `json:"thumb"`
	Community         Community     `json:"community"`
	Companies         []Company     `json:"companies" xml:"companies>company"`
	Country           string        `json:"country" xml:"country"`
	DateAdded         string        `json:"date_added"`
	DateChanged       string        `json:"date_changed"`
	EstimatedWeight   int           `json:"estimated_weight,omitempty" xml:"estimated_weight"`
	Extraartists      []ArtistEntry `json:"extraartists" xml:"extraartists>artist"`
	FormatQuantity    int           `json:"format_quantity"`
	Formats           []Format      `json:"formats" xml:"formats>format"`
	Genres            []string      `json:"genres" xml:"genres>genre"`
	Identifiers       []Identifier  `json:"identifiers" xml:"identifiers>identifier"`
	Images            []Image       `json:"images" xml:"images>image"`
	Labels            []LabelEntry  `json:"labels" xml:"labels>label"`
	LowestPrice       float32       `json:"lowest_price"`
	MasterID          int64         `json:"master_id" xml:"master_id"`
	MasterURL         string        `json:"master_url"`
	Notes             string        `json:"notes" xml:"notes"`
	NumForSale        int           `json:"num_for_sale"`
	Released          string        `json:"released" xml:"released"`
	ReleasedFormatted string        `json:"released_formatted"`
	ResourceURL       string        `json:"resource_url"`
	Series            []string      `json:"series"`
	Status            string        `json:"status" xml:"status,attr"`
	Styles            []string      `json:"styles" xml:"styles>style"`
	Tracklist         []TrackEntry  `json:"tracklist" xml:"tracklist>track"`
	URI               string        `json:"uri"`
	Videos            []Video       `json:"videos" xml:"videos>video"`
	Year              int           `json:"year"`
}
