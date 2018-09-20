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
	Raw   interface{} `json:"raw" bson:"raw"` // The value should be of the same type that is defined in asset’s profile.
	Value struct {
		Boolean bool          `json:"boolean,omitempty" bson:"boolean,omitempty"`
		Number  float64       `json:"number,omitempty" bson:"number,omitempty"`
		String  string        `json:"string,omitempty" bson:"string,omitempty"`
		Object  interface{}   `json:"object,omitempty" bson:"object,omitempty"`
		Array   []interface{} `json:"array,omitempty" bson:"array,omitempty"`
	} `json:"value" bson:"value"`
	At time.Time `json:"at" bson:"at"` // 2018-09-19T03:04:10Z

	Project string `json:"project" bson:"project"`   // Project Name
	ThingID string `json:"thing_id" bson:"thing_id"` // Thing Identification
	Asset   string `json:"asset" bson:"asset"`       // Asset Name
}
