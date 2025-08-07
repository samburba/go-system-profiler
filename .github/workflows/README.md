# GitHub Workflows

This directory contains GitHub Actions workflows for automated testing, building, and releasing the go-system-profiler library.

## Workflows

### 1. Main Branch CI/CD (`main.yml`)

**Triggers:**
- Push to `main` branch
- Pull requests to `main` branch

**What it does:**
- Runs tests on macOS latest with Go 1.23
- Builds the project
- Runs the audio example
- **On push to main only:** Creates a new versioned release

**Versioning:**
- Automatically increments the patch version (e.g., v1.2.3 → v1.2.4)
- Creates a GitHub release with the new version
- Includes release notes with features and breaking changes

### 2. Nightly Release (`nightly.yml`)

**Triggers:**
- Daily at midnight UTC (cron: '0 0 * * *')
- Manual trigger via workflow_dispatch

**What it does:**
- Tests on multiple macOS versions:
  - macOS Latest (Go 1.21, 1.22, 1.23)
  - macOS 13 (Go 1.22, 1.23)
  - macOS 12 (Go 1.22, 1.23)
- Creates a nightly release with suffix `-nightly.YYYYMMDD`

**Versioning:**
- Format: `v0.0.0-nightly.20241201`
- Marked as pre-release
- Includes test matrix information

### 3. Pull Request Checks (`pr.yml`)

**Triggers:**
- Pull requests to `main` branch

**What it does:**
- Runs tests on macOS latest
- Checks code formatting with `gofmt`
- Runs `go vet` for static analysis
- Checks for race conditions with `-race` flag
- Builds the project and runs examples

## Versioning Strategy

### Stable Releases (main.yml)
- **Format:** `vX.Y.Z` (semantic versioning)
- **Increment:** Patch version automatically
- **Tag:** Latest stable release
- **Release:** Full GitHub release with notes

### Nightly Releases (nightly.yml)
- **Format:** `v0.0.0-nightly.YYYYMMDD`
- **Purpose:** Testing and development
- **Tag:** Pre-release
- **Release:** Marked as pre-release

## Usage

### For Contributors
1. Create a pull request to `main`
2. The PR workflow will automatically run tests and checks
3. Ensure all checks pass before merging

### For Maintainers
1. Push to `main` branch
2. The main workflow will automatically:
   - Run tests
   - Create a new versioned release
   - Tag the release

### Manual Nightly Release
1. Go to Actions tab in GitHub
2. Select "Nightly Release" workflow
3. Click "Run workflow" button
4. The workflow will run tests on all macOS versions and create a nightly release

## Requirements

- **macOS:** All workflows run on macOS runners (required for `system_profiler`)
- **Go:** Multiple Go versions tested (1.21, 1.22, 1.23)
- **Permissions:** Workflows require `GITHUB_TOKEN` for creating releases

## Notes

- All workflows use the latest GitHub Actions versions
- Go module caching is enabled for faster builds
- Tests are run with verbose output (`-v` flag)
- Race condition detection is enabled for PR checks
- Code formatting is enforced via `gofmt`
