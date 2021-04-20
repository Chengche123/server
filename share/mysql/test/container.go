package test

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/docker/docker/api/types"
	_ "github.com/docker/docker/client"
	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlPort = 3307
)

// RunWithMongoInDocker create a temp container for mongo test
func RunWithMysqlInDocker(m *testing.M, dsn *string) int {
	// ctx := context.Background()

	// c, err := client.NewClientWithOpts()
	// if err != nil {
	// 	panic(err)
	// }

	// err = c.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	// if err != nil {
	// 	panic(err)
	// }

	// defer func() {
	// 	duration := 1 * time.Second
	// 	err := c.ContainerStop(ctx, containerID, &duration)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }()

	*dsn = fmt.Sprintf("root:root@tcp(127.0.0.1:%d)/test?charset=utf8mb4&parseTime=True&loc=Local", mysqlPort)

	db, err := sql.Open("mysql", *dsn)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("DROP DATABASE IF EXISTS test")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE DATABASE test")
	if err != nil {
		panic(err)
	}

	db.Close()

	return m.Run()
}
