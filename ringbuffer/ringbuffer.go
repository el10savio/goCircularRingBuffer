package ringbuffer

import (
	"errors"
	"log"
)

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
	Buffer []interface{}

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
	log.Println("Creating new ring buffer with capacity:", capacity)

	if capacity <= 0 {
		log.Println("Unable to create new ring buffer with negative capacity:", capacity)
		return RingBuffer{}, ErrNotPositiveCapacity
	}

	ringBuffer := RingBuffer{
		Buffer:        make([]interface{}, capacity),
		Capacity:      capacity,
		FillCount:     0,
		WritePosition: 0,
	}

	log.Println("Successfully created new ring buffer:", ringBuffer)

	return ringBuffer, nil
}

// Clear resets the ring buffer
// by resetting both FillCount
// WritePosition to the
// start of the buffer
func (ringBuffer *RingBuffer) Clear() {
	log.Println("Clearing ring buffer:", ringBuffer)

	ringBuffer.FillCount = 0
	ringBuffer.WritePosition = 0

	log.Println("Successfully Cleared ring buffer:", ringBuffer)
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
func (ringBuffer *RingBuffer) Enqueue(value interface{}) (RingBuffer, error) {
	log.Println("Enqueueing value:", value, "to ring buffer:", ringBuffer)

	if ringBuffer.AvailableCapacity() > 0 {
		if ringBuffer.WritePosition >= ringBuffer.Capacity {
			ringBuffer.WritePosition = 0
		}

		ringBuffer.Buffer[ringBuffer.WritePosition] = value

		ringBuffer.WritePosition++
		log.Println("Updated WritePosition:", ringBuffer.WritePosition)

		ringBuffer.FillCount++
		log.Println("Updated FillCount:", ringBuffer.FillCount)

		log.Println("Successfully enqueued value:", value, "Buffer:", ringBuffer)
		return *ringBuffer, nil
	}

	log.Println("Unable to enqueue value:", value, "Buffer is full:", ringBuffer)
	return *ringBuffer, ErrFullBuffer
}

// Dequeue removes the oldest value from the ring buffer
// It does so by reducing the FillCount
// In the case when the buffer is
// empty it returns an error
func (ringBuffer *RingBuffer) Dequeue() (interface{}, error) {
	log.Println("Dequeueing ring buffer:", ringBuffer)

	if ringBuffer.FillCount == 0 {
		log.Println("Unable to dequeue buffer is empty:", ringBuffer)
		return 0, ErrEmptyBuffer
	}

	readPosition := ringBuffer.WritePosition - ringBuffer.FillCount
	if readPosition < 0 {
		readPosition += ringBuffer.Capacity
	}

	value := ringBuffer.Buffer[readPosition]

	ringBuffer.FillCount--
	log.Println("Updated FillCount:", ringBuffer.FillCount)

	log.Println("Successfully dequeued value:", value, "Buffer:", ringBuffer)
	return value, nil
}
