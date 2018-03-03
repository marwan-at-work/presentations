function makeSandwich() {
    const t1 = new Date;
    const pbPromise = getPeanutButter();
    const jellyPromise = getJelly();
    Promise.all([pbPromise, jellyPromise]).then(ingredients => {
        console.log(
            `Sandwich is ready! Putting together ${ingredients[0]} and ${ingredients[1]}.`
        );
        const t2 = new Date;
        const diff = (t2.getTime() - t1.getTime()) / 1000;
        console.log(`The sandwich took ${diff} seconds to make`);
    });
}

function getPeanutButter() {
    console.log('Butler 1 says: I am getting the peanut butter');
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            simulateWork();
            resolve('peanut butter');
        }, 0);
    });
}

function getJelly() {
    console.log('Butler 2 says: I am getting the jelly');
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            simulateWork();
            resolve('jelly');
        }, 0);
    });
}

function simulateWork() {
    const t1 = new Date;
    for (let t2 = new Date; ((t2.getTime() - t1.getTime()) / 1000) < 1; t2 = new Date) { }
}

makeSandwich();