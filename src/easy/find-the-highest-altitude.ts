function largestAltitude(gains: number[]): number {
    let altitude = 0
    let max = 0

    for (const gain of gains) {
        altitude += gain
        if (altitude > max) max = altitude
    }

    return max
};

import {describe, expect, test} from "@jest/globals";

const cases: [number[], number][] = [
    [[-5, 1, 5, 0, -7], 1],
    [[-4, -3, -2, -1, 4, 3, 2], 0],
];


describe("'largestAltitude'", () => {
    test.each(cases)(
        "given %p as arguments, returns %p",
        (gains, expected) => {
            const result = largestAltitude(gains);
            expect(result).toEqual(expected);
        },
    );
});
