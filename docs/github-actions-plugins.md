# GitHub Actions - Plugin Build and Release

## Overview

The repository is configured with GitHub Actions workflows that automatically build and push Docker images for all plugins to GitHub Container Registry (ghcr.io).

## Workflows

### 1. Pull Request Workflow (`plugins-source-pullrequest.yaml`)

**Triggers:** On pull requests to `main` branch

**What it does:**
- Runs tests for all plugins in the workspace
- Detects which plugins have changed
- Builds Docker images (without pushing) for changed plugins
- Automatically discovers all `*-plugin` directories

### 2. Tag/Release Workflow (`plugins-source-tag.yaml`)

**Triggers:** On git tags matching pattern: `*-plugin/[0-9]+.[0-9]+.[0-9]+`

**What it does:**
- Runs tests for all plugins
- Parses the plugin name and version from the tag
- Builds and pushes multi-arch Docker images (linux/amd64, linux/arm64)
- Tags the image as: `ghcr.io/arubacloud/arubacloud-provider-kog/{plugin-name}:{version}`

## How to Release a Plugin

### Step 1: Commit and Push Your Changes

```bash
git add .
git commit -m "feat: add compute plugin with cloudserver and keypair resources"
git push origin main
```

### Step 2: Create and Push a Tag

To release a specific plugin, create a tag with the format: `{plugin-name}/{version}`

**Examples:**

Release compute-plugin version 1.0.0:
```bash
git tag compute-plugin/1.0.0
git push origin compute-plugin/1.0.0
```

Release database-plugin version 1.0.0:
```bash
git tag database-plugin/1.0.0
git push origin database-plugin/1.0.0
```

Release all plugins at once (bash script):
```bash
VERSION="1.0.0"
for plugin in compute container database project schedule storage security network; do
  git tag ${plugin}-plugin/${VERSION}
  git push origin ${plugin}-plugin/${VERSION}
done
```

Or in PowerShell:
```powershell
$version = "1.0.0"
@('compute','container','database','project','schedule','storage','security','network') | ForEach-Object {
  git tag "$_-plugin/$version"
  git push origin "$_-plugin/$version"
}
```

### Step 3: Monitor the Workflow

1. Go to: https://github.com/Arubacloud/arubacloud-provider-kog/actions
2. Watch the `plugins-source-tag` workflow run
3. Check that the image is pushed to: https://github.com/orgs/Arubacloud/packages

### Step 4: Use the Images in Blueprints

After the images are pushed, update your blueprints:

```powershell
.\update-blueprint-images.ps1 -Version 1.0.0
```

This will update all blueprint deployments to use the newly published images.

## Plugin List

The following plugins are configured for automatic building:

- ✅ compute-plugin
- ✅ container-plugin
- ✅ database-plugin
- ✅ project-plugin
- ✅ schedule-plugin
- ✅ storage-plugin
- ✅ security-plugin
- ✅ network-plugin
- ✅ subnet-plugin (existing)

## Image Naming Convention

Images are published with the following naming pattern:

```
ghcr.io/arubacloud/arubacloud-provider-kog/{plugin-name}:{version}
```

Examples:
- `ghcr.io/arubacloud/arubacloud-provider-kog/compute-plugin:1.0.0`
- `ghcr.io/arubacloud/arubacloud-provider-kog/database-plugin:1.0.0`
- `ghcr.io/arubacloud/arubacloud-provider-kog/network-plugin:1.0.0`

## Troubleshooting

### Workflow fails on test step

Check that all plugins compile:
```bash
cd plugins
go test ./...
```

### Image not appearing in registry

1. Check workflow logs in GitHub Actions
2. Verify GITHUB_TOKEN has package write permissions
3. Ensure tag format is correct: `{plugin-name}/{semver}`

### Need to rebuild without new tag

Delete and recreate the tag:
```bash
git tag -d compute-plugin/1.0.0
git push origin :refs/tags/compute-plugin/1.0.0
git tag compute-plugin/1.0.0
git push origin compute-plugin/1.0.0
```

## Best Practices

1. **Use semantic versioning**: Follow semver (e.g., 1.0.0, 1.1.0, 2.0.0)
2. **Test locally first**: Build and test images locally before tagging
3. **Tag all plugins together**: When making workspace-wide changes, tag all plugins with the same version
4. **Update blueprints**: Always update blueprint image references after releasing new plugin versions
5. **Document changes**: Update CHANGELOG.md or release notes for each version
