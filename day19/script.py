
def parse(fileName:str):
    with open(fileName, 'r') as file:
        patternEnd = False
        toSearch = []
        made = {}
        towels:set[str] = set()
        for line in file:
            curLine = line.strip()
            if len(curLine) == 0:
                patternEnd = True
            elif patternEnd:
                toSearch.append(curLine)
            else:
                curLine = curLine.split(", ")
                for option in curLine:
                    towels.add(option)
                    made[option] = True
        return toSearch,towels,made

def canMake(toSearch: str,towels: set[str], made: dict[str,bool] ) -> bool:
    if len(toSearch) == 0: 
        return True
    if toSearch in made:
        return made[toSearch]
    for searchTo in range(0,len(toSearch)):
        curr = toSearch[0:searchTo + 1]
        if curr in towels and canMake(toSearch[searchTo + 1:],towels,made):
            return True
    
    made[toSearch] = False
    return False

def failure(toSearch: str, made: set[str], options: dict[str, list[str]]) -> bool:
    if toSearch in made:
        return True
    for option in options.get(toSearch[0], []):  # Match first character
        optionLen = len(option)
        if toSearch[:optionLen] == option:  # Check if option matches start of toSearch
            if failure(toSearch[optionLen:], made, options):  # Recur for remaining
                made.add(toSearch)  # Mark as buildable
                return True
    return False

def numMakes(toSearch: str,towels: set[str], made: dict[str,int] ) -> bool:
    if len(toSearch) == 0: 
        return 1
    if toSearch in made:
        return made[toSearch]
    numMade = 0
    for searchTo in range(0,len(toSearch)):
        curr = toSearch[0:searchTo + 1]
        if curr in towels:
            numMade += numMakes(toSearch[searchTo + 1:],towels,made)

    made[toSearch] = numMade
    return numMade

def countMade(fileName:str):
    toSearch,towels,made = parse(fileName)
    count = 0
    for search in toSearch:
        if canMake(search,towels,made):
            count += 1
    return count

def countNumMade(fileName:str):
    toSearch,towels,_= parse(fileName)
    count = 0
    made = {}
    for search in toSearch:
        count += numMakes(search,towels,made)
    return count

print(countMade("test.txt"))
print(countMade("input.txt"))
print(countNumMade("test.txt"))
print(countNumMade("input.txt"))

    
