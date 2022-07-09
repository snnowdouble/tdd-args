package tdd_args

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

/*test1：
input: -l
output: true
*/
/*test2：
input: -p 8080
output: 8080
*/
/*test3:
input: -d /usr//logs
output: /usr/logs
*/
/*test4:
input: -l -d 8080 -f /usr/local
output: true 8010 /usr/local
*/

func TestSingleArgsParseBoolReturnTrue(t *testing.T) {
	argsParser := &ArgsParser{}
	Parse(argsParser, "-l")
	assert.Equal(t, argsParser.Logging, true)
}
func TestSingleArgsParseIntReturn8080(t *testing.T) {
	argsParser := &ArgsParser{}
	Parse(argsParser, "-p", "8080")
	assert.Equal(t, argsParser.Port, 8080)
}
func TestSingleArgsParseIntReturnReal(t *testing.T) {
	argsParser := &ArgsParser{}
	Parse(argsParser, "-d", "/usr/logs")
	assert.Equal(t, argsParser.Directory, "/usr/logs")
}

func TestArgsParseBoolIntStringReturnReal(t *testing.T) {
	argsParser := &ArgsParser{}
	Parse(argsParser, "-l", "-p", "8080", "-d", "/usr/logs")
	assert.Equal(t, argsParser.Logging, true)
	assert.Equal(t, argsParser.Port, 8080)
	assert.Equal(t, argsParser.Directory, "/usr/logs")
}
