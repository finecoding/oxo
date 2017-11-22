# Notes on this Repository

On the local file system, the repository is finecoding and within it are two
sub repositories /webapp and /oxo

The full local path for oxo is:
**/home/paul/GoglandProjects/src/github.com/finecoding/oxo**

The full github url for oxo is:
**https://github.com/finecoding/oxo.git**

For any user to copy this repo into their own go workspace.

The go get command is:
**go get -x github.com/finecoding/oxo**

go get is directed by the $GOPATH environment variable.

$GOPATH is:
**/home/paul/GoglandProjects**

So **go get** conveniently clones the Github repo into $GOPATH/src/github.com/finecoding/oxo

**go get** is similar to **git clone**, but cloning the repo into the Golang workspace.
It also initialises git in this directory and tracks the remote on Github.
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



















