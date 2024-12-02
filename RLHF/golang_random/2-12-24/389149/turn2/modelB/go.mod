module test

go 1.23.3

require (
	github.com/example/package1 v1.2.3 // Pinned version
	github.com/example/package2 v2.3.4 // Pinned version
)

replace (
	github.com/example/package1 v1.2.3 => github.com/example/package1 v1.3.0 // Compatibility replacement
)