# go-cname
A CNAME discovery script written in Go

Many thanks to

    https://github.com/majek/goplayground/tree/master/resolve

Modified to parse CNAMEs instead of A records

# To Run

First, make sure you have Go installed! The easiest way to do that is to use Homebrew:

	brew install go

Then, set your GOPATH

    export GOPATH="/your/go/workspace"

Next, get this project

    go get github.com/hwyfour/go-cname

Now, you're able to use the script

    $ echo "m.allenedmonds.com" | $GOPATH/bin/go-cname
	> m.allenedmonds.com, allenedmonds.production.echocore.org.
