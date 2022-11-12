package server

import "net/url"

func NewEndpoints(ep []string) []*url.URL {
		rs := make([]*url.URL, 0)
		for _, e := range ep {
				u, err := url.Parse(e)
				if err != nil {
						panic(err)
				}
				rs = append(rs, u)
		}
		return rs

}

