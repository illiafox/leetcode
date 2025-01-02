// jest.config.js
import { Config } from "jest";
import { createDefaultPreset } from "ts-jest";

const config: Config = {
    ...createDefaultPreset(),
    testMatch: [
        "<rootDir>/src/easy/!(*.cases).ts",
        "<rootDir>/src/medium/!(*.cases).ts",
    ],
};

export default config;
