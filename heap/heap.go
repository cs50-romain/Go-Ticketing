package heap

import "github.com/cs50-romain/GoTicketing/ticket"

type Heap struct {
	Len int
	Data []*ticket.Ticket
}

func Init() *Heap {
	return &Heap{0, nil}
}

func (h *Heap) Insert(ticket *ticket.Ticket) {
	if h.Len == len(h.Data) {
		h.Data = append(h.Data, ticket)
	} else {
		h.Data[h.Len] = ticket
	}
	h.heapifyUp(h.Len)
	h.Len++
}

func (h *Heap) Pop() *ticket.Ticket {
	if h.Len == 0 {
		return &ticket.Ticket{}
	}

	head := h.Data[0]
	h.Len--
	if h.Len == 0 {
		h.Data = []*ticket.Ticket{}
		return head
	}

	h.Data[0] = h.Data[h.Len]
	h.heapifyDown(0)
	return head
}

func (h *Heap) heapifyDown(idx int) {
	if idx >= h.Len {
		return
	}

	left_idx := leftIdx(idx)
	right_idx := rightIdx(idx)

	if left_idx >= h.Len || right_idx >= h.Len {
		return
	}

	left_ticket := h.Data[left_idx]
	right_ticket := h.Data[right_idx]
	ticket := h.Data[idx]

	if left_ticket.Priority >= right_ticket.Priority && left_ticket.Priority > ticket.Priority {
		h.swap(idx, left_idx)
		h.heapifyDown(left_idx)
	} else if right_ticket.Priority > left_ticket.Priority && right_ticket.Priority > ticket.Priority {
		h.swap(idx, right_idx)
		h.heapifyDown(right_idx)
	}
}

func (h *Heap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}


	parent_idx := parentIdx(idx)
	parent_ticket := h.Data[parent_idx]
	ticket := h.Data[idx]

	if ticket.Priority > parent_ticket.Priority {
		h.swap(idx, parent_idx)
		h.heapifyUp(parent_idx)
	}
}

func (h *Heap) swap(idx, idx_2 int) {
	h.Data[idx], h.Data[idx_2] = h.Data[idx_2], h.Data[idx]	
}

func parentIdx(idx int) int {
	return (idx - 1) / 2
}

func leftIdx(idx int) int {
	return 2 * idx + 1
}

func rightIdx(idx int) int {
	return 2 * idx + 2
}
