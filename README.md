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



