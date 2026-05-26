// Derived from Inferno's utils/iyacc/yacc.c
// http://code.google.com/p/inferno-os/source/browse/utils/iyacc/yacc.c
//
// This copyright NOTICE applies to all files in this directory and
// subdirectories, unless another copyright notice appears in a given
// file or subdirectory.  If you take substantial code from this software to use in
// other programs, you must somehow include with it an appropriate
// copyright notice that includes the copyright notice and the other
// notices below.  It is fine (and often tidier) to do that in a separate
// file such as NOTICE, LICENCE or COPYING.
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors. All rights reserved.
//	Portions Copyright © 2021 The Vitess Authors.
//	Portions Copyright © 2026 The Multigres Authors.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

// yacc
// major difference is lack of stem ("y" variable)
//

import (
	"bufio"
	"bytes"
	"regexp"

	"github.com/spf13/pflag"
)

// the following are adjustable
// according to memory size
const (
	ACTSIZE  = 240000
	NSTATES  = 16000
	TEMPSIZE = 16000

	SYMINC   = 50  // increase for non-term or term
	RULEINC  = 50  // increase for max rule length prodptr[i]
	PRODINC  = 100 // increase for productions     prodptr
	WSETINC  = 50  // increase for working sets    wsets
	STATEINC = 200 // increase for states          statemem

	PRIVATE = 0xE000 // unicode private use

	// relationships which must hold:
	//	TEMPSIZE >= NTERMS + NNONTERM + 1;
	//	TEMPSIZE >= NSTATES;
	//

	NTBASE     = 0o10000
	ERRCODE    = 8190
	ACCEPTCODE = 8191
	YYLEXUNK   = 3
	TOKSTART   = 4 // index of first defined token
)

// no, left, right, binary assoc.
const (
	LASC = iota + 1
	RASC
	BASC
)

// flags for state generation
const (
	DONE = iota
	MUSTDO
	MUSTLOOKAHEAD
)

// flags for a rule having an action, and being reduced
const (
	ACTFLAG = 1 << (iota + 2)
	REDFLAG
)

// output parser flags
const yyFlag = -1000

// parse tokens
const (
	IDENTIFIER = PRIVATE + iota
	MARK
	TERM
	LEFT
	RIGHT
	BINARY
	PREC
	LCURLY
	IDENTCOLON
	NUMBER
	START
	TYPEDEF
	TYPENAME
	STRUCT
	UNION
	ERROR
)

const (
	ENDFILE  = 0
	EMPTY    = 1
	WHOKNOWS = 0
	OK       = 1
	NOMORE   = -1000
)

// macros for getting associativity and precedence levels
func ASSOC(i int) int { _ = "STUB: not implemented"; return 0 }

func PLEVEL(i int) int { _ = "STUB: not implemented"; return 0 }

func TYPE(i int) int { _ = "STUB: not implemented"; return 0 }

// macros for setting associativity and precedence levels
func SETASC(i, j int) int { _ = "STUB: not implemented"; return 0 }

func SETPLEV(i, j int) int { _ = "STUB: not implemented"; return 0 }

func SETTYPE(i, j int) int { _ = "STUB: not implemented"; return 0 }

// I/O descriptors
var (
	finput  *bufio.Reader // input file
	stderr  *bufio.Writer
	ftable  *bufio.Writer     // y.go file
	fcode   = &bytes.Buffer{} // saved code
	ftypes  = &bytes.Buffer{} // saved type definitions
	foutput *bufio.Writer     // y.output file
)

var writtenImports bool // output file has recorded an import of "fmt"

var (
	oflag           string // -o [y.go]		- y.go file
	vflag           string // -v [y.output]	- y.output file
	lflag           bool   // -l			- disable line directives
	prefix          string // name prefix for identifiers, default yy
	allowFastAppend bool
)

func init() {
	pflag.StringVarP(&oflag, "output", "o", "y.go", "parser output")
	pflag.StringVarP(&prefix, "prefix", "p", "yy", "name prefix to use in generated code")
	pflag.StringVarP(&vflag, "verbose-output", "v", "y.output", "create parsing tables")
	pflag.BoolVarP(&lflag, "disable-line-directives", "l", false, "disable line directives")
	pflag.BoolVarP(&allowFastAppend, "fast-append", "f", false, "enable fast-append optimization")
}

var initialstacksize = 16

// communication variables between various I/O routines
var (
	infile  string // input file name
	numbval int    // value of an input number
	tokname string // input token name, slop for runes and 0
	tokflag = false
)

// structure declarations
type Lkset []int

type Pitem struct {
	prod   []int
	off    int // offset within the production
	first  int // first term or non-term in item
	prodno int // production number for sorting
}

type Item struct {
	pitem Pitem
	look  Lkset
}

type Symb struct {
	name    string
	noconst bool
	value   int
}

type Wset struct {
	pitem Pitem
	flag  int
	ws    Lkset
}

// storage of types
var (
	ntypes  int                    // number of types defined
	typeset = make(map[int]string) // pointers to type tags
)

// token information

var (
	ntokens = 0 // number of tokens
	tokset  []Symb
	toklev  []int // vector with the precedence of the terminals
)

// nonterminal information

var (
	nnonter = -1 // the number of nonterminals
	nontrst []Symb
	start   int // start symbol
)

// state information

var (
	nstate   = 0                      // number of states
	pstate   = make([]int, NSTATES+2) // index into statemem to the descriptions of the states
	statemem []Item
	tystate  = make([]int, NSTATES) // contains type information about the states
	tstates  []int                  // states generated by terminal gotos
	ntstates []int                  // states generated by nonterminal gotos
	mstates  = make([]int, NSTATES) // chain of overflows of term/nonterm generation lists
	lastred  int                    // number of last reduction of a state
	defact   = make([]int, NSTATES) // default actions of states
)

// lookahead set information

var (
	nolook  = 0   // flag to turn off lookahead computations
	tbitset = 0   // size of lookahead sets
	clset   Lkset // temporary storage for lookahead computations
)

// working set information

var (
	wsets []Wset
	cwp   int
)

// storage for action table

var (
	amem  []int                  // action table storage
	memp  int                    // next free action table position
	indgo = make([]int, NSTATES) // index to the stored goto table
)

// temporary vector, indexable by states, terms, or ntokens

var (
	temp1   = make([]int, TEMPSIZE) // temporary storage, indexed by terms + ntokens or states
	lineno  = 1                     // current input line number
	fatfl   = 1                     // if on, error is fatal
	nerrors = 0                     // number of errors
)

// assigned token type values

var extval = 0

// grammar rule information

var (
	nprod  = 1     // number of productions
	prdptr [][]int // pointers to descriptions of productions
	levprd []int   // precedence levels for the productions
	rlines []int   // line number for this rule
)

// statistics collection variables

var (
	zzgoent  = 0
	zzgobest = 0
	zzacent  = 0
	zzexcp   = 0
	zzclose  = 0
	zzrrconf = 0
	zzsrconf = 0
	zzstate  = 0
)

// optimizer arrays

var (
	yypgo  [][]int
	optst  [][]int
	ggreed []int
	pgo    []int
)

var (
	maxspr int // maximum spread of any entry
	maxoff int // maximum offset into a array
	maxa   int
)

// storage for information about the nonterminals

var (
	pres   [][][]int // vector of pointers to productions yielding each nonterminal
	pfirst []Lkset
	pempty []int // vector of nonterminals nontrivially deriving e
)

// random stuff picked out from between functions

var (
	indebug = 0 // debugging flag for cpfir
	pidebug = 0 // debugging flag for putitem
	gsdebug = 0 // debugging flag for stagen
	cldebug = 0 // debugging flag for closure
	pkdebug = 0 // debugging flag for apack
	g2debug = 0 // debugging for go2gen
	adb     = 0 // debugging for callopt
)

type Resrv struct {
	name  string
	value int
}

var resrv = []Resrv{
	{"binary", BINARY},
	{"left", LEFT},
	{"nonassoc", BINARY},
	{"prec", PREC},
	{"right", RIGHT},
	{"start", START},
	{"term", TERM},
	{"token", TERM},
	{"type", TYPEDEF},
	{"union", UNION},
	{"struct", STRUCT},
	{"error", ERROR},
}

type Error struct {
	lineno int
	tokens []string
	msg    string
}

var errors []Error

type Row struct {
	actions       []int
	defaultAction int
}

var stateTable []Row

var zznewstate = 0

const EOF = -1

func main() {
	setup() // initialize and read productions

	tbitset = (ntokens + 32) / 32
	cpres()  // make table of which productions yield a given nonterminal
	cempty() // make a table of which nonterminals can match the empty string
	cpfir()  // make a table of firsts of nonterminals

	stagen() // generate the states

	yypgo = make([][]int, nnonter+1)
	optst = make([][]int, nstate)
	output() // write the states and the tables
	go2out()

	hideprod()
	summary()

	callopt()

	typeinfo()
	others()

	exit(0)
}

func setup() { _ = "STUB: not implemented"; return }

// never set so cannot happen

// tokens start in unicode 'private use'

// Do nothing.

// nonzero means new prec. and assoc.

// get identifiers so defined

// there is a type defined

// Do nothing.

// read rules
// put into prdptr array in the format
// target
// followed by id's of terminals and non-terminals
// followed by -nprod

// process a rule

// read rule body

// action within rule...

// make it a nonterminal

//
// the current rule will become rule number nprod+1
// enter null production for action
//

// update the production information

// make the action appear in the original rule

// check that default action is reasonable

// no explicit action, LHS has value

//
// end of all rules
// dump out the prefix code
//

// put out non-literal terminals

// non-literals

// put out names of tokens

// put out names of states.
// commented out to avoid a huge table just for debugging.
// re-enable to have the names in the binary.

//	for i:=TOKSTART; i<=ntokens; i++ {
//		fmt.Fprintf(ftable, "\t%q,\n", tokset[i].name);
//	}

//
// copy any postfix code
//

// allocate enough room to hold another production
func moreprod() { _ = "STUB: not implemented"; return }

// define s to be a terminal if nt==0
// or a nonterminal if nt==1
func defin(nt int, s string) int { _ = "STUB: not implemented"; return 0 }

// must be a token

// establish value for token
// single character literal

var peekline = 0

func gettok() int { _ = "STUB: not implemented"; return 0 }

// skip comment -- fix

// get, and look up, a type name (union member name)

// find a reserved word

// look ahead to distinguish IDENTIFIER from IDENTCOLON

// look for comments

func getword(c rune) { _ = "STUB: not implemented"; return }

// determine the type of a symbol
func fdtype(t int) (int, string) { _ = "STUB: not implemented"; return 0, "" }

func chfind(t int, s string) int { _ = "STUB: not implemented"; return 0 }

// cannot find name

const (
	startUnion = iota
	skippingLeadingBlanks
	readingMember
	skippingLaterBlanks
	readingType
)

type gotypeinfo struct {
	typename string
	union    bool
}

var gotypes = make(map[string]*gotypeinfo)

func typeinfo() { _ = "STUB: not implemented"; return }

// copy the union declaration to the output, and the define file if present
func parsetypes(union bool) { _ = "STUB: not implemented"; return }

// Skip line comments starting with //

// Skip to end of line

// Strip inline // comments

// Skip to end of line, then process as newline

// saves code between %{ and %}
// adds an import for __fmt__ the first time
func cpycode() { _ = "STUB: not implemented"; return }

// accumulate until %}

// emits code saved up from between %{ and %}
// called by cpycode
// adds an import for __yyfmt__ after the package clause
func emitcode(code []rune, lineno int) { _ = "STUB: not implemented"; return }

// does this line look like a package clause?  not perfect: might be confused by early comments.
func isPackageClause(line []rune) bool { _ = "STUB: not implemented"; return false }

// must be big enough.

// must start with "package"

// must have another identifier.

// eol, newline, or comment must follow

// skip initial spaces
func skipspace(line []rune) []rune { _ = "STUB: not implemented"; return nil }

// break code into lines
func lines(code []rune) [][]rune { _ = "STUB: not implemented"; return nil }

// one line per loop

// writes code to ftable
func writecode(code []rune) { _ = "STUB: not implemented"; return }

// skip over comments
// skipcom is called after reading a '/'
func skipcom() int { _ = "STUB: not implemented"; return 0 }

// lines skipped

var fastAppendRe = regexp.MustCompile(`\s+append\(\$[$1],`)

func cpyyvalaccess(fcode *bytes.Buffer, curprod []int, tok int, unionType *string) {
	_ = "STUB: not implemented"
	return
}

// copy action to the next ; or closing }
func cpyact(fcode *bytes.Buffer, curprod []int, max int, unionType *string) {
	_ = "STUB: not implemented"
	return
}

// type description

// look for $name

// put out the proper tag

// a comment

// end of // comment

// end of /* comment?

// character string or constant

func openup() { _ = "STUB: not implemented"; return }

// return a pointer to the name of symbol i
func symnam(i int) string { _ = "STUB: not implemented"; return "" }

// set elements 0 through n-1 to c
func aryfil(v []int, n, c int) { _ = "STUB: not implemented"; return }

// compute an array with the beginnings of productions yielding given nonterminals
// The array pres points to these lists
// the array pyield has the lists: the total size is only NPROD+1
func cpres() { _ = "STUB: not implemented"; return }

// make undefined symbols nonfatal

// mark nonterminals which derive the empty string
// also, look for nonterminals which don't derive any token strings
func cempty() { _ = "STUB: not implemented"; return }

// first, use the array pempty to detect productions that can never be reduced
// set pempty to WHONOWS

// now, look at productions, marking nonterminals which derive something

// production can be derived

// now, look at the nonterminals, to see if they are all OK

// the added production rises or falls as the start symbol ...

// now, compute the pempty array, to see which nonterminals derive the empty string
// set pempty to WHOKNOWS

// loop as long as we keep finding empty nonterminals

// not known to be empty

// we have a nontrivially empty nonterminal

// got one ... try for another

// compute an array with the first of nonterminals
func cpfir() { _ = "STUB: not implemented"; return }

// initially fill the sets

// now, reflect transitivity

// generate the states
func stagen() {
	_ = "STUB: not implemented"
	// initialize
	return
}

// states generated by terminal gotos
// states generated by nonterminal gotos

//
// now, the main state generation loop
// first pass generates all of the states
// later passes fix up lookahead
// could be sped up a lot by remembering
// results of the first pass rather than recomputing
//

// take state i, close it, and do gotos

// generate goto's

// do a goto on c

// this item contributes to the goto

// register new state

// generate the closure of state i
func closure(i int) {
	_ = "STUB: not implemented"

	// first, copy kernel of state i to wsets
	return
}

// this item must get closed

// now, go through the loop, closing each item

// dot is before c

// only interesting case is where . is before nonterminal

// compute the lookahead

// find items involving c

// terminal symbol

// nonterminal symbol

//
// now loop over productions derived from c
//

// initially fill the sets

//
// put these items into the closure
// is the item there
//

// yes, it is there

//  not there; make a new entry

// have computed closure; flags are reset; return

// sorts last state,and sees if it equals earlier ones. returns state number
func state(c int) int { _ = "STUB: not implemented"; return 0 }

// null state

// sort the items

// make k the biggest

// size of state

// get ith state

// found it
// delete last state

// fix up lookaheads

// state is new

func putitem(p Pitem, set Lkset) { _ = "STUB: not implemented"; return }

// creates output string for item pointed to by pp
func writem(pp Pitem) string { _ = "STUB: not implemented"; return "" }

// an item calling for a reduction

// pack state i from temp1 into amem
func apack(p []int, n int) int {
	_ = "STUB: not implemented"
	// we don't need to worry about checking because
	// we will only look at entries known to be there...
	// eliminate leading and trailing 0's
	return 0
}

// no actions

// now, find a place for the elements from p to q, inclusive

// we have found an acceptable k

// print the output for the states
func output() { _ = "STUB: not implemented"; return }

// output the stuff for state i

// output actions

// now, we have the shifts; look at the reductions

// reduction

// reduce/reduce conflict

// potential shift/reduce conflict

// decide a shift/reduce conflict by precedence.
// r is a rule number, t a token number
// the conflict is in state s
// temp1[t] is changed to reflect the action
func precftn(r, t, s int) { _ = "STUB: not implemented"; return }

// conflict

// shift

// reduce

// error action

// reduce

// output state i
// temp1 has the actions, lastred the default
func wract(i int) {
	_ = "STUB: not implemented"

	// find the best choice for lastred
	return
}

// count the number of appearances of temp1[j]

//
// for error recovery, arrange that, if there is a shift on the
// error recovery token, `error', that the default be the error action
//

// clear out entries in temp1 which equal lastred
// count entries in optst table

// writes state i
func wrstate(i int) { _ = "STUB: not implemented"; return }

// print out empty productions in closure

// check for state equal to another

// shift, error, or accept

// output the final production

// now, output nonterminal actions

// output the gotos for the nontermninals
func go2out() { _ = "STUB: not implemented"; return }

// find the best one to make default

// is j the most frequent

// is tystate[j] the most frequent

// best is now the default entry

// now, the default

// output the gotos for nonterminal c
func go2gen(c int) {
	_ = "STUB: not implemented"

	// first, find nonterminals with gotos on c
	return
}

// cc is a nonterminal with a goto on c

// thus, the left side of production i does too

// now, we have temp1[c] = 1 if a goto on c in closure of cc

// now, go through and put gotos into tystate

// goto on c is possible

// in order to free up the mem and amem arrays for the optimizer,
// and still be able to output yyr1, etc., after the sizes of
// the action array is known, we hide the nonterminals
// derived by productions in levprd.
func hideprod() { _ = "STUB: not implemented"; return }

func callopt() { _ = "STUB: not implemented"; return }

// nontrivial situation

// j is now the range
//			j -= k;			// call scj

// initialize ggreed table

// minimum entry index is always 0

// now, prepare to put the shift actions into the amem array

// print amem array

// finds the next i
func nxti() int { _ = "STUB: not implemented"; return 0 }

func gin(i int) {
	_ = "STUB: not implemented"

	// enter gotos on nonterminal i into array amem
	return
}

// now, find amem place for it

// we have found amem spot

func stin(i int) { _ = "STUB: not implemented"; return }

// enter state i into the amem array

// find an acceptable place

// check the position equals another only if the states are identical

// we have some disagreement

// states are equal

// we have some disagreement

// this version is for limbo
// write out the optimized parser
func aoutput() { _ = "STUB: not implemented"; return }

// put out other arrays, copy the parsers
func others() { _ = "STUB: not implemented"; return }

//
// yyr2 is the number of rules for each production
//

// put out token translation tables
// table 1 has 0-256

// table 2 has PRIVATE-PRIVATE+256

// table 3 has everything else

// Custom error messages.

// copy parser text

// copy yaccpar

func runMachine(tokens []string) (state, token int) { _ = "STUB: not implemented"; return 0, 0 }

// Shift to state action.

// Reduce by production -action.

func arout(s string, v []int, n int) { _ = "STUB: not implemented"; return }

// output the summary on y.output
func summary() { _ = "STUB: not implemented"; return }

// write optimizer summary
func osummary() { _ = "STUB: not implemented"; return }

// copies and protects "'s in q
func chcopy(q string) string { _ = "STUB: not implemented"; return "" }

func usage() { _ = "STUB: not implemented"; return }

func bitset(set Lkset, bit int) int { _ = "STUB: not implemented"; return 0 }

func setbit(set Lkset, bit int) { _ = "STUB: not implemented"; return }

func mkset() Lkset { _ = "STUB: not implemented"; return *new(Lkset) }

// set a to the union of a and b
// return 1 if b is not a subset of a, 0 otherwise
func setunion(a, b []int) int { _ = "STUB: not implemented"; return 0 }

func prlook(p Lkset) { _ = "STUB: not implemented"; return }

// utility routines
var peekrune rune

func isdigit(c rune) bool { _ = "STUB: not implemented"; return false }

func isword(c rune) bool { _ = "STUB: not implemented"; return false }

// return 1 if 2 arrays are equal
// return 0 if not equal
func aryeq(a []int, b []int) int { _ = "STUB: not implemented"; return 0 }

func getrune(f *bufio.Reader) rune { _ = "STUB: not implemented"; return 0 }

// fmt.Printf("rune = %v n=%v\n", string(c), n);

func ungetrune(f *bufio.Reader, c rune) { _ = "STUB: not implemented"; return }

func open(s string) *bufio.Reader { _ = "STUB: not implemented"; return nil }

// fmt.Printf("open %v\n", s);

func create(s string) *bufio.Writer { _ = "STUB: not implemented"; return nil }

// fmt.Printf("create %v mode %v\n", s);

// write out error comment
func lerrorf(lineno int, s string, v ...any) { _ = "STUB: not implemented"; return }

func errorf(s string, v ...any) { _ = "STUB: not implemented"; return }

func exit(status int) { _ = "STUB: not implemented"; return }

//nolint:forbidigo

const fastAppendHelperText = `
func $$Iaddr(v any) __yyunsafe__.Pointer {
	type h struct {
		t __yyunsafe__.Pointer
		p __yyunsafe__.Pointer
	}
	return (*h)(__yyunsafe__.Pointer(&v)).p
}
`

var yaccpar string // will be processed version of yaccpartext: s/$$/prefix/g
const yaccpartext = `
/*	parser for yacc output	*/

var (
	$$Debug        = 0
	$$ErrorVerbose = false
)

type $$Lexer interface {
	Lex(lval *$$SymType) int
	Error(s string)
}

type $$Parser interface {
	Parse($$Lexer) int
	Lookahead() int
}

type $$ParserImpl struct {
	lval  $$SymType
	stack [$$InitialStackSize]$$SymType
	char  int
}

func (p *$$ParserImpl) Lookahead() int {
	return p.char
}

func $$NewParser() $$Parser {
	return &$$ParserImpl{}
}

const $$Flag = -1000

func $$Tokname(c int) string {
	if c >= 1 && c-1 < len($$Toknames) {
		if $$Toknames[c-1] != "" {
			return $$Toknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func $$Statname(s int) string {
	if s >= 0 && s < len($$Statenames) {
		if $$Statenames[s] != "" {
			return $$Statenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func $$ErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !$$ErrorVerbose {
		return "syntax error"
	}

	for _, e := range $$ErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + $$Tokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := $$Pact[state]
	for tok := TOKSTART; tok-1 < len($$Toknames); tok++ {
		if n := base + tok; n >= 0 && n < $$Last && $$Chk[$$Act[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if $$Def[state] == -2 {
		i := 0
		for $$Exca[i] != -1 || $$Exca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; $$Exca[i] >= 0; i += 2 {
			tok := $$Exca[i]
			if tok < TOKSTART || $$Exca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if $$Exca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += $$Tokname(tok)
	}
	return res
}

func $$lex1(lex $$Lexer, lval *$$SymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = $$Tok1[0]
		goto out
	}
	if char < len($$Tok1) {
		token = $$Tok1[char]
		goto out
	}
	if char >= $$Private {
		if char < $$Private+len($$Tok2) {
			token = $$Tok2[char-$$Private]
			goto out
		}
	}
	for i := 0; i < len($$Tok3); i += 2 {
		token = $$Tok3[i+0]
		if token == char {
			token = $$Tok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = $$Tok2[1] /* unknown char */
	}
	if $$Debug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", $$Tokname(token), uint(char))
	}
	return char, token
}

func $$Parse($$lex $$Lexer) int {
	return $$NewParser().Parse($$lex)
}

func ($$rcvr *$$ParserImpl) Parse($$lex $$Lexer) int {
	var $$n int
	var $$VAL $$SymType
	var $$Dollar []$$SymType
	_ = $$Dollar // silence set and not used
	$$S := $$rcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	$$state := 0
	$$rcvr.char = -1
	$$token := -1 // $$rcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		$$state = -1
		$$rcvr.char = -1
		$$token = -1
	}()
	$$p := -1
	goto $$stack

ret0:
	return 0

ret1:
	return 1

$$stack:
	/* put a state and value onto the stack */
	if $$Debug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", $$Tokname($$token), $$Statname($$state))
	}

	$$p++
	if $$p >= len($$S) {
		nyys := make([]$$SymType, len($$S)*2)
		copy(nyys, $$S)
		$$S = nyys
	}
	$$S[$$p] = $$VAL
	$$S[$$p].yys = $$state

$$newstate:
	$$n = $$Pact[$$state]
	if $$n <= $$Flag {
		goto $$default /* simple state */
	}
	if $$rcvr.char < 0 {
		$$rcvr.char, $$token = $$lex1($$lex, &$$rcvr.lval)
	}
	$$n += $$token
	if $$n < 0 || $$n >= $$Last {
		goto $$default
	}
	$$n = $$Act[$$n]
	if $$Chk[$$n] == $$token { /* valid shift */
		$$rcvr.char = -1
		$$token = -1
		$$VAL = $$rcvr.lval
		$$state = $$n
		if Errflag > 0 {
			Errflag--
		}
		goto $$stack
	}

$$default:
	/* default state action */
	$$n = $$Def[$$state]
	if $$n == -2 {
		if $$rcvr.char < 0 {
			$$rcvr.char, $$token = $$lex1($$lex, &$$rcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if $$Exca[xi+0] == -1 && $$Exca[xi+1] == $$state {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			$$n = $$Exca[xi+0]
			if $$n < 0 || $$n == $$token {
				break
			}
		}
		$$n = $$Exca[xi+1]
		if $$n < 0 {
			goto ret0
		}
	}
	if $$n == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			$$lex.Error($$ErrorMessage($$state, $$token))
			Nerrs++
			if $$Debug >= 1 {
				__yyfmt__.Printf("%s", $$Statname($$state))
				__yyfmt__.Printf(" saw %s\n", $$Tokname($$token))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for $$p >= 0 {
				$$n = $$Pact[$$S[$$p].yys] + $$ErrCode
				if $$n >= 0 && $$n < $$Last {
					$$state = $$Act[$$n] /* simulate a shift of "error" */
					if $$Chk[$$state] == $$ErrCode {
						goto $$stack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if $$Debug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", $$S[$$p].yys)
				}
				$$p--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if $$Debug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", $$Tokname($$token))
			}
			if $$token == $$EofCode {
				goto ret1
			}
			$$rcvr.char = -1
			$$token = -1
			goto $$newstate /* try again in the same state */
		}
	}

	/* reduction by production $$n */
	if $$Debug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", $$n, $$Statname($$state))
	}

	$$nt := $$n
	$$pt := $$p
	_ = $$pt // guard against "declared and not used"

	$$p -= $$R2[$$n]
	// $$p is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if $$p+1 >= len($$S) {
		nyys := make([]$$SymType, len($$S)*2)
		copy(nyys, $$S)
		$$S = nyys
	}
	$$VAL = $$S[$$p+1]

	/* consult goto table to find next state */
	$$n = $$R1[$$n]
	$$g := $$Pgo[$$n]
	$$j := $$g + $$S[$$p].yys + 1

	if $$j >= $$Last {
		$$state = $$Act[$$g]
	} else {
		$$state = $$Act[$$j]
		if $$Chk[$$state] != -$$n {
			$$state = $$Act[$$g]
		}
	}
	// dummy call; replaced with literal code
	$$run()
	goto $$stack /* stack new state and value */
}
`
