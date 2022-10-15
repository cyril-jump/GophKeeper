package login

import "github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/login/requests"

func Setup() *requests.Requests {
	reqs := requests.New()

	return reqs
}
