package static

import "testing"

func Test_ConverURL(t *testing.T) {
	url := "https://images.dmzj1.com/webpic/4/kdjdnhai20210305a.jpg"
	t.Log(ConverURL(url))
}
