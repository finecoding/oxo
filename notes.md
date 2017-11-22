# Notes on this Repository

On the local file system, the hosting account on github is finecoding and within it are two
repositories /webapp and /oxo

Go uses a convention for naming the file path of a package on the local filesystem such that it uses part of the Version Control System (VCS) identifier.  In this way, it can derive the location of a package on a remote VCS system like Github from the local pathname and vice versa.  The Golang workspace has /src directory and that contains a set of sub-directories that refer either to locally held repositories or remote repositories on Github, Bitbucket or some other VCS.  The directories that hold repositories that are local copies of remote repositories contain the url of the remote VCS.  All copies of repositories held at Github.com are in /src/github.com.  That path is then followed by the account name and one of the remote repositories held under that account on Github.com.  A package is associated with a single directory.  A repository can have multiple packages, each held in a different sub directory.  Here we have a repository called finecoding that contains two packages: oxo and webapp.  Those packages can also be subdivided to create a structure by extending the path remotely on github.com as long as it maps to a similarly named sub directory on the local filesystem path.


This is how Golang tooling finds packages.

The full local path for oxo repository is:
**/home/paul/GoglandProjects/src/github.com/finecoding/oxo**

The full github url for oxo repository is:
**https://github.com/finecoding/oxo.git**

For any user to copy this repo into their own go workspace.

The go get command is:
**go get -x github.com/finecoding/oxo**

go get is directed by the $GOPATH environment variable.

$GOPATH is:
**/home/paul/GoglandProjects**

So **go get** conveniently clones the Github repo into the local Go into /src of the GO workspace indicated ny the $GOPATH environmetal variable

$GOPATH/src/github.com/finecoding/oxo

**go get** is a wrapper for **git clone**, but cloning the repo into the Golang workspace.
It also initialises git in this directory and tracks the remote on Github. (nb the git remote configuration has to be altered from a https:// to an ssh:// url prefix if we are accessing github using ssh and public key, see later)

**go get** downloads the packages named by the import paths, along with their dependencies. It then installs the named packages, like 'go install'.  You can point **go get** at some source code you have in your $GOPATH and it will recursively fetch any required packages by the import statements in each file. You can even have it fetch all the source code by calling **go get import/path**.


So we can now develop in this repo locally and issue git commands to ensure it is version controlled.
We can then follow the typical workflow for git when working in this directory **git add**, **git commit** and then **git push -u origin master**.

The first step was to log into git as finecoding on github and use the option to create a new repo.
Then from the local filesystem command line **go get** the repo.

One problem with this:  git asks for a username and password following the git push -u origin master.

How to get it Git authentication working automatically using ssh and a public key rather than username and password.

In ~/.ssh ensure a public and private key are generated.  id_rsa is the default name for a public key.

When the git client does a push, it logs into github by submitting this key.  If the github account has this in its list of know keys, then the git client user is authenticated.  If not, it asks for a username and password.

To get the client to submit its public key, we need the ssh-agent running on the client computer

eval "ssh-agent -s"

we add the public key to the agent:

ssh-add ~/.ssh/id_rsa

Then we need to ensure the git client is correctly configured to track the remote repo on github.

We can pick up the remote repo url to use from the copy/clone option on the github admin screen

git remote add origin https://github.com/finecoding/testweb.git

However, the githhub client was already set up by go get and

if it says "fatal: remote origin already exists" then we can remove the tracking of the remote repo

git remote rm origin

then try again:

git remote add origin https://github.com/finecoding/testweb.git


But---- that does not work with ssh, I still have to enter a password ---

Try again, this works.

within dir
/home/paul/GoglandProjects/src/github.com/finecoding/testweb

**git remote rm origin**

**git remote add origin ssh://git@github.com/finecoding/testweb.git**

**git push -u origin master**

Branch master set up to track remote branch master from origin.
Everything up-to-date

We now do the same for our second sub repo, oxo

change to dir
/home/paul/GoglandProjects/src/github.com/finecoding/oxo

**git remote rm origin**

**git remote add origin ssh://git@github.com/finecoding/oxo.git**

**git push -u origin master**

Branch master set up to track remote branch master from origin.
Everything up-to-date

We can check either of these by changing to finecoding/testweb or finecoding/oxo and

**git remote -v**

Note the important change here is in the url used to access github, it is now begins ssh:// instead of https://

Another useful flag for troubleshooting is the **go get -x** which tells you all the steps it is taking in the background.

This document is in finecoding/oxo/notes.md

**git add .**

**git commit -m "tidied up explanation text"**

[master cacd7d9] tidied up explanation text
 1 file changed, 45 insertions(+), 10 deletions(-)
**git push -u origin master**

Counting objects: 3, done.
Delta compression using up to 2 threads.
Compressing objects: 100% (3/3), done.
Writing objects: 100% (3/3), 1.60 KiB | 0 bytes/s, done.
Total 3 (delta 0), reused 0 (delta 0)
To ssh://github.com/finecoding/oxo.git
   c609c58..cacd7d9  master -> master
Branch master set up to track remote branch master from origin.

Lovely!  All works.

What next?  Lets look at the structure of the repo.

The finecoding user on Github has several repositories, here are two:

finecoding/oxo
finecoding/webapp

Each is has a few files in the repository root and there are no sub-directories.

How should we organise the a repository with sub-directories so that is has a logical structure related to its function?

What kinds of function can we identify?  Firstly Golang itself has a package structure and we can map that to sub-directories within the repository.

Golang has 'commands' where the package statement uses the reserved word main at the beginning of the a file to signify that the package is intended to create an executable file, a binary, it cannot be imported into other programs.  The name of the executable is automatically inferred from the directory name in which the file is held.  There may be other files in the directory as well as the main program????? It is common for such Golang programs to be grouped under /cmd, so we can easily find the executables.  Then to add a subdirectory name for each seperate executable, so we can give it a name.  If we also distinguish the program containing package main by calling the file main.go we have a package structure that maps to a directory structure where there is one sub-directory of /cmd for each executable program.

finecoding/oxo/cmd/programA/main.go
finecoding/oxo/cmd/programB/main.go
finecoding/oxo/cmd/programC/main.go
finecoding/oxo/cmd/programD/main.go

Where an executable is, in fact a command line program, which takes some arugments after the program name, then we might extend the structure to isolate the programs associated with each 'command'.  We have a clash of nomenclature here.  A command line program can one of several commands, each of which relate to a seperate program...

Here we have a command line program called Program A which takes several different arguments, each of which run a seperate sub program.

finecoding/oxo/cmd/ProgramA/cmd/SubProgram1
finecoding/oxo/cmd/ProgramA/cmd/SubProgram2
finecoding/oxo/cmd/ProgramA/cmd/SubProgram3
finecoding/oxo/cmd/ProgramA/cmd/SubProgram4

Some packages opt for simpler approach for command line programs. They put all program files at the root of the repo, no sub-directories.  Each file starts with package main.

Golang tooling also recognises two other directories within a package:

/internal for go code that should not be imported by other packages.

and

/vendor for copies of third party packages and their dependencies.

There is a convention to keep packages intended for import by other programs in a dedicated directory.

Often this is called /pkg (though this name is also used within the golang workspace (/pkg, /src and /bin)

finecoding  /cmd
            /pkg
            /internal
            /vendor

Now the vendor program should be reproduced for every subdirectory of /cmd

### More about /vendor and dependencies

What if the package is a very small, just enough to create a binary executable?

In that case all the files could be put in the root of the package with no sub-directories

There could be several programs within the package.  Each would start with package main.  But we would need to identify which go file calls the code in the rest of the files in the package  This is the command file.  Which one is it? One common way to identify the command go file is to name it main.go.  When this is compiled, it will make use of other code in the same directory without needing any import statements.  When the program is compiled it will create an executable with the same name as the containing directory.  Sometimes, instead of main.go we use a go file with the same name as the containing directory.

A more complicated package could have subdirectories.  One for /cmd then a subdirectory corresponding with the name of the executable containing a main.go.  If it is a CLI program, this subdirectory will sit alongside subdirectories that correspond with each of the command line arguments.

So to understand how a package is working, the first trask is the 'hunt the command go file'.  It may be at the root of the package or buried in inside a subdirectory, it may be called main.go or it might take the name of the containing directory.

Once located the main command go file should, ideally, have little in it because it is difficult to test.  It should just do a bit of initialisation and kick off some other function from the import list.   If the program is based in a CLI framework, it is best to understand how they work to be able to follow the logic further.  Hint: study cobra.

The root of a package is often littered with odd files from:
travis and circleci Continuous integration frameworks
go dep or one of the other dependency managers.

Lots of yaml and toml files from build tools, markdown md files for documentation, gitignore files.


Javascript can live in a client or ui subdirectory. Dockerfiles, continuous integration configs, or other build helpers can live in the project root or in a build subdirectory. And runtime configuration like Kubernetes manifests can have a home, too.

/client
/build

Keep test code in the same directory as the code it tests.

It is possible to create abstractions at the package level for domain objects like User, Account, Product, etc.  However one of the areas where this is a clear advantage is in abstracting the database access by creating wrap types especially for this purpose.  We can also use interfaces.  These could be grouped in a /app directory.

***need more research in this area***



Now what if the one of our binaries is not a command line program, but a server?  It may well be able to take some arguments, servers also have a structure that is often related to functional components.  The MVC structure is popular because it seperates the backend database access (Model) from the presentation logic running on the client (View) and the two are linked by the (Controller) logic in the server.

Web servers may also have static content for holding images.

**Need a discussion for each type of application framework  CLI, Webapp, etc.**

**Package by functional layer - eg MVC**
**Package by feature - eg. DDD**

How is a programmer supposed to find their way around a package?

A CLI has an fairly simple struture and it is easy to create a package stucture the correlates with how it functions.

But when it comes to servers, there may be several functional layers and this can obscure how they actually work at a higher level.  That higher level is the Domain and in order to understand what a program does it is important to understand the Domain.  High level constructs like a User, Service, Account and how they relate to each other.  Strictly functional layers consider databases, sessions, cookies, etc, used in isolation.  You can understand what it does at a low level, but not why it is being used because that is expressed in the language of the domain.









Golang also has 'code libraries' generally known as packages that are intended to be imported into other programs to make available collections of types, methods and functions.  A package is distinguished by a program starting with the package keyword followed by the name of a directory that contains the files associated with the package.  Golang associates the directory name with the package name.



















