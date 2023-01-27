package layer

type F uint64

const (
	FNone F = iota

	// FAir defines the agent can fly.
	FAir = 1 << iota
	FLand
	FSea
)
