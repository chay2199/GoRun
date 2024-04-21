### Concurrency

Race conditions involve independent parts of the program changing things that are shared.

Solutions:
- don't share
- make the shared things read-only
- allow only one writer to shared things
- make the read-modify-write operations atomic