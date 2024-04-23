### Context

The context package offers a common method to cancel requests
- explicit cancellation
- implicit cancellation based on a timeout or deadline


A context offers two controls:
- a channel that closes when the cancellation occurs
- an error that's readable once the channel closes

Cancellation or timeout applies to current context and its subtree.

Contextt values should be data specific to a request, such as:
- a trace ID or start time
- security or authorization data

A lit bit is okay but a lot is not!