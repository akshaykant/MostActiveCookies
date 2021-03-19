package search

import (
	"errors"
	"time"
)

//interface to accept different visitors
type Search interface {
	GetList() ([][]string, error)
}

type EvalList struct {
	CookieList [][]string
	Date       string
}

func (eval EvalList) GetList() ([][]string, error) {

	//range over timestamp and perform the binary search on each timestamp to find the first and last position of range
	//for the resulted range of searched elements, extract the cookies for those ids and add it to the result
	//cookie - cookieList[0], timestamp - cookieList[1]
	start, end, err := count(eval.CookieList, eval.Date, len(eval.CookieList))

	if err != nil {
		return nil, err
	}

	return eval.CookieList[start : end+1], nil

}

func count(cookieList [][]string, date string, len int) (int, int, error) {
	//index of first and last occurrence of the date element
	var start, end int

	//get the first occurrence of date
	start = first(cookieList, date, 0, len-1, len)

	// If date doesn't exist in cookieList[] then return error
	if start == -1 {
		return 0, 0, errors.New("No Occurrence of date ")
	}

	/* Else get the index of last occurrence of date.
	   Note that we are only looking in the
	   subarray after first occurrence */
	end = last(cookieList, date, start, len-1, len)

	return start, end, nil
}

/* if date is present in cookieList[] then returns the
   index of FIRST occurrence of date in cookieList[0..n-1],
   otherwise returns -1 */
func first(cookieList [][]string, date string, low int, high int, n int) int {
	if high >= low {

		mid := (low + high) / 2
		if (mid == 0 || toDateFormat(date).Before(toDateFormat(cookieList[mid-1][1]))) && cookieList[mid][1] == date {
			return mid
		} else if toDateFormat(date).Before(toDateFormat(cookieList[mid][1])) {
			return first(cookieList, date, mid+1, high, n)
		} else {
			return first(cookieList, date, low, mid-1, n)
		}
	}
	return -1
}

/* if date is present in cookieList[] then returns the
   index of LAST occurrence of date in cookieList[0..n-1],
   otherwise returns -1 */
func last(cookieList [][]string, date string, low int, high int, n int) int {
	if high >= low {

		mid := (low + high) / 2
		if (mid == n-1 || toDateFormat(date).After(toDateFormat(cookieList[mid+1][1]))) && (cookieList[mid][1] == date) {
			return mid
		} else if toDateFormat(date).After(toDateFormat(cookieList[mid][1])) {
			return last(cookieList, date, low, mid-1, n)
		} else {
			return last(cookieList, date, mid+1, high, n)
		}
	}
	return -1
}

func toDateFormat(date string) time.Time {
	var t time.Time

	//Parse a time value from string in yyyy-mm-dd format
	t, _ = time.Parse("2006-01-02", date)

	return t
}
