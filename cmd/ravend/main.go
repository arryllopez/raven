// Daemon entrypoint placeholder.
//
// Intended responsibility:
// - Run the background metric capture loop.
// - Own writes to the local SQLite database.
// - Expose a Unix socket API for the TUI.
// - Avoid terminal UI dependencies.
//
// Expected flow:
// 1. Load config.
// 2. Detect the best hardware provider.
// 3. Open the metrics database.
// 4. Start the sampler at the configured interval, defaulting to 5 seconds.
// 5. Serve current snapshots and trend queries over the Unix socket.
