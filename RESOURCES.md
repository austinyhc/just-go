# SWE Interview Prep Resources (Google / AWS / Microsoft)

## Knowledge

### Company processes
- [Interviewing at Google: best practices, advice, tips](https://www.google.com/about/careers/applications/interview-tips)
  Official Google Careers page. Use for: general interview-day expectations.
- [Our hiring process — Google Careers](https://www.google.com/about/careers/applications/how-we-hire)
  Official. Use for: understanding the screen → onsite loop → hiring-committee pipeline (Google decides via committee, not the interviewer alone).
- [Software Engineer (SWE) interview prep guide — Google Careers](https://www.google.com/about/careers/applications/candidate-prep/swe)
  Official, role-specific. Use for: what Google says it evaluates in SWE loops.
- [Leadership Principles — Amazon.jobs](https://www.amazon.jobs/en/principles)
  Official, primary source, all 16 LPs verbatim. Use for: canonical LP wording — quote this, don't paraphrase from memory.
- [Interview Loop — Amazon.jobs](https://amazon.jobs/content/en/how-we-hire/interview-loop)
  Official. Use for: how the bar-raiser/panel loop structure works.
- [Life at AWS: recruiters share 10 ways to excel in your AWS in-person interview](https://aws.amazon.com/careers/life-at-aws-recruiters-share-10-ways-to-excel-in-your-aws-in-person-interview/)
  Official AWS careers blog. Use for: current tactical advice from AWS recruiters themselves.
- [Microsoft Interview Process & Loops Explained (2026)](https://copilotinterview.com/blog/microsoft-interview-process-loops) and [IGotAnOffer: Microsoft SDE interview](https://igotanoffer.com/blogs/tech/microsoft-software-development-engineer-interview)
  Third-party, not official (Microsoft doesn't publish a process page like Google/Amazon — see Gaps). Use for: general shape only — recruiter call → Codility OA (2 problems/90 min) → 4-5 round loop mixing coding/design/behavioral → "As Appropriate" culture-fit round.

### Post-LLM interview landscape (2024-2026) — treat as time-sensitive, re-verify closer to your actual interview date
- [Can You Use AI in a Coding Interview? Google Says Yes. Amazon Says You're Disqualified. (Medium, Jul 2026)](https://medium.com/@emilyhustlenyc/can-you-use-ai-in-a-coding-interview-google-says-yes-amazon-says-youre-disqualified-88e7dd5abda9)
  Third-party but directly sourced company contrast. Use for: headline framing of the Google-vs-Amazon policy split.
- [Which companies allow vs ban AI in interviews (2026) — Assistly](https://tryassistly.com/blog/companies-that-allow-or-ban-ai-in-interviews-2026) and [InterviewMan: same topic](https://interviewman.com/blog/companies-allow-ban-ai-interviews-2026)
  Two independent trackers, broadly agree (~64% of companies still ban AI in interviews). Use for: defaulting to "assume AI use is banned unless told otherwise" going into any loop.
- [Interview Cheating in 2026: Cluely, Interview Coder... (Fabric)](https://fabrichq.ai/blogs/interview-cheating-in-2026-the-rise-of-ai-tools-like-cluely-and-interview-coder) + [companion data post](https://fabrichq.ai/blogs/state-of-cheating-in-interviews-in-2026-tools-trends-and-prevention)
  Largest dataset found (19,368 interviews analyzed): 38.5% flagged for AI use, 48% in technical roles, 61% of flagged candidates still passed. Use for: the actual numbers behind "interviews have changed."
- [Google's AI-Assisted Coding Interview (2026 Guide)](https://customcareer.miami.edu/blog/2026/05/14/googles-ai-assisted-coding-interview-2026-guide/)
  Third-party, details Google's pilot letting some candidates use Gemini in a new "code comprehension" round (not yet universal — junior/mid roles). Use for: understanding this isn't hypothetical, it's already piloting.
- [Vinit Shahdeo: AI-assisted coding interviews in 2026](https://vinitshahdeo.substack.com/p/ai-assisted-coding-interviews-what-it-means-to-be-an-engineer)
  Practitioner take. Use for: why "explain and debug your own code live, unaided" has become the standard anti-cheat technique regardless of company policy — this is the actual skill to train for.

### DSA / algorithm practice (Go)
- [NeetCode](https://neetcode.io/)
  Free video explanations; NeetCode 150/250 lists are the current default community-recommended problem set. Pro tier (paid) adds Go-specific solutions/notes. Use for: primary problem-set spine and pattern grouping.
- [LeetCode](https://leetcode.com/)
  Accepts Go natively as a submission language; free tier is enough for practice. Use for: timed practice, company-tagged questions (Premium).
- No standalone high-trust book teaching algorithms specifically in Go was found — see Gaps.

### System design
- [ByteByteGo (Alex Xu)](https://bytebytego.com/)
  Paid, highly cited. Companion book *System Design Interview — An Insider's Guide* is a repeatedly-recommended primary text. Use for: visual/conceptual grounding on the standard building blocks.
- [Hello Interview — System Design in a Hurry](https://www.hellointerview.com/)
  Free core content, founded by ex-Meta/Amazon engineers. Use for: question-pattern breakdowns. Note: its mock-interview/mentorship marketplace was discontinued 2026-05-31 — don't route there for live mocks.

### Behavioral / Leadership Principles
- [Amazon Leadership Principles — official](https://www.amazon.jobs/en/principles)
  Canonical source, same link as above.
- [An Amazon interviewer dives deep into how she uses the Leadership Principles — About Amazon](https://www.aboutamazon.com/news/workplace/what-do-each-of-amazons-leadership-principles-really-mean)
  Official. Use for: how interviewers actually score LP answers, not just what the LPs say.
- [Amazon STAR Method guide — 4dayweek.io](https://4dayweek.io/interview-process/amazon-star)
  Third-party but solid. Use for: STAR-format walkthrough tuned to Amazon's evaluation style.

### Go language reference
- [A Tour of Go](https://go.dev/tour/)
  Official, interactive. Use for: syntax refresher if any Go idiom feels rusty.
- [Effective Go](https://go.dev/doc/effective_go)
  Official. Use for: idiomatic patterns (slices, error handling) so interview code reads as fluent Go, not translated Python/Java.

## Wisdom (Communities)

- [Blind — Interview Experiences channel](https://www.teamblind.com/channels/interview-experiences)
  Verified-employment-gated. Use for: highest-signal real loop reports and leveling/comp context. Weaker for beginner how-to content.
- r/cscareerquestions
  Use with caution — multiple sources flag it as lower signal-to-noise, generic-advice-heavy. Cross-check anything specific before trusting it.
- [interviewing.io](https://interviewing.io/)
  Anonymous mock interviews with real engineers. Use for: live feedback loop on timed problem-solving. Users report slots can be scarce — book ahead of when you actually need them.

## Gaps

- No strong *official* Microsoft interview-process page exists (unlike Google/Amazon) — Microsoft-specific process claims should be treated as lower-confidence, convergent-third-party-only.
- No standalone book/resource teaches algorithms natively in Go — practice plan is: learn the pattern language-agnostically (NeetCode), then implement in Go using go.dev/Effective Go for idiom.
