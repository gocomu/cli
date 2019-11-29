![Imgur](https://i.imgur.com/m4yuh20.png)

![](https://github.com/gocomu/cli/workflows/release/badge.svg?branch=master) [![codecov](https://codecov.io/gh/gocomu/cli/branch/master/graph/badge.svg)](https://codecov.io/gh/gocomu/cli)

# gocomu

## installation

You can a. download a compiled binary fron [releases](https://github.com/gocomu/cli/releases)
or be simply run `go get github.com/gocomu/cli/cmd/gocomu`

## use

Test that everything works by printing `--help` flag

```
$ gocomu --help

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
   serve     Hot load your composition after save 
   offline   Render the output as wav/aiff 

Flags:

  -help
        Get help on the 'gocomu' command.
```

### create a new project

gocomu provides cli & gui based templates to get you started.

#### cli

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
└── output
```

#### gui


# TODO
- [ ] Create new project
  - [x] CLI Template
  - [ ] GUI Template
- [ ] Project serve
  - [ ] markers flag for starting from certain point in compositions
- [ ] Project embedder (wav/aiff)
- [ ] Render audio output
