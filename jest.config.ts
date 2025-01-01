// jest.config.js
import { Config } from "jest";
import { createDefaultPreset } from "ts-jest";

const config: Config = {
    ...createDefaultPreset(),
    testMatch: ["<rootDir>/easy/*.ts"],
};

export default config;
