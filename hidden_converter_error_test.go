package copier_test
import("errors";"testing";"github.com/jinzhu/copier")
func TestHiddenConverterErrorPropagates(t *testing.T){type Src struct{Value string};type Dst struct{Value *string};marker:=errors.New("conversion rejected");p:="";var dst Dst;err:=copier.CopyWithOption(&dst,&Src{},copier.Option{Converters:[]copier.TypeConverter{{SrcType:"",DstType:&p,Fn:func(interface{})(interface{},error){return nil,marker}}}});if !errors.Is(err,marker){t.Fatalf("expected converter error, got %v",err)}}
