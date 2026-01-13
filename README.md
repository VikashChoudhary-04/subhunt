# Subhunt

**Subhunt** is a fast, modular **subdomain enumeration tool** focused on **active DNS bruteforce**, designed to work reliably even in **restricted network environments**.

It follows modern reconnaissance practices and avoids unreliable, noisy, or misleading techniques.

---

## ✨ Features

- 🚀 Active subdomain bruteforce
- 🌐 DNS over HTTPS (DoH) support (works when UDP/53 is blocked)
- ⚡ Concurrent scanning with configurable threads
- 📊 Live progress display (tested count shown in real time)
- 🔎 Immediate output as soon as a subdomain is found
- 🧩 Wordlist-agnostic (SecLists, Assetnote, custom lists)
- 🧼 No false positives — DNS-confirmed results only
- 📦 Clean repository (no massive wordlists bundled)

---

## 🧠 Design Philosophy

Subhunt is built on a few core principles:

- **Wordlists are data, not part of the tool**
- **Accuracy > noise**
- **Network restrictions are real**
- **Do one thing well**

Instead of embedding huge datasets or depending on flaky passive sources, Subhunt focuses on **deterministic, verifiable results**.

---

## 📦 Installation

### Requirements

- Go **1.20+**
- Internet access (HTTPS required for DNS over HTTPS)

### Clone the Repository

> ⚠️ If your network blocks GitHub DNS, use:
> - Mobile hotspot  
> - Home Wi-Fi  
> - Or GitHub Web UI (Download ZIP)

```bash
git clone https://github.com/VikashChoudhary-04/subhunt.git
cd subhunt
````

---

## 🚀 Usage

### Basic Usage

```bash
go run cmd/subhunt/main.go \
  -d example.com \
  --bruteforce /path/to/wordlist.txt
```

### With Thread Control

```bash
go run cmd/subhunt/main.go \
  -d example.com \
  --bruteforce /path/to/wordlist.txt \
  --threads 50
```

---

## 📊 Output Behavior

While running, Subhunt displays live progress:

```
[+] Tested: 1247
```

As soon as a valid subdomain is found, it is printed immediately:

```
www.example.com
api.example.com
```

After completion:

```
[✓] Finished. Total tested: 5000
```

---

## 📁 Wordlists

Subhunt **does not bundle large wordlists by design**.

Users are expected to supply **community-maintained wordlists**.

Recommended sources include:

* **SecLists**
* **Assetnote**
* **ProjectDiscovery**
* **OWASP Amass**
* **Bug bounty curated lists**

📌 See [`wordlists/README.md`](wordlists/README.md) for:

* Top 10 wordlists (as of 2026)
* Recommended files
* Usage guidance
* Warnings about numeric-only lists

### Example (SecLists)

```bash
go run cmd/subhunt/main.go \
  -d example.com \
  --bruteforce /path/to/SecLists/Discovery/DNS/subdomains-top1million-5000.txt
```

---

## ⚠️ Common Mistakes

* ❌ Using numeric-only lists directly (`1`, `01`, `0001`)
* ❌ Expecting every domain to expose many subdomains
* ❌ Assuming “no output” means the tool failed

> Numeric wordlists are intended for **permutations**, not raw bruteforce.

---

## 🧪 Sanity Check

To verify correct behavior:

```bash
echo www > test.txt
go run cmd/subhunt/main.go -d yahoo.com --bruteforce test.txt
```

Expected output:

```
www.yahoo.com
[✓] Finished. Total tested: 1
```

---

## 🧱 Project Structure

```
cmd/
 └── subhunt/        # CLI entry point
internal/
 ├── bruteforce/    # Concurrent bruteforce engine
 ├── dnsresolver/   # DNS over HTTPS resolver
 └── utils/         # Helper utilities
wordlists/
 └── README.md      # Wordlist guidance (no large files)
```

---

## 🛣️ Roadmap

Planned improvements:

* 🔁 Permutation engine (api → api1, api-dev, api-v2)
* 🔄 Recursive enumeration
* 🧠 DNS result caching
* 🚫 Wildcard DNS detection
* 🌍 Multiple DoH providers with fallback
* 🧪 Debug / verbose DNS modes

---

## 👨‍💻 Author

**Vikash Choudhary**

Built with a **professional offensive-security mindset**, focusing on correctness, clarity, and real-world usability.

---

## 📜 Disclaimer

This tool is intended for **educational purposes and authorized security testing only**.
You are responsible for complying with all applicable laws and program rules.

---

## ⭐ Final Note

Subhunt is intentionally **simple, honest, and extensible**.

It doesn’t try to do everything —
it does **one thing well**, and gives you full control over how deep you go.

```

---

If you want next, I can help you:
- prepare a **v1.0 release**
- add badges (Go version, license, status)
- add example screenshots
- add a CONTRIBUTING.md

Just tell me.
```
