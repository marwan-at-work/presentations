function makeSandwich() {
    const t1 = new Date;
    const pb = getPeanutButter();
    const j = getJelly();
    console.log(`Sandwich is ready! Putting together ${pb} and ${j}`);
    const t2 = new Date;
    const diff = (t2.getTime() - t1.getTime()) / 1000;
    console.log(`sandwich took ${diff} seconds to make`);
}

function getPeanutButter() {
    console.log('Butler 1 says: I am getting the peanut butter');
    simulateWork();
    return 'peanut butter';
}

function getJelly() {
    console.log('Butler 2 says: I am getting the jelly');
    simulateWork();
    return 'jelly';
}

function simulateWork() {
    const t1 = new Date;
    for (let t2 = new Date; ((t2.getTime() - t1.getTime()) / 1000) < 1; t2 = new Date) { }
}


makeSandwich();