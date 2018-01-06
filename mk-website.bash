#!/bin/bash

function cleanUpHTML() {
	findfile -s ".html" . | while read -r P; do
		rm "$P"
	done
}

function FindNavMD() {
	DNAME="$1"
	if [ -f "${DNAME}/nav.md" ]; then
		echo "${DNAME}/nav.md"
		return
	fi
	DNAME=$(dirname "${DNAME}")
	FindNavMD "${DNAME}"
}

# Cleanup stale HTML files
cleanUpHTML

# Look through files and build new site
TITLE=""
if [[ "${1}" != "" ]]; then
	TITLE="${1}"
fi

mkpage "nav=nav.md" "content=README.md" page.tmpl >index.html
git add index.html

mkpage "nav=nav.md" "content=markdown:$(cat LICENSE)" page.tmpl >license.html
git add license.html

mkpage "nav=nav.md" "content=INSTALL.md" page.tmpl >install.html
git add install.html

findfile -s ".md" . | while read -r P; do
	DNAME="$(dirname "$P")"
	FNAME=$(basename "$P")
	PREFIX="${DNAME:0:4}"

	if [[ "${PREFIX}" == "dist" || "${PREFIX}" == "test" || "${FNAME}" == "nav.md" || "${FNAME}" == "README.md" || "${FNAME}" == "INSTALL.md" || "${FNAME}" == "TODO.md" || "${FNAME}" == "IDEAS.md" ]]; then
		echo "Skipping $P"
	else
		HTML_NAME="${DNAME}/$(basename "$FNAME" ".md").html"
		NAV=$(FindNavMD "$DNAME")
		echo "Building $HTML_NAME from $DNAME/$FNAME and $NAV"
		mkpage "title=text:${TITLE}" "nav=${NAV}" "content=${DNAME}/${FNAME}" page.tmpl >"${HTML_NAME}"
		git add "${HTML_NAME}"
	fi
done
