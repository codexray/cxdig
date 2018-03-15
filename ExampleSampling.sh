	#!/bin/sh
	toolpath=$1
	repopath=$2
	outputname=$3
	$toolpath $repopath --json>$outputname
