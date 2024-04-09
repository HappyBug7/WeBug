package service

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type Search struct {
}

type Search_Response struct {
	AS struct {
		Results []struct {
			Suggests []struct {
				Txt string `json:"Txt"`
			} `json:"Suggests"`
		} `json:"Results"`
	} `json:"AS"`
}

// type Search_Response struct {
// 	Q string   `json:"q"`
// 	S []string `json:"s"`
// }

func (s *Search) Suggestion_get(q string) (error, []string) {
	q = url.QueryEscape(q)
	url := "https://api.bing.com/qsonhs.aspx?type=cb&q=" + q + "&cb=window.bing.sug"

	resp, err := http.Get(url)
	if err != nil {
		return err, nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}
	// fmt.Println(resp.StatusCode)

	body = body[57 : len(body)-26]

	var search Search_Response
	err = json.Unmarshal(body, &search)
	if err != nil {
		return err, nil
	}
	data := []string{}
	for i := 0; i < len(search.AS.Results); i++ {
		for j := 0; j < len(search.AS.Results[i].Suggests); j++ {
			data = append(data, search.AS.Results[i].Suggests[j].Txt)
		}
	}

	// url := "https://sp0.baidu.com/5a1Fazu8AA54nxGko9WTAnF6hhy/su?wd=" + q + "&json=1"
	// request, err := http.NewRequest("GET", url, nil)
	// if err != nil {
	// 	return err, nil
	// }

	// request.Header.Set("Accept-Charset", "utf-8")

	// client := http.DefaultClient
	// response, err := client.Do(request)
	// if err != nil {
	// 	return err, nil
	// }
	// defer response.Body.Close()

	// body, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	return err, nil
	// }

	// encodedBody, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader(body), unicode.UTF8.NewDecoder()))
	// if err != nil {
	// 	return err, nil
	// }

	// encodedBody = encodedBody[17 : len(encodedBody)-2]

	// fmt.Println(string(encodedBody))

	// var search Search_Response
	// err = json.Unmarshal(encodedBody, &search)
	// if err != nil {
	// 	return err, nil
	// }
	// data := search.S
	return nil, data
}
