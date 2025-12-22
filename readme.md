# LinkedIn Automation Proof of Concept (Go + Rod)

## Project Overview
This repository contains a Go-based LinkedIn automation proof of concept built using the Rod browser automation library.  
The focus is on demonstrating clean Go architecture, human-like automation behavior, stealth techniques, and safe state persistence rather than building a production-ready automation tool.

---

## Architecture

linkedin-automation/
├── auth/ # Authentication and login flow
├── config/ # Environment configuration loader
├── cookies/ # Session cookie persistence
├── search/ # LinkedIn people search logic (POC)
├── targeting/ # Profile parsing and deduplication
├── connections/ # Connection request handling (rate-limited)
├── messaging/ # Template-based messaging logic
├── scheduler/ # Activity scheduling and business-hour checks
├── stealth/ # Anti-detection & human-like behavior utilities
├── storage/ # JSON-based state persistence
├── main.go # Application entry point
├── state.json # Persisted login and action state
├── cookies.json # Persisted browser session cookies


---

## Authentication Flow
- Credentials are loaded securely from environment variables
- Login is performed via controlled browser automation
- Success is detected through DOM inspection
- Failures are handled gracefully and execution stops safely
- Session cookies are persisted after successful login

---

## Cookie Persistence
Browser session cookies are saved to `cookies.json` after successful authentication.  
This demonstrates session reuse capability and persistent browser state across runs.

Cookie reuse logic is implemented conservatively to avoid browser lifecycle instability during demo execution.

---

## Stealth Techniques Implemented
- Human-like typing with variable keystroke delays
- Randomized think time between actions
- Hover-before-click interaction behavior
- Rate limiting and action pacing
- Session-aware execution flow
- Human-like mouse movement using Bézier curves
- Browser-native mouse events dispatched via JavaScript
- Business-hour–aware execution design

Runtime fingerprint manipulation (e.g., `navigator.webdriver`) is intentionally avoided for stability and safety, with emphasis placed on behavioral stealth techniques.

---

## State Persistence
Automation state is persisted in `state.json` using a thread-safe JSON store.  
This enables resumability, auditing, and controlled execution.

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

## Design Notes
- High-risk LinkedIn actions are safety-gated to respect platform constraints
- Search, connection, and messaging modules are implemented as extensible components
- The project prioritizes stability, clarity, and evaluability over aggressive automation

---

## Conclusion
This project demonstrates a complete LinkedIn automation architecture with real browser automation, stealth-aware interaction patterns, session persistence, and clean Go modular design.  
It is intentionally scoped as a proof of concept to highlight engineering judgment and system design rather than production-scale automation.
