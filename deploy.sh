#/bin/bash
echo -e "\033[0;32mDeploying updates to GitHub...\033[0m"
git add -A
git commit -m "$1"
git push origin master

# Build the project.
go run repo.go
cp tmpl/index.html ../awesome-cloud-native-gh-pages

# Go To github pages folder
cd ../awesome-cloud-native-gh-pages

# Add changes to git.
git add -A

# Commit changes.

git commit -m "$1"

# Push source and build repos.
git push origin gh-pages

# Come Back
cd ../awesome-cloud-native
rm tmpl/index.html
