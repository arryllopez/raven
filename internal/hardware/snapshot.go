// Normalized hardware snapshot placeholder.
//
// Define shared structs here once implementation begins.
// The daemon stores and serves these normalized snapshots instead of exposing
// raw /sys, /proc, hwmon, or vendor-specific paths to the TUI.
//
// This prevents Bubble Tea views and SQLite schemas from depending on one
// laptop family.
