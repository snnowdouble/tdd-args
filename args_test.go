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
type ArgsParser struct {
	BooleanType bool   `tag:"l"`
	IntType     int    `tag:"p"`
	StringType  string `tag:"d"`
}

func TestArgsParseBool(t *testing.T) {
	argsParser := ArgsParser{}
	Parse(argsParser, "l")
	assert.Equal(t, argsParser.BooleanType, true)

}
