package blogo

// Generator generates the static site from the config provided.
type Generator interface {
	GenerateSite(conf Config)
	GeneratePage(md string)
	GeneratePost(md string)
}
