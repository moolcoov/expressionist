package lib

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	RPNTokenTypeOperand  = 1
	RPNTokenTypeOperator = 2
)

// RPNToken represents an abstract token object in RPN(Reverse Polish notation) which could either be an operator or operand.
type RPNToken struct {
	Type  int
	Value interface{}
}

// NewRPNOperandToken creates an instance of operand RPNToken with specified value.
func NewRPNOperandToken(val float64) *RPNToken {
	return NewRPNToken(val, RPNTokenTypeOperand)
}

// NewRPNOperatorToken creates an instance of operator RPNToken with specified value.
func NewRPNOperatorToken(val string) *RPNToken {
	return NewRPNToken(val, RPNTokenTypeOperator)
}

// NewRPNToken creates an instance of RPNToken with specified value and type.
func NewRPNToken(val interface{}, typ int) *RPNToken {
	return &RPNToken{Value: val, Type: typ}
}

// IsOperand determines whether a token is an operand with a specified value.
func (token *RPNToken) IsOperand(val float64) bool {
	return token.Type == RPNTokenTypeOperand && token.Value.(float64) == val
}

// IsOperator determines whether a token is an operator with a specified value.
func (token *RPNToken) IsOperator(val string) bool {
	return token.Type == RPNTokenTypeOperator && token.Value.(string) == val
}

// GetDescription returns a string that describes the token.
func (token *RPNToken) GetDescription() string {
	return fmt.Sprintf("(%d)%v", token.Type, token.Value)
}

var priorities = map[string]int{
	"+": 0,
	"-": 0,
	"*": 1,
	"/": 1,
}
var associativities = make(map[string]bool)

// parse parses an array of token strings and returns an array of abstract tokens
// using Shunting-yard algorithm.
func parse(tokens []string) ([]*RPNToken, error) {
	var ret []*RPNToken

	var operators []string
	for _, token := range tokens {
		operandToken := tryGetOperand(token)
		if operandToken != nil {
			ret = append(ret, operandToken)
		} else {
			// check parentheses
			if token == "(" {
				operators = append(operators, token)
			} else if token == ")" {
				foundLeftParenthesis := false
				// pop until "(" is fouund
				for len(operators) > 0 {
					oper := operators[len(operators)-1]
					operators = operators[:len(operators)-1]

					if oper == "(" {
						foundLeftParenthesis = true
						break
					} else {
						ret = append(ret, NewRPNOperatorToken(oper))
					}
				}
				if !foundLeftParenthesis {
					return nil, errors.New("mismatched parentheses found")
				}
			} else {
				// operator priority and associativity
				priority, ok := priorities[token]
				if !ok {
					return nil, fmt.Errorf("unknown operator: %v", token)
				}
				rightAssociative := associativities[token]

				for len(operators) > 0 {
					top := operators[len(operators)-1]

					if top == "(" {
						break
					}

					prevPriority := priorities[top]

					if (rightAssociative && priority < prevPriority) || (!rightAssociative && priority <= prevPriority) {
						// pop current operator
						operators = operators[:len(operators)-1]
						ret = append(ret, NewRPNOperatorToken(top))
					} else {
						break
					}
				} // end of for len(operators) > 0

				operators = append(operators, token)
			} // end of if token == "("
		} // end of if isOperand(token)
	} // end of for _, token := range tokens

	// process remaining operators
	for len(operators) > 0 {
		// pop
		operator := operators[len(operators)-1]
		operators = operators[:len(operators)-1]

		if operator == "(" {
			return nil, errors.New("mismatched parentheses found")
		}
		ret = append(ret, NewRPNOperatorToken(operator))
	}
	return ret, nil
}

// tryGetOperand determines whether a given string is an operand, if it is, an RPN operand token will be returned, otherwise nil.
func tryGetOperand(str string) *RPNToken {
	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return nil
	}
	return NewRPNOperandToken(value)
}

// Evaluate evaluates a list of RPNTokens and returns calculated value.
func evaluate(tokens []*RPNToken) (float64, error) {
	if tokens == nil {
		return 0, errors.New("tokens cannot be nil")
	}

	var stack []float64
	for _, token := range tokens {
		// push all operands to the stack
		if token.Type == RPNTokenTypeOperand {
			val := token.Value.(float64)
			stack = append(stack, val)
		} else {
			// execute current operator
			if len(stack) < 2 {
				return 0, errors.New("missing operand")
			}
			// pop 2 elements
			arg1, arg2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			val, err := evaluateOperator(token.Value.(string), arg1, arg2)
			if err != nil {
				return 0, err
			}
			// push result back to stack
			stack = append(stack, val)
		}
	}
	if len(stack) != 1 {
		return 0, errors.New("stack corrupted")
	}
	return stack[len(stack)-1], nil
}

// executes an operator
func evaluateOperator(oper string, a, b float64) (float64, error) {
	switch oper {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, errors.New("Unknown operator: " + oper)
	}
}
