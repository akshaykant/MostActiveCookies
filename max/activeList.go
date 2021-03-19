package max

import (
	"errors"
)

type Active interface {
	ActiveCookie() ([]string, error)
}

type EvalList struct {
	CookieList [][]string
}

type cookieCount struct{
	cookie string
	count int
}

type stack struct{
	list []cookieCount
	top int
}

//hash map - store count of each cookie and stack - to get most occurring cookie
//iterate the list and add each occurrence of cookie in the hash map
//for the return data, create a data structure of most visited cookie
func (evalList EvalList) ActiveCookie()([]string, error){

	count, err := cookieListToMap(evalList.CookieList)

	if err != nil{
		return nil, err
	}

	list := make([]cookieCount, len(evalList.CookieList))
	top := -1

	stack := &stack{
		list,
		top,

	}

	mostOccurrence := stack.cookieOccurrence(count)

	return mostOccurrence, nil
}

func cookieListToMap(cookieList [][]string)(map[string]int, error){

	if len(cookieList) == 0{
		return nil, errors.New("empty list")
	}

	cookieCount := make(map[string]int)

	for _, ele := range cookieList{
		cookieCount[ele[0]] += 1
	}

	return cookieCount, nil
}

//iterate each cookie and peek stack
//push - if empty
//push - if equal entry
//clean stack - add entry
func (stack *stack) cookieOccurrence(cookieMap map[string]int) []string {

	for cookie, count := range cookieMap{
		res, err := stack.peek()

		//stack is empty, add cookies and count to the stack
		if err != nil{
			stack.push(cookieCount{cookie, count})
		} else
		//if peek count entry is equal to the input entry, add to the stack
		if res.count == count {
			stack.push(cookieCount{cookie, count})
		} else
		if res.count < count{
			//set stack to empty and add the new maximum occurrence
			stack.setEmpty()
			stack.push(cookieCount{cookie, count})
		}
	}

	//stack contains all the maximum entries
	if stack.top == -1{
		return []string{}
	}
	occurrence := make([]string, stack.top + 1)
	for i := 0; i <= stack.top; i ++{
		occurrence[i] = stack.list[i].cookie
	}

	return occurrence
}

func (stack *stack) push(count cookieCount){

	stack.top += 1
	stack.list[stack.top] = count

}

func (stack *stack) peek() (cookieCount, error){
	if stack.top < 0 {
		return cookieCount{}, errors.New("empty stack")
	}
	count := stack.list[stack.top]

	return count, nil
}

func (stack *stack) setEmpty(){
	stack.top = -1
}