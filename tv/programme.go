package programme

type Programme struct {
	title     string
	subtitle  string
	synopsis  string
	pid       string
	thumbnail string
	url       string
	index     int
}

func NewProgramme(title, subtitle, synopsis, pid,
	thumbnail, url string) *Programme {
	return &Programme{title, subtitle, synopsis, pid,
		thumbnail, url, 0}
}
