# OMISE-GO

### 2015-11-19

* **NEW!** Experimental api package. [![GoDoc](https://godoc.org/github.com/omise/omise-go/api?status.svg)](https://godoc.org/github.com/omise/omise-go/api)
* **NEW!** String() and ColorString() method for pretty-printing model structs in one line.
* **CHANGED** `operations.List.Order` now uses the new `Ordering` set of constants. This
  should not break anything since constant lieterals use the target field's type.
* **IMPROVED** Internal generator improvements.
* **IMPROVED** Added some missing fields and removed some redundant ones.

### 2015-11-13

* Initial release.

