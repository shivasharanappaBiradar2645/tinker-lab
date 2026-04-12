def tokenize(text):
    tokens = []
    for char in text:
        if char.isdigit():
            tokens.append(("NUMBER",int(char)))
        elif char == "+":
            tokens.append(("PLUS",char))
        elif char == "-":
            tokens.append(("SUBSTRACT",char))
        elif char == "*":
            tokens.append(("MULTIPLY", char))
        elif char == "/":
            tokens.append(("DIVIDE",char))

    return tokens

text = input("Enter: ")
tokens= tokenize(text)

class Parser:
    def __init__(self,tokens):
        self.tokens = tokens
        self.pos = 0

    def eat(self, token_type):
        token = self.tokens[self.pos]
        
        if isinstance(token_type,list):
            valid = token[0] in token_type
        else:
            valid = token[0] == token_type


        if valid:
            self.pos += 1
            return token

        raise Exception("unexpected token")

    def parse(self):
        left = self.eat("NUMBER")
        op = self.eat(["PLUS","SUBSTRACT","MULTIPLY","DIVIDE"])
        right = self.eat("NUMBER")

        return (op[0], left[1], right[1])

parser = Parser(tokens)
ast = parser.parse()

def evaluate(ast):
    op,left,right = ast
    if op == "PLUS":
        return left + right
    elif op == "SUBSTRACT":
        return left - right
    elif op == "MULTIPLY":
        return left * right
    elif op == "DIVIDE":
        return left / right


print(evaluate(ast))
