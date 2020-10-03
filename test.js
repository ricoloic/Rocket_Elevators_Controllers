var nb = [[10, 100, 1000], 20, 30, 40, 50]

/* for (var i of nb) {
    for (var j of i) {
        console.log(j + "\n")
    }
} */

var i = 0;
while (i < nb.length) {
    i++
}

for (var i = 0; i < nb.length; i++) {
    if (i == 0) {
        for (var e of nb[i]) {
            console.log(e + "\n")
        }
    } else {
        console.log(nb[i] + "\n")
    }
}

for (var i of nb) {
    console.log(i + "\n")
}
/
var i;
var j = 0;
while (nb.length != 0) {
    i = nb.length
    nb[0].splice(0, 1)
    console.log(nb[j])
}