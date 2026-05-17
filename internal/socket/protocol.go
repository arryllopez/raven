// Unix socket protocol placeholder.
//
// The daemon owns the socket.
// The TUI connects as a client.
//
// Suggested requests:
// - Get latest snapshot.
// - Subscribe to snapshot updates.
// - Query trend series.
// - Request supported capabilities.
// - Request a privileged control action.
//
// Keep messages versioned so the daemon and TUI can evolve independently.
