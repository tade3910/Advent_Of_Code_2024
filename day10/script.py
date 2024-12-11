#!/usr/bin/python3
def findTrailHeads(input:list[list[str]]):
    trailHeadLocations = []
    for row,line in enumerate(input):
        for col,val in enumerate(line):
            if val == '0':
                trailHeadLocations.append((row,col))
    return trailHeadLocations

def isInWorld(row,col,world:list[list]):
    return row >= 0 and row < len(world) and col >= 0 and col < len(world[row])

def find9s(row:int,col:int,input:list[list[str]],found:set[str],paths:set[str],prev:int=-1,path=""):
    if not isInWorld(row,col,input):
        return 
    curVal = input[row][col]
    if curVal == '.':
        return 
    locString = f"{row}:{col}"
    path = path + locString
    curVal = int(curVal)
    if prev + 1 != curVal:
        return 
    if curVal == 9:
        found.add(locString)
        paths.add(path)
        return
    find9s(row,col -1,input,found,paths,curVal,path)
    find9s(row,col + 1,input,found,paths,curVal,path)
    find9s(row - 1,col,input,found,paths,curVal,path)
    find9s(row + 1,col,input,found,paths,curVal,path)
   

def parseInput(fileName:str):
     input = []
     with open(fileName, 'r') as file:
        for line in file:
            input.append(line.strip())
        trailHeads = findTrailHeads(input)
        part1Count = 0
        part2Count = 0
        for trailHead in trailHeads:
            row,col = trailHead
            found = set()
            paths = set()
            find9s(row,col,input,found,paths)
            part1Count = part1Count + len(found)
            part2Count = part2Count + len(paths)
        return part1Count,part2Count

def main():
    print(parseInput('input.txt'))
    print(parseInput('test6.txt'))

if __name__ == "__main__":
    main()

    
            
