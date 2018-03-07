#!/bin/bash
#Script creating a svn repository and committing all the base actions.

# make bash stop and fail if any command fails
set -e

#Defining repository name
GITREPO=git-repository_test

#Initialisation of the repository
mkdir $GITREPO
cd $GITREPO
git init
touch README.md
git add README.md
git commit -m 'Initialisating testing Git repository'

#Commit of the action 'add'
touch git-file_test.go
cat << FIN >> git-file_test.go
git text : test
FIN

touch git-file_test-delete.go
cat << FIN >> git-file_test-delete.go
git text : test
FIN

touch git-file_test-modify.go
cat << FIN >> git-file_test-modify.go
git text : test modify
FIN

touch git-file_test-rename.go
cat << FIN >> git-file_test-rename.go
git text : test rename
FIN

touch git-file_test-move.go
cat << FIN >> git-file_test-move.go
git text : test move
FIN
mkdir git-dir_test-move

touch git-file_test-rename-and-move.go
cat << FIN >> git-file_test-rename-and-move.go
git text : test rename and move
FIN
mkdir git-dir_test-rename-and-move

touch git-file_test-delete-and-restore.go
cat << FIN >> git-file_test-delete-and-restore.go
git text : test delete and restore
FIN

git add *
git commit -m 'Adding files'

#Commit of the action 'delete'
rm ./git-file_test-delete.go
git add -A
git commit -m 'Deleting file'

#Commit of the action 'modify'
cat << FIN >> git-file_test-modify.go
git text : modifying tested
FIN
git add -A
git commit -m 'Modifying file'

#commit of the action 'renaming'
mv ./git-file_test-rename.go ./git-file_renaming-tested.go
git add -A
git commit -m 'Renaming file'

#commit of the action 'moving'
mv ./git-file_test-move.go ./git-dir_test-move/
git add -A
git commit -m 'Moving file'

#commit of the action 'renaming and moving'
mv ./git-file_test-rename-and-move.go ./git-dir_test-rename-and-move/git-file_renaming-and-moving-tested.go
git add -A
git commit -m 'Moving and renaming file'

#commit of the action 'deleting and restoring'
rm git-file_test-delete-and-restore.go
git rm  git-file_test-delete-and-restore.go
git commit -m 'Deleting and restoring file - Deleting'

touch git-file_test-delete-and-restore.go
cat << FIN >> git-file_test-delete-and-restore.go
git text : test delete and restore
FIN

git add  git-file_test-delete-and-restore.go
git commit -m 'Deleting and restoring file - Restoring'