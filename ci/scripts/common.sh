#!/usr/bin/env bash

function abort() {
  echo $'\e[31m'"$@"$'\e[0m' >&2
  exit 1
}

ran_progress="false"

function progress() {
  if [ "$ran_progress" = "true" ]; then
    echo ""
  fi

  ran_progress="true"

  echo $'\e[1m'"$@"$'\e[0m'
}

function list-pr-branches() {
  git branch -r | grep pr/  | sed -e 's|\s*origin/||'
}

function create-branch-track-upstream-release() {
  git fetch upstream tag $1 --no-tags
  git checkout -b $1 $1
}

set +x

mkdir ~/.ssh
ssh-keyscan -H github.com >> ~/.ssh/known_hosts
echo "$CONCOURSE_DEX_DEPLOY_KEY" > ~/.ssh/id_rsa
chmod 400 ~/.ssh/id_rsa

git config --global user.email "ci@localhost"
git config --global user.name "CI Bot"

if ! git remote | grep upstream >/dev/null; then
  git remote add upstream https://github.com/dexidp/dex
fi

git config remote.origin.fetch "+refs/heads/*:refs/remotes/origin/*"
git remote update origin --prune >/dev/null

set -x
