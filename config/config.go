package config

// Site
// Profile configs for domains
type Site struct {
	Domain  string
	Profile string
}

// Profile
// Google chrome profile info
type Profile struct {
	Profile     string
	DisplayName string
}

// Settings
// The application settings
type Settings struct {
	ChromePath string
	Site       []Site
	Profile    []Profile
}
