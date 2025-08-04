package bounce_parser

import (
	"errors"
	"io"
	"strings"

	"github.com/emersion/go-message/mail"
)

func Parse(subject string, mr *mail.Reader) (BounceInfo, error) {
	if contains(subject, HardBounceSubjectList) {
		return BounceInfo{Type: BounceHard, Reason: "Mailbox unreachable", Subject: subject, Mailbox: extractMailbox(mr)}, nil
	}

	for {
		p, err := mr.NextPart()
		if err != nil {
			break
		}

		switch p.Header.(type) {
		case *mail.InlineHeader, *mail.Header:
			body, err := readBody(p.Body)
			if err != nil {
				return BounceInfo{}, err
			}

			if contains(body, HardBounceBodyList) {
				return BounceInfo{Type: BounceHard, Reason: "Hard bounce body", Subject: subject, Mailbox: extractMailbox(mr)}, nil
			}

			if contains(body, SoftBounceBodyList) {
				return BounceInfo{Type: BounceSoft, Reason: "Soft bounce body", Subject: subject, Mailbox: extractMailbox(mr)}, nil
			}
		}
	}

	return BounceInfo{Type: BounceNone, Subject: subject}, nil
}

// contains checks if the string contains any of the patterns
func contains(s string, patterns []string) bool {
	s = strings.ToLower(s)
	for _, p := range patterns {
		if strings.Contains(s, p) {
			return true
		}
	}
	return false
}

// readBody read body
func readBody(r io.Reader) (string, error) {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, r)
	if err != nil {
		return "", errors.New(err.Error() + ": " + buf.String())
	}
	return strings.ToLower(buf.String()), nil
}

func extractMailbox(mr *mail.Reader) string {
	// Global header for gmail X-Failed-Recipients
	if failed := mr.Header.Get("X-Failed-Recipients"); failed != "" {
		return strings.TrimSpace(failed)
	}

	// for other (Delivery-Status, RFC822)
	for {
		p, err := mr.NextPart()
		if err != nil {
			break
		}

		contentType := strings.ToLower(p.Header.Get("Content-Type"))

		// message/delivery-status
		if strings.HasPrefix(contentType, "message/delivery-status") {
			body, err := readBody(p.Body)
			if err != nil {
				continue
			}
			lines := strings.Split(body, "\n")
			for _, line := range lines {
				lowerLine := strings.ToLower(line)
				if strings.HasPrefix(lowerLine, "final-recipient:") ||
					strings.HasPrefix(lowerLine, "original-recipient:") ||
					strings.HasPrefix(lowerLine, "x-failed-recipients:") {
					parts := strings.Split(line, ";")
					if len(parts) > 1 {
						return strings.TrimSpace(parts[len(parts)-1])
					}
				}
			}
		}

		// message/rfc822
		if strings.HasPrefix(contentType, "message/rfc822") {
			subReader, err := mail.CreateReader(p.Body)
			if err != nil {
				continue
			}
			if to := subReader.Header.Get("To"); to != "" {
				return strings.TrimSpace(to)
			}
		}
	}

	return ""
}
