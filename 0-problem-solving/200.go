package main

import "fmt"

// https://leetcode.com/problems/encode-and-decode-tinyurl/description/?envType=problem-list-v2&envId=hash-table

type Codec struct {

}


func Constructor() Codec {
    return Codec{}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	return longUrl
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
    return shortUrl
}


/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */


func main () {

}