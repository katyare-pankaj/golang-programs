module test

go 1.23.3



require (
        example.com/dep1 v1.0.0 // Platform-independent dependency
        example.com/dep2 v2.0.0 // Platform-independent dependency
        example.com/dep3 v3.0.0 // Platform-specific dependency for linux
        example.com/dep4 v4.0.0 // Platform-specific dependency for windows
)

// Replace directives for platform-specific dependencies
//go:build linux
// +build linux
replace example.com/dep3 => example.com/dep3/linux

//go:build windows
// +build windows
replace example.com/dep4 => example.com/dep4/windows