package main

import (
    "errors"
)
 
// 定义一个简单的栈结构
type Stack struct {
    elements []interface{}
}
 
// 入栈操作
func (s *Stack) Push(element interface{}) {
    s.elements = append(s.elements, element)
}
 
// 出栈操作
func (s *Stack) Pop() (interface{}, error) {
    if len(s.elements) == 0 {
        return nil, errors.New("stack is empty")
    }
    index := len(s.elements) - 1
    element := s.elements[index]
    s.elements = s.elements[:index]
    return element, nil
}
 
// 查看栈顶元素
func (s *Stack) Peek() interface{} {
    if len(s.elements) == 0 {
        return nil
    }
    return s.elements[len(s.elements)-1]
}