package keystore

type key struct {
	privatePEM string
	publicPEM  string
}

type KeyStory struct {
	store map[string]key
}
