package models

type File struct {
	Folder  string
	Content string
}

type RequestData struct {
	Index   string
	Records []File
}

type SearchRequest struct {
	SearchType string `json:"search_type"`
	Query      Query  `json:"query"`
	//SortFields string   `json:"sort_fields"`
	SortFields string   `json:"-"`
	From       int      `json:"from"`
	MaxResults int      `json:"max_results"`
	Source     []string `json:"_source"`
}

type Query struct {
	Term string `json:"term"`
	//StartTime string `json:"start_time"`
	//EndTime   string `json:"end_time"`
	StartTime string `json:"-"` //
	EndTime   string `json:"-"` //`json:"end_time"`
}

type ClientRequest struct {
	Search string
	Type   string
}
