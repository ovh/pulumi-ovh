# Upgrading/Sync with existing Terraform provider

* install gh CLI
* Install upgrade-provider CLI
* create and retrieve your GitHub Personal access token (PAT)
* export it:

```bash
$ export GITHUB_TOKEN="<your-gh-pat>"
```

* execute the upgrade-provider CLI:

If the provider is already up to date:

```bash
$ upgrade-provider scraly/pulumi-ovh
---- Setting Up Environment ----
- ✓ GOWORK="off": done
- ✓ PULUMI_MISSING_DOCS_ERROR="false": done
- ✓ PULUMI_CONVERT_EXAMPLES_CACHE_DIR="": done
---- Discovering Repository ----
- Ensure 'github.com/scraly/pulumi-ovh'
  - ✓ Expected Location: /workspace/pulumi-ovh
  - ✓ Downloading: skipped - already exists
  - ✓ Validating: done
- pull default branch
  - ✓ /usr/bin/git ls-remote --heads origin: done
  - ✓ finding default branch: main
  - ✓ /usr/bin/git fetch: done
  - ✓ /usr/bin/git checkout main: done
  - ✓ /usr/bin/git pull origin: done
- ✓ Repo kind: plain
- X Planning Provider Update: current upstream version 0.32.0 is greater than/ equal to the target version 0.32.0
```

In this case, OK nothing to do :)

If the provider need to be upgraded:
```bash
$ upgrade-provider scraly/pulumi-ovh
---- Setting Up Environment ----
- ✓ GOWORK="off": done
- ✓ PULUMI_MISSING_DOCS_ERROR="false": done
- ✓ PULUMI_CONVERT_EXAMPLES_CACHE_DIR="": done
---- Discovering Repository ----
- Ensure 'github.com/scraly/pulumi-ovh'
  - ✓ Expected Location: /workspaces/pulumi-ovh
  - ✓ Downloading: skipped - already exists
  - ✓ Validating: done
- pull default branch
  - ✓ /usr/bin/git ls-remote --heads origin: done
  - ✓ finding default branch: main
  - ✓ /usr/bin/git fetch: done
  - ✓ /usr/bin/git checkout main: done
  - ✓ /usr/bin/git pull origin: done
- ✓ Repo kind: plain
- ✓ Planning Provider Update: 0.32.0 -> 0.33.0
- ✓ Planning Bridge Update: v3.53.0 -> v3.57.0
- ✓ Planning Plugin SDK Update: v2.26.1 -> 2.27.0
- ✓ Planning Plugin Framework Update:  -> 0.15.2
---- Update Artifacts ----
- Ensure Branch
  - ✓ /usr/bin/git branch: done
  - ✓ Already exists: no
  - ✓ /usr/bin/git checkout -b upgrade-terraform-provider-ovh-to-v0.33.0: done
  - ✓ /usr/bin/git checkout upgrade-terraform-provider-ovh-to-v0.33.0: done
- Update TF Provider
  - ✓ Update TF Plugin SDK Fork: updated
  - ✓ /usr/local/go/bin/go mod tidy: done
  - ✓ Lookup Tag SHA: 6f0cc4b7d791d1682b613e2c2cbd9a48d4267686
  - ✓ /usr/local/go/bin/go get github.com/ovh/terraform-provider-ovh@6f0cc4b7d791d1682...: done
- ✓ /usr/local/go/bin/go get github.com/pulumi/pulumi-terraform-bridge/v3@v3.57.0: done
- ✓ /usr/local/go/bin/go get github.com/pulumi/pulumi-terraform-bridge/pf@v0.15.2: done
- ✓ /usr/local/go/bin/go mod tidy: done
- ✓ /usr/local/go/bin/go mod tidy: done
- ✓ /bin/echo Plugins not removed.: done
- X /usr/bin/make tfgen: exit status 2:
[resource plugin random-4.8.2] installing
warning: A new version of Pulumi is available. To upgrade from version '3.78.1' to '3.79.0', run 
   $ curl -sSL https://get.pulumi.com | sh
or visit https://pulumi.com/docs/install/ for manual instructions and release notes.
...
- Update TF Provider
  - ✓ Update TF Plugin SDK Fork: already up to date
  - ✓ Lookup Tag SHA: 6f0cc4b7d791d1682b613e2c2cbd9a48d4267686
  - ✓ /usr/local/go/bin/go get github.com/ovh/terraform-provider-ovh@6f0cc4b7d791d1682...: done
- ✓ /usr/local/go/bin/go get github.com/pulumi/pulumi-terraform-bridge/v3@v3.57.0: done
- ✓ /usr/local/go/bin/go get github.com/pulumi/pulumi-terraform-bridge/pf@v0.15.2: done
- ✓ /usr/local/go/bin/go mod tidy: done
- ✓ /usr/local/go/bin/go mod tidy: done
- ✓ /bin/echo Plugins not removed.: done
- ✓ /usr/bin/make tfgen: done
- ✓ /usr/bin/git add --all: done
- ✓ /usr/bin/git commit -m make tfgen: done
- ✓ /usr/bin/make build_sdks: done
- ✓ /usr/bin/git add --all: done
- ✓ /usr/bin/git commit -m make build_sdks: done
- GitHub
  - ✓ /usr/bin/git push --set-upstream origin upgrade-terraform-provider-ovh-to-v0.33....: done
  - ✓ /usr/bin/gh pr create --assignee @me --base main --head upgrade-terraform-provid...: done
  - Assign Issues
```

Add new TF resources in `provider/resources.go` file and execute the command again and again.
A new Pull Request will be created at this step.