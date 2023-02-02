package models

type File struct {
	Folder  string
	Content string
}

type RequestData struct {
	Index   string
	Records []File
}
