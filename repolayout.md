# Layout of a Repository


Lets look at the structure of the repository.

The finecoding user on the Github repository hosting service has several repos.  Here are two that contain Golang code:

- finecoding/oxo          
- finecoding/webapp    

Each is has a few files in the repository root and there are no sub-directories.  They are almost empty.  However, as code is developed more files will be created and we need to think about how they should be organised.  A package is simply a directory containing files within a repository.

Packages themselves can have structure, they can be divided with nested sub-directories.  Maybe give it a logical structure where each package groups files that are used together to implement some functional purpose?

What kinds of shared function or purpose can we identify?

The Golang tools like **go get**, **go build** and **go install** are package oriented and they work on package names, they also recognise sub-directories within a package.

Golang distinguishes between two kinds of Golang file. 'Command' files, where the package statement uses the reserved word **main** at the beginning of the a file to signify that the package is intended to create an executable file, a binary.  Such command files cannot be imported into other programs and used as a code library.  The name of the executable is automatically inferred from the directory name in which the file is held.

The other kind of Golang file are 'Packages', intended to imported by other golang programs that will have access to all the package types, methods and functions.  But first let us look at Command  files, whose first line contains 'package main'.

A command file may share its directory with many other command files.  Because they all share the same directory, they can access each others types, methods and functions.   But which one of these files is the main calling file that calls the rest?

Also, what if we have several executable programs we wish to create?  If there are several command files which one should be compiled first and derive the name of the executable from the containing directory?  It is not obvious on first inspection if there are several command files, all with package main, sharing the same directory.  There can be only that calls the rest and compiles to an executable.  If we want a second executable, it and its supporting files should be in different directory so that on compilation it will take its own executable name from the name of the containing directory.  If we have a package that contains many command files, it makes sense to create a sub directory for each, so all the executables get different names.  Seperating them in this way also means the go tools can work on just part of a the package, rather than it all.  Useful if we wanted to compile just one executeable from the many that might  be in the package.

It is common for such command programs and their supporting gofiles to be grouped under individual subdirectories under /cmd, so we can easily find the executables.  We add a subdirectory for the files associated with each executable, and we give it a unique name.

/cmd/prog1/ (sub directory contains a command file any any supporting files for executable prog1)
/cmd/prog2/
/cmd/prog3/
/cmd/prog4/

Each of these command subdirectories may contain several golang files, all with package main at the beginning.  Which one should be compiled to create the executable.  Which one calls the rest to use their types, methods and functions?

A common practice is to distinguish the command file by calling the file main.go.  Some programmers consider it more logical to give the command file same name as the containing directory and thus the name of the resulting executable.

 So here we have the beginnings of package structure, at least as far as command files are concerned:

finecoding/oxo/cmd/progA/main.go + supporting files
finecoding/oxo/cmd/progB/main.go
finecoding/oxo/cmd/progC/main.go
finecoding/oxo/cmd/progD/main.go

or alternatively:

finecoding/oxo/cmd/progA/progA.go + supporting files
finecoding/oxo/cmd/progB/progB.go
finecoding/oxo/cmd/progC/progC.go
finecoding/oxo/cmd/progD/progD.go



## Package structure for a Command Line Interface



Consider the case, where the program being developed is a command line interface. Where an executable is a command line program, which takes some arguments after the program name, each of which may be handled by a seperate executable.

Here we have a command line program (cli) which takes several different arguments (create,review,update and delete), each of which run a seperate sub program.  This results in this kind of directory structure.

<!--From tree -L 5 -a -F finecoding/-->

```
finecoding/
└── oxo/
    └── cmd/
        ├── cli/
        │   ├── cli.go
        │   ├── create/
        │   │   └── create.go
        │   ├── delete/
        │   │   └── delete.go
        │   ├── review/
        │   │   └── review.go
        │   └── update/
        │       └── update.go
        └── server/

```

Here we a command file for each of the cli actions (create.go, delete.go, review.go and update.go).  They have been given the same name as their enclosing directory, but this is not required, they could all have been called main.go or any other name.  This name was chosen to distinguish the command program from any other supporting programs.  There is also a cli.go command program to create the executable for the command line program itself.  It is common to see lots of main.go programs in this kind of directory structure, but this is more consistent.  

This package design allows go tooling to operate on either the cli program and everything below it or the tooling build just part of the package, (eg the create program.)

`go build cmd/cli`

or

`go build cmd/cli/create` 

Some programmers opt for simpler approach for command line programs. They put all program files at the root of the repo, no sub-directories.  Each file starts with package main.  There can be only one command file to give a name to executable.  The subprograms cannot be called in isolation, which makes testing more difficult.

Many popular CLI frameworks introduce their own opinions about package structure.

## Other directories recognised by Golang tooling 

`/internal` for go code that should not be imported by other packages.

and

`/vendor` for copies of third party packages and their dependencies.

There is also a convention to keep the second common type of golang file, a package or library intended for import by other programs, in a dedicated directory.   While any subdirectory in a repository is a package, only those that contain golang files that start with the statement `package "packagename"` can be imported so that the types and functions they contain can be used in the importing program. 

Packages to be imported are specified in an import statement at the beginning of a Golang file and they lend their name whenever the imported functions and types are used in the calling program. The package name will be used as part of the names of functions and types all over the code.  Consequently the naming of packages should be considered carefully.  Lower case, not very long, not too general.  They should be designed from the perspective of the consumer, the programmer who will use them in their program.  Choose package names that lead to self descriptive code when they are used.  It is a convention that importable packages are all contained in the same top level directory in the repository.  For each package file, the `package "packagename"` statement must match the directory name.

Packages are often kept seperately and a common choice for a top level directory name to hold them within a repository is /pkg (though this name is also used within the golang workspace (/pkg, /src and /bin).  Sometimes it is called /package or /lib.

So, taking the clues from Golang itself the top level repository layout looks like this:

finecoding  /cmd
            /pkg
            /internal
            /vendor

The vendor directory should be reproduced for every level of subdirectory of /pkg

### More about /vendor and dependencies

What if the package is a very small, just enough to create a binary executable?

In that case all the files could be put in the root of the package with no sub-directories

To understand how a package is working, the first task is the 'hunt the command go file'.  It may be at the root of the package or buried in inside a subdirectory, it may be called main.go or it might take the name of the containing directory.

Once located the main command go file should, ideally, have little in it because it is difficult to test.  It should just do a bit of initialisation and kick off some other function from the import list.   If the program is based in a CLI framework, it is best to understand how they work to be able to follow the logic further.  Hint: study cobra.

## Files found in the root of a repository


The root of a package is often littered with odd files from:
travis and circleci Continuous integration frameworks
go dep or one of the other dependency managers.

Lots of yaml and toml files from build tools, markdown md files for documentation, gitignore files.

https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2

Javascript can live in a client or ui subdirectory. Dockerfiles, continuous integration configs, or other build helpers can live in the project root or in a build subdirectory. And runtime configuration like Kubernetes manifests can have a home, too.

/client
/build

Keep test code in the same directory as the code it tests.

It is possible to create abstractions at the package level for domain objects like User, Account, Product, etc.  However one of the areas where this is a clear advantage is in abstracting the database access by creating wrap types especially for this purpose.  We can also use interfaces.  These could be grouped in a /app directory.

***need more research in this area***

## Repository layout for servers

Now what if the one of our binaries is not a command line program, but a server?  It may well be able to take some arguments, servers also have a structure that is often related to functional components.  The MVC structure is popular because it seperates the backend database access (Model) from the presentation logic running on the client (View) and the two are linked by the (Controller) logic in the server.

Web servers may also have static content for holding images.

**Need a discussion for each type of application framework  CLI, Webapp, etc.**

**Package by functional layer - eg MVC**
**Package by feature - eg. DDD**

How is a programmer supposed to find their way around a package?

A CLI has an fairly simple structure and it is easy to create a package stucture the correlates with how it functions. But when it comes to servers, there may be several functional layers and this can obscure how they actually work at a higher level.  That higher level is the Domain and in order to understand what a program does it is important to understand the Domain.  High level constructs like a User, Service, Account and how they relate to each other.  Strictly functional layers consider databases, sessions, cookies, etc, used in isolation.  You can understand what it does at a low level, but not why it is being used because that is expressed in the language of the domain.


What sort of package structure would work for a simple web server?

