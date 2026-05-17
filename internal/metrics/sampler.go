// Metric sampler placeholder.
//
// The sampler should:
// - Run at a configurable interval, defaulting to 5 seconds.
// - Ask the active hardware provider for a normalized snapshot.
// - Send snapshots to storage.
// - Keep the latest snapshot in memory for fast socket responses.
//
// Sampling should continue when one field fails.
// Prefer partial snapshots with field-level errors over dropping the entire tick.
