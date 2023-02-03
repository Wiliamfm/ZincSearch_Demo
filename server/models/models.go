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
	SearchType  string
	Query       map[string]string
	Sort_fields string
	From        int
	Max_results int
	Source      []string
}

type ClientRequest struct {
	Search string
	Type   string
}
