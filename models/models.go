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
