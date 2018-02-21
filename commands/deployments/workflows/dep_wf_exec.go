package workflows

import (
	"fmt"
	"net/http"

	"path"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/ystia/yorc/commands/deployments"
	"github.com/ystia/yorc/commands/httputil"
)

func init() {
	var shouldStreamLogs bool
	var shouldStreamEvents bool
	var continueOnError bool
	var workflowName string
	var wfExecCmd = &cobra.Command{
		Use:     "execute <id>",
		Short:   "Trigger a custom workflow on deployment <id>",
		Aliases: []string{"exec"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.Errorf("Expecting an id (got %d parameters)", len(args))
			}
			client, err := httputil.GetClient()
			if err != nil {
				httputil.ErrExit(err)
			}
			if workflowName == "" {
				return errors.New("Missing mandatory \"workflow-name\" parameter")
			}
			url := fmt.Sprintf("/deployments/%s/workflows/%s", args[0], workflowName)
			if continueOnError {
				url = url + "?continueOnError"
			}
			request, err := client.NewRequest("POST", url, nil)
			if err != nil {
				httputil.ErrExit(err)
			}
			request.Header.Add("Content-Type", "application/json")
			response, err := client.Do(request)
			defer response.Body.Close()
			if err != nil {
				httputil.ErrExit(err)
			}
			ids := args[0] + "/" + workflowName
			httputil.HandleHTTPStatusCode(response, ids, "deployment/workflow", http.StatusAccepted, http.StatusCreated)

			fmt.Println("New task ", path.Base(response.Header.Get("Location")), " created to execute ", workflowName)
			if shouldStreamLogs && !shouldStreamEvents {
				deployments.StreamsLogs(client, args[0], !deployments.NoColor, false, false)
			} else if !shouldStreamLogs && shouldStreamEvents {
				deployments.StreamsEvents(client, args[0], !deployments.NoColor, false, false)
			} else if shouldStreamLogs && shouldStreamEvents {
				return errors.Errorf("You can't provide stream-events and stream-logs flags at same time")
			}
			return nil
		},
	}
	wfExecCmd.PersistentFlags().StringVarP(&workflowName, "workflow-name", "w", "", "The workflows name")
	wfExecCmd.PersistentFlags().BoolVarP(&continueOnError, "continue-on-error", "", false, "By default if an error occurs in a step of a workflow then other running steps are cancelled and the workflow is stopped. This flag allows to continue to the next steps even if an error occurs.")
	wfExecCmd.PersistentFlags().BoolVarP(&shouldStreamLogs, "stream-logs", "l", false, "Stream logs after triggering a workflow. In this mode logs can't be filtered, to use this feature see the \"log\" command.")
	wfExecCmd.PersistentFlags().BoolVarP(&shouldStreamEvents, "stream-events", "e", false, "Stream events after triggering a workflow.")
	workflowsCmd.AddCommand(wfExecCmd)
}
