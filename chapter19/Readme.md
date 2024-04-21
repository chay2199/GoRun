### CSP, Goroutines and Channels

A channel is like a one-way socket or a Unix pipe but it allows multiple readers and writers.

It's also a vehicle for transferring ownership of data, so that only one goroutine at a time is writing the data (avoid race conditions)

You can only close the channel at once.