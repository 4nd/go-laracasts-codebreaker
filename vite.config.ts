import {defineConfig, Plugin, ResolvedConfig} from "vite";
import * as fs from "node:fs";
import fullReload from 'vite-plugin-full-reload'

function hotFilePlugin(): Plugin[] {
    let exitHandlersBound = false
    let resolvedConfig: ResolvedConfig
    return [{
        name: 'hotFile-on-serve',
        apply: 'serve',
        configResolved(config) {
            resolvedConfig = config
        },
        configureServer(server) {
            let hotFile = resolvedConfig.root + '/vite-hot';
            server.httpServer?.once('listening', () => {
                const address = server.httpServer?.address()
                if (typeof address === 'object') {

                    fs.writeFileSync(hotFile, "vite is hot")
                    if (!exitHandlersBound) {
                        const clean = () => {
                            if (fs.existsSync(hotFile)) {
                                fs.rmSync(hotFile)
                            }
                        }

                        process.on('exit', clean)
                        process.on('SIGINT', () => process.exit())
                        process.on('SIGTERM', () => process.exit())
                        process.on('SIGHUP', () => process.exit())

                        exitHandlersBound = true
                    }
                }
            })
        }
    }];
}

export default defineConfig({
    build: {
        manifest: true,
        rollupOptions: {
            input: ['assets/src/app.js'],
        },
        copyPublicDir: false,
        outDir: 'assets/dist',
        assetsDir: '',
    },
    plugins: [
        hotFilePlugin(),
        fullReload(['views/**/*'], { delay: 1000 })
    ]
})