package handler

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"streaming/common"
	"streaming/scheduler/model"
)

func VidDelRecHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	if len(vid) == 0 {
		common.SendResponse(w, 400, "video id should not be empty")
		return
	}
	err := model.AddVideoDeletionRecord(vid)
	if err != nil {
		common.SendResponse(w, 500, "Internal server error")
		return
	}
	common.SendResponse(w, 200, "")
	return
}
