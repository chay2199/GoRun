### Select

select allows any "Ready" alternative to proceed among:
- a channel we can read from
- a channel we can write to
- a default action that's always ready

Most often select rus a loop so we keep trying

We can put a timeout or "done" channel into the select
