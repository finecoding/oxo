package oxo

import (
	"errors"
	"fmt"
)

/* A simple command line interface programme.

What will it do

A simple collection of methods that work on a datatype to hold a noughts and crosses grid.
So our datatype will be a nine byte slice

Lets call it....board

This datatype will be used to create instances */

//import "github.com/spf13/"
const s = " "
const x = "X"
const o = "O"

//S is a space, could use and 'S' instead.
const S = 32

//X is a space, could use and 'S' instead.
const X = 88

//O is a space, could use and 'S' instead.
const O = 79

//these error variables are not used anymore, I created am error type instead create a
//struc to hold useful error information and this is used by a swithc statement in main,
//ErrInvalidParam = fmt.Errorf("invalid parameter [%s]", string(b))

//ErrIllegalMarker ...lots of sttuff
var ErrIllegalMarker = errors.New("bad marker, only O, X or S permitted on Board")

//ErrIllegalLength ...lots  of stuff
var ErrIllegalLength = errors.New("bad length, Board must be of length 9")

// ErrTooManyX ...lots of stuff comment
var ErrTooManyX = errors.New("too many X")

// ErrTooManyO comment
var ErrTooManyO = errors.New("too many O")

// board is the datastructure used to hold the contents of the nine postitions on
// the 3 x 3 grid used in the game.  Enumerated from top left to bottom right, starting
// at 0.
type board []byte

//boardError, field t is for type of error from 1 to 5.  This helps select the switch in main()
//Errors are formatted a bit differently depending on what was amiss.
//Some of the fields in the struct are relevant to an error, others are not.
//Eg.  for a bad character error, the offending char is held in the struct, but for
// a too many x error, this is irrelevant and is not used.
// The boardError stuct holds all data on all errors.
type boardError struct {
	t                 int
	message           string
	len, pos, x, o, s int
	bad               byte
}

//Error method makes boardError consistent with the Error interface.
//The function doesn't do much execept return a formatted string
//However, the boardError struct that is passed is used to carry useful
//contextual information in its fields that can be used by the
//calling program.
func (e *boardError) Error() string {
	format := "%s"
	return fmt.Sprint(format, e.message)
}

//old method
func (b board) countX() int {
	var count int
	fmt.Println("hello from method")
	for _, v := range b {
		if v == 88 {
			count++
		}
	}
	return count
}

//countXt - an old function
func countXt(b board) int {
	var count int
	fmt.Println("hello from function")

	//	fmt.Println(string(b))
	for _, v := range b {
		if v == 79 {
			count++
		}
	}
	return count
}

//countThisA simple counter function -  now incorporated in countThem()
func (b board) countThis(marker byte) int {
	var count int
	for _, v := range b {
		if v == marker {
			count++
		}
	}
	return count
}

//Counthem takes a board struct repsresenting of 9 squares of a 3 x 3 grid and
// returns any error information in a boardError struct
//This is the latest error detection method.  It detects several types of errors
//while looping through the elements of the board, setting values in the fields of
//e, which has the boardError struct.  This error checking is for really basic errors:
//Whether the difference in the number of Xs and Os is legal, it checks the length of board and
//reports on illegal markers (it has to be a O, X or space.
//This error reporting will be useful in whittling down big lookup tables and sanitising
//imported data files of historic games.
//For the former, we can later add detection of various illegal states that the game
//should never reach.
//If a lookup table is generated of all possible permutations of game, a large number
//of them will be illegal states: eg.  All Xs or all Os.  The test that detects an illegal imbalance
//between the number of Os will filter a lot of these.  However, there are more subtle illegal states
//to consider:
//Eg.  You can't have a state where both X and O win at the same time.
//Detecting error conditions is contained in this method.  Another method will be used
//to identify the state of a game concluding to a winner,loser or draw state.

func (b board) countThem() (err boardError) {
	var e boardError
	if len(b) > 9 {
		e.len = len(b)
		e.message = "too long"
		e.t = 1

	}
	if len(b) < 9 {
		e.len = len(b)
		e.message = "too short"
		e.t = 2

	}
	//ranges over board counting markers
	for i, v := range b {
		switch v {

		case X:
			{
				e.x++
			}
		case O:
			{
				e.o++
			}
		case S:
			{

				e.s++
			}
		default: //error if not an X, O or space.  Return position and value of bad marker.
			{
				e.message = "illegal marker"
				e.pos = i
				e.bad = v
				e.t = 3
			}
		}
	}
	//test for too many Xs on the board
	if e.x-e.o > 1 {
		e.message = "too many X"
		e.t = 4

	}
	//test for too may Os on the board
	if e.o-e.x > 1 {
		e.message = "too many O"
		e.t = 5

	}

	return e
}

//String method makes type board consistent with the Stringer interface so
//it can be used by fmt.format prints a grid for the game with the markers in board used for each
//square.
//
//fmt.Printf(b.String()) pretty prints the contents of an instance of board in b.
//This is used by the cases in the switch in main()

func (b board) String() string {
	format := "\n      -------------\n      | %s | %s | %s |\n      -------------" + "\n      | %s | %s | %s |\n      -------------" + "\n      | %s | %s | %s |\n      -------------\n\n"
	return fmt.Sprintf(format, string(b[0]), string(b[1]), string(b[2]), string(b[3]), string(b[4]), string(b[5]), string(b[6]), string(b[7]), string(b[8]))

}

func (b board) error() string {
	//what to put in here?
	//
	//we need the values of o, x and s
	//
	//if (x-o) > 1 then too many x
	//if (o-x) > 1 then too many o
	//if len(b) !=9 then panic
	//if marker is not o, x or s then panic
	//
	return "boo"
}
func cmd() {
	var b board
	//b contains a record the relative positions of each marker on a 3 x 3 board state
	b = []byte("O  OXXX O")

	/*
		fmt.Printf("type: %T, value: %v\n, %s, a space is: '%d',an X is '%d',an O is '%d'\n", b, b, string(b), S, X, O)
		fmt.Println()
		fmt.Printf(" count of x from method %d\n", b.countX())
		fmt.Printf(" count of o from function %d\n\n\n", countXt(b))
		fmt.Printf(" count of x from method counthis %d\n", b.countThis(X))
		fmt.Printf(" count of o from method counthis %d\n", b.countThis(O))
		fmt.Printf(" count of space from method counthis %d\n\n\n", b.countThis(S))
	*/
	//countThem is the error checking method returns a error struct
	//the fields are used to report on different types of error
	//The switch works on the t field of the boardError struct held in instance e
	//
	e := b.countThem()
	//print a message and information for each error type
	switch e.t {

	case 0: //all good
		{
			fmt.Printf("\n        all good")
			fmt.Printf(b.String())
		}
	case 1: //too long
		{
			fmt.Printf("\n        error - %s length = %d", e.message, e.len)
		}
	case 2: //too short
		{
			fmt.Printf("\n        error - %s length = %d", e.message, e.len)
		}
	case 3: //illegal marker
		{
			fmt.Printf("\n        error - %s (%s at position %d) [%s]", e.message, string(e.bad), e.pos+1, string(b))
			fmt.Printf(b.String())
		}
	case 4: //too many x
		{
			fmt.Printf("\n        error - %s (x = %d o = %d)", e.message, e.x, e.o)
			fmt.Printf(b.String())
		}
	case 5: //too many o
		{
			fmt.Printf("\n        error - %s (x = %d o = %d)", e.message, e.x, e.o)
			fmt.Printf(b.String())
		}
	}

} /*

    we now need a method to extract a slice representing the locations of spaces
    on the board.  This will be used to provide players will a list of possible moves

    we also need some methods to test for a winner, for that we will need a lookup table
    this should be prepared before the program runs in an init function
    we need to turn this into a CLI program
    what about testing?  Loading test files.
    a test file should a key/value store.  A value for board and an expected result.
    what about logging.  Some of these error messages could go in a log file.
    how about a print method to print a grid?

    need to organise the structure of the repo and use github

	a package lives in a specific directory that may have sub directories

	choose good names for users to understand, make them short

	In this example, the program is a game of noughts and crosses.
	I call the package oxo.  But how to organise it internally?

	I can think of several different parts, some mentioned already:

	lookup table generator
	error detection
	command line interface
	print to text or web

	Organising workspace is very important.  This needs to be written down with a nice
	diagram showing directory trees.

	How vendoring is organised is particularly important because of the changes between
	Go versions and the various vendoring tools in use.

	Jetbrains! How do I map its projects to seperate workspaces?

	Keep the source code for all projects in the same workspace?

	From:

	https://golang.org/doc/code.html#Workspaces

    Go programmers typically keep all their Go code in a single workspace.
    A workspace contains many version control repositories (managed by Git, for example).
    Each repository contains one or more packages.
    Each package consists of one or more Go source files in a single directory.
    The pat A workspace is a directory hierarchy with three directories at its root:

    src contains Go source files,
    pkg contains package objects, and
    bin contains executable commands.
	Path to a package's directory determines its import path.

	 The go tool builds source packages and installs the resulting binaries to the pkg and bin directories.

The src subdirectory typically contains multiple version control repositories (such as for Git or Mercurial) that track the development of one or more source packages.

To give you an idea of how a workspace looks in practice, here's an example:

bin/
    hello                          # command executable
    outyet                         # command executable
pkg/
    linux_amd64/
        github.com/golang/example/
            stringutil.a           # package object
src/
    github.com/golang/example/
        .git/                      # Git repository metadata
	hello/
	    hello.go               # command source
	outyet/
	    main.go                # command source
	    main_test.go           # test source
	stringutil/
	    reverse.go             # package source
	    reverse_test.go        # test source
    golang.org/x/image/
        .git/                      # Git repository metadata
	bmp/
	    reader.go              # package source
	    writer.go              # package source
    ... (many more repositories and packages omitted) ...

The tree above shows a workspace containing two repositories (example and image). The example repository contains two commands (hello and outyet) and one library (stringutil). The image repository contains the bmp package and several others.

A typical workspace contains many source repositories containing many packages and commands. Most Go programmers keep all their Go source code and dependencies in a single workspace.

Commands and libraries are built from different kinds of source packages. We will discuss the distinction later.

	A package is a directory inside your $GOPATH/src directory containing, amongst other things, .go source files.

	package names should be lower case with no punctuation
c




	Notes:

	A workspace is a directory that holds all your files
	It contains many source repositories
	This arrangement replaces a Makefile
	Change the file layout, change the build
	GOPATH environmental variable points to the current workspace
	go get copies repos from the Internet to your local workspace
	Paths are a direct indication of the source of a package in the Internet
	go install builds a binary and places in your workspace bin directory, it will
	use all the repos downloaed by go get to build a single binary
	To change to a different project, just change the directory that GOPATH points to
	Comments should be above each function of method.
	Always start with function or method name so it is more understandable when
	documentation is generated by extracting these comments
















*/
