# Notes on this Repository

On the local file system, the hosting account on github is finecoding and within it are two repositories /webapp and /oxo

Go uses a convention for naming the file path of a package on the local filesystem.  It uses part of the Version Control System (VCS) identifier in the directory.  By this means it can derive the location of a package on a remote VCS system like Github from the local pathname and vice versa.  This makes finding code in public repository hosting services very easy.

The Golang workspace has /src directory and that contains a set of sub-directories that refer either to locally held repositories or remote repositories on Github, Bitbucket or some other VCS.  The directory names that hold repositories that are local copies of remote repositories are formed from the url of the remote VCS.

Taking Github as an example, all copies of repositories held at Github.com are stored locally in the Go workspace at /src/github.com.  That path is then extended to include the account name and one of the remote repositories held under that account on Github.com.

A repository can hold multiple packages and each package is associated with a single directory.  Here we have a repository called finecoding that contains two packages: oxo and webapp.  Packages can be further subdivided by extending the path with subdirectories.


## How Golang tooling finds packages

The full local path for oxo repository is:
**/home/paul/GoglandProjects/src/github.com/finecoding/oxo**

The full github url for oxo repository is:
**https://github.com/finecoding/oxo.git**

For any user to copy this repo into their own go workspace.

The go get command is:

**go get -x github.com/finecoding/oxo**

**go get** is directed to the Go workspace on the local filesystem by the $GOPATH environment variable.

$GOPATH is:
**/home/paul/GoglandProjects**

source code packages are held in the /src directory.

**go get 'repo url' ** conveniently clones the Github repo into /src of the GO workspace indicated by the $GOPATH environmental variable.

**go get -x github.com/finecoding/oxo**

*note the -x option gives a log of all the steps and git commands **go get** does in the background.

creates the following local copy:

$GOPATH/src/github.com/finecoding/oxo

**go get** is a wrapper for **git clone**, but oriented toward cloning the repo into the Golang workspace.

It also initialises git in this directory and tracks the remote on Github. (nb the git remote configuration has to be altered from a https:// to an ssh:// url prefix if we are accessing github using ssh and public key, see later)

**go get** also downloads the packages named by the import paths, along with their dependencies. It then installs the named packages, like 'go install'.  You can point **go get** at some source code you have in your $GOPATH and it will recursively fetch any required packages by the import statements in each file. You can even have it fetch all the source code by calling **go get import/path**.


So we can now develop in this repo locally and issue git commands to ensure it is version controlled.
We can then follow the typical workflow for git when working in this directory: add changed files to staging, commit with a message, then push to the remote repository.


**git add .**

**git commit**

 and then

 **git push -u origin master**

##  How to create a repository from scratch


To create a repository from scratch, log onto the Github web interface as account finecoding.  Then use create repo option to create a new repo.

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

## Using SSH and a public key to automatically login to Github

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

* note this document is in finecoding/oxo/notes.md

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

Lovely!  All working!

