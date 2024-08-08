package internal

type Artist struct {
	ID             int      `json:"id"`
	Image          string   `json:"image"`
	Name           string   `json:"name"`
	Members        []string `json:"members"`
	CreationDate   int      `json:"creationDate"`
	FirstAlbum     string   `json:"firstAlbum"`
	Locations      string   `json:"locations"`
	ConcertDates   string   `json:"concertDates"`
	Relations      string   `json:"relations"`
	DatesLocations map[string][]string
}

var (
	artists   = []Artist{}
	relations = Relation{}
)

type Relation struct {
	Index []struct {
		ID             int `json:"id"`
		DatesLocations map[string][]string
	}
}

type ErrorPage struct {
	Code    int
	Message string
}
