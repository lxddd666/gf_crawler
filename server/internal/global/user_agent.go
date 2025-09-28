package global

import "math/rand"

func RandomUserAgent() string {
	return UserAgentList[rand.Intn(len(UserAgentList))]
}
