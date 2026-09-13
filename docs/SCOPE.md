# Linistic Scope

## What is Linistic?

Linistic is a Linux troubleshooting assistant that helps users diagnose and understand system issues by automatically collecting evidence, analyzing system state, and presenting likely causes in plain English.

The goal is not to replace Linux debugging tools, but to reduce the time required to identify the root cause of common Linux problems.

---

# Problem Statement

Linux provides powerful debugging and observability tools such as:

* journalctl
* dmesg
* ss
* ip
* systemctl
* lsof
* strace
* tcpdump

However, troubleshooting often requires manually correlating information across multiple subsystems including:

* Processes
* Services
* Networking
* Drivers
* Hardware
* Kernel logs
* Power management

For many users, the difficult part is not collecting information but understanding what the information means.

Linistic aims to bridge that gap.

---

# Target Audience

## Primary Audience

* Linux developers
* Students learning Linux
* Homelab users
* Linux enthusiasts
* Junior system administrators

These users frequently encounter issues but may not know where to begin troubleshooting.

Examples:

* WiFi connected but internet unavailable
* Service running but unreachable
* VPN connected but not working
* Unexpected reboot
* Random system slowdown
* Driver or firmware problems

---

## Secondary Audience

* System administrators
* DevOps engineers
* SREs

These users already know the debugging tools but may benefit from faster first-level diagnosis and incident summaries.

---

# What Linistic Is

Linistic is:

* A Linux troubleshooting assistant
* A diagnostic engine
* A post-mortem investigation tool
* A system health analyzer
* A plain-English explanation layer over Linux diagnostics

---

# What Linistic Is Not

Linistic is not:

* A monitoring dashboard
* A SIEM platform
* A security product
* A cloud observability platform
* A log aggregation system
* A replacement for Coroot, Grafana, Prometheus, or Falco
* A replacement for strace, perf, gdb, tcpdump, or bpftrace

Instead, Linistic uses information from Linux and existing tools to explain what is likely happening and why.

---

# Core Goals

## 1. Live Diagnostics

Help users diagnose current problems.

Example:

```bash
linistic diagnose
```

Potential output:

```text
Issue Detected:
WiFi Connected but Internet Unreachable

Likely Cause:
DNS configuration issue

Confidence:
82%
```

---

## 2. Incident Investigation

Help users understand what happened after a failure.

Example:

```bash
linistic report
```

Potential output:

```text
Previous Boot Analysis

Detected:
Unexpected Shutdown

Possible Causes:
- Power loss
- Kernel panic
- Watchdog reset
```

---

## 3. System Health Analysis

Detect recurring or hidden issues before they become major problems.

Examples:

* Repeated driver resets
* Thermal throttling
* Frequent network disconnects
* Excessive interrupts
* Hardware instability

---

# Initial Diagnostic Scope

The initial MVP focuses on the following categories:

## Networking

* Connected to WiFi but no Internet
* WiFi authentication failures
* Random WiFi disconnects
* VPN connected but no traffic
* Driver and firmware networking issues

## Services

* Service running but unreachable
* Port conflicts
* Service startup failures

## System Performance

* High load average with low CPU usage
* Thermal throttling
* Interrupt storms

## Process Analysis

* Process refuses to exit
* Hung processes
* Blocked tasks

## Storage

* Disk full despite file deletion
* Filesystem-related stalls

## Incident Analysis

* Unexpected reboot investigation
* Freeze before reboot analysis
* Driver and firmware reset detection

---

# Long-Term Vision

Linistic should act like an experienced Linux engineer performing first-level troubleshooting.

Given a symptom, it should:

1. Gather relevant evidence.
2. Correlate data from multiple sources.
3. Generate likely explanations.
4. Present findings in plain English.
5. Suggest next investigative steps.
6. Reduce debugging time from hours to minutes.

The primary objective is not monitoring Linux systems.

The primary objective is understanding why Linux systems behave unexpectedly.
