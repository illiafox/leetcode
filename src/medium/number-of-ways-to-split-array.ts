// https://leetcode.com/problems/number-of-ways-to-split-array

// [DEPRECATED] Look for the rust solution
function waysToSplitArrayIneffective(nums: number[]): number {
    let suffixSum = new Array<number>(nums.length).fill(0);
    for (let i = suffixSum.length - 1; i >= 0; i--) {
        suffixSum[i] = nums[i] + (i < nums.length - 1 ? suffixSum[i + 1] : 0);
    }

    let prefixSum = new Array<number>(nums.length).fill(0);
    for (let i = 0; i < nums.length; i++) {
        prefixSum[i] = nums[i] + (i > 0 ? prefixSum[i - 1] : 0);
    }

    let validSplits = 0;
    for (let i = 0; i < nums.length - 1; i++) {
        if (prefixSum[i] >= suffixSum[i + 1]) validSplits++
    }

    return validSplits;
};

const cases: [number[], number][] = [
    [[10, 4, -8, 7], 2],
    [[2, 3, 1, 0], 2],
];

import {describe, expect, test} from "@jest/globals";

describe("'waysToSplitArrayIneffective'", () => {
    test.each(cases)(
        "given %p as arguments, returns %p",
        (nums, expected) => {
            const result = waysToSplitArrayIneffective(nums);
            expect(result).toEqual(expected);
        },
    );
});