void moveZeroes(int* nums, int numsSize) {
    int id = 0;

    for (int i = 0; i < numsSize; i++) {
        if (nums[i] != 0) {
            nums[id] = nums[i];
            id++;
        }
    }

    for (int i = id; i < numsSize; i++) {
       nums[i] = 0;
    }
}