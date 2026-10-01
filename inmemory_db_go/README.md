# Practice: In-Memory Database (CodeSignal ICA style)

**Format:** 90 minutes, 4 levels, one codebase. Each level unlocks only after the
previous level passes. Level 4 needs snapshots of the whole database, so think
about how you store TTLs before you get there.

Implement `InMemoryDB` in `in_memory_db.go`. Run tests one level at a time:

```bash
cd inmemory_db_go
go test -run Level1 -v
go test -run Level2 -v
go test -run Level3 -v
go test -run Level4 -v
```

**General rules**
- The database holds **records**. Each record has a `key` and holds any number
  of `field -> value` pairs. All keys, fields and values are strings.
- Every method takes a `timestamp` (milliseconds). Timestamps across calls are
  strictly increasing.
- A `nil` return means "not found". Never panic.

---

## Level 1: Basic operations (target: ~10 min)

- `Set(timestamp int, key, field, value string)`
  Sets `field` in record `key` to `value`, creating the record if needed.
  Overwrites an existing value.
- `Get(timestamp int, key, field string) *string`
  Returns the value, or `nil` if the record or field doesn't exist.
- `Delete(timestamp int, key, field string) bool`
  Removes the field. Returns `true` if it existed, `false` otherwise.

## Level 2: Scanning (target: ~15 min)

- `Scan(timestamp int, key string) []string`
  Returns every field of the record as `"field(value)"`, sorted by field name
  (plain string order, so `"B10"` comes before `"B2"`). Returns an empty slice
  (not `nil`) if the record doesn't exist or has no fields.
- `ScanByPrefix(timestamp int, key, prefix string) []string`
  Same as `Scan`, but only fields whose name starts with `prefix`.

## Level 3: TTL (target: ~25 min)

- `SetWithTTL(timestamp int, key, field, value string, ttl int)`
  Like `Set`, but the field only exists during `[timestamp, timestamp + ttl)`.
  At time `timestamp + ttl` it's gone.
- An expired field behaves exactly as if it was deleted: `Get` returns `nil`,
  `Delete` returns `false`, and scans skip it.
- A plain `Set` on a field removes any TTL it had (the field no longer expires).
  A `SetWithTTL` on a field replaces its old TTL.

## Level 4: Backup and restore (target: ~30 min)

- `Backup(timestamp int) int`
  Saves a snapshot of the database. For each field with a TTL, the snapshot
  stores its **remaining** time to live (`expiry - timestamp`), not its expiry.
  Returns how many records have at least one live field.
- `Restore(timestamp int, timestampToRestore int)`
  Replaces the whole database with the latest backup taken at or before
  `timestampToRestore`. Each restored TTL field gets a new expiry:
  `timestamp + remaining TTL`. A backup always exists when this is called.
- Later changes to the database must not change a saved backup, and a backup
  can be restored more than once.
