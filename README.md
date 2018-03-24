# GoFuck
Brainfuck interpreter written in Go. This is the faster version of my previous brainfuck interpreter:
https://github.com/harkce/GoFuckYourBrain

## What the F is Brainfuck?
From https://en.wikipedia.org/wiki/Brainfuck

Brainfuck is an esoteric programming language created in 1993 by Urban Müller, and notable for its extreme minimalism.

The language consists of only eight simple commands and an instruction pointer. While it is fully Turing-complete, it is not intended for practical use, but to challenge and amuse programmers. Brainfuck simply requires one to break commands into microscopic steps.

The language's name is a reference to the slang term brainfuck, which refers to things so complicated or unusual that they exceed the limits of one's understanding.

## Usage
1. Clone this repository
2. On project directory, execute:
```
$ go build
```
OR
If you don't want to do step 3, execute:
```
$ go run main.go examples
```
3. To interpret a brainfuck file:
```
$ ./GoFuck examples/HelloWorld.bf
```
