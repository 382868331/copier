package copier_test
import("testing";"github.com/jinzhu/copier")
func TestHiddenNilEmbeddedPointer(t *testing.T){type Inner struct{ID int};type Src struct{Title string;*Inner};type Dst struct{Title string};src:=Src{Title:"kept"};var dst Dst;if err:=copier.Copy(&dst,&src);err!=nil||dst.Title!="kept"{t.Fatalf("got %#v err %v",dst,err)}}
