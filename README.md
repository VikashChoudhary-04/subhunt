# Subhunt

**Subhunt** is a fast, minimal, and reliable **active subdomain enumeration tool** built using **DNS over HTTPS (DoH)**.

It is designed for **real-world reconnaissance** where traditional DNS (UDP/53), passive data sources, or large recon frameworks are unreliable, blocked, or noisy.

Subhunt focuses on **correctness, clarity, and clean CLI behavior**.

---

## ✨ Features

- 🚀 Active subdomain bruteforce enumeration  
- 🌐 DNS over HTTPS (DoH) with multiple resolver failover
  - Cloudflare
  - Google
  - Quad9
- ⚡ Concurrent scanning with configurable threads  
- 📊 Single live status line (no spam, no flicker)  
- 🔎 Results printed immediately when found  
- 🧩 Wordlist-agnostic (SecLists, Assetnote, custom lists)  
- 🧼 Zero false positives (live DNS verification)  
- 🤫 `--quiet` mode for automation and pipelines  
- 🧪 Meaningful exit codes for scripting
- 🧠 DNS result caching (faster scans, fewer network requests)
- 🎯 Wildcard DNS detection with explicit warning
- 📄 `--json` mode for machine-readable output
- 📊 Smoothed rate calculation (sliding window)

---

## 🧠 Design Philosophy

Subhunt is intentionally **simple and opinionated**.

- **Active enumeration only**
- **No passive data sources** (no crt.sh, APIs, or third-party datasets)
- **Wordlists are external data**, not part of the tool
- **Accuracy over noise**
- **Verification-first**: every result is DNS-confirmed at scan time

Subhunt favors deterministic, verifiable results over maximum coverage.
If a subdomain is printed, it was resolvable during the scan.

---

## 📦 Installation

### Requirements

- Go **1.20+**
- Internet access (HTTPS required for DNS over HTTPS)

### Clone the Repository

```bash
git clone https://github.com/VikashChoudhary-04/subhunt.git
cd subhunt
```

> ⚠️ If `git clone` is slow or fails on restricted networks, use:
> - a different network (mobile hotspot / home Wi-Fi)
> - or GitHub Web UI → **Download ZIP**

---

## 🚀 Usage

### Basic Usage

```bash
go run ./cmd/subhunt \
  -d example.com \
  --bruteforce /path/to/wordlist.txt
```

### Increase Concurrency

```bash
go run ./cmd/subhunt \
  -d example.com \
  --bruteforce /path/to/wordlist.txt \
  --threads 100
```

### Quiet Mode (Results Only)

```bash
go run ./cmd/subhunt \
  -d example.com \
  --bruteforce /path/to/wordlist.txt \
  --quiet
```

### JSON Output Mode (Automation)

```bash
go run ./cmd/subhunt \
  -d example.com \
  --bruteforce /path/to/wordlist.txt \
  --json
```

Example JSON output:

{
  "domain": "example.com",
  "found": ["www.example.com", "api.example.com"],
  "tested": 5000,
  "found_count": 2,
  "duration": "2.3s"
}

---

## 📊 Output Behavior

### Live Status Line (stderr)

```text
[RUNNING] Tested: 4214 | Found: 19 | Rate: 247/s
```

- Single line
- Updated in place
- Never duplicated
- Never mixed with results
- Smoothed rate (sliding window, not total average)

### Results (stdout)

```text
[+] www.example.com
[+] api.example.com
```

- Printed immediately when found
- Always start on a new line
- Safe for piping into other tools

---

## 🔄 Exit Codes

Subhunt uses **automation-friendly exit codes**:

| Condition | Exit Code |
|--------|-----------|
| At least one subdomain found | `0` |
| No subdomains found | `1` |
| Invalid usage | `1` |

### Example

```bash
subhunt ... && echo "Subdomains found"
```

---

## 📁 Project Structure

```text
cmd/
 └── subhunt/
     └── main.go        # CLI entry point
internal/
 ├── bruteforce/
 │   └── dns.go         # Concurrent bruteforce engine
 ├── dnsresolver/
 │   └── doh.go         # DNS over HTTPS resolver
 └── ui/
     └── ui.go          # CLI UI helpers
wordlists/
 └── README.md          # Wordlist guidance (no lists bundled)
.gitignore
README.md
go.mod
```

---

## 📁 Wordlists

Subhunt **does not bundle wordlists**.

You are expected to use **community-maintained wordlists**, such as:

- SecLists  
- Assetnote  
- ProjectDiscovery  
- OWASP Amass  
- Bug bounty curated lists  

See [README](wordlists/README.md) for:
- Recommended wordlists (as of 2026)
- Usage guidance
- Warnings about numeric-only lists

### Example (SecLists)

```bash
go run ./cmd/subhunt \
  -d example.com \
  --bruteforce /usr/share/wordlists/seclists/Discovery/DNS/subdomains-top1million-5000.txt
```

---

## Notes on Accuracy

- Wildcard DNS is detected and explicitly reported
- Results are not suppressed to avoid false negatives
- Large wordlists are handled safely (scanner buffer adjusted)
- DNS results are cached in memory for performance

---

## ⚠️ Common Mistakes

- ❌ Expecting every domain to expose many subdomains  
- ❌ Using numeric-only lists directly  
- ❌ Assuming “no output” means the tool failed  

> Numeric lists are typically used for **permutations**, not raw bruteforce.

---

## 🧪 Quick Sanity Test

```bash
echo www > test.txt
go run ./cmd/subhunt -d yahoo.com --bruteforce test.txt
```

Expected output:

```text
[+] www.yahoo.com
```

Check exit code:

```bash
echo $?
# 0
```

---

## 👨‍💻 Author

**Vikash Choudhary**

Built with a **professional offensive-security mindset**, focusing on correctness, clean UX, and real-world constraints.

---

## 📜 Disclaimer

This tool is intended for **educational purposes and authorized security testing only**.  
You are responsible for complying with all applicable laws and program rules.

---

## ⭐ Final Note

Subhunt is intentionally **minimal, honest, and predictable**.

It does not try to replace large recon frameworks.  
It provides **clean, verifiable results** you can trust and build upon.
