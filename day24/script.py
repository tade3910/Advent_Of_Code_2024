def doOperation(gates:dict[str,int],firstWire:str,secondWire:str,operation:str):
    first = gates[firstWire]
    second = gates[secondWire]
    if operation == "AND":
        return first == 1 and second == 1
    elif operation == "OR":
        return first == 1 or second == 1
    elif operation == "XOR":
        return first != second
    raise Exception("I must've done something wrong")

def solveGates(gates:dict[str,int],inputs:list[tuple[str,str,str,str]],numZ:int):
    solved:set[int] = set()
    while len(solved) < len(inputs):
        for i, input in enumerate(inputs):
            if i not in solved:
                firstWire = input[0]
                secondWire = input[2]
                if firstWire in gates and secondWire in gates:
                    operation = input[1]
                    gate = input[3]
                    gates[gate] = doOperation(gates,firstWire,secondWire,operation)
                    solved.add(i)
    
    zBits = ""
    for i in range(numZ + 1):
        formatedI = f"{i}"
        if i < 10:
            formatedI = f"0{i}"
        if gates[f"z{formatedI}"]:
            zBits = f"1{zBits}"
        else:
            zBits = f"0{zBits}"
    print(zBits)
    decimalZ = int(zBits,2)
    print(decimalZ)

def parseInput(fileName:str):
    numZ = 0
    gates:dict[str,int] = {}
    inputs:list[tuple[str,str,str,str]] = []
    parseInput = False
    with open(fileName, 'r') as file:
        for line in file:
            line = line.strip()
            if len(line) == 0:
                parseInput = True
            elif parseInput:
                cur = line.split(" ")
                firstWire = cur[0]
                secondWire = cur[1]
                operation = cur[2]
                savedGate = cur[4]
                input = firstWire,secondWire,operation,savedGate
                for i in [0,1,4]:
                    if cur[i][0] == 'z':
                        val = int(cur[i][1:])
                        if val > numZ:
                            numZ = val
                inputs.append(input)
            else:
                cur = line.split(": ")
                gate = cur[0]
                val = int(cur[1])
                gates[gate] = val
    return gates,inputs,numZ

def partOne(fileName:str):
    gates,inputs,numZ = parseInput(fileName)
    solveGates(gates,inputs,numZ)

partOne("test.txt")
partOne("test2.txt")
partOne("input.txt")
    
        