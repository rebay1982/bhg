package shodan

// BaseURL The Shodan base URL
const BaseURL = "https://api.shodan.io"

type Client struct {
	apiKey string
}

// New Creates a new client based on the API key passed as a parameter.
func New(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}

func (c *Client) HostSearch() {

}
