    npm install shelljs
    vi -e -c "20,75wq! mw-bootstrap.js" README.md
    sed -ie "s/    //" mw-bootstrap.js
    chmod 770 mw-bootstrap.js
    ./mw-bootstrap.js Markdown-Weave.md
