/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 19-09-2018
 * |
 * | File Name:     asset.go
 * +===============================================
 */

package types

// Asset is sensor or actuator that is attached to a thing.
// Asset data type is only applies when data is retried from database.
// Database only recognizes following types:
// boolean, number, string, object and array
type Asset struct {
	Title string `json:"title" bson:"title"`
	Type  string `json:"type" bson:"type"` // boolean, number, string, object, location and array
	Kind  string `json:"kind" bson:"kind"` // sensor or actuator
}
