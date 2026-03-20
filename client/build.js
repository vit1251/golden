
import * as esbuild from 'esbuild';
import { copy } from 'esbuild-plugin-copy';
import path from 'node:path';

const isWatch = process.argv.includes('--watch');

const reloadScript = `new EventSource('/esbuild').addEventListener('change', () => location.reload());`;

const config = {
  entryPoints: ['src/main.tsx'],
  entryNames: 'public/[name]',
  bundle: true,
  minify: !isWatch,          // Минифицируем только для продакшена
  sourcemap: isWatch,        // Карты кода только в дев-режиме
  outdir: 'dist',
  banner: isWatch ? { js: reloadScript } : {},
  target: ['esnext'],        // Ориентируемся на современные браузеры
  plugins: [
    copy({
      assets: [
        { from: ['./src/assets/sprite.svg'], to: ['./sprite.svg'] },
        { from: ['./src/index.html'], to: ['./index.html'] }
      ],
    }),
  ],
};

if (isWatch) {
  const ctx = await esbuild.context(config);
  // Включаем слежку за файлами
  await ctx.watch();
  // Запускаем сервер
  const { host, port } = await ctx.serve({
    servedir: 'dist',  // Какую папку отдавать в браузер
    host: '127.0.0.1', // Желаемый интерфейс
    port: 8000         // Желаемый порт
  });
  console.log(`🚀 Server running at http://${host}:${port}`);
} else {
  await esbuild.build(config);
  console.log('✅ Production build finished!');
}
