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
			lexeme:    `a`,
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
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      12,
		},
		{
			tokenType: While,
			lexeme:    `while`,
			literal:   nil,
			line:      13,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      13,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      13,
		},
		{
			tokenType: Less,
			lexeme:    `<`,
			literal:   nil,
			line:      13,
		},
		{
			tokenType: Number,
			lexeme:    `10`,
			literal:   10.0,
			line:      13,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      13,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      13,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
			literal:   nil,
			line:      14,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      14,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      14,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
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
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      15,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      15,
		},
		{
			tokenType: Number,
			lexeme:    `1`,
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
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      16,
		},
		{
			tokenType: For,
			lexeme:    `for`,
			literal:   nil,
			line:      18,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      18,
		},
		{
			tokenType: Var,
			lexeme:    `var`,
			literal:   nil,
			line:      18,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
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
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      18,
		},
		{
			tokenType: Less,
			lexeme:    `<`,
			literal:   nil,
			line:      18,
		},
		{
			tokenType: Number,
			lexeme:    `10`,
			literal:   10.0,
			line:      18,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      18,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
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
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      18,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
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
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      18,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      18,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
			literal:   nil,
			line:      19,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      19,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      19,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      20,
		},
		{
			tokenType: Fun,
			lexeme:    `fun`,
			literal:   nil,
			line:      22,
		},
		{
			tokenType: Identifier,
			lexeme:    `printSum`,
			literal:   nil,
			line:      22,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      22,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      22,
		},
		{
			tokenType: Comma,
			lexeme:    `,`,
			literal:   nil,
			line:      22,
		},
		{
			tokenType: Identifier,
			lexeme:    `b`,
			literal:   nil,
			line:      22,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      22,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      22,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
			literal:   nil,
			line:      23,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      23,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      23,
		},
		{
			tokenType: Identifier,
			lexeme:    `b`,
			literal:   nil,
			line:      23,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      23,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      24,
		},
		{
			tokenType: Fun,
			lexeme:    `fun`,
			literal:   nil,
			line:      26,
		},
		{
			tokenType: Identifier,
			lexeme:    `returnSum`,
			literal:   nil,
			line:      26,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      26,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      26,
		},
		{
			tokenType: Comma,
			lexeme:    `,`,
			literal:   nil,
			line:      26,
		},
		{
			tokenType: Identifier,
			lexeme:    `b`,
			literal:   nil,
			line:      26,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      26,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      26,
		},
		{
			tokenType: Return,
			lexeme:    `return`,
			literal:   nil,
			line:      27,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      27,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      27,
		},
		{
			tokenType: Identifier,
			lexeme:    `b`,
			literal:   nil,
			line:      27,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      27,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      28,
		},
		{
			tokenType: Fun,
			lexeme:    `fun`,
			literal:   nil,
			line:      30,
		},
		{
			tokenType: Identifier,
			lexeme:    `addPair`,
			literal:   nil,
			line:      30,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      30,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      30,
		},
		{
			tokenType: Comma,
			lexeme:    `,`,
			literal:   nil,
			line:      30,
		},
		{
			tokenType: Identifier,
			lexeme:    `b`,
			literal:   nil,
			line:      30,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      30,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      30,
		},
		{
			tokenType: Return,
			lexeme:    `return`,
			literal:   nil,
			line:      31,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      31,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      31,
		},
		{
			tokenType: Identifier,
			lexeme:    `b`,
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
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      32,
		},
		{
			tokenType: Fun,
			lexeme:    `fun`,
			literal:   nil,
			line:      34,
		},
		{
			tokenType: Identifier,
			lexeme:    `identity`,
			literal:   nil,
			line:      34,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      34,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
			literal:   nil,
			line:      34,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      34,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      34,
		},
		{
			tokenType: Return,
			lexeme:    `return`,
			literal:   nil,
			line:      35,
		},
		{
			tokenType: Identifier,
			lexeme:    `a`,
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
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      36,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
			literal:   nil,
			line:      38,
		},
		{
			tokenType: Identifier,
			lexeme:    `identity`,
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
			tokenType: Identifier,
			lexeme:    `addPair`,
			literal:   nil,
			line:      38,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
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
			tokenType: Comma,
			lexeme:    `,`,
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
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      38,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      38,
		},
		{
			tokenType: Fun,
			lexeme:    `fun`,
			literal:   nil,
			line:      40,
		},
		{
			tokenType: Identifier,
			lexeme:    `outerFunction`,
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
			tokenType: Fun,
			lexeme:    `fun`,
			literal:   nil,
			line:      41,
		},
		{
			tokenType: Identifier,
			lexeme:    `localFunction`,
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
			lexeme:    `"I'm local!"`,
			literal:   "I'm local!",
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
			tokenType: Identifier,
			lexeme:    `localFunction`,
			literal:   nil,
			line:      45,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      45,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      45,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      45,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      46,
		},
		{
			tokenType: Class,
			lexeme:    `class`,
			literal:   nil,
			line:      48,
		},
		{
			tokenType: Identifier,
			lexeme:    `Breakfast`,
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
			tokenType: Identifier,
			lexeme:    `cook`,
			literal:   nil,
			line:      49,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      49,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      49,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      49,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
			literal:   nil,
			line:      50,
		},
		{
			tokenType: String,
			lexeme:    `"Eggs a-fryin'!"`,
			literal:   "Eggs a-fryin'!",
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
			tokenType: Identifier,
			lexeme:    `serve`,
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
			tokenType: Identifier,
			lexeme:    `who`,
			literal:   nil,
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
			tokenType: String,
			lexeme:    `"Enjoy your breakfast, "`,
			literal:   "Enjoy your breakfast, ",
			line:      54,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      54,
		},
		{
			tokenType: Identifier,
			lexeme:    `who`,
			literal:   nil,
			line:      54,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      54,
		},
		{
			tokenType: String,
			lexeme:    `"."`,
			literal:   ".",
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
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      56,
		},
		{
			tokenType: Var,
			lexeme:    `var`,
			literal:   nil,
			line:      59,
		},
		{
			tokenType: Identifier,
			lexeme:    `someVariable`,
			literal:   nil,
			line:      59,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      59,
		},
		{
			tokenType: Identifier,
			lexeme:    `Breakfast`,
			literal:   nil,
			line:      59,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      59,
		},
		{
			tokenType: Identifier,
			lexeme:    `someFunction`,
			literal:   nil,
			line:      62,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      62,
		},
		{
			tokenType: Identifier,
			lexeme:    `Breakfast`,
			literal:   nil,
			line:      62,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
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
			tokenType: Var,
			lexeme:    `var`,
			literal:   nil,
			line:      64,
		},
		{
			tokenType: Identifier,
			lexeme:    `breakfast`,
			literal:   nil,
			line:      64,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      64,
		},
		{
			tokenType: Identifier,
			lexeme:    `Breakfast`,
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
			tokenType: RightParen,
			lexeme:    `)`,
			literal:   nil,
			line:      64,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      64,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
			literal:   nil,
			line:      65,
		},
		{
			tokenType: Identifier,
			lexeme:    `breakfast`,
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
			tokenType: Class,
			lexeme:    `class`,
			literal:   nil,
			line:      67,
		},
		{
			tokenType: Identifier,
			lexeme:    `Breakfast`,
			literal:   nil,
			line:      67,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      67,
		},
		{
			tokenType: Identifier,
			lexeme:    `init`,
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
			lexeme:    `meat`,
			literal:   nil,
			line:      68,
		},
		{
			tokenType: Comma,
			lexeme:    `,`,
			literal:   nil,
			line:      68,
		},
		{
			tokenType: Identifier,
			lexeme:    `bread`,
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
			tokenType: This,
			lexeme:    `this`,
			literal:   nil,
			line:      69,
		},
		{
			tokenType: Dot,
			lexeme:    `.`,
			literal:   nil,
			line:      69,
		},
		{
			tokenType: Identifier,
			lexeme:    `meat`,
			literal:   nil,
			line:      69,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      69,
		},
		{
			tokenType: Identifier,
			lexeme:    `meat`,
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
			tokenType: This,
			lexeme:    `this`,
			literal:   nil,
			line:      70,
		},
		{
			tokenType: Dot,
			lexeme:    `.`,
			literal:   nil,
			line:      70,
		},
		{
			tokenType: Identifier,
			lexeme:    `bread`,
			literal:   nil,
			line:      70,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      70,
		},
		{
			tokenType: Identifier,
			lexeme:    `bread`,
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
			tokenType: Identifier,
			lexeme:    `serve`,
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
			lexeme:    `who`,
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
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      73,
		},
		{
			tokenType: Print,
			lexeme:    `print`,
			literal:   nil,
			line:      74,
		},
		{
			tokenType: String,
			lexeme:    `"Enjoy your "`,
			literal:   "Enjoy your ",
			line:      74,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      74,
		},
		{
			tokenType: This,
			lexeme:    `this`,
			literal:   nil,
			line:      74,
		},
		{
			tokenType: Dot,
			lexeme:    `.`,
			literal:   nil,
			line:      74,
		},
		{
			tokenType: Identifier,
			lexeme:    `meat`,
			literal:   nil,
			line:      74,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      74,
		},
		{
			tokenType: String,
			lexeme:    `" and "`,
			literal:   " and ",
			line:      74,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      74,
		},
		{
			tokenType: This,
			lexeme:    `this`,
			literal:   nil,
			line:      75,
		},
		{
			tokenType: Dot,
			lexeme:    `.`,
			literal:   nil,
			line:      75,
		},
		{
			tokenType: Identifier,
			lexeme:    `bread`,
			literal:   nil,
			line:      75,
		},

		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      75,
		},
		{
			tokenType: String,
			lexeme:    `", "`,
			literal:   ", ",
			line:      75,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      75,
		},
		{
			tokenType: Identifier,
			lexeme:    `who`,
			literal:   nil,
			line:      75,
		},
		{
			tokenType: Plus,
			lexeme:    `+`,
			literal:   nil,
			line:      75,
		},
		{
			tokenType: String,
			lexeme:    `"."`,
			literal:   ".",
			line:      75,
		},
		{
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      75,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      76,
		},
		{
			tokenType: RightBrace,
			lexeme:    `}`,
			literal:   nil,
			line:      79,
		},
		{
			tokenType: Class,
			lexeme:    `class`,
			literal:   nil,
			line:      81,
		},
		{
			tokenType: Identifier,
			lexeme:    `Brunch`,
			literal:   nil,
			line:      81,
		},
		{
			tokenType: Less,
			lexeme:    `<`,
			literal:   nil,
			line:      81,
		},
		{
			tokenType: Identifier,
			lexeme:    `Breakfast`,
			literal:   nil,
			line:      81,
		},
		{
			tokenType: LeftBrace,
			lexeme:    `{`,
			literal:   nil,
			line:      81,
		},
		{
			tokenType: Identifier,
			lexeme:    `init`,
			literal:   nil,
			line:      82,
		},
		{
			tokenType: LeftParen,
			lexeme:    `(`,
			literal:   nil,
			line:      82,
		},
		{
			tokenType: Identifier,
			lexeme:    `meat`,
			literal:   nil,
			line:      82,
		},
		{
			tokenType: Comma,
			lexeme:    `,`,
			literal:   nil,
			line:      82,
		},
		{
			tokenType: Identifier,
			lexeme:    `bread`,
			literal:   nil,
			line:      82,
		},
		{
			tokenType: Comma,
			lexeme:    `,`,
			literal:   nil,
			line:      82,
		},
		{
			tokenType: Identifier,
			lexeme:    `drink`,
			literal:   nil,
			line:      82,
		},
		{
			tokenType: RightParen,
			lexeme:    `)`,
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
			tokenType: Super,
			lexeme:    `super`,
			literal:   nil,
			line:      83,
		},
		{
			tokenType: Dot,
			lexeme:    `.`,
			literal:   nil,
			line:      83,
		},
		{
			tokenType: Identifier,
			lexeme:    `init`,
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
			tokenType: Identifier,
			lexeme:    `meat`,
			literal:   nil,
			line:      83,
		},
		{
			tokenType: Comma,
			lexeme:    `,`,
			literal:   nil,
			line:      83,
		},
		{
			tokenType: Identifier,
			lexeme:    `bread`,
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
			tokenType: Semicolon,
			lexeme:    `;`,
			literal:   nil,
			line:      83,
		},
		{
			tokenType: This,
			lexeme:    `this`,
			literal:   nil,
			line:      84,
		},
		{
			tokenType: Dot,
			lexeme:    `.`,
			literal:   nil,
			line:      84,
		},
		{
			tokenType: Identifier,
			lexeme:    `drink`,
			literal:   nil,
			line:      84,
		},
		{
			tokenType: Equal,
			lexeme:    `=`,
			literal:   nil,
			line:      84,
		},
		{
			tokenType: Identifier,
			lexeme:    `drink`,
			literal:   nil,
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
			lexeme:    `drink`,
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
			lexeme:    `"How about a Bloody Mary?"`,
			literal:   "How about a Bloody Mary?",
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
			tokenType: Eof,
			lexeme:    ``,
			literal:   nil,
			line:      91,
		},
	}

	for i, tkn := range tokens {
		assert.Equal(t, expectedTokens[i], tkn)
	}
}
