// Generic Linux provider placeholder.
//
// This provider should cover best-effort reads that work across many laptops:
// - Battery data from /sys/class/power_supply.
// - Thermal zones and hwmon sensors.
// - CPU usage from /proc/stat.
// - CPU model from /proc/cpuinfo.
// - Governor and clocks from /sys/devices/system/cpu.
//
// It should not assume fan controls, battery thresholds, or vendor-specific
// performance modes are writable.
