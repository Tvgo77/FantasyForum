import path from "path"
import {defineConfig} from "vite"

export default defineConfig({
    resolve: {
        alias: {
            '@': "/workspaces/FantasyForum/frontend/src"
        }
    }
})