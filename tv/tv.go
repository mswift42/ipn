package tv

type Programme struct {
	Title     string
	Subtitle  string
	Synopsis  string
	Pid       string
	Thumbnail string
	Url       string
	Index     int
}

func NewProgramme(title, subtitle, synopsis, pid,
	thumbnail, url string) *Programme {
	return &Programme{title, subtitle, synopsis, pid,
		thumbnail, url, 0}
}
