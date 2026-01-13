# Wordlists

Subhunt does **not** bundle large wordlists by default.

This is intentional.

Professional reconnaissance tools expect users to supply **community-maintained, battle-tested wordlists** rather than embedding massive datasets inside the repository.

Below are the **Top 10 most widely used and trusted wordlists in the world (as of 2026)** for subdomain enumeration and web reconnaissance.

---

## 🔝 Top 10 Wordlists (2026)

### 1. SecLists (Daniel Miessler)
**The industry standard**

Repository:
https://github.com/danielmiessler/SecLists

Recommended subdomain files:
- `Discovery/DNS/subdomains-top1million-5000.txt`
- `Discovery/DNS/subdomains-top1million-20000.txt`
- `Discovery/DNS/bitquark-subdomains-top100000.txt`

Why it matters:
- Maintained for years
- Used by most recon tools
- Best starting point for almost all targets

---

### 2. Assetnote Wordlists
**High-signal, bug-bounty focused**

Repository:
https://wordlists.assetnote.io/

Recommended:
- `best-dns-wordlist.txt`
- `dns.txt`

Why it matters:
- Curated from real bug bounty findings
- Excellent signal-to-noise ratio
- Frequently updated

---

### 3. ProjectDiscovery Wordlists
**Built for modern recon tooling**

Repository:
https://github.com/projectdiscovery/fuzzing-templates

Why it matters:
- Created by the team behind `subfinder`, `nuclei`, `httpx`
- Cloud, API, SaaS focused
- Modern naming patterns

---

### 4. Jason Haddix Wordlists
**Methodology-driven recon**

Source:
https://github.com/jhaddix/domain

Why it matters:
- Based on real recon workflows
- Excellent for phased enumeration
- Often used in bug bounty playbooks

---

### 5. Amass Wordlists
**Research-grade enumeration**

Repository:
https://github.com/owasp-amass/amass

Why it matters:
- Used internally by Amass
- Structured and research-backed
- Good for deep enumeration

---

### 6. RapidDNS / Recon Datasets
**Passive-data derived lists**

Source:
- RapidDNS datasets
- Certificate Transparency derived lists

Why it matters:
- Derived from real DNS observations
- Good for validating large infrastructures

---

### 7. Bug Bounty Community Curated Lists
**Crowd-sourced, practical**

Sources:
- HackerOne disclosed reports
- Bugcrowd writeups
- Independent recon blogs

Why it matters:
- Reflects real-world attack surface
- Often uncovers non-obvious subdomains

---

### 8. FuzzDB
**Classic but still useful**

Repository:
https://github.com/fuzzdb-project/fuzzdb

Why it matters:
- Older but stable
- Useful for legacy and enterprise targets
- Complements modern lists

---

### 9. DNS Permutation / Numeric Lists
**NOT standalone — used with permutations**

Examples:
- Numeric-only lists
- Suffix/prefix lists

Why it matters:
- Used to generate names like `api1`, `web02`, `node003`
- Only useful with a permutation engine
- Not meant for direct bruteforce

---

### 10. Custom Target-Specific Wordlists
**Highest ROI**

How they are created:
- From discovered subdomains
- From application names
- From company-specific terminology

Why it matters:
- Outperforms generic lists
- Essential for serious recon
- Used by top bug bounty hunters

---

## ⚠️ Important Notes

- Bigger wordlists do **not** guarantee better results
- Start small, then expand intelligently
- Numeric-only lists require permutation logic
- Wordlists are **data**, not part of the tool

---

## Example Usage with SecLists

```bash
subhunt -d example.com \
  --bruteforce /path/to/SecLists/Discovery/DNS/subdomains-top1million-5000.txt --threads 150
```

---

## Recommended Workflow

- Start with a small or top-5000 list
- Analyze discovered subdomains
- Generate permutations
- Use larger lists only if necessary
