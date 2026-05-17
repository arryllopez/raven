// Lenovo provider placeholder.
//
// v1 should treat Lenovo and ThinkPad as first-class targets.
//
// Likely read sources:
// - /sys/class/dmi/id/product_name for model.
// - /sys/class/dmi/id/product_serial for serial.
// - /sys/class/power_supply for battery data.
// - /sys/class/thermal and /sys/class/hwmon for temperatures and fans.
// - /proc/cpuinfo and /sys/devices/system/cpu for CPU details.
//
// Likely control sources:
// - thinkpad_acpi for fan mode and ThinkPad-specific controls.
// - lenovo_ideapad_extras where relevant.
//
// Keep every sysfs path behind small reader helpers so tests can point at
// fixture directories later.
