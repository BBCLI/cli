package create

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
	Use:   "create",
	Short: "create comment on PR",
	Long:  "usage: bbc pr comment create <pr-id> [<path> <line>]",
	Args:  cobra.RangeArgs(1, 3),
	RunE: func(cmd *cobra.Command, args []string) error {
		prId, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		workspace, repo, err := git.GetGitRemoteDetails()
		if err != nil {
			return err
		}

		fmt.Print("enter comment content: ")
		in := bufio.NewReader(os.Stdin)
		content, err := in.ReadString('\n')
		if err != nil {
			return errors.New("error Reading your comment")
		}
		content = strings.TrimSpace(content)
		fmt.Println(content)
		if len(args) == 1 {
			err := createWithoutFileContext(workspace, repo, prId, &content)
			if err != nil {
				return err
			}
		}

		if len(args) == 3 {
			line, err := strconv.Atoi(args[2])
			if err != nil {
				return err
			}
			err = createWithFileContext(workspace, repo, prId, args[1], &line, &content)
			if err != nil {
				return err
			}
		}

		fmt.Println("comment created")
		return nil
	},
}

func createWithFileContext(workspace string, repo string, prId int, path string, line *int, content *string) error {
	res, err := bbclient.BbClient.PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsWithResponse(context.TODO(), workspace, repo, prId, bbclient.PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsJSONRequestBody{
		Content: &struct {
			// Html The user's content rendered as HTML.
			Html *string `json:"html,omitempty"`

			// Markup The type of markup language the raw content is to be interpreted in.
			Markup *bbclient.PullrequestCommentContentMarkup `json:"markup,omitempty"`

			// Raw The text as it was typed by a user.
			Raw *string `json:"raw,omitempty"`
		}{
			Raw: content,
		},
		Inline: &struct {
			// From The comment's anchor line in the old version of the file.
			From *int `json:"from,omitempty"`

			// Path The path of the file this comment is anchored to.
			Path string `json:"path"`

			// To The comment's anchor line in the new version of the file. If the 'from' line is also provided, this value will be removed.
			To *int `json:"to,omitempty"`
		}{
			To:   line,
			Path: path,
		},
	})
	if err != nil || res.StatusCode() != http.StatusOK {
		return err
	}
	return nil
}

func createWithoutFileContext(workspace string, repo string, prId int, content *string) error {
	res, err := bbclient.BbClient.PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsWithResponse(context.TODO(), workspace, repo, prId, bbclient.PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsJSONRequestBody{
		Content: &struct {
			// Html The user's content rendered as HTML.
			Html *string `json:"html,omitempty"`

			// Markup The type of markup language the raw content is to be interpreted in.
			Markup *bbclient.PullrequestCommentContentMarkup `json:"markup,omitempty"`

			// Raw The text as it was typed by a user.
			Raw *string `json:"raw,omitempty"`
		}{
			Raw: content,
		},
	})
	fmt.Printf("%s", string(res.Body))
	if err != nil || res.StatusCode() != http.StatusOK {
		return err
	}
	return nil
}
