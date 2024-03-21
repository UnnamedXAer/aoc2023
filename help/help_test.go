package help_test

import (
	"testing"

	"github.com/unnamedxaer/aoc2023/help"
)

func TestPQ(t *testing.T) {

	q := help.NewPQAny[int](3)

	if !q.IsEmpty() {
		t.Fatal("queue should be initially empty")
	}

	firstElement := help.PQItem[int]{Key: 3, Priority: 5}

	q.EnqueueItem(firstElement)
	if q.IsEmpty() {
		t.Fatal("queue should not be empty")
	}

	top := q.Top()

	if top != firstElement {
		t.Fatalf("top should return the only element we have: want: %+v, got: %+v", firstElement, top)
	}

	topPriorityElement := help.PQItem[int]{Key: -10, Priority: 11}

	q.EnqueueItem(topPriorityElement)
	q.EnqueueItem(help.PQItem[int]{})
	q.EnqueueItem(help.PQItem[int]{Key: 100, Priority: 10})
	q.EnqueueItem(help.PQItem[int]{Key: 88, Priority: 11})

	greatest := q.Dequeue()
	if greatest != topPriorityElement {
		t.Fatalf("expected element with greatest priority (in case on tie, first in first out), want: %+v, got: %+v", topPriorityElement, greatest)
	}

	greatest = q.Dequeue()
	if greatest.Priority != 11 {
		t.Fatalf("expected to get second element with priority 11, got: %+v", greatest)
	}

	top = q.Top()
	greatest = q.Dequeue()
	if top != greatest {
		t.Fatalf("expected Dequeue to yield the result as Top, top: %+v, dequeued: %+v", top, greatest)
	}
}
