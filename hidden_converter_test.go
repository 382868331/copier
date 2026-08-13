package copier_test

import (
	"reflect"
	"testing"

	"github.com/jinzhu/copier"
)

type hiddenSource struct{ Value int }
type hiddenTarget struct{ Value int; private int }

func hiddenOptions() copier.Option {
	want := hiddenTarget{Value: 91, private: 37}
	return copier.Option{Converters: []copier.TypeConverter{{
		SrcType: hiddenSource{}, DstType: hiddenTarget{},
		Fn: func(interface{}) (interface{}, error) { return want, nil },
	}}}
}

func TestHiddenConverterAcrossContainers(t *testing.T) {
	want := hiddenTarget{Value: 91, private: 37}
	var direct hiddenTarget
	if err := copier.CopyWithOption(&direct, &hiddenSource{Value: 1}, hiddenOptions()); err != nil || !reflect.DeepEqual(direct, want) {
		t.Fatalf("direct conversion: got %#v, err %v", direct, err)
	}
	src := []*hiddenSource{{Value: 2}}
	dst := []*hiddenTarget{{Value: 3}}
	if err := copier.CopyWithOption(&dst, &src, hiddenOptions()); err != nil || len(dst) != 1 || !reflect.DeepEqual(*dst[0], want) {
		t.Fatalf("slice conversion: got %#v, err %v", dst, err)
	}
	mapSrc := map[string]*hiddenSource{"k": {Value: 4}}
	mapDst := map[string]*hiddenTarget{}
	if err := copier.CopyWithOption(&mapDst, &mapSrc, hiddenOptions()); err != nil || !reflect.DeepEqual(*mapDst["k"], want) {
		t.Fatalf("map conversion: got %#v, err %v", mapDst, err)
	}
}
