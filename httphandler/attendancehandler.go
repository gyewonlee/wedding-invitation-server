package httphandler

import (
	"encoding/json"
	"net/http"

	"github.com/juhonamnam/wedding-invitation-server/sqldb"
	"github.com/juhonamnam/wedding-invitation-server/types"
)

type AttendanceHandler struct {
	http.Handler
}

func (h *AttendanceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		attendances, err := sqldb.GetAllAttendance()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}

		pbytes, err := json.Marshal(attendances)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(pbytes)
	} else if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		var attendance types.AttendanceCreate
		err := decoder.Decode(&attendance)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("BadRequest"))
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("InternalServerError"))
			return
		}

		err = sqldb.CreateAttendance(attendance.Side, attendance.Name, attendance.Meal, attendance.Count)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("InternalServerError"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
	}
}
