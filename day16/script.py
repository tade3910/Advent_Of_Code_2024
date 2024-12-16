
import heapq

WALL = '#'
START = 'S'
END = 'E'

def getOpositeDirection(direction:str):
     if direction == 'E':
          return 'W'
     elif direction == 'W':
          return 'E'
     elif direction == 'N':
          return 'S'
     else:
          return 'N'

def exploreNeigbors(world:list[list[str]],currentEntry:tuple[int,int,str],currentDistance,queu:list[tuple[int,int,str]],distances:dict[tuple[int,int],int],prevs:dict[tuple[int,int],tuple[int,int]]):
        curentRow,currentCol,currentDirection = currentEntry
        displacements = [(-1,0,'N'),(1,0,'S'),(0,-1,'W'),(0,1,'E')]
        opositeDirection = getOpositeDirection(currentDirection)
        for displacement in displacements:
             neigborRow = curentRow + displacement[0]
             neigborCol = currentCol + displacement[1]
             if world[neigborRow][neigborCol] == WALL: 
                  continue
             neigborDirection = displacement[2]
             scoreChange = 1001
             if currentDirection == neigborDirection: scoreChange = 1
             elif currentDirection == opositeDirection: scoreChange = 2001
             newDistance = currentDistance + scoreChange
             oldDistance = distances.get((neigborRow,neigborCol),float('inf'))
             if newDistance < oldDistance:
                  neigbor = neigborRow,neigborCol
                  neigborEntry = neigborRow,neigborCol,neigborDirection
                  distances[neigbor] = newDistance
                  prevs[neigbor] = (curentRow,currentCol)
                  heapq.heappush(queu,(newDistance,neigborEntry))
    

def dikstras(world:list[list[str]],start:tuple[int,int]):
    queu = []
    startRow,startCol = start
    distances = {start:0}
    visited = set()
    prevs = {}
    heapq.heappush(queu,(0,(startRow,startCol,'E')))
    while queu:
        currentDistance,currentEntry = heapq.heappop(queu)
        curentRow,currentCol,_ = currentEntry
        curentLocation = curentRow,currentCol
        if curentLocation in visited: continue
        visited.add(curentLocation)
        exploreNeigbors(world,currentEntry,currentDistance,queu,distances,prevs)
    return distances,prevs

def debug(prevs:dict[tuple[int,int],tuple[int,int]],endLocation:tuple[int,int],startLocation:tuple[int,int],distances:dict[tuple[int,int]]):
     curLocation = endLocation
     while curLocation != startLocation:
          print(f"curLocation is {curLocation} and distance is {distances[curLocation]}")
          curLocation = prevs[curLocation]

def parseInput(fileName:str):
    with open(fileName, 'r') as file:
        world:list[list[str]] = []
        endLocation = (0,0)
        startLocation = (0,0)
        for i,line in enumerate(file):
            row = []
            for j,char in enumerate(line.strip()):
                row.append(char)
                if char == START:
                      startLocation = (i,j)
                elif char == END:
                     endLocation = (i,j)
            world.append(row)
        distances,prevs = dikstras(world,startLocation)
        # debug(prevs,endLocation,startLocation,distances)
        print(distances[endLocation])

def main():
    parseInput('test.txt')
    parseInput('test2.txt')
    parseInput('input.txt')

if __name__ == "__main__":
     main()



