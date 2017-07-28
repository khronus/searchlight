package check_node_exists

import (
	"fmt"

	"github.com/appscode/go/flags"
	"github.com/appscode/searchlight/pkg/icinga"
	"github.com/appscode/searchlight/pkg/util"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Request struct {
	Selector string
	Name     string
	Count    int
}

func CheckNodeExists(req *Request, isCountSet bool) (icinga.State, interface{}) {
	kubeClient, err := util.NewClient()
	if err != nil {
		return icinga.UNKNOWN, err
	}

	total_node := 0
	if req.Name != "" {
		node, err := kubeClient.Client.CoreV1().Nodes().Get(req.Name, metav1.GetOptions{})
		if err != nil {
			return icinga.UNKNOWN, err
		}
		if node != nil {
			total_node = 1
		}
	} else {
		nodeList, err := kubeClient.Client.CoreV1().Nodes().List(metav1.ListOptions{
			LabelSelector: req.Selector,
		},
		)
		if err != nil {
			return icinga.UNKNOWN, err
		}

		total_node = len(nodeList.Items)
	}

	if isCountSet {
		if req.Count != total_node {
			return icinga.CRITICAL, fmt.Sprintf("Found %d node(s) instead of %d", total_node, req.Count)
		} else {
			return icinga.OK, "Found all nodes"
		}
	} else {
		if total_node == 0 {
			return icinga.CRITICAL, "No node found"
		} else {
			return icinga.OK, fmt.Sprintf("Found %d node(s)", total_node)
		}
	}
}

func NewCmd() *cobra.Command {
	var req Request

	cmd := &cobra.Command{
		Use:     "check_node_exists",
		Short:   "Count Kubernetes Nodes",
		Example: "",

		Run: func(c *cobra.Command, args []string) {
			flags.EnsureRequiredFlags(c, "count")
			isCountSet := c.Flag("count").Changed
			icinga.Output(CheckNodeExists(&req, isCountSet))
		},
	}

	cmd.Flags().StringVarP(&req.Selector, "selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.")
	cmd.Flags().StringVarP(&req.Name, "name", "n", "", "Name of node whose existence is checked")
	cmd.Flags().IntVarP(&req.Count, "count", "c", 0, "Number of expected Kubernetes Node")
	return cmd
}
