package simpleresty

// Option is a functional option for configuring the API client.
type Option func(*HttpClient) error

// parseOptions parses the supplied options functions and returns a configured *HttpClient instance.
func (c *HttpClient) parseOptions(opts ...Option) error {
	// Range over each options function and apply it to our API type to
	// configure it. Options functions are applied in order, with any
	// conflicting options overriding earlier calls.
	for _, option := range opts {
		err := option(c)
		if err != nil {
			return err
		}
	}

	return nil
}