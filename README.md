# Screenshot
A screenshot CLI application built in Golang. This app uses [`Cobra`](https://github.com/spf13/cobra), as well as a bunch of other libraries I created. 

This app will allow you to take a screenshot of your whole screen, or just a selected region. It will copy the screenshot to your clipboard, and with the `--upload` flag will upload it to [imgur](https://imgur.com/).

## Installation
Make sure you have your `$GOPATH` setup properly. Clone the repo:
```
git clone https://github.com/ZacJoffe/screenshot.git
```
Then install it:
```
cd screenshot; go install
```

## Use
Invoke the app with the `screenshot` command. Using the `--select` (or `-s` for short) flag will allow you to select a region. Using the `--upload` (or `-u`) flag will upload the screenshot to imgur, output the link and copy it to your clipboard.

So, to select a region of the screen and upload it, you can use:
```
screenshot -su
```