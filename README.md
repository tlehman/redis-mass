# redis-mass

[Go](https://golang.org/) package to encode a sequence of [Redis commands](http://redis.io/commands) into [Redis protocol](http://redis.io/topics/protocol), suitable for [mass insertion](http://redis.io/topics/mass-insert).

*A port of [almeida's redis-mass](https://github.com/almeida/redis-mass)*

## Installation

Via git (or downloaded tarball):

```
$ git clone git://github.com/tlehman/redis-mass.git
$ cd redis-mass/
$ go test
$ go build
```

Via `go get`

```
$ go get github.com/tlehman/redis-mass
```

## Usage

### Mass insertion on Redis

```
$ redis-mass /path/to/input-file | redis-cli --pipe
```

### Output to console

```
$ redis-mass /path/to/input-file
```

### Output to file

```
$ redis-mass /path/to/input-file /path/to/output-file
```

## Running tests

Use the built in `go test`:

```
$ go test
```

## Issues

You can find list of issues using **[this link](http://github.com/tlehman/redis-mass/issues)**.

## Requirements

 - **[Go](https://golang.org)**

## Dependencies

 - None.


## Examples

### Input file ([Redis Commands](http://redis.io/commands))

```
SET key1 value1
SADD key2 value1 "value2" "value3"
ZADD "key3" 1 "value3"
```

### Output ([Redis Protocol](http://redis.io/topics/protocol))

```
*3
$3
SET
$4
key1
$6
value1
*5
$4
SADD
$4
key2
$6
value1
$6
value2
$6
value3
*4
$4
ZADD
$4
key3
$1
1
$6
value3
```

## License

(The MIT License)

Copyright (c) 2016 Tobi Lehman <mail@tobilehman.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the **Software**), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED **AS IS**, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
