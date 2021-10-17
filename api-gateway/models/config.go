package models

// Config for config.json
type Config struct {
	NewsServerCount                int    `json:"newsServerCount"`
	CommentsServerCount            int    `json:"commentsServerCount"`
	NewsServerEndpointTemplate     string `json:"newsServerEndpointTemplate"`
	CommentsServerEndpointTemplate string `json:"commentsServerEndpointTemplate"`
}
