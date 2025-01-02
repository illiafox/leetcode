// https://leetcode.com/problems/count-vowel-strings-in-ranges/
import { long_cases } from "./count-vowel-strings-in-ranges.cases";

function vowelStrings(words: string[], queries: number[][]): number[] {
    const vowels = ["a", "e", "i", "o", "u"];

    let prefixSum = new Array<number>(words.length).fill(0);

    words.forEach((word, i) => {
        const match =
            vowels.includes(word[0]) && vowels.includes(word[word.length - 1]);

        prefixSum[i] = (match ? 1 : 0) + (i > 0 ? prefixSum[i - 1] : 0);
    });

    return queries.map(([from, to]) => {
        if (from === 0) {
            return prefixSum[to];
        }
        return prefixSum[to] - prefixSum[from - 1];
    });
}

import { describe, expect, test } from "@jest/globals";

const cases: [string[], number[][], number[]][] = [
    [
        ["aba", "bcb", "ece", "aa", "e"],
        [
            [0, 2],
            [1, 4],
            [1, 1],
        ],
        [2, 3, 0],
    ],
    [
        ["a", "e", "i"],
        [
            [0, 2],
            [0, 1],
            [2, 2],
        ],
        [3, 2, 1],
    ],
    [
        ["aba", "bcb", "ece", "aa", "e"],
        [
            [0, 2],
            [1, 4],
            [1, 1],
        ],
        [2, 3, 0],
    ],

    ...long_cases,
];
describe("'maxScore'", () => {
    test.each(cases)(
        "given %p as arguments, returns %p",
        (words, queries, expected) => {
            const result = vowelStrings(words, queries);
            expect(result).toEqual(expected);
        },
    );
});
