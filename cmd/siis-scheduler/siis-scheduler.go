package siis_scheduler

import (
	. "flag"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/sushshring/kubernetes/cmd/siis-scheduler/app"
	"github.com/sushshring/kubernetes/pkg/kubectl/util/logs"
	. "k8s.io/apiserver/pkg/util/flag"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	command := app.NewSiisSchedulerCommand()

	pflag.CommandLine.SetNormalizeFunc(WordSepNormalizeFunc)
	pflag.CommandLine.AddGoFlagSet(CommandLine)
	logs.InitLogs()
	defer logs.FlushLogs()

	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
