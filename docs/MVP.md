<!--
Raven MVP scope placeholder.

Ship first:
Core dashboard with system info, battery, thermals, fans, and CPU panels.

Next:
Daemon that samples every 5 seconds, stores snapshots in SQLite, and exposes a
Unix socket for the TUI.

Differentiator:
Trend analysis with Bubble Tea sparklines for battery discharge, CPU
temperature, fan speed versus temperature, and battery health over time.

Controls:
Privileged write operations for battery charge threshold, fan mode, CPU
governor, and TDP where supported.

Warranty:
For v1, detect Lenovo serial number from /sys/class/dmi/id/product_serial and
open the Lenovo warranty lookup URL from the TUI.

Hardware support:
v1 is Lenovo first, with a generic Linux read-only fallback. ASUS and HP should
be added as new provider packages later, without changing TUI panel code.
-->
