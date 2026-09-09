# Linistic

**Linistic** is a Linux diagnostic and root-cause analysis platform that helps users understand **what changed, what failed, and why it likely failed**.

Unlike traditional monitoring tools that focus on metrics, logs, or alerts, Linistic focuses on **diagnosis**.

Its goal is to reduce the time between:

```text
Something is broken
```

and

```text
I understand why it is broken
```

---

# Problem Statement

Linux systems generate enormous amounts of information:

* Logs
* Metrics
* System events
* Network events
* Service events
* Kernel events

The problem is not lack of data.

The problem is that users must manually investigate:

* Why did Wi-Fi stop working?
* Why is CPU usage abnormal?
* Why is memory growing?
* Why is a service crashing repeatedly?
* Why did latency suddenly increase?
* What changed before the incident?

Linistic aims to automatically correlate system events and generate diagnostic explanations.

---

# Vision

Instead of:

```text
CPU = 95%
```

Linistic should explain:

```text
CPU usage increased from 4% to 95%.

Primary process:
postgres

Observed after:
database migration

Likely cause:
migration query causing high CPU utilization

Confidence:
81%
```

---

# Core Principles

## 1. Diagnosis First

Not another monitoring dashboard.

Not another log viewer.

Focus on:

```text
What changed?
What failed?
Why?
```

---

## 2. Correlation Over Collection

Many tools already collect events.

Linistic focuses on:

```text
event correlation
timeline generation
causal relationships
root cause hypotheses
```

---

## 3. Human Readable

Users should not need to search through:

```bash
journalctl
dmesg
top
systemctl
ip
ss
```

Linistic should summarize findings in plain English.

---

# Architecture

## Layer 1 – Observation

Collect system facts.

No diagnosis.

No AI.

Only evidence.

Sources:

* Processes
* Memory
* CPU
* Disk
* Network
* Services
* System logs

---

## Layer 2 – Correlation Engine

Build relationships:

```text
File Change
      ↓
Service Restart
      ↓
Connection Failures
      ↓
CPU Spike
      ↓
Crash
```

---

## Layer 3 – Diagnosis Engine

Generate hypotheses:

```text
Possible Causes

88% nginx configuration change

53% upstream service failure

17% memory pressure
```

---

## Layer 4 – Explanation Engine

Convert diagnostic findings into human-readable incident reports.

---

# Development Phases

## Phase 1 – Userspace Observer

Goal:

Create a local system observer.

Collect:

### Process Information

* PID
* PPID
* Command
* CPU usage
* Memory usage
* Start time

### Memory Information

* RSS
* Virtual memory
* Swap usage

### CPU Information

* Per-process CPU
* System CPU
* Load average

### Network Information

* Open sockets
* Listening ports
* Connection counts

### Service Information

* systemd service state
* restart count
* failure count

Output:

```text
System Snapshot
```

---

## Phase 2 – Event Timeline

Create historical timelines.

Example:

```text
22:11 nginx restarted

22:12 CPU spike detected

22:13 memory growth started

22:15 service unavailable
```

Output:

```text
Chronological Incident Timeline
```

---

## Phase 3 – Diagnostic Rules

Create first diagnosis engine.

Examples:

### CPU

```text
CPU > 90%

Find top process

Explain process ownership
```

---

### Memory

```text
Memory growth

Identify responsible process

Track growth trend
```

---

### Services

```text
Repeated crashes

Detect restart loops

Identify crash frequency
```

---

### Network

```text
DNS failures

Gateway failures

DHCP failures

TLS failures
```

---

## Phase 4 – Change Detection

Track:

### File Changes

Critical files:

```text
/etc
/usr/lib/systemd
/etc/nginx
/etc/ssh
```

### Permission Changes

### Ownership Changes

### Service Configuration Changes

Output:

```text
What changed before the incident?
```

---

## Phase 5 – eBPF Integration

Move into kernel-level observability.

Collect:

### Process Events

```text
execve
fork
clone
exit
```

### File Events

```text
open
rename
unlink
chmod
```

### Network Events

```text
connect
accept
bind
```

### Signal Events

```text
SIGKILL
SIGTERM
SIGSEGV
```

Output:

```text
Kernel Event Stream
```

---

## Phase 6 – Causality Graph

Create relationships between events.

Example:

```text
nginx.conf modified
        ↓
nginx restarted
        ↓
worker crashes
        ↓
502 errors
```

Output:

```text
Incident Dependency Graph
```

---

## Phase 7 – Root Cause Engine

Generate probable explanations.

Example:

```text
Most Likely Cause

nginx configuration change

Confidence:
87%
```

Multiple hypotheses supported.

---

## Phase 8 – AI-Assisted Explanations

Optional.

AI should never be the source of truth.

AI only converts diagnostic results into explanations.

Example:

```text
Nginx failures began immediately after a configuration update.
The update is the most likely trigger based on observed events.
```

---

# What Makes Linistic Different?

## Existing Tools

### top / htop

Show resource usage.

Do not explain causes.

---

### journalctl

Shows logs.

Does not correlate events.

---

### Falco

Detects suspicious events.

Does not perform system diagnosis.

---

### Tetragon

Provides runtime observability.

Does not generate root-cause explanations.

---

### Coroot

Provides observability and some RCA capabilities.

Primarily cloud and service focused.

---

# Linistic's Goal

Move beyond:

```text
What happened?
```

towards:

```text
Why did it happen?
```

and eventually:

```text
What changed that caused it?
```

---

# Long-Term Vision

A Linux user should be able to run:

```bash
linistic diagnose
```

and receive:

```text
Issue Detected

Service:
nginx

Observed:
Repeated worker crashes

Recent Changes:
nginx.conf modified 4 minutes before first crash

Likely Cause:
configuration error

Confidence:
84%
```

instead of spending hours manually searching logs, metrics, and system state.
