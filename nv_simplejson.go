/* Nahanni extensions to simplejson */

package simplejson

func FromInterface(i interface{}) *Json {
	return &Json{i}
}
