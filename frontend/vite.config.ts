import { defineConfig } from 'vite';

export default defineConfig({
  server: {
    proxy: {
      '/api/v1/url': {
        target: 'http://localhost:8101',
        changeOrigin: true,
      }
    }
  }
});
