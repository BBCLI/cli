package show

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"path"
	"strconv"

	"github.com/spf13/cobra"

	"cli/app/bbclient"
	"cli/app/lib/git"
)

type PRComment struct {
	Id            int
	Base          bool
	Comment       string
	Resolved      bool
	Path          *string
	Commenter     string
	ChildComments []*PRComment
}

var Cmd = &cobra.Command{
	Use:   "show",
	Short: "Show pull request details",
	Long:  "usage: bbc pr show <pullrequest-id>",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		prId, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		workspace, repo, err := git.GetGitRemoteDetails()
		if err != nil {
			return err
		}

		sort := "-created_on"
		response, err := bbclient.BbClient.GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesWithResponse(context.TODO(), workspace, repo, prId, &bbclient.GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesParams{
			Sort: &sort,
		})
		if err != nil || response.StatusCode() != http.StatusOK {
			return errors.New("error on bbc api")
		}

		latestCommits := *response.JSON200.Values
		if len(latestCommits) == 0 {
			return errors.New("no commit statuses found")
		}
		symbol := ""
		commit := latestCommits[0]
		commitState := *commit.State
		if commitState == bbclient.CommitstatusStateSUCCESSFUL {
			symbol = "✅ "
		}
		if commitState == bbclient.CommitstatusStateFAILED {
			symbol = "❌ "
		}
		if commitState == bbclient.CommitstatusStateINPROGRESS {
			symbol = "🚧 "
		}

		prResponse, err := bbclient.BbClient.GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdWithResponse(context.TODO(), workspace, repo, prId)
		if err != nil || prResponse.StatusCode() != http.StatusOK {
			return errors.New("error on bbc api")
		}

		commentResponse, err := bbclient.BbClient.GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsWithResponse(context.TODO(), workspace, repo, prId)
		if err != nil || commentResponse.StatusCode() != http.StatusOK {
			return errors.New("error on bbc api")
		}

		repoPath, err := git.GetGitAbsolutePath()
		if err != nil {
			return err
		}

		inlineComments := make(map[int]*PRComment)
		comments := make(map[int]*PRComment)
		for len(comments)+len(inlineComments) < len(*commentResponse.JSON200.Values) {
			for _, comment := range *commentResponse.JSON200.Values {
				if comment.Inline != nil {
					if comment.Parent != nil {
						if _, ok := inlineComments[*comment.Parent.Id]; ok {
							if _, ok := inlineComments[*comment.Id]; !ok {
								p := path.Join(repoPath, comment.Inline.Path) + fmt.Sprintf(":%v", *comment.Inline.To)
								inlineComments[*comment.Id] = &PRComment{
									Id:            *comment.Id,
									Base:          comment.Parent == nil,
									Comment:       *comment.Content.Raw,
									Resolved:      comment.Resolution != nil,
									Path:          &p,
									Commenter:     *comment.User.DisplayName,
									ChildComments: make([]*PRComment, 0),
								}
							}
							inlineComments[*comment.Parent.Id].ChildComments = append(inlineComments[*comment.Parent.Id].ChildComments, inlineComments[*comment.Id])
						} else {
							continue
						}
					}
					if _, ok := inlineComments[*comment.Id]; !ok {
						p := path.Join(repoPath, comment.Inline.Path) + fmt.Sprintf(":%v", *comment.Inline.To)
						inlineComments[*comment.Id] = &PRComment{
							Id:            *comment.Id,
							Base:          comment.Parent == nil,
							Comment:       *comment.Content.Raw,
							Resolved:      comment.Resolution != nil,
							Path:          &p,
							Commenter:     *comment.User.DisplayName,
							ChildComments: make([]*PRComment, 0),
						}
					}
				} else {
					if _, ok := comments[*comment.Id]; !ok {
						if comment.Parent != nil {
							if _, ok := comments[*comment.Parent.Id]; ok {
								if _, ok := comments[*comment.Id]; !ok {
									comments[*comment.Id] = &PRComment{
										Id:            *comment.Id,
										Base:          comment.Parent == nil,
										Comment:       *comment.Content.Raw,
										Resolved:      comment.Resolution != nil,
										Path:          nil,
										Commenter:     *comment.User.DisplayName,
										ChildComments: make([]*PRComment, 0),
									}
								}
								comments[*comment.Parent.Id].ChildComments = append(comments[*comment.Parent.Id].ChildComments, comments[*comment.Id])
							} else {
								continue
							}
						}
						if _, ok := comments[*comment.Id]; !ok {
							comments[*comment.Id] = &PRComment{
								Id:            *comment.Id,
								Base:          comment.Parent == nil,
								Comment:       *comment.Content.Raw,
								Resolved:      comment.Resolution != nil,
								Path:          nil,
								Commenter:     *comment.User.DisplayName,
								ChildComments: make([]*PRComment, 0),
							}
						}
					}
				}
			}
		}

		fmt.Printf("from %s to %s\n", *prResponse.JSON200.Source.Branch.Name, *prResponse.JSON200.Destination.Branch.Name)
		fmt.Printf("Title: %s\n", *prResponse.JSON200.Title)
		fmt.Printf("Created By: %s\n", *prResponse.JSON200.Author.DisplayName)
		if *prResponse.JSON200.Summary.Raw != "" {
			fmt.Printf("\n%s\n\n", *prResponse.JSON200.Summary.Raw)
		}

		fmt.Printf("%s: %s %s \n", *commit.Name, commitState, symbol)

		fmt.Printf("latest commit created at: %s \n\n", *prResponse.JSON200.UpdatedOn)

		fmt.Printf("Comments: \n")

		fmt.Println("General Comments: ")
		for _, comment := range comments {
			if comment.Base {
				printCommentOverview(*comment, 1)
			} else {
				continue
			}
			var status string
			if comment.Resolved {
				status = "resolved"
			} else {
				status = "pending"
			}
			fmt.Printf("  status: %s\n", status)
			fmt.Println("  -------------------------------------------------")
		}
		fmt.Println("Inline Comments: ")
		for _, comment := range inlineComments {
			if comment.Base {
				fmt.Printf("  file: %s\n", *comment.Path)
				printCommentOverview(*comment, 1)
			} else {
				continue
			}
			var status string
			if comment.Resolved {
				status = "resolved"
			} else {
				status = "pending"
			}
			fmt.Printf("  status: %s\n", status)
			fmt.Println("  -------------------------------------------------")
		}
		return nil
	},
}

func printCommentOverview(comment PRComment, level int) {
	for i := 0; i < level; i++ {
		fmt.Printf("  ")
	}
	fmt.Printf("%v - %s: %s\n", comment.Id, comment.Commenter, comment.Comment)
	for _, child := range comment.ChildComments {
		printCommentOverview(*child, level+1)
	}
}
