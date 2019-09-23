# Rigetti Classical Computer implementation

## Author

* Tom Mornini (tmornini)

## Installation

``` bash
go get github.com/tmornini/rigetti-computing/rcc
```

## The Rigetti Classical Computer Coding Challenge

Rigetti is on a mission to build the world's most powerful classical computer, called the Rigetti Classical Computer (RCC). In order to do this, they've developed an entirely new ISA, which they are in the process of implementing in hardware. Below is a description of the ISA.

### Four general-purpose 8-bit registers

* X: #x00
* Y: #x01
* Z: #x02
* W: #x03

### Two flag registers

* C: #x04 Whether a comparison was successful.
* E: #x05 Whether an error occured in execution.

### One special purpose register

* PC: The program counter.

### Memory

* 256 bytes of program memory.
* 256 bytes of main memory.

Program memory is set upon initialization, and isn't otherwise accessible. Main memory is initialized to 0.

Both memories are addressed by bytes #x00 to #xFF.

### Sixteen variable-width instructions

NOP - #x00

* [No-op instruction. Do nothing.](https://github.com/tmornini/rigetti-computing/blob/master/processor/instruction-set.go#L11-L13)

ADD - #x01 \<r1\> \<r2\> \<r3\>

* [Add registers r1 and r2, and deposit the result into register r3.](https://github.com/tmornini/rigetti-computing/blob/master/processor/instruction-set.go#L15-L23)

SUB - #x02 \<r1\> \<r2\> \<r3\>

* [Subtract register r2 from r1, and deposit the result into register r3.](https://github.com/tmornini/rigetti-computing/blob/master/processor/instruction-set.go#L25-L33)

MUL - #x03 \<r1\> \<r2\> \<r3\>

* [Multiply r1 and r2, and deposit the result into register r3.](https://github.com/tmornini/rigetti-computing/blob/master/processor/instruction-set.go#L35-43)

DIV - #x04 \<r1\> \<r2\> \<r3\>

* [Integer divide r1 by r2, and deposit the result into register r3. If r2 is zero, then set the E flag.](https://github.com/tmornini/rigetti-computing/blob/master/processor/instruction-set.go#L45-L60)

LDM - #x05 \<r1\> \<r2\>

* [Load the contents located at the address in register r1 into the register r2.](https://github.com/tmornini/rigetti-computing/blob/master/processor/instruction-set.go#L62-L77)

LDI - #x06 \<imm\> \<r1\>

* [Load the constant byte specified by imm into the register r1.](https://github.com/tmornini/rigetti-computing/blob/master/processor/instruction-set.go#L79-L87)

STR - #x07 \<r1\> \<r2\>

* [Store the contents of register r1 into the address stored in register r2.](https://github.com/tmornini/rigetti-computing/blob/master/processor/instruction-set.go#L102)

SWP - #x08 \<r1\> \<r2\>

* [Swap the contents of registers r1 and r2.](https://github.com/tmornini/rigetti-computing/blob/master/processor/instruction-set.go#L112)

EQL - #x09 \<r1\> \<r2\>

* [Check if registers r1 and r2 contain the same value. If so, set the C flag.](https://github.com/tmornini/rigetti-computing/blob/master/processor/instruction-set.go#L114-122)

NQL - #x0A \<r1\> \<r2\>

* [Check if registers r1 and r2 contain different values. If so, set the C flag.](https://github.com/tmornini/rigetti-computing/blob/master/processor/instruction-set.go#L124-L132)

JMP - #x0B \<imm\>

* [Jump unconditionally to the address designated by imm.](https://github.com/tmornini/rigetti-computing/blob/master/processor/instruction-set.go#L134-L138)

JMC - #x0C \<imm\>

* [Jump conditionally on the C flag to the address designated by imm. Clear the C flag after.](https://github.com/tmornini/rigetti-computing/blob/master/processor/instruction-set.go#L140-L148)

JME - #x0D \<imm\>

* [Jump conditionally on the E flag to the address designated by imm. Clear the E flag after.](https://github.com/tmornini/rigetti-computing/blob/master/processor/instruction-set.go#L150-158)

PRN = #x0E \<r1\>

* [Print to standard output the character represented by the ASCII value in register r1.](https://github.com/tmornini/rigetti-computing/blob/master/processor/instruction-set.go#L160-L172)

HLT - #x0F

* [Halt all execution.](https://github.com/tmornini/rigetti-computing/blob/master/processor/instruction-set.go#L174-L176)

### Caveats

* arithmetic is "wrap-around"; all operations are essentially done modulo 256.

### Rational

Unfortunately, building a new processor is a very time- and resource-intensive process, and Rigetti will need to provide software and applications as soon as the machine becomes available. As such, Rigetti has opted to build software tools to assist writing and debugging programs written for the RCC.

### Assignment

Your task is to help make these tools by writing C/C++ code for GNU/Linux using Clang or GCC, using best practices. _(Rigetti allowed me to write this in the Google Go language instead of C/C++)_

1. [Write a small program for the RCC to swap the contents of two memory locations M1 and M2.](spec-successes/assignment-1.hex) Translate it into a binary file containing the machine code. __(see 4.2 below)__ - results: [STDERR](spec-successes/assignment-1.stderr) & [STDOUT](spec-successes/assignment-1.stdout)
2. [Write a small program for the RCC to compute the length of a 0-terminated alphanumeric ASCII string, which is known to start at address M.](spec-successes/assignment-2.hex) - results: [STDERR](spec-successes/assignment-2.stderr) & [STDOUT](spec-successes/assignment-2.stdout)
3. Critique the ISA. What features might be changed or added to make it more practical for more general-purpose computations?
   * An increment register instruction, or add-immediate instruction would be a godsend.
   * A Z (zero) register would be terrific for fixed count loops and null terminated string handling.
   * You can never be too rich, too thin, too tan, or have too many registers.
   * It's a small thing, but for consistency it would be nice if W was register 0x00 and the others were re-ordered to accomodate this. 😮 I'll have to assume that this is RCC ISA 2.0 and W was added after 1.0 and backward compatibility was required. 😀
   * The SWP instruction appears to be unnecessary. All other instructions treat registers orthogonally, it's unclear why you'd ever need to swap register values.
   * If additional instructions could be added, a `JSR` (jump to subroutine) and `RET` (return from subroutine) instruction pair would be highly useful. They would require a return address stack to be allocated in main memory, which is extremely precious on the RCC, but the [Atari 2600 VCS](https://en.wikipedia.org/wiki/Atari_2600_hardware) used a MOS Technology 6502 with such an instruction pair and had just 128 bytes of RAM, half that of the RCC and the 6502 required 2 bytes per return address wheras the RCC requires just one.
4. [Write a simulator for the RCC](rcc/main.go). It should be able to be built from a [makefile](Makefile), and should be callable in the following way: `rcc file.bin` where `file.bin` contains binary machine code to be executed on the simulator.
   * This implementation can read program from ``STDIN`` in addition to the specified CLI file argument.
   * I found it more convenient to write machine code in a hex file with assembly language encoded in the ASCII section. `rcc-hex file.hex` converts the hex to binary and pipes it into `go run rcc/main.go` 😎
5. [If the environment variable DEBUG is set to anything, then print out the instructions being executed as they get executed.](https://github.com/tmornini/rigetti-computing/blob/master/rcc/main.go#L41-L47)
   * for sanity, I implemented a separate, depedent NONOP environment variable to suppress NOP debugging.
6. [At the end of execution, all bytes of memory should be printed to standard out on a single line.](https://github.com/tmornini/rigetti-computing/blob/master/processor/processor.go#L40)
   * The words "all bytes of memory" are not specific enough to guarantee that I've implemented this properly, so I've chosen to implemented it to mean "all bytes of main memory." I do output all bytes of memory, including program memory and all registers and flags as well. 😊
7. Discuss how this ISA could be modified to allow for interrupts.
   * My experience with processor interrupts was that they are implemented as pointers at fixed addresses in the program memory, one for each type of interrupt. When interrupted, the pointer would be placed into the program counter to execute the interrupt code.
   * With interrupts, you'd need a way to keep a separate register context for, at minimum, the program counter and, ideally, the entire register set to allow the main line of execution to resume when the interrupt is complete. It's possible, though inefficiewnt, inconvenient and error prone, to handle register swapping upon entry and exit of the interrupt code manually.
8. Discuss how this ISA might be implemented in hardware.
   * I'm neither qualified nor able to answer this question. 🤷‍

## Implementation

* This implementation has 4 possible exit codes
  1. 0 - Execution exited normally
  2. 1 - Error opening binary code file supplied as argument to the executable
  3. 2 - Too many arguments were supplied. A usage message provided
  4. 3 - There was an error reading binary code from STDIN or file supplied as argument
  5. 4 - Execution exited unexpectedly
* All DEBUG output is written to `STDERR` to keep `STDOUT` tidy

## Testing

  1. The [`./test`](https://github.com/tmornini/rigetti-computing/blob/master/test) script builds [`rcc/main.go`](https://github.com/tmornini/rigetti-computing/blob/master/rcc/main.go) then converts all `spec-*/*.hex` files to `spec-*/*.bin` files using `xxd`.
  2. It then runs the `rcc` executable (with `DEBUG=true`) against all the `.bin` files, makes certain the spec-failures/ do fail and spec-successes/ do succeed.
  3. In addition it compares the actual `STDOUT` and `STDERR` against corresponding `.stdout` and `.stderr` spec files. These form very complete integration tests to make certain that the code behaves as it is intended to.
  4. On failure of step 2, `./test` overwrites the  `.stdout` and `.stderr`  files with the actual output, then executes `git diff` to conveniently highlight the difference(s). This aided debugging enormously.
  5. Due to step 4, you are cautioned against accidentally commiting spec failures. 👀

### I'm super happy with the way this came together, particularly with respect to the readability of the verb and noun set. Hope you enjoy reading the code as much as I enjoyed writing it! I'm also quite curious to know how it performs compared to other efforts at the same stage of development
