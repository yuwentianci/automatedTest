package google

import (
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"os"
)

func EmailCode() {
	ctx := context.Background()

	// 使用 OAuth 2.0 认证获取 Gmail 服务客户端
	client, err := getClient(ctx, "YOUR_CREDENTIALS_JSON_FILE.json", gmail.MailGoogleComScope) //要运行此示例，需要先在 Google 开发者控制台创建一个项目，启用 Gmail API，并下载 OAuth 2.0 凭据。
	if err != nil {
		log.Fatalf("无法获取 Gmail 客户端: %v", err)
	}

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("无法创建 Gmail 服务: %v", err)
	}

	// 获取未读邮件列表
	user := "me"
	query := "is:unread"
	messages, err := srv.Users.Messages.List(user).Q(query).Do()
	if err != nil {
		log.Fatalf("无法获取未读邮件: %v", err)
	}

	// 打印邮件主题
	for _, msg := range messages.Messages {
		email, err := srv.Users.Messages.Get(user, msg.Id).Do()
		if err != nil {
			log.Printf("无法获取邮件 %s: %v", msg.Id, err)
		} else {
			subject := getEmailSubject(email)
			fmt.Printf("邮件主题: %s\n", subject)
		}
	}
}

func getClient(ctx context.Context, credentialsFile string, scope ...string) (*http.Client, error) {
	data, err := os.ReadFile(credentialsFile)
	if err != nil {
		return nil, err
	}

	config, err := google.ConfigFromJSON(data, scope...)
	if err != nil {
		return nil, err
	}

	client := config.Client(ctx, &oauth2.Token{
		AccessToken: "YOUR_ACCESS_TOKEN", // AccessToken(访问令牌) 通常是从 OAuth 2.0 授权流程中获取的。
	})

	return client, nil
}

func getEmailSubject(email *gmail.Message) string {
	for _, header := range email.Payload.Headers {
		if header.Name == "Subject" {
			return header.Value
		}
	}
	return "无主题"
}
