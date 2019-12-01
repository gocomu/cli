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
   embed     Embed all *.wav/*.aiff files as []byte 
   serve     Hot load your composition while working
   offline   Render the output as wav/aiff 
   build     Build a stand-alone binary 

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
├── cmd/
│   └── sampleProject
│       └── main.go
├── embed/
│   ├── embedded.go
│   └── embed.go
├── gocomu.yml
├── go.mod
└── output/
```

For gocomu itself as well as new created cli projects we use [`clir` library](https://github.com/leaanthony/clir). You can find detailed instructions on how to use it [on author's website](https://clir.leaanthony.com/).


#### GUI

**NOTE: in order to use the gui template you need `wails` on your system.
For instructions on how to install refer to [project's documentation](https://github.com/wailsapp/wails#installation).**

### Embed

### Serve

### Offline render

### Build


# TODO
- [ ] Create new project
  - [x] CLI Template
  - [ ] GUI Template
- [ ] Project serve
  - [ ] markers flag for starting from certain point in compositions
- [ ] Project embedder (wav/aiff)
- [ ] Render audio output
