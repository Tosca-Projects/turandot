package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	puccinicommon "github.com/tliron/puccini/common"
	"github.com/tliron/puccini/common/terminal"
)

func init() {
	serviceCommand.AddCommand(serviceOutputCommand)
}

var serviceOutputCommand = &cobra.Command{
	Use:   "output [SERVICE NAME] [OUTPUT NAME]",
	Short: "Get a deployed service's output",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		serviceName := args[0]
		outputName := args[1]
		ServiceOutput(serviceName, outputName)
	},
}

func ServiceOutput(serviceName string, outputName string) {
	// TODO: in cluster mode we must specify the namespace
	namespace := ""

	service, err := NewClient().Turandot().GetService(namespace, serviceName)
	puccinicommon.FailOnError(err)

	if service.Status.Outputs != nil {
		if output, ok := service.Status.Outputs[outputName]; ok {
			// TODO: unpack the YAML
			// TODO: support output in various formats
			fmt.Fprintln(terminal.Stdout, output)
			return
		}
	}

	puccinicommon.Failf("output %q not found in service %q", outputName, serviceName)
}
