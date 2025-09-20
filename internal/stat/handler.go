package stat

import (
	"GolangAdvanced/configs"
	"GolangAdvanced/pkg/event"
	"GolangAdvanced/pkg/middleware"
	"fmt"
	"net/http"
	"time"
)

const (
	FilterByDay   = "day"
	FilterByMonth = "month"
)

type StatHandler struct {
	StatRepository *StatRepository
	EventBus       *event.EventBus
}

type StatHandlerDeps struct {
	StatRepository *StatRepository
	Config         *configs.Config
	EventBus       *event.EventBus
}

func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps) {
	handler := &StatHandler{
		StatRepository: deps.StatRepository,
		EventBus:       deps.EventBus,
	}
	router.Handle("Get /stat", middleware.IsAuthed(handler.GetStat(), deps.Config))
}

func (handler *StatHandler) GetStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		from, err := time.Parse("2006-01-02", r.URL.Query().Get("from"))
		if err != nil {
			http.Error(w, "Invalid from param", http.StatusBadRequest)
			return
		}

		to, err := time.Parse("2006-01-02", r.URL.Query().Get("to"))
		if err != nil {
			http.Error(w, "Invalid to param", http.StatusBadRequest)
			return
		}
		by := r.URL.Query().Get("by")
		if by != FilterByDay && by != FilterByMonth {
			http.Error(w, "Invalid by param", http.StatusBadRequest)
			return
		}
		fmt.Println(from, to, by)
	}
}
