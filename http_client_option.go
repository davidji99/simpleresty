package simpleresty

// Option is a functional option for configuring a client.
type Option func(interface{}) error

// parseOptions parses the supplied options functions and returns a configured *Client instance.
func ParseOptions(c interface{}, opts ...Option) error {
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