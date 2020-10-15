package types

const (
	// ModuleName defines the module name
	ModuleName = "step3"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_step3"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	PostKey = "Post"
)

const (
	CommentKey = "Comment"
)
