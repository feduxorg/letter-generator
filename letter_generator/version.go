package letter_generator

type AppAuthor struct {
	Name  string
	Email string
}

type AppMetadata struct {
	Version string
	Authors []AppAuthor
	License string
}
