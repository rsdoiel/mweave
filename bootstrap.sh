#!/bin/bash
npm install shelljs
vi -e -c "20,79wq! mw-bootstrap.js" README.md
sed -i -e "s/    //" mw-bootstrap.js
chmod 770 mw-bootstrap.js
./mw-bootstrap.js
npm install
npm test
