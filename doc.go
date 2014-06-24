// Package pools implements an interface similar to sync.Pool, but allows
// for Getting and Putting slices of an object rather than an individual object.
//
// Get allows you to request a slice of a certain size. Each slice size is backed
// by its own pool.
//
// For example:
//
//    var s Suite
//    s.New = func(){return &MyObject{}}
//    s.Get(100) // returns []MyObject of length 100
package pools
