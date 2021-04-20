package token

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const publicKey = `
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnmQlbIfjtp6Krra8oT00
XGFE2y6TnwPNcK3csTwXJDZ5Dv/t2YS9PyQDFsVQGxWdOPSnZ8iy8cvcA4SnJZMz
oFI+LXWuMCPxCzKWzdikaS+iSOrVImQMyS+p9Ec+VeFlaBKJeESAmaHgHLGGE3L/
L7yw2koRpeg77RLtN1geiJ14qT6vw57lCalO1bn7RD7jnsaTnUs1QKP0af1psZvB
Da3ZNLA7k9+AJImV/FJ3VrJ6BFwyIWoD1vDQB0DZQpusEzHt2LR5Htp+a7jMKNV9
YvBhNgLIJ2ykvvahnmOFspZI62Kh/fsLZiY4Sp3dXaLCkGfIbAS8knSgz4W+Rsf8
iQIDAQAB
-----END PUBLIC KEY-----
`

func TestVerifyToken(t *testing.T) {
	cases := []struct {
		name              string
		now               time.Time
		publickey         string
		token             string
		want              string
		verifierCreateErr bool
		wantErr           bool
	}{
		{
			name:              "invalid_publickey",
			now:               time.Unix(1516239122, 0),
			publickey:         "",
			token:             `eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29taWMvYXV0aCIsInN1YiI6IjYwNTVhOTZlNDg1MDk4MTMwODJmYjkxMyJ9.Yz6rfzVGCdSl3_wy-xl7Cmy3tPFkdITrO_ji3SDB_oTYEJyP3VgIC_1SDS3fgAD04Dvj1YdHAMfzKy8k-RHctaGUu0ZzwUosAVLzwN2OztWAIZlG6-JYayeDWV82qu_hpOLAvrvNftHL01Y_m6qB3Z--0pbErg6CLB6dd51RZ4mor0l3CJiRvbJVgapWD7YGCPFD883Ywia8tlqJ3dNKHWgTu7obRwfUS9B-_GCc0SBO2M-2xqw9nhzdz-7jYYzdnIXsKt5_-GZVPQnQJw19-GlfpQVtslZPAEfqH_Amqeio_vCnPdOwQsbaOVMVupmwwTJ54_WZCb--wOyZmlBUGw`,
			want:              "",
			verifierCreateErr: true,
			wantErr:           true,
		},
		{
			name:      "valid_token",
			now:       time.Unix(1516239122, 0),
			publickey: publicKey,
			token:     `eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29taWMvYXV0aCIsInN1YiI6IjYwNTVhOTZlNDg1MDk4MTMwODJmYjkxMyJ9.Yz6rfzVGCdSl3_wy-xl7Cmy3tPFkdITrO_ji3SDB_oTYEJyP3VgIC_1SDS3fgAD04Dvj1YdHAMfzKy8k-RHctaGUu0ZzwUosAVLzwN2OztWAIZlG6-JYayeDWV82qu_hpOLAvrvNftHL01Y_m6qB3Z--0pbErg6CLB6dd51RZ4mor0l3CJiRvbJVgapWD7YGCPFD883Ywia8tlqJ3dNKHWgTu7obRwfUS9B-_GCc0SBO2M-2xqw9nhzdz-7jYYzdnIXsKt5_-GZVPQnQJw19-GlfpQVtslZPAEfqH_Amqeio_vCnPdOwQsbaOVMVupmwwTJ54_WZCb--wOyZmlBUGw`,
			want:      "6055a96e48509813082fb913",
			wantErr:   false,
		},
		{
			name:      "expired_token",
			now:       time.Unix(1516246223, 0),
			publickey: publicKey,
			token:     `eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29taWMvYXV0aCIsInN1YiI6IjYwNTVhOTZlNDg1MDk4MTMwODJmYjkxMyJ9.Yz6rfzVGCdSl3_wy-xl7Cmy3tPFkdITrO_ji3SDB_oTYEJyP3VgIC_1SDS3fgAD04Dvj1YdHAMfzKy8k-RHctaGUu0ZzwUosAVLzwN2OztWAIZlG6-JYayeDWV82qu_hpOLAvrvNftHL01Y_m6qB3Z--0pbErg6CLB6dd51RZ4mor0l3CJiRvbJVgapWD7YGCPFD883Ywia8tlqJ3dNKHWgTu7obRwfUS9B-_GCc0SBO2M-2xqw9nhzdz-7jYYzdnIXsKt5_-GZVPQnQJw19-GlfpQVtslZPAEfqH_Amqeio_vCnPdOwQsbaOVMVupmwwTJ54_WZCb--wOyZmlBUGw`,
			want:      "",
			wantErr:   true,
		},
		{
			name:      "bad_token",
			now:       time.Unix(1516246223, 0),
			publickey: publicKey,
			token:     "",
			want:      "",
			wantErr:   true,
		},
		{
			name:      "fake_token",
			now:       time.Unix(1516246223, 0),
			publickey: publicKey,
			token:     `eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29taWMvYXV0aCIsInN1YiI6IjYwNTVhOTZlNDg1MDk4MTMwODJmYjkxNCJ9.kENrWciIrG1DVl_ERYcutrgQDAhszTxnPJbxnVT-OxcTfnz3AOcn2Rkm-Brkv80obaEknZvzQpAXcDoWswhgNf4ErPn1VKKlsU-Qa5y3wrWUhdrLnBON1LxpySpCm4Byoi21lwYeZnRpPc-tHYuz2cwAFNsYKzwa9lXtvSiuA6eyV-6zgIPMQmgv_h4tv-u4hWahelKTQp5O4RC-tjga5K_glmQEQz-hkMaH4ffTNXyqTQNKah563_8Ta5qgW1q50pwVq5LbA_gMLL78ZzM4shTzV_A1XJwIXegbvbj4W1iqU0rmnZ2EDWeWtR_loQ5gDuObVgcv2awg8MMAcI99hw`,
			want:      "",
			wantErr:   true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			verifier, err := NewJWTVerifier(c.publickey)
			if err != nil {
				if c.verifierCreateErr {
					return
				}
				t.Fatalf("cannot creater verifier: %v", err)
			}

			jwt.TimeFunc = func() time.Time {
				return c.now
			}

			accountID, err := verifier.VerifyToken(c.token)
			if !c.wantErr && err != nil {
				t.Errorf("failed to verify token: %v", err)
			}

			if c.wantErr && err == nil {
				t.Errorf("want error,got no error")
			}

			if accountID != c.want {
				t.Errorf("wront accountID: want: %q; got: %q", c.want, accountID)
			}
		})
	}
}
