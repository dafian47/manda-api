package model

type MandaSocialAccount struct {
	UserID      string `json:"user_id" gorm:"unique"`
	FacebookID  string `json:"facebook_id"`
	TwitterID   string `json:"twitter_id"`
	InstagramID string `json:"instagram_id"`
	YouTubeID   string `json:"you_tube_id"`
	GitHubID    string `json:"git_hub_id"`
	LinkedInID  string `json:"linked_in_id"`
}
