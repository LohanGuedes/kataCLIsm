package kata

type Kata struct { // TODO: Read from a JSON
	reviewers  map[string]string
	tags       map[int]int
	title      string
	readme     string
	link       string
	difficulty int
}

func (k Kata) FilterValue() string {
	return k.title
}

func (k Kata) Title() string {
	return k.title
}

func (k Kata) Readme() string {
	return k.readme
}
