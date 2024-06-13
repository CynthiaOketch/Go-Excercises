# ASCII Art Renderer

## Introduction

This project is an ASCII art renderer, designed to take a string input as an argument and output a graphical representation using ASCII characters. The program produces a visual representation of the input string, converting it into ASCII art.

## What is ASCII Art?
ASCII art involves creating visual images using characters from the ASCII character set. In this project, each character in the input string is represented by a corresponding ASCII art pattern, as exemplified below:
`
@@@@@@BB@@@@``^^``^^``@@BB$$@@BB$$
@@%%$$$$^^^^WW&&8888&&^^""BBBB@@@@
[@@@@@@""WW8888&&WW888888WW``@@@@$$
BB$$``&&&&WWWW8888&&&&8888&&``@@@@
$$``&&WW88&&88&&&&8888&&88WW88``$$
@@""&&&&&&&&88888888&&&&&&88&&``$$
``````^^``^^^^^^````""^^``^^``^^``
...`](https://)


## Prerequisites

- Go programming language installed on your system.

## Usage

To use the tool, follow these steps:

1. Prepare a text file with the content you want to modify.
2. Run the tool by providing the input file name and the output file name as arguments.

```bash
$ go run . input.txt output.txt