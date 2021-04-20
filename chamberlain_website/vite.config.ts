import path from 'path'

module.exports = {
    // 导入别名
    alias: {
        '/@/': path.resolve(__dirname, './src'),
        '/@views/': path.resolve(__dirname, './src/views'),
        '/@components/': path.resolve(__dirname, './src/components'),
        '/@api/': path.resolve(__dirname, './src/api'),
    },
    // 配置Dep优化行为
    optimizeDeps: {
        include: ["lodash"]
    },
    // 为开发服务器配置自定义代理规则。
    proxy: {
        '/rest': {
            target: 'http://localhost:8080',
            changeOrigin: true,
        }
    }
}