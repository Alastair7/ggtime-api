package domain

// TODO: Add more platforms
type PlatformEnum int

const (
	PS3 PlatformEnum = 9
	PS4 PlatformEnum = 48
	PS5 PlatformEnum = 167
)

func (p PlatformEnum) String() string {
	switch p {
	case PS3:
		return "ps3"
	case PS4:
		return "ps4--1"
	case PS5:
		return "ps5"
	default:
		return "unknown"
	}
}
