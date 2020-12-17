/*
 * A Blog System
 *
 * send, manage and share your blog 
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type UserInfo struct {

	ID string `json:"ID,omitempty"`

	Name string `json:"Name,omitempty"`

	Bio string `json:"Bio,omitempty"`

	Level float32 `json:"Level,omitempty"`
}
