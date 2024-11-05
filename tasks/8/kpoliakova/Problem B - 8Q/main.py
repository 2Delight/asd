import asyncio

global chess_board

async def placer(placement_list):
    global chess_board
    chess_board = [['*', '*', '*', '*', '*', '*', '*', '*'],
                   ['*', '*', '*', '*', '*', '*', '*', '*'],
                   ['*', '*', '*', '*', '*', '*', '*', '*'],
                   ['*', '*', '*', '*', '*', '*', '*', '*'],
                   ['*', '*', '*', '*', '*', '*', '*', '*'],
                   ['*', '*', '*', '*', '*', '*', '*', '*'],
                   ['*', '*', '*', '*', '*', '*', '*', '*'],
                   ['*', '*', '*', '*', '*', '*', '*', '*']]
    for i in range(8):       
        chess_board[i][int(placement_list[i])] = 'Q'
        
async def placement_checker(line):
    if await vertical_checker(line):
        await placer(line)
        if await diagonal_checker():
            return True
        else:
            return False
    else:
        return False

async def vertical_checker(line):
    if len(set(line)) != len(line):
        return False
    else:
        return True

async def diagonal_checker():
    global chess_board
    for i in range(8):
        index2check = chess_board[i].index('Q')
        for j in range(8):
            if (j-i) == 0:
                continue
            else:        
                #left        
                if (index2check-abs(j-i)) >= 0:
                    if chess_board[j][index2check-abs(j-i)] == 'Q':
                        return False
                #right    
                if (index2check+abs(j-i)) <= 7:
                    if chess_board[j][index2check+abs(j-i)] == 'Q':
                        return False
    return True


async def eight_queens(filename):
    result_placements = list()
    i = 0
    placements_file = open(filename + ".txt", "r")
    for line in placements_file:
        if await placement_checker(line):
            result_placements.append(line)
            for i in range(8):
                print(chess_board[i])
            print("\n")

    print("Amount of solutions:", len(result_placements))
    
loop = asyncio.get_event_loop()
loop.run_until_complete(eight_queens("all_possible_placements"))
loop.close()