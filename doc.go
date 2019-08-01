/**
Package phonedata 内地手机号归属信息.

Data file from: https://github.com/ls0f/phone

data file format:
	  | 4 bytes |                  <- phone.dat 版本号
	  ------------
	  | 4 bytes |                  <-  第一个索引的偏移
	  -----------------------
	  |  offset - 8         |      <-  记录区
	  -----------------------
	  |  index              |      <-  索引区
	  -----------------------

*/
package phonedata
