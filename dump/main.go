package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var c int // current count of the number of elements in the list

type LinkedList struct {
	number int
	// count  int // Adding a count means always traversing the list to find the value
	next *LinkedList
}

/* ##############################
######## error handling #########
#################################
*/

type customErrors struct {
	num  int
	part string
	msg  string
}

func (err customErrors) NilUtil(head ...**LinkedList) error {
	if *head[0] == nil {
		log.Printf("Nil error: %d, %s, %s", err.num, err.part, err.msg)
		return errors.New("Empty List Error:")
	}
	return nil
}

func (err customErrors) SpecError(head ...**LinkedList) {
	log.Println("Error: ", err.num, err.part, err.msg)
}

/* ##############################
#### end of error handling ######
#################################
*/

func insert(head **LinkedList, num int) {

	if *head == nil {
		*head = &LinkedList{number: num}
		c = 1
		log.Printf("Inserted the first item: %d", num)
		displayCount()

	} else {
		ptr := *head
		for ptr.next != nil {
			ptr = ptr.next // &*x will be simplified to x. It will not copy x.
		}

		ptr.next = &LinkedList{
			number: num,
		}
		c++
		log.Printf("Added item: %d", num)
		displayCount()
	}
}

func display(head **LinkedList) {
	if *head == nil {
		log.Println("There was nothing to display")
		displayCount()
		return
	}

	ptr := *head
	for ptr != nil {
		fmt.Printf("%d -> ", (*ptr).number)
		ptr = ptr.next
	}

	fmt.Println()
	displayCount()
}

func displayCount() {
	log.Printf("Current count: %d", c)
}

func deleteFirst(head **LinkedList) {
	if *head != nil {
		*head = (*head).next
		c--
		fmt.Println("Deleted head of the list, count = ", c)

	} else {
		fmt.Println("The list is empty")
	}
}

func deleteLast(head **LinkedList) {
	// for an empty list, the type will be nil
	if *head == nil {
		fmt.Println("List is empty")
		return
	}

	// delete when there is only one
	if (*head).next == nil {
		*head = nil
		return
	}

	// delete from second to that last element
	p := *head
	for p.next.next != nil {
		fmt.Println(p, p.next)
		p = p.next
	}
	p.next = nil

}

func insertAtLocation(head **LinkedList, num int, position int) {
	if *head == nil {
		fmt.Println("The list is empty")
		return
	}

	if position < c { // indices start from 0
		for i := 0; i <= position; i++ { // Struct has no indexes, establish that the index starts from 0
			head = &(*head).next // traversing with the pointer without changing the values, address used
		}
		// we are at the position where to insert
		newNode := &LinkedList{
			number: num,
		}
		newNode.next = *head
		*head = newNode
		c = c + 1
		fmt.Println("New count:", c)

	} else {
		fmt.Printf("There is no position %d in the list", position)
		return
	}

}

func deleteAtLocation(head **LinkedList, position int) {
	if *head == nil {
		fmt.Println("The List is empty")
		return
	}

	if position < c {
		for i := 0; i < position; i++ {
			head = &(*head).next
		}
		*head = (*head).next
		c = c - 1
		displayCount()

	} else {
		fmt.Printf("There is no position %d in the list", position)
	}
}

func insertAtBegining(head **LinkedList, num int) {
	newNode := &LinkedList{
		number: num,
	}
	newNode.next = *head // point to new head
	*head = newNode      // update head to new node
	displayCount()
}

func updateValue(head **LinkedList, num int, position int) {
	if *head == nil {
		fmt.Println("The list is still empty")
		return
	}

	for i := 0; i < position; i++ {
		head = &(*head).next
	}

	(*head).number = num
	displayCount()
}

func updateSameValues(head **LinkedList, oldNum int, newNum int) {
	if *head == nil {
		log.Println("The list is still nil")
		return
	}

	for *head != nil {
		if (*head).number == oldNum {
			(*head).number = newNum
		}
		head = &(*head).next
	}
	log.Printf("Updated %d to %d", oldNum, newNum)

}

func parseMaps(head **LinkedList, part int, mp string) map[int]int {
	// check nil list with custom error handler...
	nils := customErrors{
		num: part,
		msg: errors.New("parseMaps").Error(),
	}
	nils.NilUtil(head)          // replace with if *head != nil
	newMap := make(map[int]int) // instantiate map

	// check if the string is empty
	if strings.TrimSpace(mp) == "" {
		log.Println("Nothing received")
		return nil
	}

	// split at the space this return a slice and we extract pairs
	// pairs := strings.Split(mp, "\n")
	for _, pair := range strings.Split(mp, "\n") {
		ex := strings.Split(strings.TrimSpace(pair), ",")
		tmp := strings.Split(ex[0], ":")

		if len(tmp) != 2 {
			log.Printf("%v is not a pair", tmp)
			continue
		}

		key := strings.TrimSpace(tmp[0])
		k, err := strconv.Atoi(key)
		if err != nil {
			log.Printf("Key Error: %v %v", k, err.Error())
		}

		val := strings.TrimSpace(tmp[1])
		v, err := strconv.Atoi(val)
		if err != nil {
			log.Printf("Value error: %v %v", v, err.Error())
		}

		newMap[k] = v
	}
	return newMap
}

func updateValuesAtIndex(head **LinkedList, part int, mp map[int]int) {
	nils := &customErrors{
		num: part,
		msg: errors.New("updateValuesAtIndex").Error(),
	}
	nils.NilUtil(head)

	count := 0

	for *head != nil {
		if val, ok := mp[count]; ok {
			(*head).number = val
		}
		head = &(*head).next
		count++
	}
}

func deleteValuesAtIndex(head **LinkedList, part int, mp map[int]int) {
	nils := &customErrors{
		num: part,
		msg: errors.New("deleteValuesAtIndex").Error(),
	}
	nils.NilUtil(head)

	count := 0

	for *head != nil {
		if _, ok := mp[count]; ok {
			*head = (*head).next
			c--
			log.Printf("Deleted %d", mp[count])
			continue
		}

		head = &(*head).next
		count++
	}
	displayCount()
}

func deleteSimilarValues(head **LinkedList, value int) {
	for *head != nil {
		if (*head).number == value {
			*head = (*head).next
			fmt.Println("Deleted", value)
			c--
		} else {
			head = &(*head).next
		}

	}
}

func parseArray(head **LinkedList, part int, list string) []int {
	nils := &customErrors{
		num: part,
		msg: errors.New("parseArray").Error(),
	}
	nils.NilUtil(head)

	if strings.TrimSpace(list) == "" {
		log.Println("Nothing received")
		return nil
	}

	var items []int

	tmp := strings.TrimSpace(list)
	l := strings.Split(tmp, ",")
	for _, k := range l {
		v, _ := strconv.Atoi(strings.TrimSpace(k))
		items = append(items, v)
	}
	return items

}

func getValuesAtPositions(head **LinkedList, part int, del []int) {
	// downside hear (On^2) when both del and head are large
	nils := &customErrors{
		num: part,
		msg: errors.New("getValuesAtPositions").Error(),
	}
	nils.NilUtil(head)

	i := 0
	chead := *head // making a copy of the head

	for chead != nil {
		for _, v := range del {
			if v == i {
				fmt.Printf("Key: %T Value: %v \n", v, chead.number)
			}
		}
		i++
		chead = chead.next
	}
}

func printAndDestroy(head **LinkedList) {
	for *head != nil {
		fmt.Println((*head).number)
		*head = (*head).next
		c--
	}
	// fmt.Println((*head).number) this is a panic, no need to stretch the search
	log.Println("---Nothing to print----")
	displayCount()
}

func main() {
	var head *LinkedList
	nils := customErrors{} // instantiate errors struct

	for true {
		fmt.Println("\nEnter your choice")
		fmt.Println("1. Insert value in linked list (rear end)")
		fmt.Println("2. Display linked list")
		fmt.Println("3. Deleting from begining")
		fmt.Println("4. Delete from last node")
		fmt.Println("5. Deleting from specific location")
		fmt.Println("6. insert at specific location")
		fmt.Println("7. insert at begining (front end)")
		fmt.Println("8. Update node")
		fmt.Println("9. Update similar values")
		fmt.Println("10. Display count")
		fmt.Println("11. Update multiple values with indexes")
		fmt.Println("12. Delete multiple values with indexes")
		fmt.Println("13. Get multiple values with indexes")
		fmt.Println("14. Delete Similar values")
		fmt.Println("15. Indepotency")
		fmt.Println("0. Exit")

		reader := bufio.NewReader(os.Stdin) // fmt.Scan(&input) negated for its way of handling spaces
		raw, err := reader.ReadString('\n')
		choice, err := strconv.Atoi(strings.TrimSpace(raw))

		if err != nil {
			nils = customErrors{
				num:  -1,
				part: "Choises",
				msg:  err.Error(),
			}
			nils.SpecError(&head)
			continue // continue the loop after the error is known
		}
		fmt.Println("Your choice:", choice)

		switch choice {
		case 1:
			var data string
			fmt.Println("Enter your value for linked list node:")
			fmt.Scanln(&data)
			if num, err := strconv.Atoi(data); err == nil {
				insert(&head, num)
				continue
			} else {
				log.Printf("Error: %v", err)
				continue
			}

		case 2:
			display(&head)
			continue

		case 3:
			deleteFirst(&head)
			continue

		case 4:
			deleteLast(&head)
			continue

		case 5:
			var pos string
			fmt.Println("Enter index to delete")
			fmt.Scan(&pos)
			num, _ := strconv.Atoi(pos)

			deleteAtLocation(&head, num)
			continue

		case 6:
			var val, pos string
			fmt.Println("Type a new value")
			fmt.Scan(&val)
			fmt.Println("Type the index to insert at")
			fmt.Scan(&pos)
			num, _ := strconv.Atoi(val)
			position, _ := strconv.Atoi(pos)

			insertAtLocation(&head, num, position)
			continue

		case 7:
			var data string
			fmt.Println("Enter number you want to insert at the begining of the list")
			fmt.Scan(&data)
			num, _ := strconv.Atoi(data)

			insertAtBegining(&head, num)
			continue

		case 8:
			var number, position string
			fmt.Println("Enter the new value")
			fmt.Scan(&number)
			num, err := strconv.Atoi(number)
			if err != nil {
				log.Println("Value error:", err.Error())
				continue
			}

			fmt.Println("Enter the position in which to insert")
			fmt.Scan(&position)
			pos, err := strconv.Atoi(position)
			if err != nil {
				log.Println("Position error:", err.Error())
				continue
			}

			updateValue(&head, num, pos)
			continue

		case 9:
			var oldData, newData string
			fmt.Println("Enter the old value")
			fmt.Scan(&oldData)
			oldNum, err := strconv.Atoi(oldData)
			if err != nil {
				log.Println("Old Value error:", err.Error())
				continue
			}

			fmt.Println("Enter the new value")
			fmt.Scan(&newData)
			newNum, err := strconv.Atoi(newData)
			if err != nil {
				log.Println("New Value error:", err.Error())
				continue
			}

			updateSameValues(&head, oldNum, newNum)
			continue
		case 10:
			displayCount()
			continue

		case 11:
			fmt.Println("Update values 1:4234, 2:4354:")
			reader := bufio.NewReader(os.Stdin)
			raw, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Value Error: ", err.Error()) // make it custom
				continue
			}
			raw = strings.TrimSpace(raw)

			mp := parseMaps(&head, choice, raw)
			fmt.Println("map==", mp)
			updateValuesAtIndex(&head, choice, mp)
			continue

		case 12:
			fmt.Println("Delete values 1:4234, 2:4354:")
			reader := bufio.NewReader(os.Stdin)
			raw, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Value Error: ", err.Error())
				continue
			}
			raw = strings.TrimSpace(raw)

			mp := parseMaps(&head, choice, raw)
			deleteValuesAtIndex(&head, choice, mp)
			continue

		case 13:
			fmt.Println("Get multiple values")
			reader := bufio.NewReader(os.Stdin)
			raw, _ := reader.ReadString('\n')

			list := parseArray(&head, choice, raw)
			getValuesAtPositions(&head, choice, list)
			continue

		case 14:
			fmt.Println("Delete a value from the list")
			reader := bufio.NewReader(os.Stdin)
			key, _ := reader.ReadString('\n')
			num, _ := strconv.Atoi(strings.TrimSpace(key))
			deleteSimilarValues(&head, num)
			continue

		case 15:
			fmt.Println("Self desctruction... y/n")
			reader := bufio.NewReader(os.Stdin)
			val, _ := reader.ReadString('\n')
			decision := strings.TrimSpace(val)

			if decision == "y" {
				fmt.Println("Starting the nuke...")
				printAndDestroy(&head)
			} else {
				fmt.Println("Continuing with previous state...")
				continue
			}
			continue

		case 0:
			fmt.Println("Gracefully exiting...")
			os.Exit(0)

		default:
			// keeps some variables in memory with the restart
			fmt.Printf("Really, %d, that's not a choice, begin, AGAIN!!! :)", choice)
			continue

		}
		break
	}
}
