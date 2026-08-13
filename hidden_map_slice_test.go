package copier_test
import("testing";"github.com/jinzhu/copier")
func TestHiddenMapSliceValue(t *testing.T){src:=map[string][]map[string]int{"k":{{"n":7}}};var dst map[string][]map[string]int;if err:=copier.CopyWithOption(&dst,src,copier.Option{DeepCopy:true});err!=nil{t.Fatalf("copy failed: %v",err)};if len(dst["k"])!=1||dst["k"][0]["n"]!=7{t.Fatalf("bad result %#v",dst)}}
