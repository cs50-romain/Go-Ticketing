package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/cs50-romain/GoTicketing/heap"
	"github.com/cs50-romain/GoTicketing/ticket"
	//"github.com/cs50-romain/GoTicketing/ticket"
)

var ticket_heap = heap.Init()

func parseCmdLine(cmd string) {
	if cmd == "add" {
		gatherTicketInfo()
	} else if cmd == "show" {
		showTickets()
	}
}

func showTickets() {
	fmt.Println("Company|\tSummary|\tPriority|\tEntered|")
	for _, ticket := range ticket_heap.Data {
		fmt.Printf("%s|\t%s|\t%d|\t%d|\n", ticket.Company, ticket.Summary, ticket.Priority, ticket.Entered.Day())
	}
}

func gatherTicketInfo() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter ticket Priority: ")
	scanner.Scan()
	priority := scanner.Text()
	fmt.Print("Enter ticket summary: ")
	scanner.Scan()
	summary := scanner.Text()
	fmt.Print("Enter company name: ")
	scanner.Scan()
	company := scanner.Text()

	fmt.Printf("Entering ticket for: %s, Priority level: %s. Issue: %s\n", company, priority, summary)

	addTicketToHeap(company, priority, summary)
}

func addTicketToHeap(company, priority, summary string) {
	priorityInt, err := strconv.Atoi(priority)
	if err != nil {
		fmt.Println("[ERROR] Converting ASCII to Int -> ", err)
	}

	ticket := ticket.CreateTicket("Service Desk", company, summary, priorityInt)	
	ticket_heap.Insert(ticket)
}

func main() {
	fmt.Println("Commands:\n- add -> Add a ticket\n- show -> Show list of all tickets")
	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		parseCmdLine(scanner.Text())
	}
}
