package googlemap

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	Key       string
	ClientID  string
	Signature string
	Url       = "https://maps.googleapis.com/maps/api/"
	Format    = "json"
)

func Do(method string, args url.Values) ([]byte, error) {
	u := Url + method + "/" + Format

	if len(Key) > 0 {
		args.Set("key", Key)
	}

	if len(ClientID) > 0 {
		args.Set("client", ClientID)
	}

	if len(Signature) > 0 {
		args.Set("signature", Signature)
	}

	rsp, err := http.Get(u + "?" + args.Encode())
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
