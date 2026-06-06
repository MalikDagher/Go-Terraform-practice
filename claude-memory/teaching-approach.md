---
name: teaching-approach
description: How Malik wants to be taught Go/Terraform
metadata: 
  node_type: memory
  type: feedback
  originSessionId: a7d9d3da-809a-4d97-a0b9-92586c1c434a
---

Malik wants to **alternate Go and Terraform** across sessions, and wants me to **vary teaching style and adapt to whatever makes concepts click for him** (rather than fixing one format up front).

**Why:** He learns slowly and wants deep mastery, not speed. He explicitly asked me to experiment and notice the pattern of what he picks up best.

**How to apply:** One concept at a time. Make him do the typing where possible (muscle memory). Always connect concepts back to how HashiCorp/Terraform uses them in real Go. Be patient; invite "wait, why?" questions. Observe which formats (small-bites-practice vs worked-example) land best and lean into them. See [[user-profile]].

**What's working (observed sessions 1-2, 2026-06-04 to 06):** Small-bites + hands-on REPS is landing well. Malik explicitly values repetition ("I need to keep doing it for it to stick"). He learns strongly from **guided debugging** — when stuck, walking through his own buggy code issue-by-issue (not just giving the answer) produced real understanding. He restates concepts in his own words when they click ("declare first and print them in main" = producer/consumer). He started adding his own test cases unprompted.

**Go progress:** Mastered functions, multiple return values, the `err != nil` error pattern (producer vs consumer). Solid.

**Terraform progress / IMPORTANT calibration:** Syntax-first did NOT work for TF — he ran init/plan/apply/destroy successfully but said it "looks like gibberish" and rated his understanding ~20%. Root cause: **he has no lived experience of the pain IaC solves** — has an AWS CCP (mostly forgotten) but has NEVER worked with servers or IaC. Abstract benefits don't land. What helped: grounding in the **surprise-AWS-bill fear** (CCP holders relate), `terraform destroy` = "no forgotten resources" button, and the blueprint/construction-crew/logbook analogy. LESSON: for TF, teach the WHY and real-world problem viscerally BEFORE syntax; he may need a foundational "what even is a server / EC2" explainer since that base is missing. Completed the change→plan→apply→destroy demo. See [[github-practice-repo]].
