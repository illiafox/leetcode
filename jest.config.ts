// jest.config.js
import { Config } from "jest";
import { createDefaultPreset } from "ts-jest";

const config: Config = {
    ...createDefaultPreset(),
    testMatch: [
        "<rootDir>/easy/!(*.cases).ts",
        "<rootDir>/medium/!(*.cases).ts",
    ],
};

export default config;
