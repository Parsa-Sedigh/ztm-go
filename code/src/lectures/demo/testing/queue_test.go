package queue

import "testing"

func TestAddQueue(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		if len(q.items) != i {
			t.Errorf("Incorrect queue element count: %v, want %v", len(q.items), i)
		}

		// since we have a capacity of 3 in queue and we're only iterating 3 times, this should be always be ture and if not, it means sth is wrong
		if !q.Append(i) {
			t.Errorf("failed to append item %v to queue", i)
		}
	}

	// capacity is 3 and at this point we have already appended 3 items. So we shouldn't be able to append another one to the queue:
	if q.Append(4) {
		t.Errorf("should not be able to add to a full queue")
	}
}

func TestNext(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		q.Append(i)
	}

	for i := 0; i < 3; i++ {
		item, ok := q.Next()

		// we should be able to retrieve all the items because we know there are 3 items in there and the queue has a capacity of 3:
		if !ok {
			t.Errorf("should be able to get item from queue")
		}

		// the item should be equal to i each time it comes out:
		if item != i {
			t.Errorf("got item in wrong order: %v, want %v", item, i)
		}
	}

	// queue should be empty at this point
	item, ok := q.Next()
	if ok {
		t.Errorf("should not be any more items in queue, got: %v", item)
	}
}
