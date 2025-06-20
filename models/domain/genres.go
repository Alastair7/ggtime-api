package domain

type GenreEnum int

const (
	PointAndClick        GenreEnum = 2
	Fighting             GenreEnum = 4
	Shooter              GenreEnum = 5
	Music                GenreEnum = 7
	PlatformGenre        GenreEnum = 8
	Puzzle               GenreEnum = 9
	Racing               GenreEnum = 10
	RealTimeStrategy     GenreEnum = 11
	RolePlaying          GenreEnum = 12
	Simulator            GenreEnum = 13
	Sport                GenreEnum = 14
	Strategy             GenreEnum = 15
	TurnBasedStrategy    GenreEnum = 16
	Tactical             GenreEnum = 24
	HackAndSlashBeatEmUp GenreEnum = 25
	QuizTrivia           GenreEnum = 26
	Pinball              GenreEnum = 30
	Adventure            GenreEnum = 31
	Indie                GenreEnum = 32
	Arcade               GenreEnum = 33
	VisualNovel          GenreEnum = 34
	CardAndBoardGame     GenreEnum = 35
	MOBA                 GenreEnum = 36
)

func (g GenreEnum) String() string {
	switch g {
	case PointAndClick:
		return "point-and-click"
	case Fighting:
		return "fighting"
	case Shooter:
		return "shooter"
	case Music:
		return "music"
	case PlatformGenre:
		return "platform"
	case Puzzle:
		return "puzzle"
	case Racing:
		return "racing"
	case RealTimeStrategy:
		return "real-time-strategy-rts"
	case RolePlaying:
		return "role-playing-rpg"
	case Simulator:
		return "simulator"
	case Sport:
		return "sport"
	case Strategy:
		return "strategy"
	case TurnBasedStrategy:
		return "turn-based-strategy-tbs"
	case Tactical:
		return "tactical"
	case HackAndSlashBeatEmUp:
		return "hack-and-slash-beat-em-up"
	case QuizTrivia:
		return "quiz-trivia"
	case Pinball:
		return "pinball"
	case Adventure:
		return "adventure"
	case Indie:
		return "indie"
	case Arcade:
		return "arcade"
	case VisualNovel:
		return "visual-novel"
	case CardAndBoardGame:
		return "card-and-board-game"
	case MOBA:
		return "moba"
	default:
		return "unknown"
	}
}
