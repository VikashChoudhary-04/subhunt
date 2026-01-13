# SubHunt

**SubHunt** is a modular, fast, and clean **subdomain enumeration framework** written in **Go**.

It is designed to replace outdated and fragmented tools like:
- Sublist3r
- Knock
- Turbolist3r
- Racoon (partial)
- dnsx (basic resolution logic)

SubHunt focuses on **quality over quantity**:
- Clean output
- DNS-validated results
- Modular enumeration strategies
- Pipeline-friendly usage

---

## ✨ Features

- Passive subdomain enumeration (crt.sh)
- Active DNS bruteforce
- DNS resolution & validation
- Deduplication
- Clean CLI output (stdout)
- Designed for bug bounty & pentesting workflows

---

## 📁 Project Structure

subhunt/
├── cmd/subhunt/main.go # CLI entry point
├── internal/
│ ├── passive/ # Passive sources
│ ├── bruteforce/ # DNS bruteforce engine
│ ├── resolver/ # DNS resolution & validation
│ ├── utils/ # Helpers (dedupe, etc.)
├── wordlists/ # Bruteforce wordlists
├── go.mod
└── README.md


---

## 🚀 Usage

### Enumerate subdomains (passive + DNS validation)

```bash
go run cmd/subhunt/main.go -d example.com --passive --resolve
