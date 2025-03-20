char* mergeAlternately(char* word1, char* word2) {
    const int str1_len = strlen(word1);
    const int str2_len = strlen(word2);

    char* out = malloc((str1_len + str2_len + 1) * sizeof(char)); // + 1 for \0

    size_t i_1 = 0, i_2 = 0;
    size_t i = 0;

    while (i_1 < str1_len || i_2 < str2_len) {
        if (i_1 < str1_len) {
            out[i] = word1[i_1];
            i_1++;
            i++;
        }
        if (i_2 < str2_len) {
            out[i] = word2[i_2];
            i_2++;
            i++;
        }
    }

    out[i] = '\0';

    return out;
}