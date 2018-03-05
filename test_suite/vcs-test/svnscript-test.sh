#!/bin/bash
#Script creating a svn repository and committing all the base actions.

# make bash stop and fail if any command fails
set -e

#Defining repository name
SVNREPO=svn-server_test
SVNCHECKOUT=svn-repository_test

# we need a SVN server root dir to be able to run svn client commands
svnadmin create $SVNREPO

#Initialisation of the repository
svn checkout file:///$PWD/$SVNREPO/ $SVNCHECKOUT
cd $SVNCHECKOUT
touch README.md
svn add README.md --force

svn commit -m 'Initialisating testing svn repository'


#Commit of the action 'add'
touch svn-file_test.go
cat << FIN >> svn-file_test.go
svn text : test
FIN

touch svn-file_test-delete.go
cat << FIN >> svn-file_test-delete.go
svn text : test
FIN

touch svn-file_test-modify.go
cat << FIN >> svn-file_test-modify.go
svn text : test modify
FIN

touch svn-file_test-rename.go
cat << FIN >> svn-file_test-rename.go
svn text : test rename
FIN

touch svn-file_test-move.go
cat << FIN >> svn-file_test-move.go
svn text : test move
FIN
mkdir svn-dir_test-move

touch svn-file_test-rename-and-move.go
cat << FIN >> svn-file_test-rename-and-move.go
svn text : test rename and move
FIN
mkdir svn-dir_test-rename-and-move

touch svn-file_test-delete-and-restore.go
cat << FIN >> svn-file_test-delete-and-restore.go
svn text : test delete and restore
FIN

svn add * --force

svn commit -m 'Adding files'

#Commit of the action 'delete'
svn rm ./svn-file_test-delete.go
svn commit -m 'Deleting file'

#Commit of the action 'modify'
cat << FIN >> svn-file_test-modify.go
svn text : modifying tested
FIN
svn commit -m 'Modifying file'

#commit of the action 'renaming'
svn mv ./svn-file_test-rename.go ./svn-file_renaming-tested.go
svn commit -m 'Renaming file'

#commit of the action 'moving'
svn mv ./svn-file_test-move.go ./svn-dir_test-move/
svn commit -m 'Moving file'

#commit of the action 'renaming and moving'
svn mv ./svn-file_test-rename-and-move.go ./svn-dir_test-rename-and-move/svn-file_renaming-and-moving-tested.go
svn commit -m 'Moving and renaming file'

#commit of the action 'deleting and restoring'
svn rm svn-file_test-delete-and-restore.go
svn commit -m 'Deleting and restoring file - Deleting'

touch svn-file_test-delete-and-restore.go
cat << FIN >> svn-file_test-delete-and-restore.go
svn text : test delete and restore
FIN
svn add svn-file_test-delete-and-restore.go --force
svn commit -m 'Deleting and restoring file - Restoring'

svn update