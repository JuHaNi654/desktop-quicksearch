# README
Basic scuffed quick search bar. That works somehow in ubuntu 24.04

## Required things
- gtk-launch
- golang (1.21)
- node (22.13)
- https://wails.io/docs/gettingstarted/installation

## Links 
- https://specifications.freedesktop.org/desktop-entry-spec/latest/index.html

## Run development mode 
```bash
wails dev -tags="webkit2_41"
```

## Build application 
```bash
wails build -tags="webkit2_41" -o="quicksearch" -clean
```
