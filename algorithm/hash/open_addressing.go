package main

import (
	"errors"
	"fmt"
)

const TableSize = 10

// 해시 테이블의 개별 슬롯
type Entry struct {
	Key      string
	Value    string
	IsFilled bool
}

// 개방 주소법 테이블 구조체
type OpenAddressingTable struct {
	Array [TableSize]*Entry
	Size  int
}

// 간단한 해시 함수
func hashFunc(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % TableSize
}

func (t *OpenAddressingTable) Insert(key, value string) error {
	if t.Size >= TableSize {
		return errors.New("테이블이 가득 참")
	}

	index := hashFunc(key)

	// 빈 슬롯을 찾을 때까지 선형적으로 탐색한다.
	for t.Array[index] != nil && t.Array[index].IsFilled || t.Array[index].Key == "<deleted>" {
		index = (index + 1) % TableSize
	}

	t.Array[index] = &Entry{Key: key, Value: value, IsFilled: true}
	t.Size++
	return nil
}

func (t *OpenAddressingTable) Search(key string) (string, bool) {
	index := hashFunc(key)
	startIndex := index

	// 선형 탐색
	for t.Array[index] != nil {
		if t.Array[index].Key == key && t.Array[index].IsFilled {
			return t.Array[index].Value, true
		}
		index = (index + 1) % TableSize

		// 한 바퀴 다 돌았는데 못 찾은 경우
		if index == startIndex {
			break
		}
	}

	return "", false
}

func (t *OpenAddressingTable) Delete(key string) error {
	index := hashFunc(key)
	startIndex := index

	for t.Array[index] != nil {
		if t.Array[index].Key == key && t.Array[index].IsFilled && t.Array[index].Key != "<deleted>" {
			t.Array[index].Key = "<deleted>"
			t.Array[index].Value = ""
			return nil
		}

		index = (index + 1) % TableSize

		if index == startIndex {
			break
		}
	}

	return errors.New("삭제할 데이터가 없습니다.")
}

// 연결리스트를 사용하지 않고, 해시 테이블 자체의 반공간을 활용한다.
// 충돌이 발생하면 빈공간을 찾아간다.
// 테이블이 어느 정도 차면 성능이 떨어지고, 배열 크기를 늘리고 전체를 다시 해싱하는 리해싱이 필수적이다
func main() {
	table := &OpenAddressingTable{}

	table.Insert("name", "isName")
	table.Insert("omg", "mooooooooo")
	table.Insert("name", "noName")

	if val, found := table.Search("name"); found {
		fmt.Printf("Found 'name': %s\n", val)
	}

	if val, found := table.Search("omg"); found {
		fmt.Printf("Found 'omg': %s\n", val)
	}
}
