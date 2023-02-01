package models

type Emails struct {
	Emails []Email
}

type Email struct {
	Username    string
	MailFolders map[string][]File
}

type File struct {
	FileName string
	Content  string
}

type FileV2 struct {
	Folder  string
	Content string
}

type RequestDataV2 struct {
	Index   string
	Records []FileV2
}

type RequestData struct {
	Index   string
	Records []Email
}
