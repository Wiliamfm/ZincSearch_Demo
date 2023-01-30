package models

type Emails struct {
	Emails []Email
}

type Email struct {
	Username    string
	MailFolders map[string][]string
}

type File struct {
	FileName string
	Content  string
}
