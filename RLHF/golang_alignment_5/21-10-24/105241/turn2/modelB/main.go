name: CI with Code Review

on:
  push:
    branches:
      - main
  pull_request:
    branches:
      - main

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout code
        uses: actions/checkout@v2

      - name: Set up Go
        uses: actions/setup-go@v2
        with:
          go-version: 1.16

      - name: Install dependencies
        run: go install -v ./...

      - name: Run tests
        run: |
          go test -v ./...
          if [[ $? -ne 0 ]]; then
            exit 1
          fi

      - name: Perform code review (Example: Using 'golint')
        run: |
          golint .
          if [[ $? -ne 0 ]]; then
            exit 1
          fi

      - name: Check for documentation
        run: |
          find . -name "*.go" -exec grep -q "//" {} \; && echo "No documentation found." || true