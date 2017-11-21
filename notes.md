# Notes on this Repository

The full local path is:
**/home/paul/GoglandProjects/src/github.com/finecoding/oxo**

The full github url is:
**https://github.com/finecoding/oxo.git**

The go get command is:
**go get -x github.com/finecoding/oxo**

$GOPATH is:
**/home/paul/GoglandProjects**

So **go get** conveniently clones the Github repo into $GOPATH/src/github.com/finecoding/oxo

Similar to git clone, but cloning the repo into the Golang workspace.  It also initialises git in this directory and tracks the remote on Github.  So we can now develop in this repo locally and issue git commands to ensure it is version controlled.  We can **git add**, **git commit** and then **git push -u origin masterg**, follow a Git workflow.

An empty repo was created within finecoding on github using the git web interface and then cloned locally using **go get**

To do:  get git working without a username and password.

Next: structure the repo for commands and libraries, make an example of each.













