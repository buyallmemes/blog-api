package github

// Config holds the configuration for the GitHub repository
type Config struct {
	// Owner is the GitHub repository owner
	Owner string

	// Repo is the GitHub repository name
	Repo string

	// Path is the path to the blog posts directory in the repository
	Path string

	// Token is the GitHub API token
	Token string
}

// NewConfig creates a new Config instance with the given parameters
func NewConfig(owner, repo, path, token string) *Config {
	return &Config{
		Owner: owner,
		Repo:  repo,
		Path:  path,
		Token: token,
	}
}
