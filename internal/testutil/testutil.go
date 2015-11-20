package testutil

import (
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"

	a "github.com/stretchr/testify/assert"
)

func Main(m *testing.M) {
	// never test against live key.
	pkey, skey := Keys()
	switch {
	case pkey == "" || skey == "":
		log.Fatalln("no test key specified, please set both $OMISE_PUBKEY and $OMISE_KEY")
	case !strings.HasPrefix(pkey, "pkey_test_") || !strings.HasPrefix(skey, "skey_test_"):
		log.Fatalln("specified key is invalid or is not a test key!!! You might lose money!!!")
	}

	flag.Parse()
	os.Exit(m.Run())
}

func LogObj(t *testing.T, obj interface{}) {
	t.Log(fmt.Sprintf("%#v", obj))
}

func Require(t *testing.T, env string) {
	env = strings.ToUpper(env)
	val := strings.ToLower(os.Getenv(env))

	switch val {
	case "", "0", "false", "no":
		t.Skip("requires " + env + "=1")
	}
}

func AssertJSONEquals(t *testing.T, m1 map[string]interface{}, m2 map[string]interface{}) bool {
	return assertJSONEquals(t, "", m1, m2)
}

func assertJSONEquals(t *testing.T, prefix string, m1 map[string]interface{}, m2 map[string]interface{}) bool {
	for k, v := range m1 {
		v2, ok := m2[k]
		if !a.True(t, ok, "missing `"+prefix+k+"` key") {
			return false
		}
		if !a.IsType(t, v, v2, "mismatched type for `"+prefix+k+"` key") {
			return false
		} // postcond: v.(type) == v2.(type)

		rv1, rv2 := reflect.ValueOf(v), reflect.ValueOf(v2)
		switch rv1.Kind() {
		case reflect.Array, reflect.Slice:
			// Only support slice of maps, for now.
			if !a.Equal(t, rv1.Len(), rv2.Len(), "key `"+prefix+k+"` has different length") {
				return false
			}
			for i := 0; i < rv1.Len(); i++ {
				if !assertJSONEquals(t, prefix+"["+strconv.Itoa(i)+"]",
					rv1.Index(i).Interface().(map[string]interface{}),
					rv2.Index(i).Interface().(map[string]interface{})) {
					return false
				}
			}

		case reflect.Map:
			vmap, vmap2 := v.(map[string]interface{}), v2.(map[string]interface{})
			if !assertJSONEquals(t, prefix+k+".", vmap, vmap2) {
				return false
			}

		default:
			if !a.Equal(t, v, v2, "mismatched value for`"+prefix+k+"` key") {
				return false
			}
		}
	}

	return true
}
