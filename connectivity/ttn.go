/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 24-09-2018
 * |
 * | File Name:     ttn.go
 * +===============================================
 */

package connectivity

// TTN contains information for TheThingsNetwork connectivity
// these information are stored in connectivity section of things and
// used in link component.
type TTN struct {
	ApplicatioID string `json:"applicatio_id"`
	DeviceEUI    string `json:"device_eui"`
}
