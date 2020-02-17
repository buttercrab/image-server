package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"time"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	a := rightCount[h[i]] + wrongCount[h[i]] + tempCount[h[i]]
	b := rightCount[h[j]] + wrongCount[h[j]] + tempCount[h[j]]
	if a == b {
		return h[i] < h[j]
	}
	return a < b
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	location[h[i]] = i
	location[h[j]] = j
}

func (h *IntHeap) Push(element interface{}) {
	location[element.(int)] = len(*h)
	*h = append(*h, element.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	element := old[n-1]
	*h = old[0 : n-1]
	location[element] = -1
	return element
}

var list []string
var rightCount []int
var wrongCount []int
var tempCount []int
var timeToPop []time.Time
var pq IntHeap
var location []int
var rank map[string]int
var rankSecret map[string]string
var q Queue
var size int

func Init() {
	file, err := os.Open("data/url.txt")
	if err != nil {
		panic(err)
		return
	}
	defer file.Close()

	list = make([]string, 0)
	pq = make(IntHeap, 0)

	cnt := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		list = append(list, scanner.Text())
		pq = append(pq, cnt)
		cnt++
	}

	size = cnt

	rank = make(map[string]int)
	rankSecret = make(map[string]string)
	location = make([]int, cnt)
	rightCount = make([]int, cnt)
	wrongCount = make([]int, cnt)
	tempCount = make([]int, cnt)
	timeToPop = make([]time.Time, cnt)
	for i := 0; i < cnt; i++ {
		location[i] = i
	}
	q = Queue{len: 0, head: nil, tail: nil}

	rank["Hello"] = 1
	rank["World"] = 2
	rank["BABA"] = 1
	rank["sss"] = 0
	rank["sssa"] = 0
	rank["sssb"] = 0
	rank["sssc"] = 0
	rank["sssd"] = 0
	rank["ssse"] = 0
	rank["sssf"] = 0
	rank["sssg"] = 0
	rank["sssh"] = 0
	rank["sssi"] = 0
	rank["sssj"] = 0
	rank["sssk"] = 0
	rank["sssl"] = 0
	rank["sssm"] = 0
}

// in goroutine
func Update() {
	for {
		if q.len != 0 {
			x := q.Pop()
			timeToPop[x] = timeToPop[x].Add(time.Minute * 10)
			time.Sleep(time.Until(timeToPop[x]))
			tempCount[x]--
			heap.Fix(&pq, location[x])
		}
	}
}

func GetNew() (int, string) {
	x := pq[0]
	tempCount[x]++
	timeToPop[x] = time.Now()
	q.Push(x)
	heap.Fix(&pq, location[x])
	return x, list[x]
}

func SetValue(id, value int, name string) {
	if id < 0 || id >= size {
		return
	}
	if name != "" {
		rank[name]++
	}
	if value == 1 {
		rightCount[id]++
	} else {
		wrongCount[id]++
	}
	heap.Fix(&pq, location[id])
}

func saveResult() {
	file, err := os.OpenFile("data/output.txt", os.O_WRONLY, os.FileMode(0644))
	if err != nil {
		log.Printf("Save Error: %s\n", err.Error())
		return
	}
	defer file.Close()

	for i := 0; i < size; i++ {
		_, err = file.WriteString(fmt.Sprintf("%d %d\n", rightCount[i], wrongCount[i]))
	}

	log.Println("Result Saved")
}

func saveRanking() {
	file, err := os.OpenFile("data/ranking.txt", os.O_WRONLY, os.FileMode(0644))
	if err != nil {
		log.Printf("Save Error: %s\n", err.Error())
		return
	}
	defer file.Close()

	for k, v := range rank {
		_, _ = file.WriteString(fmt.Sprintf("%s %s %d\n", k, rankSecret[k], v))
	}

	log.Println("Ranking Saved")
}

func Save() {
	for {
		go saveResult()
		go saveRanking()
		time.Sleep(time.Minute * 5)
	}
}

func Register(name, secret string) bool {
	if rankSecret[name] == "" {
		rankSecret[name] = secret
		return true
	}
	return rankSecret[name] == secret
}
