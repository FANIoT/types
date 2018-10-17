/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 17-10-2018
 * |
 * | File Name:     error.go
 * +===============================================
 */

package types

// Error is a global error type in I1820 platfrom and all components tries to
// reqturn their error in this format
type Error struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
	Trace string `json:"trace"`
}
