/* Nahanni extensions to simplejson */

package simplejson

func FromInterface(i interface{}) *Json {
	return &Json{i}
}

func (j Json) IsNil() bool {
	return j.data == nil
}

func (j Json) Data() interface{} {
	return j.data
}
