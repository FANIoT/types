/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 07-02-2018
 * |
 * | File Name:     thing/thing.go
 * +===============================================
 */

package types

// Thing contains identification and parent project of a thing
// In this platform, a thing is every device that can connect to the platform directly using the link component provided interfaces.
type Thing struct {
	ID     string   `json:"id" bson:"_id,omitempty"` // thing unique identifier
	Name   string   `json:"name" bson:"name"`        // thing human readable name
	Status bool     `json:"status" bson:"status"`    // active/inactive
	Model  string   `json:"model" bson:"model"`      // model describes how to decode an incoming payload
	Tokens []string `json:"tokens" bson:"tokens"`    // thing access tokens that are generated based on K-Sortable Globally Unique IDs

	Assets         map[string]Asset       `json:"assets" bson:"assets"`                 // thing assets (sensors and actuators)
	Connectivities map[string]interface{} `json:"connectivities" bson:"connectivities"` // supported connectivities (like TheThingsNetwork and etc.)
	Tags           []string               `json:"tags" bson:"tags"`                     // tags associated with things
	Location       struct {
		Type        string    `json:"type" bson:"type"`               // GeoJSON type eg. "Point"
		Coordinates []float64 `json:"coordinates" bson:"coordinates"` // condinates eg. [-73.856077, 40.848447]
	} `json:"location" bson:"location"`

	Project string `json:"project" bson:"project"`
}
