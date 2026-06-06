---
name: github-practice-repo
description: "Where Malik's Go/Terraform learning work is version-controlled"
metadata: 
  node_type: memory
  type: reference
  originSessionId: a7d9d3da-809a-4d97-a0b9-92586c1c434a
---

Malik pushes his learning playground (`c:\Users\malik\learning`) to GitHub: **https://github.com/MalikDagher/Go-Terraform-practice**

Git identity on this machine: name `MalikDagher`, email `mdagh002@fiu.edu`. No `gh` CLI installed; uses plain `git` over HTTPS.

The repo has a `.gitignore` that excludes Terraform state (`*.tfstate*`), the `.terraform/` provider-cache dir, and Go build artifacts — but KEEPS `.terraform.lock.hcl` (provider version lock, meant to be committed). Reason state is excluded: it can contain secrets and is the #1 Terraform-in-git mistake — worth reinforcing as a real HashiCorp practice. See [[teaching-approach]].

**Memory mirroring (IMPORTANT for continuity):** Malik wanted memory portable across machines, so the live memory folder is mirrored into the repo at `learning\claude-memory\`. RULE: whenever I update any live memory file in `.claude\...\memory\`, I must also copy it to `learning\claude-memory\` so they don't drift, and remind/offer to commit+push. The live copy (in `.claude`) is what Claude auto-loads; the repo copy is the portable backup restored on new machines. See `learning\claude-memory\README.md` for the restore steps.
