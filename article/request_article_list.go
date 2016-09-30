package article

import (
	"net/http"
	"strconv"
)

type articleListRequest struct {
	UserID int64
}

func (req *articleListRequest) Parse(r *http.Request) error {
	userID, err := strconv.ParseInt(r.FormValue("user_id"), 10, 64)
	if err != nil {
		return err
	}
	req.UserID = userID
	return nil
}
