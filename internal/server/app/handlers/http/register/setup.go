package register

import "github.com/cyril-jump/gophkeeper/internal/server/app/handlers/http/register/requests"

func Setup() *requests.Requests {
	reqs := requests.New()

	return reqs
}
