package simpleresty

import "github.com/go-resty/resty/v2"

// Request represents a HTTP request.
//
// This struct embeds resty.Request
type Request struct {
	*resty.Request
}
