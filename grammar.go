//line grammar.y:4
package gdb

import __yyfmt__ "fmt"

//line grammar.y:4
const (
	terminator  = "(gdb) " // yes there's the trailing space
	typeKey     = "type"
	classKey    = "class"
	payloadKey  = "payload"
	sequenceKey = "sequence"
)

// avoid DRY due to a poor lexer
func newClassResult(typeString, class string, payload map[string]interface{}) map[string]interface{} {
	out := map[string]interface{}{
		typeKey:  typeString,
		classKey: class,
	}
	if payload != nil {
		out[payloadKey] = payload
	}
	return out
}

//line grammar.y:27
type yySymType struct {
	yys          int
	text         string
	record       map[string]interface{}
	class_result struct {
		class   string
		payload map[string]interface{}
	}
	result_pair struct {
		variable string
		value    interface{}
	}
	value interface{}
	list  []interface{}
}

const text = 57346

var yyToknames = []string{
	"text",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 29
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 53

var yyAct = []int{

	24, 29, 13, 23, 42, 20, 46, 26, 30, 21,
	15, 16, 27, 33, 47, 34, 39, 33, 44, 34,
	25, 25, 26, 22, 14, 19, 45, 28, 43, 18,
	36, 17, 1, 38, 37, 41, 40, 35, 12, 6,
	7, 8, 32, 9, 10, 11, 48, 49, 31, 5,
	4, 3, 2,
}
var yyPact = []int{

	34, -1000, -1000, -1000, -1000, -1000, 20, 20, 20, 27,
	25, 21, -7, -1000, 1, -1000, -1000, -1000, -1000, -1000,
	20, 17, -1000, -1, -1000, 5, 17, 4, -1000, -1000,
	-1000, -1000, -1000, 16, 0, 14, -1000, 10, -2, -1000,
	-1000, -1000, 5, -1000, 4, -1000, 17, -1000, -1000, -1000,
}
var yyPgo = []int{

	0, 3, 52, 51, 50, 49, 48, 2, 0, 1,
	42, 34, 33, 32,
}
var yyR1 = []int{

	0, 13, 2, 2, 2, 3, 3, 3, 7, 7,
	4, 4, 4, 5, 1, 1, 8, 9, 9, 9,
	11, 11, 6, 6, 12, 12, 10, 10, 10,
}
var yyR2 = []int{

	0, 1, 1, 1, 1, 2, 2, 2, 3, 1,
	2, 2, 2, 3, 3, 1, 3, 1, 1, 1,
	3, 1, 3, 2, 3, 1, 3, 3, 2,
}
var yyChk = []int{

	-1000, -13, -2, -3, -4, -5, 5, 6, 7, 9,
	10, 11, 4, -7, 4, -7, -7, 4, 4, 4,
	12, 8, -7, -1, -8, 4, 8, 7, -8, -9,
	4, -6, -10, 13, 15, -1, 14, -11, -12, 16,
	-9, -8, 4, 14, 8, 16, 8, 16, -9, -8,
}
var yyDef = []int{

	0, -2, 1, 2, 3, 4, 0, 0, 0, 0,
	0, 0, 0, 5, 9, 6, 7, 10, 11, 12,
	0, 0, 13, 8, 15, 0, 0, 0, 14, 16,
	17, 18, 19, 0, 0, 0, 23, 0, 0, 28,
	21, 25, 17, 22, 0, 26, 0, 27, 20, 24,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 11, 3,
	3, 3, 5, 6, 8, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 7, 3, 3, 10, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 15, 3, 16, 12, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 13, 3, 14, 9,
}
var yyTok2 = []int{

	2, 3, 4,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line grammar.y:48
		{
			yylex.(*parser).output = yyS[yypt-0].record
		}
	case 2:
		yyVAL.record = yyS[yypt-0].record
	case 3:
		yyVAL.record = yyS[yypt-0].record
	case 4:
		yyVAL.record = yyS[yypt-0].record
	case 5:
		//line grammar.y:56
		{
			yyVAL.record = newClassResult("exec", yyS[yypt-0].class_result.class, yyS[yypt-0].class_result.payload)
		}
	case 6:
		//line grammar.y:57
		{
			yyVAL.record = newClassResult("status", yyS[yypt-0].class_result.class, yyS[yypt-0].class_result.payload)
		}
	case 7:
		//line grammar.y:58
		{
			yyVAL.record = newClassResult("notify", yyS[yypt-0].class_result.class, yyS[yypt-0].class_result.payload)
		}
	case 8:
		//line grammar.y:61
		{
			yyVAL.class_result.class, yyVAL.class_result.payload = yyS[yypt-2].text, yyS[yypt-0].record
		}
	case 9:
		//line grammar.y:62
		{
			yyVAL.class_result.class, yyVAL.class_result.payload = yyS[yypt-0].text, nil
		}
	case 10:
		//line grammar.y:65
		{
			yyVAL.record = map[string]interface{}{typeKey: "console", payloadKey: yyS[yypt-0].text}
		}
	case 11:
		//line grammar.y:66
		{
			yyVAL.record = map[string]interface{}{typeKey: "target", payloadKey: yyS[yypt-0].text}
		}
	case 12:
		//line grammar.y:67
		{
			yyVAL.record = map[string]interface{}{typeKey: "log", payloadKey: yyS[yypt-0].text}
		}
	case 13:
		//line grammar.y:71
		{
			yyVAL.record = map[string]interface{}{sequenceKey: yyS[yypt-2].text, classKey: yyS[yypt-0].class_result.class}
			if yyS[yypt-0].class_result.payload != nil {
				yyVAL.record[payloadKey] = yyS[yypt-0].class_result.payload
			}
		}
	case 14:
		//line grammar.y:77
		{
			yyVAL.record[yyS[yypt-0].result_pair.variable] = yyS[yypt-0].result_pair.value
		}
	case 15:
		//line grammar.y:78
		{
			yyVAL.record = map[string]interface{}{yyS[yypt-0].result_pair.variable: yyS[yypt-0].result_pair.value}
		}
	case 16:
		//line grammar.y:81
		{
			yyVAL.result_pair.variable, yyVAL.result_pair.value = yyS[yypt-2].text, yyS[yypt-0].value
		}
	case 17:
		//line grammar.y:84
		{
			yyVAL.value = yyS[yypt-0].text
		}
	case 18:
		//line grammar.y:85
		{
			yyVAL.value = yyS[yypt-0].record
		}
	case 19:
		//line grammar.y:86
		{
			yyVAL.value = yyS[yypt-0].list
		}
	case 20:
		//line grammar.y:89
		{
			yyVAL.list = append(yyVAL.list, yyS[yypt-0].value)
		}
	case 21:
		//line grammar.y:90
		{
			yyVAL.list = []interface{}{yyS[yypt-0].value}
		}
	case 22:
		//line grammar.y:93
		{
			yyVAL.record = yyS[yypt-1].record
		}
	case 23:
		//line grammar.y:94
		{
			yyVAL.record = map[string]interface{}{}
		}
	case 24:
		//line grammar.y:97
		{
			yyVAL.list = append(yyVAL.list, map[string]interface{}{yyS[yypt-0].result_pair.variable: yyS[yypt-0].result_pair.value})
		}
	case 25:
		//line grammar.y:98
		{
			yyVAL.list = []interface{}{map[string]interface{}{yyS[yypt-0].result_pair.variable: yyS[yypt-0].result_pair.value}}
		}
	case 26:
		//line grammar.y:101
		{
			yyVAL.list = yyS[yypt-1].list
		}
	case 27:
		//line grammar.y:102
		{
			yyVAL.list = yyS[yypt-1].list
		}
	case 28:
		//line grammar.y:103
		{
			yyVAL.list = []interface{}{}
		}
	}
	goto yystack /* stack new state and value */
}
