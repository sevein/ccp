set shell := ["bash", "-uc"]

[private]
default:
  @just --list --unsorted

e2e-dump:
  dagger call --progress=plain --source=".:default" generate-dumps export --path=e2e/testdata/dumps

e2e:
  dagger call --progress=plain --source=".:default" etoe

amflow:
  amflow edit --file ./internal/workflow/assets/workflow.json

grpcui:
  grpcui -plaintext -H "Authorization: ApiKey test:test" localhost:63030

run:
  make -C hack run

transfer:
  ./hack/helpers/transfer-via-api.sh

# Tag and release new version.
release:
    #!/usr/bin/env bash
    set -euo pipefail
    branch=qa/2.x
    git checkout ${branch} > /dev/null 2>&1
    git diff-index --quiet HEAD || (echo "Git directory is dirty" && exit 1)
    version=v$(semver bump prerelease beta.. $(git describe --abbrev=0))
    echo "Detected version: ${version}"
    read -n 1 -p "Is that correct (y/N)? " answer
    echo
    case ${answer:0:1} in
        y|Y )
            echo "Tagging release with version ${version}"
        ;;
        * )
            echo "Aborting"
            exit 1
        ;;
    esac
    git tag -m "Release ${version}" $version
    git push origin refs/tags/$version

git-log-recent-upstream:  # Show recent commits in upstream (qa/1.x).
  #!/usr/bin/env bash
  if ! git remote get-url upstream > /dev/null 2>&1; then
      git remote add -f upstream https://github.com/artefactual/archivematica.git
  else
      git fetch upstream
  fi
  git log --oneline upstream/qa/1.x ^HEAD
