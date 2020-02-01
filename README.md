# goCircularRingBuffer
A bounded circular ring buffer queue implemented in Go

Here we can create a bounded circular ring buffer queue that can quickly add and read values as it is implemented in a ring like fashion.

**Methodology**

It works by creating a bounded slice that is initially filled ith zeros. When elements are added they replace the old values.

A writePosition marker and FillCount is also stored which tracks the elements to be read and inserted next.
