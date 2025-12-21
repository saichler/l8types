/*
© 2025 Sharon Aicler (saichler@gmail.com)

Layer 8 Ecosystem is licensed under the Apache License, Version 2.0.
You may obtain a copy of the License at:

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ifs

type LogLevel int

const (
	Trace_Level   LogLevel = 1
	Debug_Level   LogLevel = 2
	Info_Level    LogLevel = 3
	Warning_Level LogLevel = 4
	Error_Level   LogLevel = 5
)

func (l LogLevel) String() string {
	switch l {
	case Trace_Level:
		return "(Trace)"
	case Debug_Level:
		return "(Debug)"
	case Info_Level:
		return "(Info) "
	case Warning_Level:
		return "(Warn )"
	case Error_Level:
		return "(Error)"
	}
	return ""
}

type ILogger interface {
	Trace(...interface{})
	Debug(...interface{})
	Info(...interface{})
	Warning(...interface{})
	Error(...interface{}) error
	Empty() bool
	Fail(interface{}, ...interface{})
	SetLogLevel(LogLevel)
}
