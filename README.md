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
```

### Passive enumeration only

```bash
go run cmd/subhunt/main.go -d example.com --passive
```

### DNS bruteforce + validation

```bash
go run cmd/subhunt/main.go -d example.com --bruteforce wordlists/small.txt --resolve
```

### 🔌 Pipeline Example

```bash
subhunt -d example.com --passive --resolve | httpx
```

---

## 🛡️ Philosophy

(i) Enumeration without validation is noise

(ii) DNS resolution is mandatory

(iii) One responsibility per module

(iv) No external tool wrapping

---

## 📌 Roadmap

(i) Concurrent DNS resolution

(ii) Permutation-based enumeration

(iii) JSON output

(iv) Wildcard DNS detection

(v) HTTP probing module

---

## ⚠️ Disclaimer

This tool is intended for authorized security testing only.

---

## ⭐ Author

Built with a professional red-team mindset.
