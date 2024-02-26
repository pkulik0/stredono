package tips

import (
	"bytes"
	"context"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform/mocks"
	"github.com/pkulik0/stredono/cloud/platform/modules"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestTipSend(t *testing.T) {
	validTip := &pb.Tip{
		SenderId:    "senderid",
		DisplayName: "sender",
		RecipientId: "recipient",
		Status:      pb.TipStatus_INITIATED,
		Amount:      1.0,
		Currency:    pb.Currency_PLN,
		Email:       "email",
	}

	invalidTipSender := *validTip
	invalidTipSender.DisplayName = ""

	invalidTipSenderId := *validTip
	invalidTipSenderId.SenderId = ""

	invalidTipRecipientId := *validTip
	invalidTipRecipientId.RecipientId = ""

	invalidTipStatus := *validTip
	invalidTipStatus.Status = pb.TipStatus_PAYMENT_PENDING

	invalidTipAmount := *validTip
	invalidTipAmount.Amount = -1.0

	invalidTipCurrency := *validTip
	invalidTipCurrency.Currency = pb.Currency_UNKNOWN

	invalidTipEmail := *validTip
	invalidTipEmail.Email = ""

	invalidTipTimestamp := *validTip
	invalidTipTimestamp.Timestamp = time.Now().Unix()

	tests := []struct {
		name   string
		tip    *pb.Tip
		status int
	}{
		{
			name:   "valid",
			tip:    validTip,
			status: http.StatusOK,
		},
		{
			name:   "empty sender",
			tip:    &invalidTipSender,
			status: http.StatusBadRequest,
		},
		{
			name:   "empty sender id",
			tip:    &invalidTipSenderId,
			status: http.StatusBadRequest,
		},
		{
			name:   "empty recipient id",
			tip:    &invalidTipRecipientId,
			status: http.StatusBadRequest,
		},
		{
			name:   "empty status",
			tip:    &invalidTipStatus,
			status: http.StatusBadRequest,
		},
		{
			name:   "invalid amount",
			tip:    &invalidTipAmount,
			status: http.StatusBadRequest,
		},
		{
			name:   "empty currency",
			tip:    &invalidTipCurrency,
			status: http.StatusBadRequest,
		},
		{
			name:   "empty email",
			tip:    &invalidTipEmail,
			status: http.StatusBadRequest,
		},
		{
			name:   "timestamp included",
			tip:    &invalidTipTimestamp,
			status: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, err := providers.CreateContextMock(context.Background(), &providers.Config{
				DocDb: true,
			}, t)
			if err != nil {
				t.Errorf("Failed to create context: %s", err)
			}

			if tt.status == http.StatusOK {
				db, ok := providers.GetDocDb(ctx)
				if !ok {
					t.Errorf("Failed to get doc db")
				}
				dbMock := db.(*mocks.MockNoSqlDb)

				col := mocks.NewMockCollection(t)
				doc := mocks.NewMockDocument(t)

				dbMock.EXPECT().Collection("tips").Return(col)
				col.EXPECT().Add(mock.Anything, mock.Anything).Return(&modules.AddResult{
					Doc: doc,
					WriteResult: modules.WriteResult{
						Time: time.Now(),
					},
				}, nil)
				doc.EXPECT().Id().Return("1")
			}

			reqBody, err := proto.Marshal(tt.tip)
			if err != nil {
				t.Errorf("Failed to marshal request: %s", err)
			}

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/octet-stream")
			req = req.WithContext(ctx)

			w := httptest.NewRecorder()
			send(w, req)

			res := w.Result()
			defer res.Body.Close()

			if res.StatusCode != tt.status {
				t.Errorf("Want status: %d, got status: %d", tt.status, res.StatusCode)
			}

			if res.StatusCode != http.StatusOK {
				return
			}

			data, err := io.ReadAll(res.Body)
			if err != nil {
				t.Errorf("Failed to read response: %s", err)
			}
			defer res.Body.Close()

			if string(data) != "http://google.com" { // TODO: replace with actual URL
				t.Errorf("Want response: %s, got response: %s", "http://google.com", string(data))
			}
		})
	}
}

func TestTipSendMalformed(t *testing.T) {
	tip := &pb.Tip{
		SenderId:    "senderid",
		DisplayName: "sender",
		RecipientId: "recipient",
		Status:      pb.TipStatus_INITIATED,
		Amount:      1.0,
		Currency:    pb.Currency_PLN,
		Email:       "email",
	}
	data, err := proto.Marshal(tip)
	if err != nil {
		t.Errorf("Failed to marshal test tip: %s", err)
	}

	tests := []struct {
		name string
		body string
		code int
	}{
		{
			name: "empty",
			body: "",
			code: http.StatusBadRequest,
		},
		{
			name: "invalid",
			body: "invalid",
			code: http.StatusBadRequest,
		},
		{
			name: "missing context",
			body: bytes.NewBuffer(data).String(),
			code: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/octet-stream")

			w := httptest.NewRecorder()
			send(w, req)

			res := w.Result()
			defer res.Body.Close()

			if res.StatusCode != tt.code {
				t.Errorf("Want status: %d, got status: %d", tt.code, res.StatusCode)
			}
		})
	}
}
