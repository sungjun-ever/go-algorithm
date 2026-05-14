package main

import "fmt"

const ArraySize = 10

// 연결 리스트의 개별 요소
type node struct {
	key   string
	value string
	next  *node
}

// 연결 리스트의 헤드
type bucket struct {
	head *node
}

// 버킷의 배열
type HashTable struct {
	array [ArraySize]*bucket
}

// 간단한 해싱 함수
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

func (h *HashTable) Insert(key string, value string) {
	// hashmap의 인덱스를 구함
	index := hash(key)

	// 비어있다면 초기화
	if h.array[index] == nil {
		h.array[index] = &bucket{}
	}

	// 넣어줌
	h.array[index].insert(key, value)
}

func (h *HashTable) Search(key string) (string, bool) {
	index := hash(key)
	if h.array[index] == nil {
		return "", false
	}

	return h.array[index].search(key)
}

// bucker insert 충돌 시에 앞에 추가함
func (b *bucket) insert(k, v string) {
	if !b.exists(k) {
		newNode := &node{key: k, value: v}
		// 새 노드 다음에 기존 헤드를 넣어주고
		newNode.next = b.head
		// 헤드를 새 노드로 바꿔줌
		b.head = newNode
	} else {
		currentNode := b.head
		for currentNode != nil {
			if currentNode.key == k {
				currentNode.value = v
				return
			}
			currentNode = currentNode.next
		}
	}
}

func (b *bucket) search(k string) (string, bool) {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return currentNode.value, true
		}
		currentNode = currentNode.next
	}
	return "", false
}

func (b *bucket) exists(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// 데이터 충돌이 발생하면 연결 리스트 형식으로 관리한다
func main() {
	myHashTable := &HashTable{}

	// 데이터 삽입
	myHashTable.Insert("my", "isMy")
	myHashTable.Insert("name", "isName")
	myHashTable.Insert("is", "isIs")

	// 데이터 검색
	val, found := myHashTable.Search("name")
	if found {
		fmt.Printf("Found: %s\n", val)
	} else {
		fmt.Println("Not Found")
	}

	// 존재하지 않는 키 검색
	_, found = myHashTable.Search("hello")
	fmt.Printf("hello found? %v\n", found)
}
