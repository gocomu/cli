![Imgur](https://i.imgur.com/m4yuh20.png)

![](https://github.com/gocomu/cli/workflows/release/badge.svg?branch=master) [![codecov](https://codecov.io/gh/gocomu/cli/branch/master/graph/badge.svg)](https://codecov.io/gh/gocomu/cli)

# gocomu

`gocomu` is a command user interface designed to help speed things up when working with [`comu` library](http://github.com/gocomu/comu).  

## Installation

You can download the latest pre-compiled binary from [releases](https://github.com/gocomu/cli/releases)

or simply run `go get github.com/gocomu/cli/cmd/gocomu`

## Use

Test that everything works by printing `-help` flag

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

#### CLI

`gocomu new cli -name sampleProject`

```
.
├── cmd
│   └── sampleProject
│       └── main.go
├── embed
│   ├── embedded.go
│   └── embed.go
├── gocomu.yml
├── go.mod
└── output/
```

Your newly created project uses [`clir` library](https://github.com/leaanthony/clir). You can find detailed instructions on how to use it at [author's website](https://clir.leaanthony.com/).


#### GUI

_This template uses [`Fyne`](https://github.com/fyne-io/fyne)._

_From project's README:_

_In order to use it you will need Go version 1.12 or later. As Fyne uses CGo you will require a C compiler (typically gcc). If you don't have one set up the instructions at Compiling may help.
By default Fyne uses the gl golang bindings which means you need a working OpenGL configuration (or GLES for ARM or mobile devices). Debian/Ubuntu based systems may also need to install the `libegl1-mesa-dev` and `xorg-dev` packages._

_For more information refer to [project's documentation](https://github.com/fyne-io/fyne#prerequisites)._

### Serve

`gocomu serve`

### Record 

### Offline render

### Embed

### Build

`gocomu build`


# TODO
- [x] Create new project
  - [x] CLI Template
  - [ ] GUI Template
- [x] Project serve
  - [ ] comu's timeline starting markers/ques flag
- [ ] Project embedder (wav/aiff)
- [ ] Record audio output in real-time (wav/aiff)
- [ ] Render offline
- [x] Build stand-alone app
