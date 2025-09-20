package stat

import (
	"GolangAdvanced/configs"
	"GolangAdvanced/pkg/event"
	"GolangAdvanced/pkg/middleware"
	res "GolangAdvanced/pkg/response"
	"net/http"
	"time"
)

const (
	GroupByDay   = "day"
	GroupByMonth = "month"
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
		if by != GroupByDay && by != GroupByMonth {
			http.Error(w, "Invalid by param", http.StatusBadRequest)
			return
		}
		stats := handler.StatRepository.GetStats(by, from, to)
		res.JsonRes(w, stats, 200)
	}
}
