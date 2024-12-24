def getTriangles(key:str,networkMap:dict[str,set[str]]):
    curList = list(networkMap[key])
    triangles = []
    for i in range(len(curList) - 1):
        curAdd = curList[i]
        curAddSet = networkMap[curAdd]
        for j in range(i + 1,len(curList)):
            val = curList[j]
            if val in curAddSet:
                triangle = tuple(sorted([key, curAdd, val]))
                triangles.append(triangle)
    return triangles

def getThreeComputerParties(fileName:str):
    networkMap = parseInput(fileName)
    threeComputerParties = []
    seen = set()
    for key in networkMap.keys():
        if key[0] == "t":
            triangles = getTriangles(key,networkMap)
            for triangle in triangles:
                if triangle not in seen:
                    threeComputerParties.append(triangle)
                    seen.add(triangle)
    # print(threeComputerParties)
    print(len(threeComputerParties))
    return threeComputerParties

def parseInput(fileName:str):
    networkMap:dict[str,set[str]] = {}
    with open(fileName, 'r') as file:
        for line in file:
            computers = line.strip().split('-')
            first = computers[0]
            second = computers[1]
            curFirst = networkMap.get(first,set())
            curFirst.add(second)
            networkMap[first] = curFirst
            curSecond = networkMap.get(second,set())
            curSecond.add(first)
            networkMap[second] = curSecond
    return networkMap

getThreeComputerParties("test.txt")
getThreeComputerParties("input.txt")