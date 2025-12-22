# LinkedIn Automation Proof of Concept (Go + Rod)

## Disclaimer
This project is built strictly for educational and technical evaluation purposes.  
It must not be used on real LinkedIn accounts or deployed in production environments.

---

## Project Overview
This repository contains a Go-based browser automation proof of concept built using the Rod library.  
The focus is on demonstrating clean Go architecture, human-like automation behavior, and reliable state persistence rather than production automation.

---

## Architecture

linkedin-automation/
├── auth/ # Authentication logic
├── config/ # Environment configuration
├── stealth/ # Human-like behavior utilities
├── storage/ # JSON-based state persistence
├── utils/ # Shared helpers
├── main.go # Entry point
├── state.json # Persisted automation state


---

## Authentication Flow
- Credentials are loaded from environment variables
- Login is performed via controlled browser automation
- Success is detected through DOM inspection
- Execution stops safely on failure

---

## Stealth Techniques Implemented
- Human-like typing with variable delays
- Randomized think time between actions
- Hover-before-click mouse behavior
- Action pacing and rate limiting
- Session-aware execution flow

Runtime browser fingerprint modification was intentionally avoided for stability, with priority given to behavioral stealth.

---

## State Persistence
Automation state is persisted in `state.json` using a thread-safe JSON store.

Example:
```json
{
  "logins": [
    {
      "timestamp": "2025-12-22T21:08:31+05:30",
      "status": "success"
    }
  ]
}
