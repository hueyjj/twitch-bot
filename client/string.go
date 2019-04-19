package client

const helpStr = `
!help (displays this help menu)
!roll (roll a number between 0 and 100)
!hello (says hi)
!bestgirl (?)
!random (this bot says something random)
!links (list my links)
`

var hello = map[int]string{
	0:  "hello person of earth",
	1:  "good day to you",
	2:  "konnichiwa",
	3:  "salutations",
	4:  "g'day m'lady",
	5:  "sup",
	6:  "wuzzup",
	7:  "yo",
	8:  "hello",
	9:  "hi",
	10: "how ya doin",
}

var bestgirlStr = `
you
`

var random = map[int]string{
	0:  "a cow has multiple stoma",
	1:  "cat heart beats 2x the rate of human's",
	2:  "all seafood are required by law to be flash frozen",
	3:  "is a hamburger a sandwich?",
	4:  "how many times have you-",
	5:  "what even",
	6:  "absolut unit",
	7:  "attac but most importantly it protec",
	8:  "chonkers. oh lawd",
	9:  "did you fall from heaven",
	10: "sometimes life gives you a box of chocolate",
}

var links = []string{
	"jasperjeng.com",
	"github.com/hueyjj",
	"linkedin.com/in/jasperjeng",
	"medium.com/@hueyjj",
	"twitter.com/jasperjeng",
	"twitch.tv/muddycoacoa",
}
