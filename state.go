/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 13-08-2018
 * |
 * | File Name:     state.go
 * +===============================================
 */

package types

import "time"

// State represents the asset's state after the decode phase that is done by models.
// This structure is created and remains in the platform for each incoming data.
type State struct {
	Raw   interface{} // The value should be of the same type that is defined in asset’s profile.
	Value struct {
		Boolean bool
		Number  float64
		String  string
		Object  interface{}
		Array   []interface{}
	}
	At time.Time // 2018-09-19T03:04:10Z

	Project string // Project Name
	ThingID string // Thing Identification
	Asset   string // Asset Name
}
