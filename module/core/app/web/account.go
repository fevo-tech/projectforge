package web

import (
	"sort"
	"strings"

	"{{{ .Package}}}/app/util"
)

type Account struct {
	Provider string `json:"provider"`
	Email    string `json:"email"`
}

func accountFromString(s string) *Account {
	p, e := util.SplitString(s, ':', true)
	return &Account{Provider: p, Email: e}
}

func (a Account) String() string {
	return a.Provider + ":" + a.Email
}

type Accounts []*Account

func (a Accounts) String() string {
	ret := make([]string, 0, len(a))
	for _, x := range a {
		ret = append(ret, x.String())
	}
	return strings.Join(ret, ",")
}

func (a Accounts) Sort() {
	sort.Slice(a, func(i, j int) bool {
		l := a[i]
		r := a[j]
		if l.Provider == r.Provider {
			return l.Email < r.Email
		}
		return l.Provider < r.Provider
	})
}

func (a Accounts) GetByProvider(p string) Accounts {
	var ret Accounts
	for _, x := range a {
		if x.Provider == p {
			ret = append(ret, x)
		}
	}
	return ret
}

func (a Accounts) Matches(m string) bool {
	return true
}

func (a Accounts) Purge(keys ...string) Accounts {
	ret := make(Accounts, 0, len(a))
	for _, ss := range a {
		hit := false
		for _, key := range keys {
			if ss.Provider == key {
				hit = true
			}
		}
		if !hit {
			ret = append(ret, ss)
		}
	}
	return ret
}

func AccountsFromString(s string) Accounts {
	split := util.SplitAndTrim(s, ",")
	ret := make(Accounts, 0, len(split))
	for _, x := range split {
		ret = append(ret, accountFromString(x))
	}
	return ret
}