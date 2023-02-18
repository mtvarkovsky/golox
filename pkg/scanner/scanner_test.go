package scanner

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testLoxCode = `
print "Hello, world!";

var imAVariable = "here is my value";
var iAmNil;

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
			tokenType: Var,
			lexeme:    `var`,
			literal:   nil,
			line:      7,
		},
		{
			tokenType: Identifier,
			lexeme:    `breakfast`,
			literal:   nil,
			line:      7,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      7,
		},
		{
			tokenType: String,
			lexeme:    `"bagels"`,
			literal:   "bagels",
			line:      7,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      7,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
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
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      8,
		},
		{
			tokenType: Identifier,
			lexeme:    `breakfast`,
			literal:   nil,
			line:      9,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      9,
		},
		{
			tokenType: String,
			lexeme:    `"beignets"`,
			literal:   "beignets",
			line:      9,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      9,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
			literal:   nil,
			line:      10,
		},
		{
			tokenType: Identifier,
			lexeme:    `breakfast`,
			literal:   nil,
			line:      10,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      10,
		},
		{
			tokenType: Var,
			lexeme:    `var`,
			literal:   nil,
			line:      12,
		},
		{
			tokenType: Identifier,
			lexeme:    `addition`,
			literal:   nil,
			line:      12,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      12,
		},
		{
			tokenType: Number,
			lexeme:    `1`,
			literal:   1.0,
			line:      12,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      12,
		},
		{
			tokenType: Number,
			lexeme:    `1`,
			literal:   1.0,
			line:      12,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      12,
		},
		{
			tokenType: Var,
			lexeme:    `var`,
			literal:   nil,
			line:      13,
		},
		{
			tokenType: Identifier,
			lexeme:    `subtraction`,
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
			tokenType: Minus,
			lexeme:    `-`,
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
			lexeme:    `multiplication`,
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
			lexeme:    `1.0`,
			literal:   1.0,
			line:      14,
		},
		{
			tokenType: Star,
			lexeme:    `*`,
			literal:   nil,
			line:      14,
		},
		{
			tokenType: Number,
			lexeme:    `1.0`,
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
			lexeme:    `division`,
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
			lexeme:    `2.0`,
			literal:   2.0,
			line:      15,
		},
		{
			tokenType: Slash,
			lexeme:    `/`,
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
			line:      17,
		},
		{
			tokenType: Identifier,
			lexeme:    `negation`,
			literal:   nil,
			line:      17,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      17,
		},
		{
			tokenType: Minus,
			lexeme:    `-`,
			literal:   nil,
			line:      17,
		},
		{
			tokenType: Number,
			lexeme:    `1`,
			literal:   1.0,
			line:      17,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      17,
		},
		{
			tokenType: Number,
			lexeme:    `100`,
			literal:   100.0,
			line:      19,
		},
		{
			tokenType: Less,
			lexeme:    `<`,
			literal:   nil,
			line:      19,
		},
		{
			tokenType: Number,
			lexeme:    `101`,
			literal:   101.0,
			line:      19,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      19,
		},
		{
			tokenType: Number,
			lexeme:    `101`,
			literal:   101.0,
			line:      20,
		},
		{
			tokenType: LessEqual,
			lexeme:    `<=`,
			literal:   nil,
			line:      20,
		},
		{
			tokenType: Number,
			lexeme:    `100`,
			literal:   100.0,
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
			lexeme:    `100`,
			literal:   100.0,
			line:      21,
		},
		{
			tokenType: Greater,
			lexeme:    `>`,
			literal:   nil,
			line:      21,
		},
		{
			tokenType: Number,
			lexeme:    `101`,
			literal:   101.0,
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
			lexeme:    `101`,
			literal:   101.0,
			line:      22,
		},
		{
			tokenType: GreaterEqual,
			lexeme:    `>=`,
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
			lexeme:    `1`,
			literal:   1.0,
			line:      24,
		},
		{
			tokenType: EqualEqual,
			lexeme:    `==`,
			literal:   nil,
			line:      24,
		},
		{
			tokenType: Number,
			lexeme:    `2`,
			literal:   2.0,
			line:      24,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      24,
		},
		{
			tokenType: String,
			lexeme:    `"cat"`,
			literal:   "cat",
			line:      25,
		},
		{
			tokenType: BangEqual,
			lexeme:    `!=`,
			literal:   nil,
			line:      25,
		},
		{
			tokenType: String,
			lexeme:    `"dog"`,
			literal:   "dog",
			line:      25,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      25,
		},
		{
			tokenType: Number,
			lexeme:    `123`,
			literal:   123.0,
			line:      27,
		},
		{
			tokenType: BangEqual,
			lexeme:    `!=`,
			literal:   nil,
			line:      27,
		},
		{
			tokenType: String,
			lexeme:    `"123"`,
			literal:   "123",
			line:      27,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      27,
		},
		{
			tokenType: Bang,
			lexeme:    `!`,
			literal:   nil,
			line:      29,
		},
		{
			tokenType: True,
			lexeme:    `true`,
			literal:   nil,
			line:      29,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      29,
		},
		{
			tokenType: Bang,
			lexeme:    `!`,
			literal:   nil,
			line:      30,
		},
		{
			tokenType: False,
			lexeme:    `false`,
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
			tokenType: True,
			lexeme:    `true`,
			literal:   nil,
			line:      32,
		},
		{
			tokenType: And,
			lexeme:    `and`,
			literal:   nil,
			line:      32,
		},
		{
			tokenType: False,
			lexeme:    `false`,
			literal:   nil,
			line:      32,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      32,
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
			tokenType: True,
			lexeme:    `true`,
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
			tokenType: False,
			lexeme:    `false`,
			literal:   nil,
			line:      35,
		},
		{
			tokenType: Or,
			lexeme:    `or`,
			literal:   nil,
			line:      35,
		},
		{
			tokenType: False,
			lexeme:    `false`,
			literal:   nil,
			line:      35,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      35,
		},
		{
			tokenType: True,
			lexeme:    `true`,
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
			tokenType: Var,
			lexeme:    `var`,
			literal:   nil,
			line:      38,
		},
		{
			tokenType: Identifier,
			lexeme:    `average`,
			literal:   nil,
			line:      38,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      38,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      38,
		},
		{
			tokenType: Number,
			lexeme:    `1`,
			literal:   1.0,
			line:      38,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      38,
		},
		{
			tokenType: Number,
			lexeme:    `1`,
			literal:   1.0,
			line:      38,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      38,
		},
		{
			tokenType: Slash,
			lexeme:    `/`,
			literal:   nil,
			line:      38,
		},
		{
			tokenType: Number,
			lexeme:    `2`,
			literal:   2.0,
			line:      38,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      38,
		},
		{
			tokenType: If,
			lexeme:    `if`,
			literal:   nil,
			line:      40,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      40,
		},
		{
			tokenType: True,
			lexeme:    `true`,
			literal:   nil,
			line:      40,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      40,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      40,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
			literal:   nil,
			line:      41,
		},
		{
			tokenType: String,
			lexeme:    `"yes"`,
			literal:   "yes",
			line:      41,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      41,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      42,
		},
		{
			tokenType: Else,
			lexeme:    `else`,
			literal:   nil,
			line:      42,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      42,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
			literal:   nil,
			line:      43,
		},
		{
			tokenType: String,
			lexeme:    `"no"`,
			literal:   "no",
			line:      43,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      43,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      44,
		},
		{
			tokenType: Var,
			lexeme:    `var`,
			literal:   nil,
			line:      46,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      46,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      46,
		},
		{
			tokenType: Number,
			lexeme:    `1`,
			literal:   1.0,
			line:      46,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      46,
		},
		{
			tokenType: While,
			lexeme:    `while`,
			literal:   nil,
			line:      47,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
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
			tokenType: Less,
			lexeme:    `<`,
			literal:   nil,
			line:      47,
		},
		{
			tokenType: Number,
			lexeme:    `10`,
			literal:   10.0,
			line:      47,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      47,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      47,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
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
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      48,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      49,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
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
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      49,
		},
		{
			tokenType: Number,
			lexeme:    `1`,
			literal:   1.0,
			line:      49,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      49,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      50,
		},
		{
			tokenType: For,
			lexeme:    `for`,
			literal:   nil,
			line:      52,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      52,
		},
		{
			tokenType: Var,
			lexeme:    `var`,
			literal:   nil,
			line:      52,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      52,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      52,
		},
		{
			tokenType: Number,
			lexeme:    `1`,
			literal:   1.0,
			line:      52,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      52,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      52,
		},
		{
			tokenType: Less,
			lexeme:    `<`,
			literal:   nil,
			line:      52,
		},
		{
			tokenType: Number,
			lexeme:    `10`,
			literal:   10.0,
			line:      52,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      52,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      52,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      52,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      52,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      52,
		},
		{
			tokenType: Number,
			lexeme:    `1`,
			literal:   1.0,
			line:      52,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      52,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      52,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
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
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      53,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      54,
		},
		{
			tokenType: Fun,
			lexeme:    `fun`,
			literal:   nil,
			line:      56,
		},
		{
			tokenType: Identifier,
			lexeme:    `printSum`,
			literal:   nil,
			line:      56,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      56,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      56,
		},
		{
			tokenType: Comma,
			lexeme:    `,`,
			literal:   nil,
			line:      56,
		},
		{
			tokenType: Identifier,
			lexeme:    `b`,
			literal:   nil,
			line:      56,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      56,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      56,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
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
			tokenType: Plus,
			lexeme:    `+`,
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
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      57,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      58,
		},
		{
			tokenType: Fun,
			lexeme:    `fun`,
			literal:   nil,
			line:      60,
		},
		{
			tokenType: Identifier,
			lexeme:    `returnSum`,
			literal:   nil,
			line:      60,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      60,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      60,
		},
		{
			tokenType: Comma,
			lexeme:    `,`,
			literal:   nil,
			line:      60,
		},
		{
			tokenType: Identifier,
			lexeme:    `b`,
			literal:   nil,
			line:      60,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      60,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      60,
		},
		{
			tokenType: Return,
			lexeme:    `return`,
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
			tokenType: Plus,
			lexeme:    `+`,
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
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      61,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      62,
		},
		{
			tokenType: Fun,
			lexeme:    `fun`,
			literal:   nil,
			line:      64,
		},
		{
			tokenType: Identifier,
			lexeme:    `addPair`,
			literal:   nil,
			line:      64,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      64,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      64,
		},
		{
			tokenType: Comma,
			lexeme:    `,`,
			literal:   nil,
			line:      64,
		},
		{
			tokenType: Identifier,
			lexeme:    `b`,
			literal:   nil,
			line:      64,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      64,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      64,
		},
		{
			tokenType: Return,
			lexeme:    `return`,
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
			tokenType: Plus,
			lexeme:    `+`,
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
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      65,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      66,
		},
		{
			tokenType: Fun,
			lexeme:    `fun`,
			literal:   nil,
			line:      68,
		},
		{
			tokenType: Identifier,
			lexeme:    `identity`,
			literal:   nil,
			line:      68,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      68,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      68,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      68,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      68,
		},
		{
			tokenType: Return,
			lexeme:    `return`,
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
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      69,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      70,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
			literal:   nil,
			line:      72,
		},
		{
			tokenType: Identifier,
			lexeme:    `identity`,
			literal:   nil,
			line:      72,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      72,
		},
		{
			tokenType: Identifier,
			lexeme:    `addPair`,
			literal:   nil,
			line:      72,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      72,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      72,
		},
		{
			tokenType: Number,
			lexeme:    `1`,
			literal:   1.0,
			line:      72,
		},
		{
			tokenType: Comma,
			lexeme:    `,`,
			literal:   nil,
			line:      72,
		},
		{
			tokenType: Number,
			lexeme:    `2`,
			literal:   2.0,
			line:      72,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      72,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      72,
		},
		{
			tokenType: Fun,
			lexeme:    `fun`,
			literal:   nil,
			line:      74,
		},
		{
			tokenType: Identifier,
			lexeme:    `outerFunction`,
			literal:   nil,
			line:      74,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      74,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      74,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      74,
		},
		{
			tokenType: Fun,
			lexeme:    `fun`,
			literal:   nil,
			line:      75,
		},
		{
			tokenType: Identifier,
			lexeme:    `localFunction`,
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
			tokenType: Print,
			lexeme:    `print`,
			literal:   nil,
			line:      76,
		},
		{
			tokenType: String,
			lexeme:    `"I'm local!"`,
			literal:   "I'm local!",
			line:      76,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      76,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      77,
		},
		{
			tokenType: Identifier,
			lexeme:    `localFunction`,
			literal:   nil,
			line:      79,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      79,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      79,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      79,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      80,
		},
		{
			tokenType: Class,
			lexeme:    `class`,
			literal:   nil,
			line:      82,
		},
		{
			tokenType: Identifier,
			lexeme:    `Breakfast`,
			literal:   nil,
			line:      82,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      82,
		},
		{
			tokenType: Identifier,
			lexeme:    `cook`,
			literal:   nil,
			line:      83,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      83,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
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
			tokenType: Print,
			lexeme:    `print`,
			literal:   nil,
			line:      84,
		},
		{
			tokenType: String,
			lexeme:    `"Eggs a-fryin'!"`,
			literal:   "Eggs a-fryin'!",
			line:      84,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      84,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      85,
		},
		{
			tokenType: Identifier,
			lexeme:    `serve`,
			literal:   nil,
			line:      87,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      87,
		},
		{
			tokenType: Identifier,
			lexeme:    `who`,
			literal:   nil,
			line:      87,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      87,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      87,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
			literal:   nil,
			line:      88,
		},
		{
			tokenType: String,
			lexeme:    `"Enjoy your breakfast, "`,
			literal:   "Enjoy your breakfast, ",
			line:      88,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
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
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      88,
		},
		{
			tokenType: String,
			lexeme:    `"."`,
			literal:   ".",
			line:      88,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      88,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
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
			tokenType: Var,
			lexeme:    `var`,
			literal:   nil,
			line:      93,
		},
		{
			tokenType: Identifier,
			lexeme:    `someVariable`,
			literal:   nil,
			line:      93,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      93,
		},
		{
			tokenType: Identifier,
			lexeme:    `Breakfast`,
			literal:   nil,
			line:      93,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      93,
		},
		{
			tokenType: Identifier,
			lexeme:    `someFunction`,
			literal:   nil,
			line:      96,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      96,
		},
		{
			tokenType: Identifier,
			lexeme:    `Breakfast`,
			literal:   nil,
			line:      96,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      96,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      96,
		},
		{
			tokenType: Var,
			lexeme:    `var`,
			literal:   nil,
			line:      98,
		},
		{
			tokenType: Identifier,
			lexeme:    `breakfast`,
			literal:   nil,
			line:      98,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      98,
		},
		{
			tokenType: Identifier,
			lexeme:    `Breakfast`,
			literal:   nil,
			line:      98,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      98,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      98,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      98,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
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
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      99,
		},
		{
			tokenType: Class,
			lexeme:    `class`,
			literal:   nil,
			line:      101,
		},
		{
			tokenType: Identifier,
			lexeme:    `Breakfast`,
			literal:   nil,
			line:      101,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      101,
		},
		{
			tokenType: Identifier,
			lexeme:    `init`,
			literal:   nil,
			line:      102,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      102,
		},
		{
			tokenType: Identifier,
			lexeme:    `meat`,
			literal:   nil,
			line:      102,
		},
		{
			tokenType: Comma,
			lexeme:    `,`,
			literal:   nil,
			line:      102,
		},
		{
			tokenType: Identifier,
			lexeme:    `bread`,
			literal:   nil,
			line:      102,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
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
			tokenType: This,
			lexeme:    `this`,
			literal:   nil,
			line:      103,
		},
		{
			tokenType: Dot,
			lexeme:    `.`,
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
			tokenType: Equal,
			lexeme:    `=`,
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
			tokenType: Semicolon,
			lexeme:    `;`,
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
			lexeme:    `bread`,
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
			lexeme:    `bread`,
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
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      105,
		},
		{
			tokenType: Identifier,
			lexeme:    `serve`,
			literal:   nil,
			line:      107,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      107,
		},
		{
			tokenType: Identifier,
			lexeme:    `who`,
			literal:   nil,
			line:      107,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      107,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      107,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
			literal:   nil,
			line:      108,
		},
		{
			tokenType: String,
			lexeme:    `"Enjoy your "`,
			literal:   "Enjoy your ",
			line:      108,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      108,
		},
		{
			tokenType: This,
			lexeme:    `this`,
			literal:   nil,
			line:      108,
		},
		{
			tokenType: Dot,
			lexeme:    `.`,
			literal:   nil,
			line:      108,
		},
		{
			tokenType: Identifier,
			lexeme:    `meat`,
			literal:   nil,
			line:      108,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      108,
		},
		{
			tokenType: String,
			lexeme:    `" and "`,
			literal:   " and ",
			line:      108,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      108,
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
			lexeme:    `bread`,
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
			lexeme:    `", "`,
			literal:   ", ",
			line:      109,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      109,
		},
		{
			tokenType: Identifier,
			lexeme:    `who`,
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
			lexeme:    `"."`,
			literal:   ".",
			line:      109,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      109,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      110,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      113,
		},
		{
			tokenType: Class,
			lexeme:    `class`,
			literal:   nil,
			line:      115,
		},
		{
			tokenType: Identifier,
			lexeme:    `Brunch`,
			literal:   nil,
			line:      115,
		},
		{
			tokenType: Less,
			lexeme:    `<`,
			literal:   nil,
			line:      115,
		},
		{
			tokenType: Identifier,
			lexeme:    `Breakfast`,
			literal:   nil,
			line:      115,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      115,
		},
		{
			tokenType: Identifier,
			lexeme:    `init`,
			literal:   nil,
			line:      116,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      116,
		},
		{
			tokenType: Identifier,
			lexeme:    `meat`,
			literal:   nil,
			line:      116,
		},
		{
			tokenType: Comma,
			lexeme:    `,`,
			literal:   nil,
			line:      116,
		},
		{
			tokenType: Identifier,
			lexeme:    `bread`,
			literal:   nil,
			line:      116,
		},
		{
			tokenType: Comma,
			lexeme:    `,`,
			literal:   nil,
			line:      116,
		},
		{
			tokenType: Identifier,
			lexeme:    `drink`,
			literal:   nil,
			line:      116,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
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
			tokenType: Super,
			lexeme:    `super`,
			literal:   nil,
			line:      117,
		},
		{
			tokenType: Dot,
			lexeme:    `.`,
			literal:   nil,
			line:      117,
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
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      117,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      117,
		},
		{
			tokenType: This,
			lexeme:    `this`,
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
			lexeme:    `drink`,
			literal:   nil,
			line:      118,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      118,
		},
		{
			tokenType: Identifier,
			lexeme:    `drink`,
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
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      119,
		},
		{
			tokenType: Identifier,
			lexeme:    `drink`,
			literal:   nil,
			line:      121,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      121,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      121,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      121,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
			literal:   nil,
			line:      122,
		},
		{
			tokenType: String,
			lexeme:    `"How about a Bloody Mary?"`,
			literal:   "How about a Bloody Mary?",
			line:      122,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      122,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
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
			tokenType: EOF,
			lexeme:    ``,
			literal:   nil,
			line:      125,
		},
	}

	for i, tkn := range tokens {
		assert.Equal(t, expectedTokens[i], tkn)
	}
}
