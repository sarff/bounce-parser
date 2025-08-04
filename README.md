# Bounce-Parser

Bounce-Parser — lightweight Go library for detecting and parsing email bounce notifications (hard and soft).  
Supports extracting failed recipient address (`Mailbox`) from multiple email formats used by Gmail, Outlook, Yahoo, Exchange, Postfix, SendGrid, SES, etc.

---

## 🚀 Features
- Detect **Hard bounce** (permanent delivery failure)
- Detect **Soft bounce** (temporary delivery failure)
- Extract **Mailbox** (failed recipient) from:
    - `X-Failed-Recipients` header
    - `message/delivery-status` part (`Final-Recipient`, `Original-Recipient`)
    - `message/rfc822` part (`To` header of original message)
- Compatible with Gmail, Outlook, Yahoo, Postfix, Exchange, SendGrid, SES bounce formats
- Minimal dependencies — built on top of [emersion/go-message](https://github.com/emersion/go-message)

---

## 📦 Installation
```bash
go get github.com/sarff/bounce-parser
```

---

## 🛠 Usage Example
See the example folder

---

## 📋 BounceInfo
`Parse()` returns a `BounceInfo` struct:
```go
type BounceInfo struct {
	Type     BounceType // hard, soft, none
	Reason   string     // human-readable reason
	Subject  string     // original subject
	Mailbox  string     // failed recipient address
}
```

---

## ⚡ Supported Bounce Patterns
### Subjects
- `Undelivered Mail`
- `Mail Delivery Failed`
- `Delivery Status Notification (Failure)`
- `Failure Notice`
- `Undeliverable:`
- `Your message wasn't delivered`

### Body keywords (Hard)
- `550`, `553`, `User unknown`, `No such user`, `Mailbox does not exist`
- `Address not found`, `Recipient address rejected`, `Invalid recipient`

### Body keywords (Soft)
- `Mailbox full`, `Quota exceeded`, `Temporarily deferred`
- `Greylisted`, `Server busy`, `Connection timed out`

---
# 📝 TODO List for Bounce-Parser

## ✅ Core Functionality (done)
- [x] Parse Hard/Soft bounce (basic patterns)
- [x] Extract `Mailbox` from:
    - [x] `X-Failed-Recipients` header
    - [x] `Final-Recipient` / `Original-Recipient` in `message/delivery-status`
    - [x] `To` header in `message/rfc822`
- [x] `BounceInfo` struct with `Type`, `Reason`, `Subject`, `Mailbox`

---

## 🧪 Tests
- [ ] Write unit tests for `Parse()` with:
    - [ ] Hard bounce detection by subject
    - [ ] Hard bounce detection by body
    - [ ] Soft bounce detection by body
    - [ ] Mailbox extraction from `X-Failed-Recipients`
    - [ ] Mailbox extraction from `Final-Recipient` / `Original-Recipient`
    - [ ] Mailbox extraction from `To` in `message/rfc822`
- [ ] Add table-driven tests for various providers
- [ ] Add golden files for sample DSN messages

---

## 🔍 Providers to Test
- [x] **Gmail** (X-Failed-Recipients, Final-Recipient)
- [ ] **Outlook / Exchange** (Original-Recipient, Undeliverable subjects)
- [ ] **Yahoo Mail** (Final-Recipient, localized subjects)
- [ ] **ProtonMail** (check DSN structure — often RFC822 + localized)
- [ ] **iCloud / Apple Mail** (Apple DSN format)
- [ ] **Zoho Mail** (Final-Recipient and policy-based rejections)
- [ ] **Postfix** (typical `Undelivered Mail Returned to Sender`)
- [ ] **Exim** (Final-Recipient in delivery-status, no X-Failed-Recipients)
- [ ] **SendGrid** (Custom DSN formats + X-Failed-Recipients)
- [ ] **Amazon SES** (delivery-status only, no X-Failed-Recipients)
- [ ] **Mailgun** (Custom DSN JSON — check MIME fallback)
- [ ] **Fastmail** (Final-Recipient + localized bounce messages)

---

## 🚀 Future Improvements
- [ ] Add language detection for localized subjects (RU, FR, DE, ES)
- [ ] Expand bounce keyword patterns for additional languages
- [ ] Implement performance benchmarks (large mailbox scanning)
- [ ] Add CI/CD pipeline with automated tests
- [ ] Add examples for common providers in `examples/`

---
# 🤝 Contributing to BounceparserBounce-Parser

Thank you for considering contributing to **Bounce-Parser**!  
We welcome bug reports, feature requests, test cases, and pull requests for improving bounce detection and mailbox extraction.

---

## 📌 How to Contribute

### 1️⃣ Fork and Clone
```bash
git clone https://github.com/<your-username>/bounce-parser.git
cd bounce-parser
```

### 2️⃣ Create a Feature Branch
```bash
git checkout -b feature/my-improvement
```

### 3️⃣ Make Your Changes
- Add or update patterns for **Hard/Soft bounce detection**.
- Add mailbox extraction logic for **new email providers**.
- Improve parsing for `X-Failed-Recipients`, `Final-Recipient`, `Original-Recipient`, or `message/rfc822`.

### 4️⃣ Add Tests
- Add table-driven tests in `parser_test.go`.
- Add sample `.eml` test files for new providers (in `/testdata`).
- Ensure all tests pass:
```bash
go test ./...
```

### 5️⃣ Commit and Push
Use [Conventional Commits](https://www.conventionalcommits.org/) for commit messages:
```bash
feat(parser): add ProtonMail DSN mailbox extraction
fix(parser): correct Yahoo subject pattern
test(parser): add Amazon SES bounce example
```

Push your branch:
```bash
git push origin feature/my-improvement
```

### 6️⃣ Open a Pull Request
- Go to the main repository
- Open a **Pull Request** to `main`
- Provide a clear description of your changes and testing results

---

## 📂 Project Structure
```
bounceparser/
  parser.go           # Main parsing logic
  patterns.go         # Subject/Body bounce patterns
  mailbox.go          # Mailbox extraction logic
  types.go            # Types and constants
  parser_test.go      # Unit tests
  /examples           # Usage examples
  /testdata           # Sample .eml files for testing
```

---

## 🧪 Provider Test Coverage
When adding a new provider:
- [ ] Add `.eml` test message to `/testdata`
- [ ] Add pattern(s) to `patterns.go`
- [ ] Add test case to `parser_test.go`

---

## ✅ Code Style
- Follow idiomatic Go
- Use `go fmt` before committing
- Keep patterns **lowercased** for case-insensitive matching

---

Thank you for contributing! 🎉  
Your improvements make Bounce-Parser more reliable across different providers.
---

## 📜 License
MIT License — see [LICENSE](LICENSE)