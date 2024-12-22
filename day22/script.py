def parseInput(fileName:str):
    secretNumbers = []
    with open(fileName, 'r') as file:
        for line in file:
            secretNumbers.append(int(line.strip()))
    return secretNumbers

def mix(value:int,secretNumber:int):
    return value ^ secretNumber

def prune(secretNumber:int):
    return secretNumber % 16777216

def stepOne(secretNumber:int):
    result = secretNumber * 64
    result = mix(result,secretNumber)
    return prune(result)

def stepTwo(secretNumber:int):
    result = secretNumber // 32
    result = mix(result,secretNumber)
    return prune(result)

def stepThree(secretNumber:int):
    result = secretNumber * 2048
    result = mix(result,secretNumber)
    return prune(result)

def generateNextSecret(secretNumber:int):
    secretNumber = stepOne(secretNumber)
    secretNumber = stepTwo(secretNumber)
    secretNumber = stepThree(secretNumber)
    return secretNumber

def generateXSecretNumbers(secretNumber:int,x:int):
    for _ in range(x):
        secretNumber = generateNextSecret(secretNumber)
    return secretNumber

def simulate(fileName:int,x:int):
    secretNumbers = parseInput(fileName)
    sum = 0
    for secretNumber in secretNumbers:
        sum += generateXSecretNumbers(secretNumber,x)
    print(f"sum is {sum}")

simulate("test.txt",2000)
simulate("input.txt",2000)