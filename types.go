package bounce_parser

type BounceType string

var HardBounceBodyList = []string{
	"550", "551", "552", "553", "user unknown", "unknown user", "no such user",
	"mailbox unavailable", "mailbox does not exist", "address doesn't exist", "address not found",
	"recipient address rejected", "recipient not found", "user not found", "account disabled", "account not found",
	"invalid recipient", "unrouteable address",
}

var SoftBounceBodyList = []string{
	"mailbox full", "over quota", "quota exceeded", "temporarily deferred", "temporary failure",
	"deferred", "greylisted", "server busy", "connection timed out", "resources temporarily unavailable",
	"try again later", "temporary local problem",
}

var HardBounceSubjectList = []string{
	"undelivered mail", "mail delivery failed", "failure notice", "delivery failure",
	"delivery status notification (failure)", "undeliverable:", "wasn't delivered",
	"returned mail", "message delivery failure", "your message wasn't delivered",
}

const (
	BounceHard BounceType = "hard"
	BounceSoft BounceType = "soft"
	BounceNone BounceType = "none"
)

type BounceInfo struct {
	Type    BounceType
	Reason  string
	Subject string
	Mailbox string
}
