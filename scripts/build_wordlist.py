#!/usr/bin/env python3
import re
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
SOURCES_FILE = ROOT / "wordlists" / "sources.txt"
CACHE_DIR = ROOT / ".cache_wordlists"
OUT_DIR = ROOT / "wordlists"

OUT_MEGA = OUT_DIR / "mega-subdomains.txt"
OUT_TOP = OUT_DIR / "mega-subdomains-top.txt"

# Keep only valid DNS label words for brute-forcing (NOT full FQDNs).
# We want a pure wordlist of subdomain labels like: api, dev, staging, etc.
LABEL_RE = re.compile(r"^[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?$")

def normalize_line(line: str) -> str:
    line = line.strip().lower()
    if not line or line.startswith("#"):
        return ""
    # remove leading wildcards like *.example.com
    line = line.lstrip("*.").strip(".")
    # if it's a full domain, keep only the left-most label
    # e.g. api.example.com -> api
    if "." in line:
        line = line.split(".")[0]
    # remove surrounding quotes/garbage
    line = line.strip("\"'`")
    return line

def is_valid_label(label: str) -> bool:
    if not label:
        return False
    if len(label) > 63:
        return False
    return bool(LABEL_RE.match(label))

def main():
    if not SOURCES_FILE.exists():
        print(f"Missing {SOURCES_FILE}", file=sys.stderr)
        sys.exit(1)

    CACHE_DIR.mkdir(parents=True, exist_ok=True)
    OUT_DIR.mkdir(parents=True, exist_ok=True)

    # Read downloaded files from cache (the workflow downloads them)
    # Anything under .cache_wordlists/*.txt will be consumed.
    candidates = []
    for p in sorted(CACHE_DIR.glob("*.txt")):
        candidates.append(p)

    if not candidates:
        print(f"No cached lists found in {CACHE_DIR}.", file=sys.stderr)
        sys.exit(1)

    seen = set()
    labels = []

    for path in candidates:
        with path.open("r", encoding="utf-8", errors="ignore") as f:
            for raw in f:
                w = normalize_line(raw)
                if not w:
                    continue
                if not is_valid_label(w):
                    continue
                if w in seen:
                    continue
                seen.add(w)
                labels.append(w)

    # Sort for stable diffs/commits
    labels.sort()

    OUT_MEGA.write_text("\n".join(labels) + "\n", encoding="utf-8")

    # Practical “top” slice for normal bruteforce runs
    # (Still big enough to be strong; keeps the tool usable.)
    TOP_N = 50000
    OUT_TOP.write_text("\n".join(labels[:TOP_N]) + "\n", encoding="utf-8")

    print(f"Wrote: {OUT_MEGA} ({len(labels)} unique labels)")
    print(f"Wrote: {OUT_TOP} ({min(TOP_N, len(labels))} labels)")

if __name__ == "__main__":
    main()
