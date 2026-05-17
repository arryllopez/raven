// TUI entrypoint placeholder.
//
// Intended responsibility:
// - Start the Bubble Tea program.
// - Connect to the daemon over the Unix socket.
// - Fall back to direct one-shot reads only if you explicitly choose that mode.
// - Keep terminal UI concerns out of hardware, storage, and daemon packages.
//
// Expected flow:
// 1. Load config.
// 2. Dial the daemon socket.
// 3. Request the latest dashboard snapshot.
// 4. Subscribe to periodic metric updates.
// 5. Render panels for system, battery, thermals, fans, CPU, and trends.
