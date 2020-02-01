package ringbuffer

import "errors"

var (
	// ErrNotPositiveCapacity error for capacity cannot be 0 or negative
	ErrNotPositiveCapacity = errors.New("capacity cannot be 0 or negative")

	// ErrFullBuffer error for buffer is full
	ErrFullBuffer = errors.New("buffer is full")

	// ErrEmptyBuffer error for buffer is empty
	ErrEmptyBuffer = errors.New("buffer is empty")
)

// RingBuffer is the type of buffer
// used to contain the ring queue
type RingBuffer struct {
	// Buffer stores the elements
	Buffer []int

	// Capacity is the max number of
	// elements the buffer can hold
	Capacity int

	// FillCount is the number of elements
	// in the buffer that are filled
	FillCount int

	// WritePosition is the last known
	// position of a written value
	// in the buffer
	WritePosition int
}

// NewRingBuffer initializes a new ring buffer
// the capacity of the buffer is given. As
// the buffer is created as an int slice
// all the values in are initialized
// to 0 by default
func NewRingBuffer(capacity int) (RingBuffer, error) {
	if capacity <= 0 {
		return RingBuffer{}, ErrNotPositiveCapacity
	}

	ringBuffer := RingBuffer{
		Buffer:        make([]int, capacity),
		Capacity:      capacity,
		FillCount:     0,
		WritePosition: 0,
	}

	return ringBuffer, nil
}

// Clear resets the ring buffer
// by resetting both FillCount
// WritePosition to the
// start of the buffer
func (ringBuffer *RingBuffer) Clear() {
	ringBuffer.FillCount = 0
	ringBuffer.WritePosition = 0
}

// AvailableCapacity returns how many elements
// in the ring buffer are free to be filled
func (ringBuffer *RingBuffer) AvailableCapacity() int {
	return ringBuffer.Capacity - ringBuffer.FillCount
}

// Enqueue adds a value to the ring buffer
// In the case when the buffer is full
// it returns an error, otherwise,
// it adds to the next available
// position and updates the
// WritePosition and FillCount
func (ringBuffer *RingBuffer) Enqueue(value int) (RingBuffer, error) {
	if ringBuffer.AvailableCapacity() > 0 {
		if ringBuffer.WritePosition >= ringBuffer.Capacity {
			ringBuffer.WritePosition = 0
		}

		ringBuffer.Buffer[ringBuffer.WritePosition] = value

		ringBuffer.WritePosition++
		ringBuffer.FillCount++

		return *ringBuffer, nil
	}

	return *ringBuffer, ErrFullBuffer
}

// Dequeue removes the oldest value from the ring buffer
// It does so by reducing the FillCount
// In the case when the buffer is
// empty it returns an error
func (ringBuffer *RingBuffer) Dequeue() (int, error) {
	if ringBuffer.FillCount == 0 {
		return 0, ErrEmptyBuffer
	}

	readPosition := ringBuffer.WritePosition - ringBuffer.FillCount
	if readPosition < 0 {
		readPosition += ringBuffer.Capacity
	}

	value := ringBuffer.Buffer[readPosition]
	ringBuffer.FillCount--

	return value, nil
}
