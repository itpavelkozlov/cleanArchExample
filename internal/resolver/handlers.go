package resolver

import (
	"cleanArch/internal/models"
	"cleanArch/internal/service"
	"cleanArch/pkg/logger"
	"encoding/json"
	"fmt"
	"net/http"
)

type Resolver struct {
	userService service.User
	logger      logger.Logger
	//	groupService service.Group
	//	blog service.Blog
}

func NewResolver(userService service.User, logger logger.Logger) *Resolver {
	return &Resolver{
		userService: userService,
		logger:      logger,
	}
}

func (r Resolver) CreateUser(w http.ResponseWriter, req *http.Request) {
	r.logger.Debug("CreateUser handler used!")
	u := new(models.User)
	err := json.NewDecoder(req.Body).Decode(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = r.userService.Create(req.Context(), u); err != nil {
		http.Error(w, fmt.Sprintf("Error where user creating %s\n", err.Error()), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Ok\n")
}
