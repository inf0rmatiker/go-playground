package queue

// Queue is our implementation of a FIFO queue
type Queue[K any] struct {

	// array - The underlying array holding our data
	array []K
}

// Size returns the number of items in the queue
func (q *Queue[K]) Size() int {
	return len(q.array)
}

// Enqueue adds an item to the tail of our list
func (q *Queue[K]) Enqueue(item K) {
	q.array = append(q.array, item)
}

// Dequeue removes and returns an item from the head of our list.
// Returns the item and true if successful, otherwise zero-value and false.
func (q *Queue[K]) Dequeue() (K, bool) {
	if len(q.array) > 0 {
		item := q.array[0]
		if len(q.array) == 1 {
			q.array = q.array[:0] // Set back to empty slice
		} else {
			q.array = q.array[1:] // Remove head
		}
		return item, true
	}
	return *new(K), false
}
