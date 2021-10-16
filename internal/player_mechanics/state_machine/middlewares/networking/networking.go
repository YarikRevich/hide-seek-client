package networking

func UseNetworkingMiddleware(c func()) {
	c()
}
