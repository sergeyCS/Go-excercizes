package notification

import (
	"encoding/json"
	"net/http"
	"os"
	"restapi-lesson/internal/domain/notification"
	"restapi-lesson/internal/handlers"
	"restapi-lesson/pkg/logging"

	"github.com/julienschmidt/httprouter"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const (
	sickURL = "/sick"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.POST(sickURL, h.Sick)
}

func (h *handler) Sick(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var sickJson SickJson
	decoder := json.NewDecoder(r.Body)
	decoder.UseNumber()
	err := decoder.Decode(&sickJson)

	sick := notification.Sick{
		Reason:      sickJson.Reason,
		StartDate:   sickJson.StartDate,
		EndDate:     sickJson.EndDate,
		Alternate:   sickJson.Alternate,
		AlternateRm: sickJson.AlternateRm,
	}

	h.logger.Infof("sending to %s", "<to_address_here>")

	from := mail.NewEmail("Fastbot", "fastbot@fastdev.se")
	subject := "[Sick leave] Name"
	to := mail.NewEmail("unavailable@fastdev.se", "ssmarkin84@gmail.com")
	plainTextContent := sick.GetMessageText()
	htmlContent := ""
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		w.WriteHeader(500)
		h.logger.Panic(err)
	} else if response.StatusCode != 200 {
		h.logger.Panicf("sending failed")
		w.WriteHeader(response.StatusCode)
	} else {
		h.logger.Infof("sending completed, status code: %s", response.StatusCode)
		w.WriteHeader(204)
	}
}

// func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
// 	w.WriteHeader(204)
// 	w.Write([]byte("this is list of users"))
// }
