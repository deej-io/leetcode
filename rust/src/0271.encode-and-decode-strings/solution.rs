// Created by Daniel J. Rollins at 2024/04/16 17:55
// leetgo: 1.4.5
// https://leetcode.com/problems/encode-and-decode-strings/

use anyhow::Result;
use leetgo_rs::*;
use std::io::{BufRead, Write};

// @lc code=begin
use std::io::{Cursor, Read};

struct Codec {

}

/** 
* `&self` means the method takes an immutable reference.
* If you need a mutable reference, change it to `&mut self` instead.
*/
impl Codec {
    fn new() -> Self {
	Codec {}
    }

    fn encode(&self, strs: Vec<String>) -> String {
	let mut cursor = Cursor::new(Vec::<u8>::with_capacity(200 * 200 + 200 * 4));
	for str in strs {
	    write!(cursor, "{} {}", str.len(), str).unwrap();
	}
	String::from_utf8(cursor.into_inner()).unwrap()
    }

    fn decode(&self, s: String) -> Vec<String> {
	let mut cursor = Cursor::new(s.as_bytes());
	let mut decoded = Vec::with_capacity(200);
	let mut count_buf = Vec::with_capacity(3);
	loop {
	    count_buf.clear();
	    let bytes_read = cursor.read_until(' ' as u8, &mut count_buf).unwrap();
	    if bytes_read == 0 {
		break;
	    }
	    let count = std::str::from_utf8(&count_buf[..bytes_read-1]).unwrap().parse().unwrap();
	    let mut buf = vec![2u8; count];
	    cursor.read_exact(&mut buf).unwrap();
	    let str = String::from_utf8(buf).unwrap();
	    println!("{}: {}", count,  str);
	    decoded.push(str);
	}
	decoded
    }
}

// @lc code=end

// Warning: this is a manual question, the generated test code may be incorrect.
fn main() -> Result<()> {
    let dummy_input: Vec<String> = deserialize(&read_line()?)?;
    let codec = Codec::new();
    let ans: Vec<String> = codec.decode(codec.encode(dummy_input));

    println!("\noutput: {}", serialize(ans)?);
    Ok(())
}
