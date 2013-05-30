package agent

import (
	"fmt"
	"strconv"
	"time"
)

import (
	"cfg"
	. "types"
)

//----------------------------------------------- timer work
func timer_work(sess *Session) {
	// check whether the user is logged in
	if !sess.LoggedIn {
		return
	}

	// check building upgrades
	sess.OpCount += sess.Estates.CheckAllCD()

	// TODO: 持久化逻辑#2： 超过一定的时间，刷入数据库
	config := cfg.Get()
	ivl, _ := strconv.Atoi(config["flush_interval"])
	if time.Now().Unix()-sess.LastFlush > int64(ivl) {
		fmt.Println("TODO: flush all to db")
		sess.LastFlush = time.Now().Unix()
	}
}
