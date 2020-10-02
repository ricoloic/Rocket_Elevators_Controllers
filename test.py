nb = [[10, 100, 1000], 20, 30, [40], [50]]

print(nb[0][2])

""" for i in nb:
    for j in i:
        print(j, "\n") """
    

for i in range(len(nb)):
    if i == 0:
        for j in nb[i]:
            print(j, "\n")

    else:
        print(nb[i], "\n")


for i in nb:
    print(i, "\n")

while len(nb) != 0:
    i = len(nb)
    print(i, nb[0])
    del nb[0]