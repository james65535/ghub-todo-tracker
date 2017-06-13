package clients

type IssueClient interface{
	SetIssue(title, body string)
	SubmitIssue()(error)
}
