// Hardware provider abstraction placeholder.
//
// The key design idea:
// expose one stable provider interface to the rest of the app, then hide
// Lenovo, generic Linux, ASUS, and HP details behind concrete providers.
//
// Suggested provider shape:
// - Identity: brand, model, serial.
// - System: OS, kernel, CPU model, RAM.
// - Battery: charge, state, watts, design capacity, full capacity, time estimate.
// - Thermals: CPU package temp, per-core temps, GPU temps when available.
// - Fans: RPMs and current mode if available.
// - CPU: per-core usage, clocks, governor.
// - Capabilities: which write controls are supported.
//
// Reads should be safe for normal users.
// Write controls should be separated so the TUI can clearly mark privileged actions.
