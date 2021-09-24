package aurrpc

type InfoRPC struct {
	Version     int    `json:"version"`
	Type        string `json:"type"`
	Resultcount int    `json:"resultcount"`
	Results     []struct {
		ID             int         `json:"ID"`
		Name           string      `json:"Name"`
		PackageBaseID  int         `json:"PackageBaseID"`
		PackageBase    string      `json:"PackageBase"`
		Version        string      `json:"Version"`
		Description    string      `json:"Description"`
		URL            string      `json:"URL"`
		NumVotes       int         `json:"NumVotes"`
		Popularity     float64     `json:"Popularity"`
		OutOfDate      interface{} `json:"OutOfDate"`
		Maintainer     string      `json:"Maintainer"`
		FirstSubmitted int         `json:"FirstSubmitted"`
		LastModified   int         `json:"LastModified"`
		URLPath        string      `json:"URLPath"`
		Depends        []string    `json:"Depends"`
		MakeDepends    []string    `json:"MakeDepends,omitempty"`
		Conflicts      []string    `json:"Conflicts,omitempty"`
		Provides       []string    `json:"Provides,omitempty"`
		Replaces       []string    `json:"Replaces,omitempty"`
		License        []string    `json:"License"`
		Keywords       []string    `json:"Keywords"`
		OptDepends     []string    `json:"OptDepends,omitempty"`
	} `json:"results"`
}

type SearchRPC struct {
	Version     int    `json:"version"`
	Type        string `json:"type"`
	Resultcount int    `json:"resultcount"`
	Results     []struct {
		ID             int         `json:"ID"`
		Name           string      `json:"Name"`
		PackageBaseID  int         `json:"PackageBaseID"`
		PackageBase    string      `json:"PackageBase"`
		Version        string      `json:"Version"`
		Description    string      `json:"Description"`
		URL            string      `json:"URL"`
		NumVotes       int         `json:"NumVotes"`
		Popularity     float64     `json:"Popularity"`
		OutOfDate      interface{} `json:"OutOfDate"`
		Maintainer     string      `json:"Maintainer"`
		FirstSubmitted int         `json:"FirstSubmitted"`
		LastModified   int         `json:"LastModified"`
		URLPath        string      `json:"URLPath"`
	} `json:"results"`
}
