# Valid Parenthesis String

> NeetCode 150 · Greedy — don't branch on each '*': carry the interval [lo, hi] of possible open counts, clamp at zero, fail when hi dips negative. Go, stdlib only, tested.

## Problem

A string of `(`, `)` and `*`, where each `*` may act as `(`, `)` or nothing. Is some assignment balanced?

## Approach

Exploring star choices is 3^k; the collapse is noticing that the set of achievable open-bracket counts at any prefix is a *contiguous range* — each character shifts it (`(` raises both ends, `)` lowers both, `*` widens by one each way). Track only the endpoints: `hi < 0` means unmatched `)` under even the most generous reading — fail; `lo` clamps at 0 because negative opens aren't real (choices making it negative are simply discarded, legal since other choices in the range remain). Balanced iff 0 is in the final range, i.e. `lo == 0`. Representing an exponential choice-set by its interval is the same trick as dsa-250's cousins in scheduling and is worth owning.

**Complexity:** O(n) time, O(1) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
