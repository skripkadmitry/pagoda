package pagoda

//declare NA constants. These are pointers to corresponding type objects
//they are effectivly constants
//every operation with NA leads to NA
var __zero int64 = 0
var na_int *int64 = &__zero

var __emptystring string = ""
var na_string *string = &__emptystring

var __emptybool bool = false
var na_bool *bool = &__emptybool
var absoluteT = true
var absoluteF = false
var T = &absoluteT //all trues and falses are the same, cause why not
var F = &absoluteF

var __float float64 = 0
var na_float *float64 = &__float

var na_categorial *Category = nil

//infinite values are like NAs, but not
//there are positive and negative infinites for both ints and floats
//here are operations with infinites(regular numbers are x)
//inf ? NA = NA
//inf + inf = inf
//inf - inf = NA
//_inf + inf = NA
//_inf - inf = _inf
//_inf - _inf = NA

//inf * inf = inf
//inf * _inf = _inf
//_inf * _inf = inf
//inf ^ inf = inf
//_inf ^ (inf, _inf, 0, float) = NA
//_int ^ even = inf
//_int ^ odd = _inf

//inf / inf = NA
//(inf, _inf) % (any) = NA
//inf == inf = true, but that's questionable
//inf == _inf = false
//inf == (x, -inf) = false
//0 / 0 = NA
//x % 0 = NA
//(x > 0) / 0 = inf
//(x < 0) / 0 = _inf

//any other function leads to NA(mb i will find way to do that better)
//maybe I`ll do overflow to inf but i doubt that
var __intposinf int64 = 1
var int_inf *int64 = &__intposinf

var __intneginf int64 = -1
var _int_inf *int64 = &__intneginf

var __floatposinf float64 = 1
var float_inf *float64 = &__floatposinf

var __floatneginf float64 = -1
var _float_inf *float64 = &__floatneginf

//possible Series types.
//To make typing more convenient
type _Type string

const ( //possible types
	String     _Type = "string"
	Int        _Type = "int"
	Float      _Type = "float"
	Bool       _Type = "bool"
	Categorial _Type = "category"
	Error      _Type = "error"
	_err             = iota - 1
	_bool
	_int
	_float
	_string
	_categorial
)

type TypeError struct {
	out string
}

func (e TypeError) Error() string {
	return e.out
}

type Ints interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

type Floats interface {
	float32 | float64
}

type Element interface {
	*Ints | *Floats | *string | *Category | *bool
}
