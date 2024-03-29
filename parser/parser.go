package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken	token.Token
	peekToken	token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens, so curToken and peekToken are both set

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program = newProgramASTNode()
	
	advanceTokens()


	for (currentToken() != EOF_TOKEN) {
		statement = null

		if (currentToken() == LET_TOKEN){
			statement = pareseLetStatement()
		} else if (currentToken() == RETURN_TOKEN){
			statement = parseReturnStatement()
		} else if (currentToken() == IF_TOKEN) {
			statement = parseIfStatement()
		}

		if (statement != null) {
			program.Statements.push(statement)
		}

		advanceTokens()
	}

	return program
}

function parseLetStatement() {
	advanceTokens()

	identifier = parseIdentifier()

	advanceTokens()

	if currentToken() != EQUAL_TOKEN {
		parseError("no equal sign!")
		return null
	}

	advanceTokens()

	value = parseExpression()

	variableStatement = newVariableStatementASTNode()
	variableStatement.identifier = identifier
	variableStatement.value = value
	return variableStatement
}

function parseIdentifier() {
	identifier = newIdentifierASTNode()
	identifier.token = currentToken()
	return identifier
}

function parseExpression() {
	if (currentToken() == INTEGER_TOKEN) {
		if (nextToken() == PLUS_TOKEN) {
			return parseOperatorExpression()
		} else if (nextToken() == SEMICOLON_TOKEN) {
			return parseIntegerLiteral()
		}
	} else if (currentToken() == LEFT_PAREN) {
		return parseGroupedExpression()
	}
}

function parseOperatorExpression() {
	operatorExpression = newOperatorExpression()

	operatorExpression.left = parseIntegerLiteral()
	advanceTokens()
	operatorExpression.operator = currentToken()
	advanceTokens()
	operatorExpression.right = parseExpression()

	return operatorExpression
}
