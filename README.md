# slack-msg

Slack webhook message sending library for Go (pretty basic).

```go
package main

import "github.com/vadviktor/slack-msg"

func main() {
    s := &slack_msg.Slack{}
    s.Create("someWebhookURL")
    s.Send("I can send plain fmt.Println.")
    s.Send("I can act like %s too.", "fmt.Sprintf")
}
```
