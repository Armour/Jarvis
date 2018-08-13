# Jarvis

[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat)](http://makeapullrequest.com)
[![Go Report Card](https://goreportcard.com/badge/github.com/Armour/jarvis)](https://goreportcard.com/report/github.com/Armour/jarvis)
[![Go Project Layout](https://img.shields.io/badge/go-layout-blue.svg)](https://github.com/golang-standards/project-layout)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/armour/jarvis)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Template from jarvis](https://img.shields.io/badge/Hi-Jarvis-ff69b4.svg)](https://github.com/Armour/Jarvis)

> "Perhaps, if you intend to visit other planets, we should improve the exosystems."

Personal assistant 🤖 [still under development]

## Projects used Jarvis

* Go:
  * [armour/jarvis](https://github.com/Armour/jarvis) - Personal assistant
  * [armour/go-node-builtins](https://github.com/Armour/go-node-builtins) - List of node.js builtin modules in Golang
  * [armour/go-validate-npm-package-name](https://github.com/Armour/go-validate-npm-package-name) - Check if the given string is an acceptable npm package name in Golang
* Python:
  * [armour/Automatic-Image-Colorization](https://github.com/Armour/Automatic-Image-Colorization) - Automatic Image Colorization using TensorFlow based on Residual Encoder Network
  * [armour/pixiv-spider](https://github.com/Armour/pixiv-spider) - A simple Pixiv Crawler
  * [armour/upass-sfu](https://github.com/Armour/upass-sfu) - Python script to auto renew monthly UPass for SFU student
  * [armour/gmail-bot](https://github.com/Armour/gmail-bot) - Gmail bot for CMPT412 TA
  * [armour/sentence-classifier-based-on-word-similarity](https://github.com/Armour/sentence-classifier-based-on-word-similarity) - A very simple sentence classifier based on word similarity with NLTK and rake_nltk package
* React:
  * [armour/express-webpack-react-redux-typescript-boilerplate](https://github.com/Armour/express-webpack-react-redux-typescript-boilerplate) - A full-stack boilerplate that using express with webpack, react and typescirpt
* Npm:
  * [armour/commitlint-config-armour](https://github.com/Armour/commitlint-config-armour) - My shareable config for commitlint
* Unity:
  * [armour/Multiplayer-FPS](https://github.com/Armour/Multiplayer-FPS) - A multiplayer first person shooter game based on Unity Game Engine
* Misc:
  * [armour/Magic-Tower-Qt](https://github.com/Armour/Magic-Tower-Qt) - A C++(with Qt) version Magic-Tower game
  * [armour/vscode-typescript-react-redux-snippets](https://github.com/Armour/vscode-typescript-react-redux-snippets) - Typescript, React and Redux snippets for VSCode (followed ES6 standard)

## Install

```bash
go get github.com/armour/jarvis
```

## Config file

Create a `jarvis.json` file under home directory like below, `jarvis` will use these informations later in the project generator.

```json
{
    "author": "Chong Guo",
    "email": "armourcy@gmail.com",
    "githubUser": "Armour",
    "dockerUser": "cguo"
}
```

## Todos

* [x] Go project generator
* [x] Python project generator
* [x] React project generator
* [x] Npm project generator
* [x] Unity project generator
* [x] Misc project generator
* [ ] Vue project generator
* [x] Dot file backup/sync
* [x] Play around with [Hitokoto](https://hitokoto.cn/) api
* [ ] Google Home integration
* [ ] And more...

## Usage examples

* Say something（#￣▽￣#)
  * <a href="https://asciinema.org/a/184121"><img src="https://asciinema.org/a/184121.png" alt="https://asciinema.org/a/184121.png" width="75%"></a>

* Sync global dot files
  * <a href="https://asciinema.org/a/185548"><img src="https://asciinema.org/a/185548.png" alt="https://asciinema.org/a/185548.png" width="75%"></a>

* Start new project using react template
  * <a href="https://asciinema.org/a/190782"><img src="https://asciinema.org/a/190782.png" alt="https://asciinema.org/a/190782.png" width="75%"></a>

* Start new project using go template
  * <a href="https://asciinema.org/a/190781"><img src="https://asciinema.org/a/190781.png" alt="https://asciinema.org/a/190781.png" width="75%"></a>

## Contributing

See [CONTRIBUTING.md](https://github.com/Armour/jarvis/blob/master/.github/CONTRIBUTING.md)

## License

[MIT License](https://github.com/Armour/jarvis/blob/master/LICENSE)
