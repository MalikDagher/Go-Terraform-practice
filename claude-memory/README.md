# claude-memory — portable Claude Code memory

This folder is a **mirror** of the Claude Code memory Claude uses to remember how
to teach me (Malik) Go + Terraform. It lives in the repo so it travels with
`git pull` to any machine.

## Why it's here
Claude only *auto-reads* memory from its config folder, not from this repo:

```
<user home>\.claude\projects\c--Users-malik-learning\memory\
```

So this repo copy is the **portable source of truth**. The copy in the config
folder is the **live copy Claude actually loads** each session.

## Restoring on a new machine
1. Clone/pull this repo to `C:\Users\<you>\learning` (same path keeps the
   project folder name identical).
2. Copy these files into Claude's memory folder:

   ```powershell
   $dst = "$env:USERPROFILE\.claude\projects\c--Users-<you>-learning\memory"
   New-Item -ItemType Directory -Force -Path $dst | Out-Null
   Copy-Item ".\claude-memory\*.md" -Destination $dst -Force
   ```
   (The project folder name is derived from the repo path; if your path differs,
   the `c--Users-...-learning` part changes to match.)
3. Or simplest: just tell Claude *"read claude-memory/ and sync it into your
   memory"* at the start of a session and it will handle the copy.

## Keeping it in sync
When Claude updates the live memory, it also updates these files. After a
learning session, run the usual `git add . && git commit && git push` to save
the latest notes.
