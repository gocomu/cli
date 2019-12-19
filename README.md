![Imgur](https://i.imgur.com/m4yuh20.png)

![](https://github.com/gocomu/cli/workflows/release/badge.svg?branch=master) [![codecov](https://codecov.io/gh/gocomu/cli/branch/master/graph/badge.svg)](https://codecov.io/gh/gocomu/cli)

# gocomu

`gocomu` is a command user interface designed to help speed things up when working with [`comu` library](http://github.com/gocomu/comu).  

## Installation

You can download the latest pre-compiled binary from [releases](https://github.com/gocomu/cli/releases)

or simply run `go get github.com/gocomu/cli/cmd/gocomu`

## Use

Test that everything works by printing `-help`

```
$ gocomu -help

      ::::::::       ::::::::       ::::::::       ::::::::         :::   :::      :::    ::: 
    :+:    :+:     :+:    :+:     :+:    :+:     :+:    :+:       :+:+: :+:+:     :+:    :+:  
   +:+            +:+    +:+     +:+            +:+    +:+      +:+ +:+:+ +:+    +:+    +:+   
  :#:            +#+    +:+     +#+            +#+    +:+      +#+  +:+  +#+    +#+    +:+    
 +#+   +#+#     +#+    +#+     +#+            +#+    +#+      +#+       +#+    +#+    +#+     
#+#    #+#     #+#    #+#     #+#    #+#     #+#    #+#      #+#       #+#    #+#    #+#      
########       ########       ########       ########       ###       ###     ########        

  v0.0.1 - GOCOMU CLI

Available commands:

   new       Create New Project 
   serve     Hot load your composition while working
   record    Record audio output in real-time as wav/aiff
   offline   Render audio output as wav/aiff 
   embed     Embed all *.wav/*.aiff files as []byte 
   build     Build stand-alone application 

Flags:

  -help
        Get help on the 'gocomu' command.
```

### New

gocomu provides cli & gui based templates to get you started.

#### Project Structure

```
.
├── cmd
│   ├── gocomu.go
│   └── sampleProject
│       └── main.go
├── embed
│   └── fs.go
├── gocomu.yml
├── go.mod
├── sine.go
└── output/
```

#### CLI

`gocomu new cli -name sampleProject`

#### GUI

_GUI template uses [`Fyne`](https://github.com/fyne-io/fyne)._

_From project's README:_

_In order to use it you will need Go version 1.12 or later. As Fyne uses CGo you will require a C compiler (typically gcc). If you don't have one set up the instructions at Compiling may help.
By default Fyne uses the gl golang bindings which means you need a working OpenGL configuration (or GLES for ARM or mobile devices). Debian/Ubuntu based systems may also need to install the `libegl1-mesa-dev` and `xorg-dev` packages._

_For more information refer to [project's documentation](https://github.com/fyne-io/fyne#prerequisites)._

`gocomu new gui -name sampleProject`


### Serve

`gocomu serve`

### Record 

WIP

### Offline render

WIP

### Embed

`gocomu embed`

### Build

`gocomu build`


# Roadmap to v1.0.0
- [x] Create new project
  - [x] CLI Template
  - [x] GUI Template
- [x] Project serve
  - [ ] comu's timeline starting markers flag
- [x] Project embedder (wav/aiff)
- [ ] Real-time audio recording (wav/aiff)
- [ ] Render offline
- [x] Build stand-alone app
