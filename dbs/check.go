package dbs

import (
	"database/sql"
	"fmt"
)

func ExecCheck(n int64, result sql.Result, err error) error {
	if err != nil {
		return err
	}

	i, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if i < n {
		return fmt.Errorf("更新校验失败(%d) want:(%d)", i, n)
	}

	return err
}

func ExecICheck(n int64, result sql.Result, err error) (int64, error) {
	if err != nil {
		return -1, err
	}

	i, err := result.RowsAffected()
	if err != nil {
		return i, err
	}
	if i < n {
		return i, fmt.Errorf("更新校验失败(%d)", i)
	}

	return i, err
}
