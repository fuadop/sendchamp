package sendchamp

import (
	"fmt"
	"net/http"
)

// add required request headers
func addHeaders(req *http.Request, c *Client) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprint("Bearer ", c.publicKey))
}
