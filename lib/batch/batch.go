package batch

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	res = make([]user, 0, n)
	errGroup := new(errgroup.Group)
	errGroup.SetLimit(int(pool))

	var (
		mutex sync.Mutex
		index int64
	)

	for index = 0; index < n; index++ {
		temp := index

		errGroup.Go(func() error {
			u := getOne(temp)
			mutex.Lock()
			defer mutex.Unlock()
			fmt.Println("Go index = ", temp)
			res = append(res, u)
			return nil
		})
	}

	if err := errGroup.Wait(); err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(fmt.Println("Done!", res))

	return res
}
