# Go Email API - Documentation

> Create and manage email addresses in GoLang. Send and receive emails and attachments in code and tests.

## Quick links

- [GitHub Repository](https://www.github.com/mailslurp/mailslurp-client-go/)
- [Full method documentation](./docs)
- [Examples](https://www.mailslurp.com/examples/)
- [Guides](https://www.mailslurp.com/guides/)

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

## Example

```go
import (
    "context"
    "github.com/antihax/optional"
    mailslurp "github.com/mailslurp/mailslurp-client-go"
    "github.com/stretchr/testify/assert"
    "os"
    "testing"
    "encoding/base64"
)

// create a client and a context 
func createClient(t *testing.T) (*mailslurp.APIClient, context.Context) {
    // read mailslurp api key from environment variable
    var apiKey = os.Getenv("API_KEY")
    assert.NotEmpty(t, apiKey, "API KEY is empty")

    // create a context with your api key
    ctx := context.WithValue(
        context.Background(),
        mailslurp.ContextAPIKey,
        mailslurp.APIKey{Key: apiKey},
    )

    // create mailslurp client
    config := mailslurp.NewConfiguration()
    return mailslurp.NewAPIClient(config), ctx
}

func Test_CanCreateInbox(t *testing.T) {
    client, ctx := createClient(t)

    // call inbox controller to create an inbox
    opts := &mailslurp.CreateInboxOpts{}
    inbox, response, err := client.InboxControllerApi.CreateInbox(ctx, opts)

    assert.Nil(t, err)
    assert.Equal(t, response.StatusCode, 201)
    assert.NotNil(t, inbox.Id)
    assert.Contains(t, inbox.EmailAddress, "@mailslurp.com")
}
```

See the [Method Documentation](./docs) for more information or see the use cases below.

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

Here is another example:

```go
sendOpts := &mailslurp.SendEmailOpts{
    SendEmailOptions: optional.NewInterface(mailslurp.SendEmailOptions{
        To:      []string{inbox2.EmailAddress},
        Subject: "Test subject",
        Body:    "Hello from inbox 1",
        Attachments: attachmentIds,
    }),
}
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

## Attachments
Attachments can be sent by first uploading an attachment then using the returned ID with send email options.

```go
func uploadAttachment(t *testing.T) []string {
    client, ctx := createClient(t)
    attachmentBytes := []byte{0}
    uploadOpts := mailslurp.UploadAttachmentOptions{
        Base64Contents: base64.URLEncoding.EncodeToString(attachmentBytes),
        ContentType:    "text/plain",
        Filename:       "test.txt",
    }
    attachmentIds, _, err := client.AttachmentControllerApi.UploadAttachment(ctx, uploadOpts)
    assert.Nil(t, err)
    return attachmentIds
}

func Test_CanSendEmail_AndReceive(t *testing.T) {
    client, ctx := createClient(t)
    opts := &mailslurp.CreateInboxOpts{}

    // create two inboxes
    inbox1, _, err := client.InboxControllerApi.CreateInbox(ctx, opts)
    inbox2, _, err := client.InboxControllerApi.CreateInbox(ctx, opts)

    assert.Nil(t, err)

    // upload attachments before send
    attachmentIds := uploadAttachment(t)

    // send a simple email from inbox 1 to inbox 2
    sendOpts := &mailslurp.SendEmailOpts{
        SendEmailOptions: optional.NewInterface(mailslurp.SendEmailOptions{
            To:      []string{inbox2.EmailAddress},
            Subject: "Test subject",
            Body:    "Hello from inbox 1",
            Attachments: attachmentIds,
        }),
    }

    _, err = client.InboxControllerApi.SendEmail(ctx, inbox1.Id, sendOpts)
    assert.Nil(t, err)

    // now receive the email at inbox 2
    waitForOpts := &mailslurp.WaitForLatestEmailOpts{
        InboxId:    optional.NewInterface(inbox2.Id),
        Timeout:    optional.NewInt64(30000),
        UnreadOnly: optional.NewBool(true),
    }
    email, _, err := client.WaitForControllerApi.WaitForLatestEmail(ctx,waitForOpts)
    assert.Nil(t, err)

    // access content
    assert.Equal(t, email.InboxId, inbox2.Id)
    assert.Contains(t, email.Subject, "Test subject")
    assert.Contains(t, email.Body, "Hello from inbox 1")

    // download the attachment as bytes
    downloadedBytes, _, err := client.EmailControllerApi.DownloadAttachment(ctx, email.Attachments[0], email.Id, nil)
    assert.Nil(t, err)
    assert.NotEmpty(t, downloadedBytes)
}
```

## Next step

For more information see the [source code on GitHub](https://github.com/mailslurp/mailslurp-client-go/).
