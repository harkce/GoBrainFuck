package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type cell struct {
	val  uint8
	next *cell
	prev *cell
}

type instruction struct {
	char byte
	ops  uint
}

func compile(input string) (compiled []instruction, err error) {
	jmpStack := make([]uint, 0)

	add := func(c byte, o uint) {
		compiled = append(compiled, instruction{char: c, ops: o})
	}

	for _, c := range input {
		if c == '>' || c == '<' || c == '+' || c == '-' || c == '.' || c == ',' {
			add(byte(c), 0)
		} else if c == '[' {
			jmpStack = append(jmpStack, uint(len(compiled)))
			add(byte(c), 0)
		} else if c == ']' {
			if len(jmpStack) == 0 {
				return nil, errors.New("compilation error")
			}
			pop := jmpStack[len(jmpStack)-1]
			compiled[pop].ops = uint(len(compiled))
			jmpStack = append(jmpStack[0 : len(jmpStack)-1])
			add(byte(c), uint(pop))
		}
	}
	return
}

func execute(program []instruction) {
	var i uint
	currCell := &cell{
		val:  0,
		next: nil,
		prev: nil,
	}
	reader := bufio.NewReader(os.Stdin)
	for i < uint(len(program)) {
		switch program[i].char {
		case '>':
			if currCell.next == nil {
				currCell.next = &cell{
					val:  0,
					next: nil,
					prev: currCell,
				}
			}
			currCell = currCell.next
		case '<':
			if currCell.prev == nil {
				currCell.prev = &cell{
					val:  0,
					next: currCell,
					prev: currCell,
				}
			}
			currCell = currCell.prev
		case '+':
			if currCell.val < 255 {
				currCell.val++
			} else {
				currCell.val = 0
			}
		case '-':
			if currCell.val > 0 {
				currCell.val--
			} else {
				currCell.val = 255
			}
		case '.':
			fmt.Printf("%c", currCell.val)
		case ',':
			read, err := reader.ReadByte()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			currCell.val = read
		case '[':
			if currCell.val == 0 {
				i = program[i].ops
			}
		case ']':
			if currCell.val != 0 {
				i = program[i].ops
			}
		default:
			panic("unknown operator")
		}
		i++
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Execute with 1 argument: the brainfuck file path")
		return
	}
	filename := os.Args[1]
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error reading", filename)
		return
	}
	program, err := compile(string(file))
	if err != nil {
		fmt.Println(err)
		return
	}
	execute(program)
}
