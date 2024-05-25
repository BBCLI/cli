package reply

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"cli/app/bbclient"
	"cli/app/lib/git"
)

var Cmd = &cobra.Command{
	Use:   "reply",
	Short: "resolve comment on PR",
	Long:  "usage: bbc pr comment reply <pr-id> <comment-id>",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		prId, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		commentId, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}

		workspace, repo, err := git.GetGitRemoteDetails()
		if err != nil {
			return err
		}

		fmt.Print("enter reply content: ")
		in := bufio.NewReader(os.Stdin)
		content, err := in.ReadString('\n')
		if err != nil {
			return errors.New("error Reading your comment")
		}
		content = strings.TrimSpace(content)

		res, err := bbclient.BbClient.PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsWithResponse(context.TODO(), workspace, repo, prId, bbclient.PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsJSONRequestBody{
			Parent: &bbclient.Comment{
				Id: &commentId,
			},
			Content: &struct {
				// Html The user's content rendered as HTML.
				Html *string `json:"html,omitempty"`

				// Markup The type of markup language the raw content is to be interpreted in.
				Markup *bbclient.PullrequestCommentContentMarkup `json:"markup,omitempty"`

				// Raw The text as it was typed by a user.
				Raw *string `json:"raw,omitempty"`
			}{
				Raw: &content,
			},
		})
		if err != nil || res.StatusCode() != http.StatusOK {
			return err
		}

		fmt.Println("reply comment successfully")
		return nil
	},
}
