function canPlaceFlowers(flowerbed: number[], n: number): boolean {
    for (let i = 0; i < flowerbed.length && n !== 0; i++) {
        if (
            flowerbed[i] === 0 &&
            flowerbed[i - 1] !== 1 &&
            flowerbed[i + 1] !== 1
        ) {
            n--
            i++
        }
    }

    return n === 0
};

import {describe, expect, test} from "@jest/globals";

const cases: [number[], number, boolean][] = [
    [[1, 0, 0, 0, 1], 1, true],
    [[1, 0, 0, 0, 1], 2, false],
];

describe("canPlaceFlowers", () => {
    test.each(cases)("given %p as arguments, returns %p", (flowerbed, n, expected) => {
        const result = canPlaceFlowers(flowerbed, n);
        expect(result).toEqual(expected);
    });
});
