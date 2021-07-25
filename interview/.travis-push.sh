#!/bin/sh

#看云版本库地址:
KY_REPO=https://git.kancloud.cn/minibear/interview-leetcode.git
#github仓库地址：
GH_REPO=https://github.com/minibear2333/interview-leetcode.git

KY_REPO_URL=https://${KANCLOUD_USER}:${KANCLOUD_PASS}@$(echo $KY_REPO | awk -F'//' '{print $2}')
GH_REPO_URL=https://${GITHUB_TOKEN}@$(echo $GH_REPO | awk -F'//' '{print $2}')
KY_REPO_NAME=$(echo $KY_REPO | awk -F'/' '{print $NF}' | awk -F '.' '{print $1}')
DEST_REPO_URL=$GH_REPO_URL
SRC_REPO_URL=$KY_REPO_URL

setup_git() {
  git show -s --format=%ct
  git config --global user.email "pzqu@qq.com"
  git config --global user.name "pzqu"
  rm -rf *
  git clone --depth=50 --branch=master $SRC_REPO_URL
  repo_dir=$(ls) && cp -rf $repo_dir/* ./ &&  rm -rf $repo_dir
}

commit_country_json_files() {
  git status
  git checkout master
  # Current month and year, e.g: Apr 2018
  dateAndMonth= `date "+%b %Y"`
  # Stage the modified files in dist/output
  git add -A
  # Create a new commit with a custom build message
  # with "[skip ci]" to avoid a build loop
  # and Travis build number for reference
  git commit -m "Travis update: $dateAndMonth (Build $TRAVIS_BUILD_NUMBER)" -m "[skip ci]"
}

upload_files() {
  # Remove existing "origin"
  git remote rm origin
  # Add new "origin" with access token in the git URL for authentication
  git remote add origin $DEST_REPO_URL > /dev/null 2>&1
  git push origin master --quiet
}

compare_new() {
   github_last_commit_time=$(git show -s --format=%ct)
   git clone $KY_REPO_URL && cd $KY_REPO_NAME
   ky_last_commit_time=$(git show -s --format=%ct)
   if [ $github_last_commit_time -gt $ky_last_commit_time ];then
       SRC_REPO_URL=$GH_REPO_URL
       DEST_REPO_URL=$KY_REPO_URL
   else
       cd ..
   fi
   echo "sync $SRC_REPO_URL to $DEST_REPO_URL"
}

compare_new

setup_git

commit_country_json_files

# Attempt to commit to git only if "git commit" succeeded
if [ $? -eq 0 ]; then
  echo "A new commit with changed country JSON files exists. Uploading to GitHub"
  upload_files
else
  echo "No changes in country JSON files. Nothing to do"
fi
