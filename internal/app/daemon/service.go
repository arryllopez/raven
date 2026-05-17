// Daemon service placeholder.
//
// This package should coordinate the daemon runtime:
// - Hardware provider selection.
// - Metric sampling schedule.
// - Database writes.
// - Socket server lifecycle.
// - Shutdown handling.
//
// Avoid placing device-specific logic here.
// Device-specific reads belong under internal/hardware.
