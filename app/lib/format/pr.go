package format

import (
	"fmt"
	"os"
	"text/tabwriter"

	"cli/app/bbclient"
)

func Prs(prs *[]bbclient.Pullrequest) error {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', tabwriter.Debug)
	_, err := fmt.Fprintln(w, fmt.Sprintf("%s\t  %s\t  %s\t  %s", "id", "repo", "source -> dest", "title"))
	for i := 0; i < len(*prs); i++ {
		pr := (*prs)[i]
		_, err := fmt.Fprintln(w, fmt.Sprintf("%v\t  %s\t  %s -> %s\t  %s", *pr.Id, *pr.Source.Repository.Name, *pr.Source.Branch.Name, *pr.Destination.Branch.Name, *pr.Title))
		if err != nil {
			return err
		}
	}
	err = w.Flush()
	if err != nil {
		return err
	}
	return nil
}
