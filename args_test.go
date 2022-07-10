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
output: false
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

/*single string happy path:
input: -d /usr//logs
output: /usr/logs
*/
func TestSingleArgsParseIntReturnReal(t *testing.T) {
	argsParser := &ArgsParser{}
	Parse(argsParser, "-d", "/usr/logs")
	assert.Equal(t, argsParser.Directory, "/usr/logs")
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
