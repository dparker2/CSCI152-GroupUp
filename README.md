# GroupUp

`/global` holds static (html/css/js) files that should be included in more than one app, and font files used throughout the app.

To include:
```
src="/css/global/toolkit.css"
```

`/node_modules` holds all npm installed packages
To include:
```
src="/js/support/vue/dist/vue/js"
```

`/pkg` holds any go packages that do *not* depend on other groupup packages or that directly manipulate a model. Packages that only include standard/3rd party packages should be placed here and be made as general as possible.

`/src/controllers` holds groupup packages that relate to back-end routing and http authentication. These packages should react to http requests and possibly manipulate models.

`/src/models` is a package that implement a backend API for various models. Everything exported from here should be methods (not functions or data structures), and all DB related things should reside here, along with any caching implementation hidden from other packages.

`/src/system` holds all other groupup packages that do not fit into the others.

`/static` holds folders that are named consistent to their app's url routing. These should only contain static (html/css/js) files.
