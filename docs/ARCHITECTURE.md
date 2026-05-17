<!--
Architecture notes placeholder.

Core idea:
Raven should have a stable domain model in the middle, with hardware providers
below it and the TUI above it.

Proposed layers:

cmd/raven
TUI binary. Bubble Tea app, rendering, keybindings, socket client.

cmd/ravend
Daemon binary. Sampling loop, SQLite writes, Unix socket server.

internal/hardware
Provider interface, normalized snapshots, discovery.

internal/hardware/lenovo
Lenovo and ThinkPad implementation. First-class v1 target.

internal/hardware/generic
Best-effort Linux fallback for common /sys and /proc data.

internal/metrics
Sampling and trend analysis.

internal/storage
SQLite persistence.

internal/socket
Daemon to TUI protocol.

internal/control
Write operations and privilege boundaries.

Why this abstraction matters:
The TUI should not know whether battery health came from Lenovo sysfs paths,
generic power_supply data, or a future ASUS/HP provider. It should only receive
a normalized BatterySnapshot and a list of supported capabilities.
-->
