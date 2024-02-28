package user

//
//import (
//	"bytes"
//	"github.com/pkulik0/stredono/cloud/pb"
//	"github.com/pkulik0/stredono/cloud/platform"
//	"github.com/pkulik0/stredono/cloud/platform/mocks"
//	"github.com/pkulik0/stredono/cloud/platform/modules"
//	"github.com/pkulik0/stredono/cloud/platform/providers"
//	"github.com/stretchr/testify/mock"
//	"google.golang.org/protobuf/proto"
//	"io"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//	"time"
//)
//
//func TestUserEdit(t *testing.T) {
//	uid := "testUser"
//
//	longStr := "1234567890"
//	for i := 0; i < 10; i++ {
//		longStr = longStr + longStr
//	}
//
//	tests := []struct {
//		name              string
//		user              *pb.User
//		code              int
//		token             string
//		isTokenValid      bool
//		isTokenForCurr    bool
//		duplicateUsername bool
//	}{
//		{
//			name: "valid-user-minimal",
//			user: &pb.User{
//				Uid:          uid,
//				Username:     "test",
//				DisplayName:  "Test",
//				PictureUrl:   "",
//				Url:          "",
//				Description:  "abc",
//				MinAmount:    2.50,
//				MinAuthLevel: pb.AuthLevel_NONE,
//				Currency:     pb.Currency_PLN,
//			},
//			code:           http.StatusOK,
//			token:          "xyz",
//			isTokenValid:   true,
//			isTokenForCurr: true,
//		},
//		{
//			name: "short-username",
//			user: &pb.User{
//				Uid:          uid,
//				Username:     "12",
//				DisplayName:  "Test",
//				PictureUrl:   "",
//				Url:          "",
//				Description:  "abc",
//				MinAmount:    2.50,
//				MinAuthLevel: pb.AuthLevel_NONE,
//				Currency:     pb.Currency_PLN,
//			},
//			code:           http.StatusBadRequest,
//			token:          "def",
//			isTokenValid:   true,
//			isTokenForCurr: true,
//		},
//		{
//			name: "long-username",
//			user: &pb.User{
//				Uid:          uid,
//				Username:     "12345678901234567123456789012345671234567890123456712345678901234567",
//				DisplayName:  "Test",
//				PictureUrl:   "",
//				Url:          "",
//				Description:  "abc",
//				MinAmount:    2.50,
//				MinAuthLevel: pb.AuthLevel_NONE,
//				Currency:     pb.Currency_PLN,
//			},
//			code:           http.StatusBadRequest,
//			token:          "abc",
//			isTokenValid:   true,
//			isTokenForCurr: true,
//		},
//		{
//			name: "invalid-token",
//			user: &pb.User{
//				Uid:          uid,
//				Username:     "test",
//				DisplayName:  "Test",
//				PictureUrl:   "",
//				Url:          "",
//				Description:  "abc",
//				MinAmount:    2.50,
//				MinAuthLevel: pb.AuthLevel_NONE,
//				Currency:     pb.Currency_PLN,
//			},
//			code:           http.StatusUnauthorized,
//			token:          "123",
//			isTokenValid:   false,
//			isTokenForCurr: true,
//		},
//		{
//			name: "wrong-user",
//			user: &pb.User{
//				Uid:          uid,
//				Username:     "test",
//				DisplayName:  "Test",
//				PictureUrl:   "",
//				Url:          "",
//				Description:  "abc",
//				MinAmount:    2.50,
//				MinAuthLevel: pb.AuthLevel_NONE,
//				Currency:     pb.Currency_PLN,
//			},
//			code:           http.StatusBadRequest,
//			token:          "123",
//			isTokenValid:   true,
//			isTokenForCurr: false,
//		},
//		{
//			name: "short-display-name",
//			user: &pb.User{
//				Uid:          uid,
//				Username:     "test",
//				DisplayName:  "1",
//				PictureUrl:   "",
//				Url:          "",
//				Description:  "abc",
//				MinAmount:    2.50,
//				MinAuthLevel: pb.AuthLevel_NONE,
//				Currency:     pb.Currency_PLN,
//			},
//			code:           http.StatusBadRequest,
//			token:          "123",
//			isTokenValid:   true,
//			isTokenForCurr: true,
//		},
//		{
//			name: "long-display-name",
//			user: &pb.User{
//				Uid:          uid,
//				Username:     "test",
//				DisplayName:  "test-test-test-test-test-test-test-test-test-test-test-test-test-test-test-test-test",
//				PictureUrl:   "",
//				Url:          "",
//				Description:  "abc",
//				MinAmount:    2.50,
//				MinAuthLevel: pb.AuthLevel_NONE,
//				Currency:     pb.Currency_PLN,
//			},
//			code:           http.StatusBadRequest,
//			token:          "123",
//			isTokenValid:   true,
//			isTokenForCurr: true,
//		},
//		{
//			name: "invalid-picture-url",
//			user: &pb.User{
//				Uid:          uid,
//				Username:     "test",
//				DisplayName:  "testUser",
//				PictureUrl:   "http",
//				Url:          "",
//				Description:  "abc",
//				MinAmount:    2.50,
//				MinAuthLevel: pb.AuthLevel_NONE,
//				Currency:     pb.Currency_PLN,
//			},
//			code:           http.StatusBadRequest,
//			token:          "123",
//			isTokenValid:   true,
//			isTokenForCurr: true,
//		},
//		{
//			name: "long-display-name",
//			user: &pb.User{
//				Uid:          uid,
//				Username:     "test",
//				DisplayName:  "User",
//				PictureUrl:   "",
//				Url:          "abc",
//				Description:  "abc",
//				MinAmount:    2.50,
//				MinAuthLevel: pb.AuthLevel_NONE,
//				Currency:     pb.Currency_PLN,
//			},
//			code:           http.StatusBadRequest,
//			token:          "123",
//			isTokenValid:   true,
//			isTokenForCurr: true,
//		},
//		{
//			name: "valid-urls",
//			user: &pb.User{
//				Uid:          uid,
//				Username:     "test",
//				DisplayName:  "User",
//				PictureUrl:   "https://storage.googleapis.com/v0/b/stredono-6394ee11.appspot.com/o/31hd1e",
//				Url:          "https://stredono.com/images/test.png",
//				Description:  "abc",
//				MinAmount:    2.50,
//				MinAuthLevel: pb.AuthLevel_NONE,
//				Currency:     pb.Currency_PLN,
//			},
//			code:           http.StatusOK,
//			token:          "123",
//			isTokenValid:   true,
//			isTokenForCurr: true,
//		},
//		{
//			name: "firebase-url",
//			user: &pb.User{
//				Uid:          uid,
//				Username:     "test",
//				DisplayName:  "User",
//				PictureUrl:   "https://firebasestorage.googleapis.com/v0/b/stredono-6394ee11.appspot.com/o/31hd1e",
//				Url:          "https://stredono.com/images/test.png",
//				Description:  "abc",
//				MinAmount:    2.50,
//				MinAuthLevel: pb.AuthLevel_NONE,
//				Currency:     pb.Currency_PLN,
//			},
//			code:           http.StatusOK,
//			token:          "123",
//			isTokenValid:   true,
//			isTokenForCurr: true,
//		},
//		{
//			name: "unknown-currency",
//			user: &pb.User{
//				Uid:          uid,
//				Username:     "test",
//				DisplayName:  "User",
//				PictureUrl:   "",
//				Url:          "",
//				Description:  "abc",
//				MinAmount:    2.50,
//				MinAuthLevel: pb.AuthLevel_NONE,
//				Currency:     pb.Currency_UNKNOWN,
//			},
//			code:           http.StatusBadRequest,
//			token:          "123",
//			isTokenValid:   true,
//			isTokenForCurr: true,
//		},
//		{
//			name: "description-too-long",
//			user: &pb.User{
//				Uid:          uid,
//				Username:     "test",
//				DisplayName:  "User",
//				PictureUrl:   "",
//				Url:          "",
//				Description:  longStr,
//				MinAmount:    2.50,
//				MinAuthLevel: pb.AuthLevel_NONE,
//				Currency:     pb.Currency_PLN,
//			},
//			code:           http.StatusBadRequest,
//			token:          "123",
//			isTokenValid:   true,
//			isTokenForCurr: true,
//		},
//		{
//			name: "min-amount-negative",
//			user: &pb.User{
//				Uid:          uid,
//				Username:     "test",
//				DisplayName:  "User",
//				PictureUrl:   "",
//				Url:          "",
//				Description:  "abc",
//				MinAmount:    -2.50,
//				MinAuthLevel: pb.AuthLevel_NONE,
//				Currency:     pb.Currency_PLN,
//			},
//			code:           http.StatusBadRequest,
//			token:          "123",
//			isTokenValid:   true,
//			isTokenForCurr: true,
//		},
//		{
//			name: "duplicate",
//			user: &pb.User{
//				Uid:          uid,
//				Username:     "test",
//				DisplayName:  "User",
//				PictureUrl:   "",
//				Url:          "",
//				Description:  "abc",
//				MinAmount:    2.50,
//				MinAuthLevel: pb.AuthLevel_NONE,
//				Currency:     pb.Currency_PLN,
//			},
//			code:              http.StatusBadRequest,
//			token:             "123",
//			isTokenValid:      true,
//			isTokenForCurr:    true,
//			duplicateUsername: true,
//		},
//		{
//			name: "invalid-min-auth-level",
//			user: &pb.User{
//				Uid:          uid,
//				Username:     "test",
//				DisplayName:  "User",
//				PictureUrl:   "",
//				Url:          "",
//				Description:  "abc",
//				MinAmount:    2.50,
//				MinAuthLevel: -1,
//				Currency:     pb.Currency_PLN,
//			},
//			code:              http.StatusBadRequest,
//			token:             "123",
//			isTokenValid:      true,
//			isTokenForCurr:    true,
//			duplicateUsername: false,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			ctx := providers.CreateContextMock(&providers.Config{
//				DocDb: true,
//				Auth:  true,
//			}, t)
//
//			auth, ok := ctx.GetAuth()
//			if !ok {
//				t.Fatal("missing auth")
//			}
//			authMock := auth.(*mocks.MockAuth)
//
//			if tt.isTokenValid {
//				token := mocks.NewMockToken(t)
//
//				tokenUid := uid
//				if !tt.isTokenForCurr {
//					tokenUid = "__other__"
//				}
//				token.EXPECT().UserId().Return(tokenUid)
//				authMock.EXPECT().VerifyToken(mock.Anything, tt.token).Return(token, nil)
//			} else {
//				authMock.EXPECT().VerifyToken(mock.Anything, mock.Anything).Return(nil, platform.ErrorUnauthorized)
//			}
//
//			db, ok := ctx.GetDocDb()
//			if !ok {
//				t.Fatal("missing db")
//			}
//			dbMock := db.(*mocks.MockDocDb)
//			coll := mocks.NewMockCollection(t)
//			doc := mocks.NewMockDocument(t)
//			q := mocks.NewMockQuery(t)
//			qSnap := mocks.NewMockQuerySnapshot(t)
//			userDoc := mocks.NewMockDocumentSnapshot(t)
//
//			if tt.code == http.StatusOK || tt.duplicateUsername {
//				dbMock.EXPECT().Collection("users").Return(coll)
//
//				coll.EXPECT().Where("Username", "==", tt.user.Username).Return(q)
//				q.EXPECT().Where("Uid", "!=", uid).Return(q)
//				q.EXPECT().Documents(mock.Anything).Return(qSnap)
//
//				if tt.duplicateUsername {
//					var docSnapshots []modules.DocumentSnapshot
//					docSnapshots = append(docSnapshots, userDoc)
//					qSnap.EXPECT().GetAll().Return(docSnapshots, nil)
//				} else {
//					coll.EXPECT().Doc(uid).Return(doc)
//
//					qSnap.EXPECT().GetAll().Return(nil, nil)
//					doc.EXPECT().Set(mock.Anything, mock.Anything, mock.Anything).Return(&modules.WriteResult{
//						Time: time.Now(),
//					}, nil)
//				}
//			}
//
//			body, err := proto.Marshal(tt.user)
//			if err != nil {
//				t.Fatal(err)
//			}
//
//			req := httptest.NewRequest("GET", "/", io.NopCloser(bytes.NewBuffer(body)))
//			req.Header.Set("Authorization", "Bearer "+tt.token)
//			req.Header.Set("Content-Type", "application/octet-stream")
//			req = req.WithContext(ctx.Ctx)
//
//			w := httptest.NewRecorder()
//			edit(ctx, w, req)
//
//			if w.Code != tt.code {
//				t.Errorf("Want status: %d, got status: %d", tt.code, w.Code)
//			}
//		})
//	}
//}
