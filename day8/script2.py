#!/usr/bin/python3

def getRowColStr(row:int,col:int):
    return f"{row}:{col}"

def strToRowCol(formatedString:str):
    strList = formatedString.split(':')
    row = int(strList[0])
    col = int(strList[1])
    return row,col

def getAntenaLocations(fileName:str):
    antenaLocations:dict[str,set[str]] = {}
    lines = []
    with open(fileName, 'r') as file:
        row = 0
        for line in file:
            line = line.strip()
            lines.append(line)
            for col in range(len(line)):
                val = line[col]
                if val != '.':
                    #We're at an antena
                    location = getRowColStr(row,col)
                    if val not in antenaLocations:
                        antenaLocations[val] = set([location])
                    else:
                        antenaLocations[val].add(location)
            row = row + 1
    return antenaLocations,lines

def isInWorld(row,col,world:list[list]):
    return row >= 0 and row < len(world) and col >= 0 and col < len(world[row])


def getDifs(row:int,col:int,locationRow:int,locationCol:int):
    rowDif = locationRow - row
    colDif = locationCol - col
    return rowDif,colDif        

def getAntiNodeLocations(coordinates:set[str],lines:list[list[str]],antiNodes:set):
    for coordinate in coordinates:
        for otherCoordinate in coordinates:
            if otherCoordinate != coordinate:
                row,col = strToRowCol(coordinate)
                otherRow,otherCol = strToRowCol(otherCoordinate)
                rowDif,colDif = getDifs(row,col,otherRow,otherCol)
                antiNodes.add(coordinate)
                antiNodes.add(otherCoordinate)
                matchingRow = otherRow + rowDif
                matchingCol = otherCol + colDif
                while isInWorld(matchingRow,matchingCol,lines):
                    antiNodes.add(getRowColStr(matchingRow,matchingCol))
                    matchingRow = matchingRow + rowDif
                    matchingCol = matchingCol + colDif

def countLocations(fileName:str):
    antenaLocations,lines = getAntenaLocations(fileName)
    antiNodes = set()
    for antenaCords in antenaLocations.values():
        getAntiNodeLocations(antenaCords,lines,antiNodes)
    return len(antiNodes)

def main():
    # print(countLocations('test4.txt'))
    print(countLocations('input.txt'))

if __name__ == "__main__":
    main()