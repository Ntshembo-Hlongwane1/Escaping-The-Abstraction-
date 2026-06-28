# Escape the Abstraction

> Rebuilding the technology we all rely on — from scratch, live.

This is not a tutorial. I'm not here to teach you.

I'm a developer who got really good at *using* technology and realized I never understood how any of it actually works. I can ship apps. I can scale a service. But ask me how the database underneath actually finds a row, how a search engine ranks results, how a cache stays consistent across machines — and I'd be guessing.

So I'm doing something about it. I'm rebuilding the systems we all stand on, one at a time, from zero — and I'm doing it **live**. Mistakes, dead ends, "wait, why doesn't this work" moments included.

This repo is the trail I'm leaving behind as I go.

If you've ever felt like you're standing on top of abstractions you couldn't rebuild if your life depended on it — follow along. We figure it out together.

---

## Why this exists

Modern development is incredibly productive *because* someone else already solved the hard problems. Frameworks think for us. Libraries hide the complexity. The cloud abstracts the infrastructure. That's not a bad thing — it's why we ship.

But it also means most of us never had to understand *why* the hard parts are hard. Data structures feel like interview trivia. Algorithms feel academic. Consensus, parsing, storage engines — they feel like someone else's job.

They stop feeling that way the moment you try to build the abstraction yourself.

A hashmap stops being a quiz question and becomes the engine of your key-value store. A tree stops being homework and becomes how your database finds things fast. Every concept suddenly has a reason to exist.

That's the whole point of this. Not to replace these tools. Not to compete with them. Just to understand them deeply enough to know how they work, what trade-offs they make, and how the people who built them think.

---

## How this works

- Each project gets built **live**, in order.
- Each one is broken into versions — small steps where each new version introduces a new problem, and each problem forces a new data structure or algorithm.
- I don't skip steps. We start at the simplest possible thing and grow it until it's real.
- The code here is the *learning artifact*, not production software. It's meant to be understood, not deployed.
- When I get something wrong, the history stays. That's part of the point.

---

## The Roadmap — 12 Projects

We go from "find a word in some text files" all the way to "build a programming language." Each project stands on the lessons of the ones before it.

### 1. Search Engine
Turn a pile of documents into something searchable. Indexing, tokenization, ranking, fuzzy matching, autocomplete, and concurrent indexing. The lesson underneath it all: search is an *indexing* problem, not a *scanning* problem.
`HashMaps · Sets · Heaps · Tries · Queues`

### 2. Redis
An in-memory key-value store. Fast reads and writes, data structures as first-class citizens, expiry, persistence. Why is it so fast, and what does it give up to get there?
`Hash tables · Linked structures · Event loops`

### 3. Git
Version control from scratch. Content-addressable storage, commits, trees, branches, diffs. Understand what's *actually* happening every time you type `git commit`.
`Hashing · DAGs · Trees · Compression`

### 4. Shell & Process Manager
Run and manage programs. Parsing commands, spawning processes, pipes, signals, job control. The layer between your keyboard and the operating system.
`Parsing · Process model · IPC`

### 5. HTTP Server & Load Balancer
Serve requests and spread traffic. Sockets, the HTTP protocol, connection handling, routing, balancing strategies. How the web actually moves bytes around.
`Networking · Concurrency · Queues`

### 6. SQL Database Engine
A real database. Tables, a query parser, an execution engine, indexes, storage on disk. The thing every CRUD app quietly depends on.
`B-Trees · Parsing · Storage engines`

### 7. Route Optimization Engine
Find the best path. Graphs, shortest-path algorithms, heuristics, optimization under constraints. The math behind maps, logistics, and routing.
`Graphs · Dijkstra · A* · Priority queues`

### 8. Kafka
Event streaming and durable logs. Append-only logs, partitions, producers and consumers, ordering and delivery guarantees. How systems talk to each other at scale.
`Logs · Partitioning · Offsets`

### 9. Distributed Cache
Caching across many machines. Consistent hashing, replication, eviction, invalidation. What changes when "the cache" isn't on one box anymore.
`Consistent hashing · Replication · Eviction`

### 10. Raft Consensus Engine
Agreement under failure. Leader election, log replication, how a cluster keeps a single source of truth when machines crash. The hard heart of distributed systems.
`Consensus · State machines · Replication`

### 11. NoSQL DB Engine
Documents at scale. Flexible schemas, different storage models, partitioning, querying without joins. The other side of the database world.
`LSM trees · Partitioning · Indexing`

### 12. Programming Language & Bytecode VM
The final boss. A language of my own — lexer, parser, compiler, and a bytecode virtual machine to run it. Everything underneath everything.
`Lexing · Parsing · Compilation · VM design`

---

## Then: from infrastructure to a real system

Once the stack exists, the journey doesn't stop at building it. The plan is to take these systems and turn them into something real:

- Assemble the projects into the **actual infrastructure of a working app** — a frontend and backend running on top of the search engine, the databases, the cache, the message broker I built.
- **Simulate live production traffic** and push it until it breaks.
- **Red team** it — attack my own stack, find the cracks, break it.
- **Blue team** it — detect, patch, harden.

Build it → run it → break it → defend it. The full lifecycle of real software, with no abstractions left to hide behind.

---

## Repo structure

Each project lives in its own directory, and each is built up in versions so you can follow the progression step by step.

```
escape-the-abstraction/
├── 01-search-engine/
├── 02-redis/
├── 03-git/
├── 04-shell-process-manager/
├── 05-http-server-load-balancer/
├── 06-sql-database-engine/
├── 07-route-optimization-engine/
├── 08-kafka/
├── 09-distributed-cache/
├── 10-raft-consensus-engine/
├── 11-nosql-db-engine/
└── 12-language-and-vm/
```

Each project folder has its own README with the versions, the goals, and notes on what I learned (and what tripped me up).

---

## A note on the code

This is learning-in-public, so expect:

- Things that aren't optimized, because clarity came first.
- Comments that are really notes-to-self.
- Commits that show the wrong turn before the right one.
- The occasional "I genuinely didn't understand this until I built it" moment.

That's not a disclaimer — that's the whole project.

---

## Follow along

The builds happen **live**. This repo is where the code lands afterward.

If understanding what's *underneath* matters to you as much as it does to me — come build it with me. We start at the bottom and we don't skip a step.

**Escape the Abstraction.**