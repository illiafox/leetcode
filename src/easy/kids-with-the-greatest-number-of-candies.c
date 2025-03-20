bool *kidsWithCandies(int *candies, int candiesSize, int extraCandies, int *returnSize) {
    int greatest = 0;
    for (int i = 0; i < candiesSize; i++) {
        if (candies[i] > greatest) {
            greatest = candies[i];
        }
    }

    bool *out = malloc((candiesSize) * sizeof(int));

    for (int i = 0; i < candiesSize; i++) {
        out[i] = candies[i] + extraCandies >= greatest;
    }
    *returnSize = candiesSize;

    return out;
}