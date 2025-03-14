package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"time"

	"demo/apps/myapp"

	"ergo.services/logger/colored"
	"ergo.services/logger/rotate"

	"ergo.services/ergo"
	"ergo.services/ergo/gen"
	"ergo.services/ergo/lib"
)

var (
	OptionNodeName   string
	OptionNodeCookie string
)

func init() {
	flag.StringVar(&OptionNodeName, "name", "demo@localhost", "node name")
	flag.StringVar(&OptionNodeCookie, "cookie", lib.RandomString(16), "a secret cookie for the network messaging")
}

func main() {
	var options gen.NodeOptions

	flag.Parse()

	// create applications that must be started
	apps := []gen.ApplicationBehavior{
		myapp.CreateMyApp(),
	}
	options.Applications = apps

	// disable default logger to get rid of multiple logging to the os.Stdout
	options.Log.DefaultLogger.Disable = true

	// add logger "colored".
	loggercolored, err := colored.CreateLogger(colored.Options{TimeFormat: time.DateTime})
	if err != nil {
		panic(err)
	}
	options.Log.Loggers = append(options.Log.Loggers, gen.Logger{Name: "colored", Logger: loggercolored})

	// add logger "rotate".
	loggerrotate, err := rotate.CreateLogger(rotate.Options{TimeFormat: time.DateTime})
	if err != nil {
		panic(err)
	}
	options.Log.Loggers = append(options.Log.Loggers, gen.Logger{Name: "rotate", Logger: loggerrotate})

	// set network options
	options.Network.Cookie = OptionNodeCookie

	// starting node
	node, err := ergo.StartNode(gen.Atom(OptionNodeName), options)
	if err != nil {
		panic(err)
	}

	// register network messages
	//if err := node.Network().RegisterMessage(demo.Messages...); err != nil {
	//	panic(err)
	//}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	node.Log().Info("terminating... recived signal %v ", <-sigCh)
	node.Stop()
	//node.Wait()
}
