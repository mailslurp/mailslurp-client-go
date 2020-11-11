
# Go Email API - Documentation

> Create and manage email addresses in GoLang. Send and receive emails and attachments in code and tests.

## Quick links

- [GitHub Repository](https://www.github.com/mailslurp/mailslurp-client-go/)
- [Full method documentation](https://www.github.com/mailslurp/mailslurp-client-go/)

## Get started

MailSlurp is an email API that let's your create email addresses on demand then send and receive emails in code and tests. **No MailServer is required**.

::: tip
This section describes how to get up and running with the Go client. To use another language or the REST API see the [developer page](https://www.mailslurp.com/developers/).

See the examples page for [code examples and use with common frameworks](https://www.mailslurp.com/examples/).

See the method documentation for a [list of all functions](https://www.github.com/mailslurp/mailslurp-client-go/).
:::

## Get an API Key

You need a free MailSlurp account to use the service. [Sign up for a free account](https://app.mailslurp.com/sign-up/) first.

## Install

```bash
go get github.com/mailslurp/mailslurp-client-go
```

You may also need these dependencies:

```bash
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

MailSlurp's Go library uses `/antihax/optional` to wrap many parameters. Example usage looks like this:

```go
waitOpts := &mailslurp.WaitForLatestEmailOpts{
    InboxId: optional.NewInterface(inbox2.Id),
    Timeout: optional.NewInt64(30000),
    UnreadOnly: optional.NewBool(true),
}
```

## Initialize

To initialize the MailSlurp client first [get an API Key](https://app.mailslurp.com). Then import the libraries:

```go
import (
	"context"
	"github.com/mailslurp/mailslurp-client-go"
)
```

Then configure the client:

```go
var apiKey = "your-mailslurp-client"

// create a context with your api key
ctx := context.WithValue(
    context.Background(),
    mailslurp.ContextAPIKey,
    mailslurp.APIKey{ Key: apiKey }
)

// create mailslurp client
config := mailslurp.NewConfiguration()
client := mailslurp.NewAPIClient(config)
```

Controllers are then available on the client instance and can be passed your context and options:

```go
inbox, response, err := client.InboxControllerApi.CreateInbox(ctx, &mailslurp.CreateInboxOpts{})
```

See the repository for [list of all methods](https://www.github.com/mailslurp/mailslurp-client-go/).

## Common use cases

### Create an inbox

You can create email addresses in MailSlurp by creating an inbox:

```go
opts := &mailslurp.CreateInboxOpts{}
inbox, response, err := client.InboxControllerApi.CreateInbox(ctx, opts)
```

The email address and ID are accessible on the Inbox:

```go
assert.NotNil(t, inbox.Id)
assert.Contains(t, inbox.EmailAddress, "@mailslurp.com")
```

### Send emails

To send emails first create an inbox then use the `SendEmailOptions` to configure the email.

```go
sendEmailOptions := mailslurp.SendEmailOptions{
    To: []string{inbox.EmailAddress},
    Subject: "Test email",
    Body: "<h1>MailSlurp supports HTML</h1>",
    IsHTML: true,
}
opts := &mailslurp.SendEmailOpts{
    SendEmailOptions: optional.NewInterface(sendEmailOptions),
}
res, err := client.InboxControllerApi.SendEmail(ctx, inbox.Id, opts)
```

### Receive emails

You can read emails in Go using the `WaitForControllerApi`:

```go
// fetch the email for inbox2
waitOpts := &mailslurp.WaitForLatestEmailOpts{
    InboxId: optional.NewInterface(inbox2.Id),
    Timeout: optional.NewInt64(30000),
    UnreadOnly: optional.NewBool(true),
}
email, response, err := client.WaitForControllerApi.WaitForLatestEmail(ctx, waitOpts)
```

You can extract the content of an email for use in your tests or application using regex and the email body:

```go
// can extract the contents
r := regexp.MustCompile(`Your code is: ([0-9]{3})`)
code := r.FindStringSubmatch(email.Body)[1]
```

## Next step

For more information see the [source code on GitHub](https://github.com/mailslurp/mailslurp-client-go/).
