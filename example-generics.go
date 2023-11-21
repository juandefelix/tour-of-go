package main
import "fmt"

type element[T any] struct {
  next *element[T]
  val  T
}

type List[T any] struct {
  head, tail *element[T]
}

func (list *List[T]) Push(value T) {
  newElement := element[T]{val: value}

  if list.tail == nil {
    list.head = &newElement
    list.tail = &newElement
  } else {
    list.tail.next = &newElement
    list.tail = &newElement
  }
}


func (list *List[T]) Trim() {
  last := list.tail
  var second *element[T]

  for item := list.head; item != nil; item = item.next {
    if last == item.next {
      second = item
      break
    }
  }

  list.tail = second
  second.next = nil
}


func (list *List[T]) GetAll() []T {
  var myArray []T
  for element := list.head; element != nil; element = element.next {
    myArray = append(myArray, element.val)
  }

  return myArray
}

func main() {
  list := List[int]{}
  list.Push(10)
  list.Push(11)
  list.Push(12)
  list.Push(13)
  list.Push(14)
  fmt.Println(list.GetAll())
  list.Trim()
  fmt.Println(list.GetAll())
}
