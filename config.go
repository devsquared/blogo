package blogo

// Config holds all config needed to parse and generate the static site.
type Config struct {
	// The Blog portion of the BlogoConfig defines the basic parts of the blog.
	Blog struct {
		SourceFolder      string `yaml:"SourceFolder"`      // Source of your blog contents.
		DestinationFolder string `yaml:"DestinationFolder"` // Destination of your generated static site.
		BaseURL           string `yaml:"BaseURL"`           // BaseURL of site. Used to create links.
		Language          string `yaml:"Language"`          // Language to use for HTML writing.

		//TODO: we can package in things like:
		//	- social links
		//	- aws deployment config
	}
}
