package main

import (
	"fmt"
	"math/rand"
)

type Job struct {
	Id      int
	RandNum int
}

type Result struct {
	job *Job
	sum int
}

func createPool(num int, jobchan <-chan *Job, resultchan chan<- *Result) {
	for i := 0; i < num; i++ {
		go func(jobchan <-chan *Job, resultchan chan<- *Result) {
			for job := range jobchan {
				r_num := job.RandNum
				var sum int
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}

				r := &Result{
					job: job,
					sum: sum,
				}
				resultchan <- r
			}
		}(jobchan, resultchan)
	}
}

func pool() {
	jobchan := make(chan *Job, 2)
	resultchan := make(chan *Result, 2)
	createPool(1, jobchan, resultchan)
	go func(resultchan <-chan *Result) {
		for result := range resultchan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id, result.job.RandNum, result.sum)
		}
	}(resultchan)

	var id int
	for {
		id++
		r_num := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: r_num,
		}
		jobchan <- job
	}
}
