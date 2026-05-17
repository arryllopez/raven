// Privileged controls placeholder.
//
// Keep write operations separate from read-only hardware collection.
//
// MVP controls:
// - Battery charge threshold.
// - Fan mode toggle.
// - CPU governor toggle.
// - TDP control when supported.
//
// Design note:
// the TUI should request a control operation, but this package should decide
// whether the operation is supported and what privilege boundary is required.
