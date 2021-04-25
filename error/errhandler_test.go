package error

import (
	"database/sql"
	"fmt"
	"testing"
)

//1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

//没有该行不算程序的错误,但是错误打印应该出来

func TestDaoError(t *testing.T) {
	err := DaoError()
	if err == sql.ErrNoRows {
		fmt.Println(err)
	} else if err != nil {
		fmt.Println("dao failed", err)
	}
	return
}

func DaoError() error {
	err := sql.ErrNoRows
	return err
}
