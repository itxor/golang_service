package server

import (
	"net/http"
	"time"
)

// HTTPClient ...
var HTTPClient = &http.Client{
	Timeout: time.Second * 10,
}
