# eg-lm-automation
Internal event and logic layer for Listmonk automation – bounce tracking, notifs, custom logic.

# Endgate Listmonk Automation

Custom Go automation layer for sovereign installs.

## Usage
```bash
curl -fsSL https://raw.githubusercontent.com/your-user/eg-lm-automation/main/install.sh | bash

---

### 🔹 3. `go.mod` and `go.sum`
Already generated from your VPS build. Copy them **as-is** from `/opt/lm-automation/`.

---

### 🔹 4. `/internal` and `/cmd`
Copy your full working Go project structure from VPS into the repo:
```bash
scp -r root@YOUR_VPS:/opt/lm-automation/* ./ # or rsync
