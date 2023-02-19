package scanner

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testLoxCode = `
print "Hello, world!";

var imAVariable = "here is my value";
var iAmNil;
iAmNil == nil // true

var breakfast = "bagels";
print breakfast; // "bagels".
breakfast = "beignets";
print breakfast; // "beignets".

var addition = 1 + 1; // 2
var subtraction = 1 - 1; // 0
var multiplication = 1.0 * 1.0; // 1.0
var division = 2.0 / 1.0; // 2.0

var negation = -1; // -1

100 < 101; // true
101 <= 100; // false
100 > 101; // false
101 >= 101; // true

1 == 2;         // false
"cat" != "dog"; // true

123 != "123"; // true

!true;  // false
!false; // true

true and false; // false
true and true;  // true

false or false; // false
true or false;  // true

var average = (1 + 1) / 2; // 1

if (true) {
    print "yes";
  } else {
    print "no";
}

var a = 1;
while (a < 10) {
  print a;
  a = a + 1;
}

for (var a = 1; a < 10; a = a + 1) {
  print a;
}

fun printSum(a, b) {
  print a + b;
}

fun returnSum(a, b) {
  return a + b;
}

fun addPair(a, b) {
  return a + b;
}

fun identity(a) {
  return a;
}

print identity(addPair)(1, 2); // Prints "3".

fun outerFunction() {
  fun localFunction() {
    print "I'm local!";
  }

  localFunction();
}

class Breakfast {
  cook() {
    print "Eggs a-fryin'!";
  }

  serve(who) {
    print "Enjoy your breakfast, " + who + ".";
  }
}

// Store it in variables.
var someVariable = Breakfast;

// Pass it to functions.
someFunction(Breakfast);

var breakfast = Breakfast();
print breakfast; // "Breakfast instance".

class Breakfast {
  init(meat, bread) {
    this.meat = meat;
    this.bread = bread;
  }

  serve(who) {
    print "Enjoy your " + this.meat + " and " +
        this.bread + ", " + who + ".";
  }

  // ...
}

class Brunch < Breakfast {
  init(meat, bread, drink) {
    super.init(meat, bread);
    this.drink = drink;
  }

  drink() {
    print "How about a Bloody Mary?";
  }
}
`

func TestScanner_Success(t *testing.T) {
	scnr := NewScanner(testLoxCode)

	isAtEnd := scnr.IsAtEnd()
	assert.False(t, isAtEnd)

	tokens, errs := scnr.ScanTokens()

	isAtEnd = scnr.IsAtEnd()
	assert.True(t, isAtEnd)

	assert.Nil(t, errs)

	expectedTokens := []*token{
		{
			tokenType: Print,
			lexeme:    `print`,
			literal:   nil,
			line:      2,
		},
		{
			tokenType: String,
			lexeme:    `"Hello, world!"`,
			literal:   "Hello, world!",
			line:      2,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      2,
		},
		{
			tokenType: Var,
			lexeme:    `var`,
			literal:   nil,
			line:      4,
		},
		{
			tokenType: Identifier,
			lexeme:    `imAVariable`,
			literal:   nil,
			line:      4,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      4,
		},
		{
			tokenType: String,
			lexeme:    `"here is my value"`,
			literal:   "here is my value",
			line:      4,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      4,
		},
		{
			tokenType: Var,
			lexeme:    `var`,
			literal:   nil,
			line:      5,
		},
		{
			tokenType: Identifier,
			lexeme:    `iAmNil`,
			literal:   nil,
			line:      5,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      5,
		},
		{
			tokenType: Identifier,
			lexeme:    `iAmNil`,
			literal:   nil,
			line:      6,
		},
		{
			tokenType: EqualEqual,
			lexeme:    `==`,
			literal:   nil,
			line:      6,
		},
		{
			tokenType: Nil,
			lexeme:    `nil`,
			literal:   nil,
			line:      6,
		},
		{
			tokenType: Var,
			lexeme:    `var`,
			literal:   nil,
			line:      8,
		},
		{
			tokenType: Identifier,
			lexeme:    `breakfast`,
			literal:   nil,
			line:      8,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      8,
		},
		{
			tokenType: String,
			lexeme:    `"bagels"`,
			literal:   "bagels",
			line:      8,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      8,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
			literal:   nil,
			line:      9,
		},
		{
			tokenType: Identifier,
			lexeme:    `breakfast`,
			literal:   nil,
			line:      9,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      9,
		},
		{
			tokenType: Identifier,
			lexeme:    `breakfast`,
			literal:   nil,
			line:      10,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      10,
		},
		{
			tokenType: String,
			lexeme:    `"beignets"`,
			literal:   "beignets",
			line:      10,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      10,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
			literal:   nil,
			line:      11,
		},
		{
			tokenType: Identifier,
			lexeme:    `breakfast`,
			literal:   nil,
			line:      11,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      11,
		},
		{
			tokenType: Var,
			lexeme:    `var`,
			literal:   nil,
			line:      13,
		},
		{
			tokenType: Identifier,
			lexeme:    `addition`,
			literal:   nil,
			line:      13,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      13,
		},
		{
			tokenType: Number,
			lexeme:    `1`,
			literal:   1.0,
			line:      13,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      13,
		},
		{
			tokenType: Number,
			lexeme:    `1`,
			literal:   1.0,
			line:      13,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      13,
		},
		{
			tokenType: Var,
			lexeme:    `var`,
			literal:   nil,
			line:      14,
		},
		{
			tokenType: Identifier,
			lexeme:    `subtraction`,
			literal:   nil,
			line:      14,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      14,
		},
		{
			tokenType: Number,
			lexeme:    `1`,
			literal:   1.0,
			line:      14,
		},
		{
			tokenType: Minus,
			lexeme:    `-`,
			literal:   nil,
			line:      14,
		},
		{
			tokenType: Number,
			lexeme:    `1`,
			literal:   1.0,
			line:      14,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      14,
		},
		{
			tokenType: Var,
			lexeme:    `var`,
			literal:   nil,
			line:      15,
		},
		{
			tokenType: Identifier,
			lexeme:    `multiplication`,
			literal:   nil,
			line:      15,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      15,
		},
		{
			tokenType: Number,
			lexeme:    `1.0`,
			literal:   1.0,
			line:      15,
		},
		{
			tokenType: Star,
			lexeme:    `*`,
			literal:   nil,
			line:      15,
		},
		{
			tokenType: Number,
			lexeme:    `1.0`,
			literal:   1.0,
			line:      15,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      15,
		},
		{
			tokenType: Var,
			lexeme:    `var`,
			literal:   nil,
			line:      16,
		},
		{
			tokenType: Identifier,
			lexeme:    `division`,
			literal:   nil,
			line:      16,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      16,
		},
		{
			tokenType: Number,
			lexeme:    `2.0`,
			literal:   2.0,
			line:      16,
		},
		{
			tokenType: Slash,
			lexeme:    `/`,
			literal:   nil,
			line:      16,
		},
		{
			tokenType: Number,
			lexeme:    `1.0`,
			literal:   1.0,
			line:      16,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      16,
		},
		{
			tokenType: Var,
			lexeme:    `var`,
			literal:   nil,
			line:      18,
		},
		{
			tokenType: Identifier,
			lexeme:    `negation`,
			literal:   nil,
			line:      18,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      18,
		},
		{
			tokenType: Minus,
			lexeme:    `-`,
			literal:   nil,
			line:      18,
		},
		{
			tokenType: Number,
			lexeme:    `1`,
			literal:   1.0,
			line:      18,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      18,
		},
		{
			tokenType: Number,
			lexeme:    `100`,
			literal:   100.0,
			line:      20,
		},
		{
			tokenType: Less,
			lexeme:    `<`,
			literal:   nil,
			line:      20,
		},
		{
			tokenType: Number,
			lexeme:    `101`,
			literal:   101.0,
			line:      20,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      20,
		},
		{
			tokenType: Number,
			lexeme:    `101`,
			literal:   101.0,
			line:      21,
		},
		{
			tokenType: LessEqual,
			lexeme:    `<=`,
			literal:   nil,
			line:      21,
		},
		{
			tokenType: Number,
			lexeme:    `100`,
			literal:   100.0,
			line:      21,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      21,
		},
		{
			tokenType: Number,
			lexeme:    `100`,
			literal:   100.0,
			line:      22,
		},
		{
			tokenType: Greater,
			lexeme:    `>`,
			literal:   nil,
			line:      22,
		},
		{
			tokenType: Number,
			lexeme:    `101`,
			literal:   101.0,
			line:      22,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      22,
		},
		{
			tokenType: Number,
			lexeme:    `101`,
			literal:   101.0,
			line:      23,
		},
		{
			tokenType: GreaterEqual,
			lexeme:    `>=`,
			literal:   nil,
			line:      23,
		},
		{
			tokenType: Number,
			lexeme:    `101`,
			literal:   101.0,
			line:      23,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      23,
		},
		{
			tokenType: Number,
			lexeme:    `1`,
			literal:   1.0,
			line:      25,
		},
		{
			tokenType: EqualEqual,
			lexeme:    `==`,
			literal:   nil,
			line:      25,
		},
		{
			tokenType: Number,
			lexeme:    `2`,
			literal:   2.0,
			line:      25,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      25,
		},
		{
			tokenType: String,
			lexeme:    `"cat"`,
			literal:   "cat",
			line:      26,
		},
		{
			tokenType: BangEqual,
			lexeme:    `!=`,
			literal:   nil,
			line:      26,
		},
		{
			tokenType: String,
			lexeme:    `"dog"`,
			literal:   "dog",
			line:      26,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      26,
		},
		{
			tokenType: Number,
			lexeme:    `123`,
			literal:   123.0,
			line:      28,
		},
		{
			tokenType: BangEqual,
			lexeme:    `!=`,
			literal:   nil,
			line:      28,
		},
		{
			tokenType: String,
			lexeme:    `"123"`,
			literal:   "123",
			line:      28,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      28,
		},
		{
			tokenType: Bang,
			lexeme:    `!`,
			literal:   nil,
			line:      30,
		},
		{
			tokenType: True,
			lexeme:    `true`,
			literal:   nil,
			line:      30,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      30,
		},
		{
			tokenType: Bang,
			lexeme:    `!`,
			literal:   nil,
			line:      31,
		},
		{
			tokenType: False,
			lexeme:    `false`,
			literal:   nil,
			line:      31,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      31,
		},
		{
			tokenType: True,
			lexeme:    `true`,
			literal:   nil,
			line:      33,
		},
		{
			tokenType: And,
			lexeme:    `and`,
			literal:   nil,
			line:      33,
		},
		{
			tokenType: False,
			lexeme:    `false`,
			literal:   nil,
			line:      33,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      33,
		},
		{
			tokenType: True,
			lexeme:    `true`,
			literal:   nil,
			line:      34,
		},
		{
			tokenType: And,
			lexeme:    `and`,
			literal:   nil,
			line:      34,
		},
		{
			tokenType: True,
			lexeme:    `true`,
			literal:   nil,
			line:      34,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      34,
		},
		{
			tokenType: False,
			lexeme:    `false`,
			literal:   nil,
			line:      36,
		},
		{
			tokenType: Or,
			lexeme:    `or`,
			literal:   nil,
			line:      36,
		},
		{
			tokenType: False,
			lexeme:    `false`,
			literal:   nil,
			line:      36,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      36,
		},
		{
			tokenType: True,
			lexeme:    `true`,
			literal:   nil,
			line:      37,
		},
		{
			tokenType: Or,
			lexeme:    `or`,
			literal:   nil,
			line:      37,
		},
		{
			tokenType: False,
			lexeme:    `false`,
			literal:   nil,
			line:      37,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      37,
		},
		{
			tokenType: Var,
			lexeme:    `var`,
			literal:   nil,
			line:      39,
		},
		{
			tokenType: Identifier,
			lexeme:    `average`,
			literal:   nil,
			line:      39,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      39,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      39,
		},
		{
			tokenType: Number,
			lexeme:    `1`,
			literal:   1.0,
			line:      39,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      39,
		},
		{
			tokenType: Number,
			lexeme:    `1`,
			literal:   1.0,
			line:      39,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      39,
		},
		{
			tokenType: Slash,
			lexeme:    `/`,
			literal:   nil,
			line:      39,
		},
		{
			tokenType: Number,
			lexeme:    `2`,
			literal:   2.0,
			line:      39,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      39,
		},
		{
			tokenType: If,
			lexeme:    `if`,
			literal:   nil,
			line:      41,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      41,
		},
		{
			tokenType: True,
			lexeme:    `true`,
			literal:   nil,
			line:      41,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      41,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      41,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
			literal:   nil,
			line:      42,
		},
		{
			tokenType: String,
			lexeme:    `"yes"`,
			literal:   "yes",
			line:      42,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      42,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      43,
		},
		{
			tokenType: Else,
			lexeme:    `else`,
			literal:   nil,
			line:      43,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      43,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
			literal:   nil,
			line:      44,
		},
		{
			tokenType: String,
			lexeme:    `"no"`,
			literal:   "no",
			line:      44,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      44,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      45,
		},
		{
			tokenType: Var,
			lexeme:    `var`,
			literal:   nil,
			line:      47,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      47,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      47,
		},
		{
			tokenType: Number,
			lexeme:    `1`,
			literal:   1.0,
			line:      47,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      47,
		},
		{
			tokenType: While,
			lexeme:    `while`,
			literal:   nil,
			line:      48,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      48,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      48,
		},
		{
			tokenType: Less,
			lexeme:    `<`,
			literal:   nil,
			line:      48,
		},
		{
			tokenType: Number,
			lexeme:    `10`,
			literal:   10.0,
			line:      48,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      48,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      48,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
			literal:   nil,
			line:      49,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      49,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      49,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      50,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      50,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      50,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      50,
		},
		{
			tokenType: Number,
			lexeme:    `1`,
			literal:   1.0,
			line:      50,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      50,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      51,
		},
		{
			tokenType: For,
			lexeme:    `for`,
			literal:   nil,
			line:      53,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      53,
		},
		{
			tokenType: Var,
			lexeme:    `var`,
			literal:   nil,
			line:      53,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      53,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      53,
		},
		{
			tokenType: Number,
			lexeme:    `1`,
			literal:   1.0,
			line:      53,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      53,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      53,
		},
		{
			tokenType: Less,
			lexeme:    `<`,
			literal:   nil,
			line:      53,
		},
		{
			tokenType: Number,
			lexeme:    `10`,
			literal:   10.0,
			line:      53,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      53,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      53,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      53,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      53,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      53,
		},
		{
			tokenType: Number,
			lexeme:    `1`,
			literal:   1.0,
			line:      53,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      53,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      53,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
			literal:   nil,
			line:      54,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      54,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      54,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      55,
		},
		{
			tokenType: Fun,
			lexeme:    `fun`,
			literal:   nil,
			line:      57,
		},
		{
			tokenType: Identifier,
			lexeme:    `printSum`,
			literal:   nil,
			line:      57,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      57,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      57,
		},
		{
			tokenType: Comma,
			lexeme:    `,`,
			literal:   nil,
			line:      57,
		},
		{
			tokenType: Identifier,
			lexeme:    `b`,
			literal:   nil,
			line:      57,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      57,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      57,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
			literal:   nil,
			line:      58,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      58,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      58,
		},
		{
			tokenType: Identifier,
			lexeme:    `b`,
			literal:   nil,
			line:      58,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      58,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      59,
		},
		{
			tokenType: Fun,
			lexeme:    `fun`,
			literal:   nil,
			line:      61,
		},
		{
			tokenType: Identifier,
			lexeme:    `returnSum`,
			literal:   nil,
			line:      61,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      61,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      61,
		},
		{
			tokenType: Comma,
			lexeme:    `,`,
			literal:   nil,
			line:      61,
		},
		{
			tokenType: Identifier,
			lexeme:    `b`,
			literal:   nil,
			line:      61,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      61,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      61,
		},
		{
			tokenType: Return,
			lexeme:    `return`,
			literal:   nil,
			line:      62,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      62,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      62,
		},
		{
			tokenType: Identifier,
			lexeme:    `b`,
			literal:   nil,
			line:      62,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      62,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      63,
		},
		{
			tokenType: Fun,
			lexeme:    `fun`,
			literal:   nil,
			line:      65,
		},
		{
			tokenType: Identifier,
			lexeme:    `addPair`,
			literal:   nil,
			line:      65,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      65,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      65,
		},
		{
			tokenType: Comma,
			lexeme:    `,`,
			literal:   nil,
			line:      65,
		},
		{
			tokenType: Identifier,
			lexeme:    `b`,
			literal:   nil,
			line:      65,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      65,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      65,
		},
		{
			tokenType: Return,
			lexeme:    `return`,
			literal:   nil,
			line:      66,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      66,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      66,
		},
		{
			tokenType: Identifier,
			lexeme:    `b`,
			literal:   nil,
			line:      66,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      66,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      67,
		},
		{
			tokenType: Fun,
			lexeme:    `fun`,
			literal:   nil,
			line:      69,
		},
		{
			tokenType: Identifier,
			lexeme:    `identity`,
			literal:   nil,
			line:      69,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      69,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      69,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      69,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      69,
		},
		{
			tokenType: Return,
			lexeme:    `return`,
			literal:   nil,
			line:      70,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      70,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      70,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      71,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
			literal:   nil,
			line:      73,
		},
		{
			tokenType: Identifier,
			lexeme:    `identity`,
			literal:   nil,
			line:      73,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      73,
		},
		{
			tokenType: Identifier,
			lexeme:    `addPair`,
			literal:   nil,
			line:      73,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      73,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      73,
		},
		{
			tokenType: Number,
			lexeme:    `1`,
			literal:   1.0,
			line:      73,
		},
		{
			tokenType: Comma,
			lexeme:    `,`,
			literal:   nil,
			line:      73,
		},
		{
			tokenType: Number,
			lexeme:    `2`,
			literal:   2.0,
			line:      73,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      73,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      73,
		},
		{
			tokenType: Fun,
			lexeme:    `fun`,
			literal:   nil,
			line:      75,
		},
		{
			tokenType: Identifier,
			lexeme:    `outerFunction`,
			literal:   nil,
			line:      75,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      75,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      75,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      75,
		},
		{
			tokenType: Fun,
			lexeme:    `fun`,
			literal:   nil,
			line:      76,
		},
		{
			tokenType: Identifier,
			lexeme:    `localFunction`,
			literal:   nil,
			line:      76,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      76,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      76,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      76,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
			literal:   nil,
			line:      77,
		},
		{
			tokenType: String,
			lexeme:    `"I'm local!"`,
			literal:   "I'm local!",
			line:      77,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      77,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      78,
		},
		{
			tokenType: Identifier,
			lexeme:    `localFunction`,
			literal:   nil,
			line:      80,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      80,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      80,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      80,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      81,
		},
		{
			tokenType: Class,
			lexeme:    `class`,
			literal:   nil,
			line:      83,
		},
		{
			tokenType: Identifier,
			lexeme:    `Breakfast`,
			literal:   nil,
			line:      83,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      83,
		},
		{
			tokenType: Identifier,
			lexeme:    `cook`,
			literal:   nil,
			line:      84,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      84,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      84,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      84,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
			literal:   nil,
			line:      85,
		},
		{
			tokenType: String,
			lexeme:    `"Eggs a-fryin'!"`,
			literal:   "Eggs a-fryin'!",
			line:      85,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      85,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      86,
		},
		{
			tokenType: Identifier,
			lexeme:    `serve`,
			literal:   nil,
			line:      88,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      88,
		},
		{
			tokenType: Identifier,
			lexeme:    `who`,
			literal:   nil,
			line:      88,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      88,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      88,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
			literal:   nil,
			line:      89,
		},
		{
			tokenType: String,
			lexeme:    `"Enjoy your breakfast, "`,
			literal:   "Enjoy your breakfast, ",
			line:      89,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      89,
		},
		{
			tokenType: Identifier,
			lexeme:    `who`,
			literal:   nil,
			line:      89,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      89,
		},
		{
			tokenType: String,
			lexeme:    `"."`,
			literal:   ".",
			line:      89,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      89,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      90,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      91,
		},
		{
			tokenType: Var,
			lexeme:    `var`,
			literal:   nil,
			line:      94,
		},
		{
			tokenType: Identifier,
			lexeme:    `someVariable`,
			literal:   nil,
			line:      94,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      94,
		},
		{
			tokenType: Identifier,
			lexeme:    `Breakfast`,
			literal:   nil,
			line:      94,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      94,
		},
		{
			tokenType: Identifier,
			lexeme:    `someFunction`,
			literal:   nil,
			line:      97,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      97,
		},
		{
			tokenType: Identifier,
			lexeme:    `Breakfast`,
			literal:   nil,
			line:      97,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      97,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      97,
		},
		{
			tokenType: Var,
			lexeme:    `var`,
			literal:   nil,
			line:      99,
		},
		{
			tokenType: Identifier,
			lexeme:    `breakfast`,
			literal:   nil,
			line:      99,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      99,
		},
		{
			tokenType: Identifier,
			lexeme:    `Breakfast`,
			literal:   nil,
			line:      99,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      99,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      99,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      99,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
			literal:   nil,
			line:      100,
		},
		{
			tokenType: Identifier,
			lexeme:    `breakfast`,
			literal:   nil,
			line:      100,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      100,
		},
		{
			tokenType: Class,
			lexeme:    `class`,
			literal:   nil,
			line:      102,
		},
		{
			tokenType: Identifier,
			lexeme:    `Breakfast`,
			literal:   nil,
			line:      102,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      102,
		},
		{
			tokenType: Identifier,
			lexeme:    `init`,
			literal:   nil,
			line:      103,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      103,
		},
		{
			tokenType: Identifier,
			lexeme:    `meat`,
			literal:   nil,
			line:      103,
		},
		{
			tokenType: Comma,
			lexeme:    `,`,
			literal:   nil,
			line:      103,
		},
		{
			tokenType: Identifier,
			lexeme:    `bread`,
			literal:   nil,
			line:      103,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      103,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      103,
		},
		{
			tokenType: This,
			lexeme:    `this`,
			literal:   nil,
			line:      104,
		},
		{
			tokenType: Dot,
			lexeme:    `.`,
			literal:   nil,
			line:      104,
		},
		{
			tokenType: Identifier,
			lexeme:    `meat`,
			literal:   nil,
			line:      104,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      104,
		},
		{
			tokenType: Identifier,
			lexeme:    `meat`,
			literal:   nil,
			line:      104,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      104,
		},
		{
			tokenType: This,
			lexeme:    `this`,
			literal:   nil,
			line:      105,
		},
		{
			tokenType: Dot,
			lexeme:    `.`,
			literal:   nil,
			line:      105,
		},
		{
			tokenType: Identifier,
			lexeme:    `bread`,
			literal:   nil,
			line:      105,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      105,
		},
		{
			tokenType: Identifier,
			lexeme:    `bread`,
			literal:   nil,
			line:      105,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      105,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      106,
		},
		{
			tokenType: Identifier,
			lexeme:    `serve`,
			literal:   nil,
			line:      108,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      108,
		},
		{
			tokenType: Identifier,
			lexeme:    `who`,
			literal:   nil,
			line:      108,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      108,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      108,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
			literal:   nil,
			line:      109,
		},
		{
			tokenType: String,
			lexeme:    `"Enjoy your "`,
			literal:   "Enjoy your ",
			line:      109,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      109,
		},
		{
			tokenType: This,
			lexeme:    `this`,
			literal:   nil,
			line:      109,
		},
		{
			tokenType: Dot,
			lexeme:    `.`,
			literal:   nil,
			line:      109,
		},
		{
			tokenType: Identifier,
			lexeme:    `meat`,
			literal:   nil,
			line:      109,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      109,
		},
		{
			tokenType: String,
			lexeme:    `" and "`,
			literal:   " and ",
			line:      109,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      109,
		},
		{
			tokenType: This,
			lexeme:    `this`,
			literal:   nil,
			line:      110,
		},
		{
			tokenType: Dot,
			lexeme:    `.`,
			literal:   nil,
			line:      110,
		},
		{
			tokenType: Identifier,
			lexeme:    `bread`,
			literal:   nil,
			line:      110,
		},

		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      110,
		},
		{
			tokenType: String,
			lexeme:    `", "`,
			literal:   ", ",
			line:      110,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      110,
		},
		{
			tokenType: Identifier,
			lexeme:    `who`,
			literal:   nil,
			line:      110,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      110,
		},
		{
			tokenType: String,
			lexeme:    `"."`,
			literal:   ".",
			line:      110,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      110,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      111,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      114,
		},
		{
			tokenType: Class,
			lexeme:    `class`,
			literal:   nil,
			line:      116,
		},
		{
			tokenType: Identifier,
			lexeme:    `Brunch`,
			literal:   nil,
			line:      116,
		},
		{
			tokenType: Less,
			lexeme:    `<`,
			literal:   nil,
			line:      116,
		},
		{
			tokenType: Identifier,
			lexeme:    `Breakfast`,
			literal:   nil,
			line:      116,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      116,
		},
		{
			tokenType: Identifier,
			lexeme:    `init`,
			literal:   nil,
			line:      117,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      117,
		},
		{
			tokenType: Identifier,
			lexeme:    `meat`,
			literal:   nil,
			line:      117,
		},
		{
			tokenType: Comma,
			lexeme:    `,`,
			literal:   nil,
			line:      117,
		},
		{
			tokenType: Identifier,
			lexeme:    `bread`,
			literal:   nil,
			line:      117,
		},
		{
			tokenType: Comma,
			lexeme:    `,`,
			literal:   nil,
			line:      117,
		},
		{
			tokenType: Identifier,
			lexeme:    `drink`,
			literal:   nil,
			line:      117,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      117,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      117,
		},
		{
			tokenType: Super,
			lexeme:    `super`,
			literal:   nil,
			line:      118,
		},
		{
			tokenType: Dot,
			lexeme:    `.`,
			literal:   nil,
			line:      118,
		},
		{
			tokenType: Identifier,
			lexeme:    `init`,
			literal:   nil,
			line:      118,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      118,
		},
		{
			tokenType: Identifier,
			lexeme:    `meat`,
			literal:   nil,
			line:      118,
		},
		{
			tokenType: Comma,
			lexeme:    `,`,
			literal:   nil,
			line:      118,
		},
		{
			tokenType: Identifier,
			lexeme:    `bread`,
			literal:   nil,
			line:      118,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      118,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      118,
		},
		{
			tokenType: This,
			lexeme:    `this`,
			literal:   nil,
			line:      119,
		},
		{
			tokenType: Dot,
			lexeme:    `.`,
			literal:   nil,
			line:      119,
		},
		{
			tokenType: Identifier,
			lexeme:    `drink`,
			literal:   nil,
			line:      119,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      119,
		},
		{
			tokenType: Identifier,
			lexeme:    `drink`,
			literal:   nil,
			line:      119,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      119,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      120,
		},
		{
			tokenType: Identifier,
			lexeme:    `drink`,
			literal:   nil,
			line:      122,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      122,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      122,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      122,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
			literal:   nil,
			line:      123,
		},
		{
			tokenType: String,
			lexeme:    `"How about a Bloody Mary?"`,
			literal:   "How about a Bloody Mary?",
			line:      123,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      123,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      124,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      125,
		},
		{
			tokenType: EOF,
			lexeme:    ``,
			literal:   nil,
			line:      126,
		},
	}

	for i, tkn := range tokens {
		assert.Equal(t, expectedTokens[i], tkn)
	}
}
