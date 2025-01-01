// https://leetcode.com/problems/maximum-score-after-splitting-a-string/description

export function maxScore(s: string): number {
    let ones = 0;
    for (let i = 0; i < s.length; i++) {
        if (s[i] === "1") ones++;
    }

    let zeros = 0;
    let max = -1;

    for (let i = 0; i < s.length - 1; i++) {
        if (s[i] === "0") {
            zeros++;
        } else {
            ones--;
        }

        const sum = zeros + ones;
        if (sum > max) max = sum;
    }

    return max;
}

import { describe, expect, test } from "@jest/globals";

const cases: [string, number][] = [
    ["011101", 5],
    ["00111", 5],
    ["1111", 3],
];

describe("'maxScore'", () => {
    test.each(cases)("given %p as arguments, returns %p", (input, expected) => {
        const result = maxScore(input);
        expect(result).toEqual(expected);
    });
});
