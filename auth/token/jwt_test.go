package token

import (
	"testing"
	"time"
)

const privateKey1 = `
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAmHCV4V+w+E+s0WtGsAy/WdvkOkx29ypLNuVDaFMVU4HTTC4m
i+/9IDGXk96xyTpZxKILLYgWrWLmXsi97r2eHVm1UOzNNtx/VrP77P9TD/WaI3vR
0EYTHZqzvZww5Hm7VWShh90ZMIqfMuvu3ZLEeDv2We4EypyaR6p6fcr4xyR/yP+r
Ap2G8eCk3ppa0AT23SZUfovQ+hJWiZohcxJhbYyROVtjwXeovkL9tKh0onBA/cuk
5cDXJ9GC3a0cbD+Ca0NPYJbdwVYfAGWVgGuV1/jI4s27Mw/Rl2Mws5H1+ETxMbJR
VYE0gvnF3czBE0WWcVFssaZBme10OvbZJE/sWQIDAQABAoIBAGk1yzwxf0LiOrSP
BcqcAbVTPsG95J3SYvcQZOWr8hwgjmUVtSUADgQaNjOoj6wCydIcxPo34u7oQ3dH
tU8BMA+xyqJ/zWL/7CM7NbTsw7dQG4JCAx1sP+U+0K8dDeGi4tidC59pyejmw9mx
XLswgMk7GVwYGtZ8Hz83OjN30xwdzxe1lO5UQUwdAKfMgAreSw6YTMbMhvGt/V7i
lo+FY6Y9CaA7bOTPAyhfxR7nVsMdrDMjsTimYT5R4wEiKbkAZ0U+rOC/9Ub88q6P
ywEgIvtaY0YEDFO5c0K3mQuTAt+dr0LBdB7eR44RqlKZyHJlcX3ApdWy5gQQEyRD
9hP/M8ECgYEA72UVWtO0OSmtXsLAxc+6kGx4I4S2rc3Wv6EGKsA5SU0yjqKbC6qt
nb+kRjiPPtBsi8ILMGK3WKg9f0gY4+jxQwsjhcec1Of4ygk5COn9Kt0J0w7HI4s+
GmsNeVip8e2m/5RLDAiY1rUeqC7kzRmax25W8FpBIWGwPeWWZl1kShcCgYEAowNy
hwmkHyTZGVjmfLDK3VthhGHrJ7XwhkOWGy41FJ6U9mEdLfq9gFADP1NHLieQb1Hw
PmFl5SlaDRNUxUNoJRUcVZlyh5ZXTO53eZDYKc5N9eqfhvMJutEst/iV5A+A+PKF
ei7UaGt16wnkFlE/BpgzmxgjUP1X+dFlu4nbMw8CgYAR7oNr+0GbqoiG+ZWQ+59g
5A8XsgAxAqMcVYZtu1pPKE4qKvnsJMu+bjVy4Rexv9DxowZNiIdtR6TbnN3obHFt
8S6m/a28sOeM0qYICebhyA7NVWULP4nWVsB6n26TZZbFvCjm7Nnt80EpneGbj1ht
qdEvTcNlKbkW2dSrNEGn3wKBgQCA/IxV38IWxqBmVxrh1T5gUxthld1tqaAtjTMs
RK7JpwW1wnhFVp819ZnsNKghFDtMamiP45+t6b+QkxhDkqfIl+C4tteSA9sGS7cL
4Rm6Ur0LdmBOdubfBP6+f5uTPtVNpxEI8h8ulMFUfpoVMqqNPEexVDM1tZE0YoV+
FIH27wKBgFhoBjkxPQyTYPyk91f5kkNzzdBxFnBlmM1xhD9g9PSpjvYtYFe0dOX3
Ke5vejH7gHs5M8CkWn5NSEFH75Gpu0BIwg540XMHJkvbi/0G7c6i3xKSqMtAd0b1
+Ye1RmB6fbLLguHYCO0qjKnyB0lBqs0r4MCsQ8quvRn+49L2dU7X
-----END RSA PRIVATE KEY-----
`

const privateKey2 = `
-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQBPLIn7yTE/r4qo3VCY1PMMNmkypSbbmcnDQlpJJrqvPYLul9Oo
7iDcqn67UAMq9g0CE3KhYrvz4FJosNvYrsHGkJMaDxokDCfn3Wv0cxtmOoU4ZC3n
IAP4+oC4QHq0Oy46y28t1p0dVlxqTH11CQFXV/qdOl93mL6vK3Bmw20QgzzlF9vo
4nHmk81YodIVAJZpp31Fr9JHwwn7biReq+WqlYNEgyuCDOMBdg9PhIZtNQimKIYO
LlWzQWKpp7FM2pdHLfb2hmajWvr1+UXEVuEzyzeJsTF7FE4xLOe10aZXSE/Zf+Cz
od+7CqLKXnlUiNa98PxuToFy7QUzR7l3D0NzAgMBAAECggEAEOdJbmA+C4LG79aQ
ZpDQc0MKoq1v5mKnGPEJpaSHJvLTCW2Q2vg8O+CKviMPChUOIetVBe+2gRLbhU59
IX3uER68yUuAIUIXIwG7Eq2go+rKloeM5VOIhZe186CQkfownIziZ77yPFk70dYS
0vHhP2yhLQBRcY9C5tM3ZDlFB8lAEUCKlsQ4/z8oeELtzVc1eTwxpRQY9HAEhJNT
DH6nRTcRN8HVnHUJbhXcuJB3Ly2YQz1bltE6xPnNiAKW8DTyeLVG9Q/t4+u3okYH
PTM4+d3WfeJLgxpyoV2SF0PBhkAvGXaXQQaFxayI9KJmJ9/B9f8mUxoSjPOt1GbO
SbjDQQKBgQCd30C4cnXKlcF0t05Ip6kbHyYRyzrG+7VFY8tyQtYzFiot2nGijiOY
2wnbYtjiIOr8XCrTHM/0oNDt3UViEYbp6GljEktXEVmR8I0jBdo0OgWS4eJ+vPbO
Rnunp1HmH7Yfw9nlAGnHsAF1igRoa8NeGHs2/QStToCRwKxEGBzD7wKBgQCAYsYW
4zLXjdrXQBMrc/uoEmpjHKz+QqsY9qHTC45+QU/YSkAg6llNBTWzu9Wnhk95UaVQ
NqjaJv5jfl3wCl5QLQHEJo5v9RFyIgHYCWIR4J5DZe1DjK4B+CeZi9ShcVsOGMOu
3Pmdv3eQd+kqCw9aroO+YCfp+BsqQvPQ/QskvQKBgQCLcuV3R1t2q9trHnOex8Ie
+9+jvhFYLIRNG75cWl9cVsrBIABaJxWYcnp9eA29iHgom70UJ5IlMlv6kAY2fubd
/HYClelIqkGlIId6lTPOCCZePPrVGlMjl6yoRaqZDA/SdSwmFLxbL8PKMDR+jBnR
JuzHgK097Gay2C8SIv5drwKBgCTvRZm0Bx5oGvst3dX0agOyOJUG3OXQ/DQUT/3N
J1M/Kt9IjSZrBQHZa9enRBIL27VFPc/TycPoz+qDoja6ZOOf7xAlH3kqAbjZ+7nR
CTunaSBupFF4mV7Lq2jLB4Lif5WQnksiG1W8jYMsXBHSWLp31kLgmtr+m6IO3OHU
ym3pAoGAeIbUNQYIxh0DxZFKAncozEAin9kDxcROMgBV7g/WmT7Rp0IdjgWzYo4m
8J6s6L4PXIw7yZhDHuyWX4qutt3a3T0GBBV3YYqZEmXqC64G+ZUZ0paKP4UZ0mzJ
XpLWiLmN1fNvaFz0SS3ouI6du1LnWpBIeAjireXTMekIv2SAQI8=
-----END RSA PRIVATE KEY-----
`

func TestGenerateToken(t *testing.T) {
	cases := []struct {
		name       string
		exp        time.Duration
		iat        int64
		iss        string
		accountID  string
		privateKey string
		wantToken  string
		tokenErr   bool
	}{
		{
			name:       "valid_token",
			exp:        2 * time.Hour,
			iat:        1516239022,
			iss:        "comic/auth",
			accountID:  "6055a96e48509813082fb913",
			privateKey: privateKey1,
			wantToken:  `eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29taWMvYXV0aCIsInN1YiI6IjYwNTVhOTZlNDg1MDk4MTMwODJmYjkxMyJ9.M2_ZUyPS9uhzFn92Lief8KvXH4NQWwhvQYJnj_9BFVygqi18WiJe4--ARLqbuNRtTC5ey2L5Hy6JtLJEOa0KBsjAc-ktMLns-vpn8jEJ1Y_UCKG5A3rZ9qjRJhiRxRcW4qE-uriuTQj_0ZtYjyO5zIUqewYxkNYMNxoK6md6TUidXYkt8ES39zvCNBbmzxVx_s6MLXUKJ59bT2-gLG1iumcXwbfLgTkObO4TCjS0D-ayvllOCUT48BpTeTmNexC1xsAeMg2uYYjxx6kK-kYHOhotUNT1lW8hz6sFxe6WmnNgqhkpsO5tk_HqIUhHa29TwLdpoh-I1SNcnLy56k9Wvw`,
			tokenErr:   false,
		},
		{
			name:       "dif_privateKey",
			exp:        2 * time.Hour,
			iat:        1516239022,
			iss:        "comic/auth",
			accountID:  "6055a96e48509813082fb913",
			privateKey: privateKey2,
			wantToken:  `eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29taWMvYXV0aCIsInN1YiI6IjYwNTVhOTZlNDg1MDk4MTMwODJmYjkxMyJ9.M2_ZUyPS9uhzFn92Lief8KvXH4NQWwhvQYJnj_9BFVygqi18WiJe4--ARLqbuNRtTC5ey2L5Hy6JtLJEOa0KBsjAc-ktMLns-vpn8jEJ1Y_UCKG5A3rZ9qjRJhiRxRcW4qE-uriuTQj_0ZtYjyO5zIUqewYxkNYMNxoK6md6TUidXYkt8ES39zvCNBbmzxVx_s6MLXUKJ59bT2-gLG1iumcXwbfLgTkObO4TCjS0D-ayvllOCUT48BpTeTmNexC1xsAeMg2uYYjxx6kK-kYHOhotUNT1lW8hz6sFxe6WmnNgqhkpsO5tk_HqIUhHa29TwLdpoh-I1SNcnLy56k9Wvw`,
			tokenErr:   true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			jwTokenGen, err := NewJWTTokenGen([]byte(c.privateKey), c.iss)
			if err != nil {
				t.Errorf("cannot create jwtTokenGenerator: %v", err)
			}

			jwTokenGen.nowFunc = func() time.Time {
				return time.Unix(c.iat, 0)
			}

			token, err := jwTokenGen.GenerateToken(c.accountID, 2*time.Hour)
			if err != nil {
				t.Errorf("failed to create token: %v", err)
			}

			if c.tokenErr && token == c.wantToken {
				t.Errorf("want error,got no error")
			}

			if !c.tokenErr && token != c.wantToken {
				t.Errorf("wrong token: want: %q; got: %q", c.wantToken, token)
			}
		})
	}
}
