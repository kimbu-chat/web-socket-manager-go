import type { Config } from "@jest/types"

const config: Config.InitialOptions = {
  preset: "ts-jest",
  testEnvironment: "node",
  verbose: true,
  automock: false,
  testTimeout: 10000,
  setupFiles: [
    'dotenv/config'
  ],
  setupFilesAfterEnv: ['<rootDir>/src/setup-tests.ts']
}
export default config
