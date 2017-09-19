/*
* @Author: wang
* @Date:   2017-09-13 11:11:41
* @Last Modified by:   wang
* @Last Modified time: 2017-09-19 18:09:04
 */
package goframe

type Config struct {
	HttpAddr string `json:"http_addr,omitempty"`
	HttpPort string `json:"http_port,omitempty"`
}
