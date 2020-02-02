package ringbuffer

import (
	"reflect"
	"testing"
)

func TestNewRingBuffer(t *testing.T) {
	capacity := 4

	expectedRingBuffer := RingBuffer{
		Buffer:        make([]interface{}, capacity),
		Capacity:      capacity,
		FillCount:     0,
		WritePosition: 0,
	}
	var expectedErr error

	actualRingBuffer, actualErr := NewRingBuffer(capacity)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedRingBuffer, actualRingBuffer) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedRingBuffer, actualRingBuffer)
	}
}

func TestNewRingBuffer_NegativeCapacity(t *testing.T) {
	capacity := -4

	expectedRingBuffer := RingBuffer{}
	expectedErr := ErrNotPositiveCapacity

	actualRingBuffer, actualErr := NewRingBuffer(capacity)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedRingBuffer, actualRingBuffer) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedRingBuffer, actualRingBuffer)
	}
}

func TestNewRingBuffer_ZeroCapacity(t *testing.T) {
	capacity := 0

	expectedRingBuffer := RingBuffer{}
	expectedErr := ErrNotPositiveCapacity

	actualRingBuffer, actualErr := NewRingBuffer(capacity)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedRingBuffer, actualRingBuffer) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedRingBuffer, actualRingBuffer)
	}
}

func TestClear(t *testing.T) {
	ringBuffer := RingBuffer{
		Buffer:        []interface{}{1, 2, 3, 4, 5},
		Capacity:      5,
		FillCount:     2,
		WritePosition: 4,
	}

	expectedRingBuffer := RingBuffer{
		Buffer:        []interface{}{1, 2, 3, 4, 5},
		Capacity:      5,
		FillCount:     0,
		WritePosition: 0,
	}

	ringBuffer.Clear()
	actualRingBuffer := ringBuffer

	if !reflect.DeepEqual(expectedRingBuffer, actualRingBuffer) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedRingBuffer, actualRingBuffer)
	}
}

func TestEnqueue(t *testing.T) {
	ringBuffer, err := NewRingBuffer(2)
	if err != nil {
		t.Fatalf("Error in creating ring buffer: %v", err)
	}

	expectedRingBuffer := RingBuffer{
		Buffer:        []interface{}{1, nil},
		Capacity:      2,
		FillCount:     1,
		WritePosition: 1,
	}
	var expectedErr error

	actualRingBuffer, actualErr := ringBuffer.Enqueue(1)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedRingBuffer, actualRingBuffer) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedRingBuffer, actualRingBuffer)
	}
}

func TestEnqueue_FullBuffer(t *testing.T) {
	ringBuffer, err := NewRingBuffer(2)
	if err != nil {
		t.Fatalf("Error in creating ring buffer: %v", err)
	}

	expectedRingBuffer := RingBuffer{
		Buffer:        []interface{}{1, "a"},
		Capacity:      2,
		FillCount:     2,
		WritePosition: 2,
	}
	expectedErr := ErrFullBuffer

	_, actualErr := ringBuffer.Enqueue(1)
	if actualErr != nil {
		t.Fatalf("Error in enqueue ring buffer: %v", err)
	}

	_, actualErr = ringBuffer.Enqueue("a")
	if actualErr != nil {
		t.Fatalf("Error in enqueue ring buffer: %v", err)
	}

	actualRingBuffer, actualErr := ringBuffer.Enqueue(1)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedRingBuffer, actualRingBuffer) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedRingBuffer, actualRingBuffer)
	}
}

func TestDequeue(t *testing.T) {
	ringBuffer, err := NewRingBuffer(2)
	if err != nil {
		t.Fatalf("Error in creating ring buffer: %v", err)
	}

	expectedRingBuffer := RingBuffer{
		Buffer:        []interface{}{1, "a"},
		Capacity:      2,
		FillCount:     1,
		WritePosition: 2,
	}
	expectedValue := 1
	var expectedErr error

	_, actualErr := ringBuffer.Enqueue(1)
	if actualErr != nil {
		t.Fatalf("Error in enqueue ring buffer: %v", err)
	}

	_, actualErr = ringBuffer.Enqueue("a")
	if actualErr != nil {
		t.Fatalf("Error in enqueue ring buffer: %v", err)
	}

	actualValue, actualErr := ringBuffer.Dequeue()
	actualRingBuffer := ringBuffer

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedRingBuffer, actualRingBuffer) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedRingBuffer, actualRingBuffer)
	}

	if !reflect.DeepEqual(expectedValue, actualValue) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedValue, actualValue)
	}

}

func TestDequeue_EmptyBuffer(t *testing.T) {
	ringBuffer, err := NewRingBuffer(2)
	if err != nil {
		t.Fatalf("Error in creating ring buffer: %v", err)
	}

	expectedRingBuffer := RingBuffer{
		Buffer:        []interface{}{nil, nil},
		Capacity:      2,
		FillCount:     0,
		WritePosition: 0,
	}
	expectedValue := 0
	expectedErr := ErrEmptyBuffer

	actualValue, actualErr := ringBuffer.Dequeue()
	actualRingBuffer := ringBuffer

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedRingBuffer, actualRingBuffer) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedRingBuffer, actualRingBuffer)
	}

	if !reflect.DeepEqual(expectedValue, actualValue) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedValue, actualValue)
	}
}
