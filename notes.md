# Notes on this Repository

There is a hosting account at  github.com called finecoding that holds several repositories.  Two that contain Golang files are /webapp and /oxo.

Golang uses a convention for naming a file path for a package on the local filesystem that links it, to the fully qualfied name of a repository on a public repository hosting service like github.   It makes the link implicitly by using part of the github Version Control System (VCS) identifier in the directory name.  By this means it can derive the location of a package on a Github or a range of other common version control system hosting services from the local filepath name and vice versa.  This linkage by convention makes finding code in public repository hosting services very easy to find and also copy to create a local version on any local filesytem.

The Golang workspace has three directories from the root: /src, /pkg and /bin.  This discussion focusses on source code is stored below the /src directory which contains a set of sub-directories that refer either to locally held repositories or remote repositories on Github, Bitbucket or some other VCS.  The directory names under .src that hold copies of remote repositories can be recognised because the are formed from the url of the repository on the remote VCS.

Taking Github as an example, all copies of repositories held at Github.com are stored locally in the Go workspace at /src/github.com.  That path is then extended to include the account name and one of the remote repositories held under that account on Github.com.  In this case there are two repositories on github at /src/github.com/finecoding/oxo and /src/github.com/finecoding/webapp.

A repository can hold multiple packages and each package is associated with a single directory.  Here we have a hosting account called called finecoding that contains holds two repositories oxo and webapp.  Each repository can be be further subdivided into subdirectories and each subdirectory within repository is a package.  The files within a package are used for code sharing.  The package name identifies that directory of shared code both on remote public repository hosting systems and also locally.  One can be determined from the other.

![gitdiag1](/home/paul/Documents/gitdiag1.svg)

## How Golang tooling finds packages

The full local filepath for oxo repository is:
**/home/paul/GoglandProjects/src/github.com/finecoding/oxo**

The full github url for oxo repository is:
**https://github.com/finecoding/oxo.git**

For any user to copy this repo into their own go workspace.

The go get command is:

**go get -x github.com/finecoding/oxo**

**go get** is directed to the Go workspace on the local filesystem by the $GOPATH environment variable.

$GOPATH is:
**/home/paul/GoglandProjects**

source code packages are held in the /src directory and go tools store packages with reference to /src.  A common mistake to avoid is to use /src in the name of the remote repository or it may appear twice in the local file system repository path once it is copied.

**go get 'repo url' ** conveniently clones the Github repo into /src of the GO workspace indicated by the $GOPATH environmental variable.

**go get -x github.com/finecoding/oxo**

> note the -x option gives a log of all the steps and git commands **go get** does in the background.
>

it creates the following local copy:

$GOPATH/src/github.com/finecoding/oxo

**go get** is a wrapper for **git clone**, but oriented toward cloning the repo into the Golang workspace.  It also does a few other things as well and the -x option will show the detail.

It initialises git in this directory and tracks the remote version on Github. (Note the git remote configuration has to be altered from a default https:// to an ssh:// url prefix if we are accessing github using ssh and public key, this will be discussed later)

**go get** also downloads the packages named by the import paths within the go files it finds in the package, along with any that those packages themselves depend on. 

It then installs the named packages, like 'go install'.  You can point **go get** at some source code you have in your $GOPATH and it will recursively fetch any required packages by the import statements in each file. You can even have it fetch all the source code by calling **go get import/path**.

So we can now develop in this repo locally and issue git commands to ensure it is version controlled.

We can then follow the typical workflow for git when working in this directory: add changed files to staging, commit with a message, then push to the remote repository.


**git add .**

**git commit**

 and then

 **git push -u origin master**

##  How to create a repository from scratch


To create a repository for the first time, log onto the Github using the web interface and account finecoding.  Then use the 'create repo' option to create a new repo on Github.  Once created (within an optional readme.md file), the fully qualified repo name is available for copying or cloning.

To copy it to the local filesystem, from the local terminal  **go get** the repo.

> Caveat:  git asks for a username and password following the git push -u origin master.



---



How to get it Git authentication working automatically using ssh and a public key rather than username and password.

In ~/.ssh ensure a public and private key are generated.  id_rsa is the default name for a public key.

When the git client does a push, it logs into github by submitting this key.  If the github account has this key in its list of know keys, the git client user is authenticated.  If not, it defaults to authentication by asking for a username and password.

To get the client to submit its public key, we need the ssh-agent running on the client computer

eval "ssh-agent -s"

we add the public key to the agent:

ssh-add ~/.ssh/id_rsa

This process can have a few complications that can be covered in a seperate document.  For now we assume the public key used by the git client has been copied to Github and stored in its list of know keys.  Basically the git client and git server (github) need to know each others public keys. The git client will learn githubs public key when it first connects using ssh.   The authenication by public key exchange is complete when there are no more username and password challenges by github.   

Next we need to ensure the git client is correctly configured to track the remote repo on github.

We can pick up the remote repo url to use from the copy/clone option on the github admin screen

git remote add origin https://github.com/finecoding/testweb.git

However, the githhub client was already set up by go get and

if it says "fatal: remote origin already exists" then we can remove the tracking of the remote repo

git remote rm origin

then try again:

git remote add origin https://github.com/finecoding/testweb.git


But---- that does not work with ssh, I still have to enter a password ---

### Using SSH and a public key to automatically login to Github

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

**Lovely!  All working!**

