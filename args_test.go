package tdd_args

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

/*single bool happy path：
input: -l
output: true
*/
func TestSingleArgsParseBoolReturnTrue(t *testing.T) {
	argsParser := &ArgsParser{}
	Parse(argsParser, "-l")
	assert.Equal(t, argsParser.Logging, true)
}

/*single bool default value：
input:
output: false
*/
func TestSingleArgsParseBoolReturnDefault(t *testing.T) {
	argsParser := &ArgsParser{}
	Parse(argsParser, "")
	assert.Equal(t, argsParser.Logging, false)
}

/*single bool sad path：
input: -l t
output: too many args error
*/
func TestSingleArgsParseBoolReturnErr(t *testing.T) {
	argsParser := &ArgsParser{}
	err := Parse(argsParser, "-l", "t")
	assert.Equal(t, err.Error(), "too many args error")
}

/*single int happy path：
input: -p 8080
output: 8080
*/

func TestSingleArgsParseIntReturn8080(t *testing.T) {
	argsParser := &ArgsParser{}
	Parse(argsParser, "-p", "8080")
	assert.Equal(t, argsParser.Port, 8080)
}

/*single bool default value：
input: -p
output: 0
*/
func TestSingleArgsParseIntReturnDefault(t *testing.T) {
	argsParser := &ArgsParser{}
	Parse(argsParser, "-p")
	assert.Equal(t, argsParser.Port, 0)
}

/*single bool sad path：
input: -p 8080 8081
output: too many args error
*/
func TestSingleArgsParseIntReturnErr(t *testing.T) {
	argsParser := &ArgsParser{}
	err := Parse(argsParser, "-p", "8080", "8081")
	assert.Equal(t, err.Error(), "too many args error")
}

/*single string happy path:
input: -d /usr//logs
output: /usr/logs
*/
func TestSingleArgsParseIntReturnReal(t *testing.T) {
	argsParser := &ArgsParser{}
	Parse(argsParser, "-d", "/usr/logs")
	assert.Equal(t, argsParser.Directory, "/usr/logs")
}

/*single string happy path:
input: -d
output: ""
*/
func TestSingleArgsParseStringReturnDefault(t *testing.T) {
	argsParser := &ArgsParser{}
	Parse(argsParser, "-d")
	assert.Equal(t, argsParser.Directory, "")
}

/*single string sad path:
input: -d
output: "too many args error"
*/
func TestSingleArgsParseStringReturnErr(t *testing.T) {
	argsParser := &ArgsParser{}
	err := Parse(argsParser, "-d", "the", "engine")
	assert.Equal(t, err.Error(), "too many args error")
}

/*all happy path
input: -l -d 8080 -f /usr/local
output: true 8010 /usr/local
*/
func TestArgsParseBoolIntStringReturnReal(t *testing.T) {
	argsParser := &ArgsParser{}
	Parse(argsParser, "-l", "-p", "8080", "-d", "/usr/logs")
	assert.Equal(t, argsParser.Logging, true)
	assert.Equal(t, argsParser.Port, 8080)
	assert.Equal(t, argsParser.Directory, "/usr/logs")
}

/*all default
input: -p -d
output: false 0 ""
*/
func TestArgsParseBoolIntStringReturnDefult(t *testing.T) {
	argsParser := &ArgsParser{}
	Parse(argsParser, "-p", "-d")
	assert.Equal(t, argsParser.Logging, false)
	assert.Equal(t, argsParser.Port, 0)
	assert.Equal(t, argsParser.Directory, "")
}

/*single int list happy path：
input: -t 1 2 -3 5
output: 1 2 -3 5
*/
func TestSingleArgsParseIntListReturnReal(t *testing.T) {
	argsParser := &ArgsParser{}
	Parse(argsParser, "-t", "1", "2", "-3", "5")
	assert.Equal(t, argsParser.IntList, []int{1, 2, -3, 5})
}

/*single int list default：
input: -t
output: 【】
*/
func TestSingleArgsParseIntListReturnDefault(t *testing.T) {
	argsParser := &ArgsParser{}
	Parse(argsParser, "-t")
	assert.Equal(t, argsParser.IntList, []int{})
}

/*single int list sad path：
input: -t 1 2 -3 u
output: "invalid args error"
*/
func TestSingleArgsParseIntListReturnErr(t *testing.T) {
	argsParser := &ArgsParser{}
	err := Parse(argsParser, "-t", "1", "2", "-3", "u")
	assert.Equal(t, err.Error(), "invalid args error")
}

/*single string list happy path：
input: -g this is a book
output: this is a book
*/
func TestSingleArgsParseStringListReturnReal(t *testing.T) {
	argsParser := &ArgsParser{}
	Parse(argsParser, "-g", "this", "is", "a", "book")
	assert.Equal(t, argsParser.StringList, []string{"this", "is", "a", "book"})
}

/*single all list happy path：
input: -g this is a book -t 1 2 -3 5
output: this is a book  1 2 -3 5
*/
func TestAllArgsParseStringIntListReturnReal(t *testing.T) {
	argsParser := &ArgsParser{}
	Parse(argsParser, "-g", "this", "is", "a", "book", "-t", "1", "2", "-3", "5")
	assert.Equal(t, argsParser.StringList, []string{"this", "is", "a", "book", "1", "2", "-3", "5"})
}
