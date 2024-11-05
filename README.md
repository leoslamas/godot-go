# godot-go
## Minimal Configuration to Start a Godot Project Using Go

This repository provides a minimal configuration to kickstart a Godot project using Golang as the game scripting language, following the tutorial from [grow-graphics/gd](https://github.com/grow-graphics/gd). 


## Prerequisites

1. Make sure you have the Godot Engine installed.
2. Install Go on your system if you haven't done so already.
3. Install *gd* tool: `go install grow.graphics/gd/cmd/gd@master`

## Project Structure

Here's an overview of the project structure:

```
godot-go
│
├── .git/
│
├── graphics/
│   ├── .godot/
│   ├── library.gdextension
│   ├── icon.svg
│   ├── main.tscn
│   └── project.godot
│
└── main.go
```

## Building and Running

1. Build the Go project using the following command:

```bash
gd build 
```

or

```bash
go build -o example.so -buildmode=c-shared
```

2. Examples

- https://github.com/grow-graphics/eg
