#!/bin/bash
set -o nounset
set -e

make contracts

sh -x
set +e
git diff --exit-code --quiet
git_diff_rc=$?
git status
set -e
if [ ${git_diff_rc} -eq 1 ]
then
    git add go/*
    LAST_SHA=$(git rev-parse --short HEAD)
    LAST_COMMIT_MESSAGE=$(git log -1 --pretty=%B)
    git commit -m "Updating contract ABIS for [${LAST_SHA}] - ${LAST_COMMIT_MESSAGE}"
    git push
else
    echo "There are no changes to commit"
fi