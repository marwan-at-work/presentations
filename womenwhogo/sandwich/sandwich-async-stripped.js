function main() {
    const pbPromise = getPeanutButter();
    const jPromise = getJelly();
    Promise.all([pbPromise, jPromise]).then(ingredients => {
        makeSandwich(ingredients[0], ingredients[1]);
    });
}

function makeSandwich(ingredient1, ingredient2) {
    console.log(`Sandwich is ready! Putting together ${ingredient1} and ${ingredient2}`);
}

function getPeanutButter() {
    console.log('Butler 1 says: I am getting the peanut butter');
    return new Promise((resolve, reject) => {
        simulateWork();
        resolve('peanut butter');
    });
}

function getJelly() {
    console.log('Butler 2 says: I am getting the jelly');
    return new Promise((resolve, reject) => {
        simulateWork();
        resolve('jelly');
    });
}

function simulateWork() {
    const t1 = new Date;
    for (let t2 = new Date; ((t2.getTime() - t1.getTime()) / 1000) < 1; t2 = new Date) { }
}


main();