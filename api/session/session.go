package session

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"streaming/api/model"
	"streaming/logger"
	"sync"
)

// 插入session到db
func InsertSession(sid string, ttl int64, uname string) error {
	stmtIns, err := model.DbConn.Prepare("insert into sessions (session_id,ttl,login_name) values(?,?,?)")
	if err != nil {
		logger.Info("insert session err:\t", sid, ttl, uname, err)
		return err
	}

	if _, err = stmtIns.Exec(sid, ttl, uname); err != nil {
		logger.Info("insert session err:\t", sid, ttl, uname, err)
		return err
	}

	defer stmtIns.Close()
	return nil
}

// 根据sid查询session
func RetrieveSession(sid string) (*model.SimpleSession, error) {
	ss := &model.SimpleSession{}
	stmtOut, err := model.DbConn.Prepare("select ttl,login_name from sessions where session_id=?")
	if err != nil {
		logger.Info("retrieve session err:\t", sid, err)
		return nil, err
	}

	var ttl string
	var uname string
	if err = stmtOut.QueryRow(sid).Scan(&ttl, &uname); err != nil && err != sql.ErrNoRows {
		logger.Info("retrieve session err:\t", sid, err)
		return nil, err
	}

	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		ss.TTL = res
		ss.Username = uname
	} else {
		return nil, err
	}

	defer stmtOut.Close()
	return ss, nil
}

// 查询出所有session存储到sync.Map
func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := model.DbConn.Prepare("select * from sessions")
	if err != nil {
		logger.Info("retrieve all sessions err:\t", err)
		return nil, err
	}

	rows, err := stmtOut.Query()
	if err != nil {
		logger.Info("retrieve all sessions err:\t", err)
		return nil, err
	}

	for rows.Next() {
		var id string
		var ttlStr string
		var loginName string

		if err := rows.Scan(&id, &ttlStr, &loginName); err != nil {
			logger.Info("retrieve session error:\t", err)
			break
		}
		if ttl, err := strconv.ParseInt(ttlStr, 10, 64); err == nil {
			ss := &model.SimpleSession{
				Username: loginName,
				TTL:      ttl,
			}
			m.Store(id, ss)
			logger.Info("session id:\t", id, "ttl:\t", ss.TTL)
		}
	}
	return m, nil
}

// 删除session
func DeleteSession(sid string) error {
	stmtOut, err := model.DbConn.Prepare("delete from sessions where session_id = ?")
	if err != nil {
		logger.Info("delete session:\t", sid, err)
		return err
	}
	if _, err := stmtOut.Query(sid); err != nil {
		logger.Info("delete session:\t", sid, err)
		return err
	}
	return nil
}
