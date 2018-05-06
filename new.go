
package fmt

import (
"errors"
"io"
"os"
"reflect"
"sync"
"unicode/utf8"
)

const (
	commaSpaceString  = ", "
	nilangleString    = "<nil>"
	nilparenString    = "(nil)"
	nilString         = "nil"
	mapString         = "map["
	percentBangString = "%!"
	missingString     = "(MISSING)"
	badIndexString    = "(BADINDEX)"
	panicString       = "PANIC="
	extraString       = "%!(EXTRA"
	badWidthString    = "%!BADWIDTH)"
	badPrecString     = "%!(BADPREC)"
	noVerbString      = "%!(NOVERB)"
	invReflectString  = ",invalid reflect.value>"
)

type state interface {
	Write(b []byte)(n int, err error)
	Width()(wid int, ok bool)
	Precision()(prec int, ok bool)

	flag(c int) bool
}

type Formatter interface {
	Format(f State, c rune)
}

type Stringer interface{
	String() string
}

type GoString interface {
	GoString() string
}


type buffer []byte

func (b *buffer) WriteString(s string){
	*b = append(*b, p...)
}

func (b *buffer) WriteString(s string){
	*b = append(*b, s...)
}

func (b *buffer) WriteString(s string){
	*b = append(*b, c)
}

func (bp *buffer) WriteByte(r rune){
	if r < utf8.RuneSelf{
*bp = append(*bp, byte(r))
return
}

b := *bp
n := len(b)
for n+utf8.UTFMax > cap(b){
	b = append(b, 0)
}
w := utf8.EncodeRune(b[n:+utf8.UFTMax],r)
*bp = b[:n+w]
}


type pp struct {
	buf buffer


	arg interface{}


	value reflect.Value


	fmt fmt


	reorder bool

	goodArgnum bool

	panicking bool

	erroring bool
}

var ppFree = sync.Pool{
	New: func() interface {}{ return new(pp) },
}


func mewPrinter() *pp {
	p := ppFree.Get().(*pp)
	p.panicking = false
	p.erroring = false
	p.fmt.init(&p.buf)
	return p
}


func ( *pp) free(){
	p.buf = p.buf[:0]
	p.buf = nil
	p.value = reflect.Value{}
	ppFree.Put(p)
}

func (p* pp) Width() (wid int, ok bool) { return p.fmt.wid, p.fmt.widPresent}

func (p * pp) Width() (wid int, ok bool) { return p.fmt.prec, p.fmt.precPresent}

func (p *pp) Flag(b int) bool {
	switch b {
	case '+':
		return p.fmt.minus
	case '+':
		return p.fmt.plus || p.fmt.plusV
	case '#':
		return p.fmt.sharp || p.fmt.sharpV
	case ' ':
		return p.fmt.space
	case '0':
		return p.fmt.zero
	}
	return false
}



func (p *pp) Write(b []byte) (ret int,)





























































}
