
import * as esbuild from 'esbuild';

const isProduction = process.argv.includes('--prod');

// Общие настройки для dev и prod
const config = {
  entryPoints: ['src/index.ts'],
  bundle: true,
  minify: isProduction,
  sourcemap: !isProduction,
  outfile: 'dist/bundle.js',
  target: ['es2022'],
};

if (isProduction) {
  // Продакшн сборка
  await esbuild.build(config);
  console.log('⚡ Сборка успешно завершена для production!');
} else {
  // Режим разработки (Сервер + Watch)
  const ctx = await esbuild.context(config);

  // Включаем отслеживание изменений файлов
  await ctx.watch();

  // Запускаем локальный веб-сервер
  const server = await ctx.serve({
    servedir: '.', // Серверим корневую папку, где лежит index.html
    port: 3000
  });

  console.log(`🚀 Сервер запущен на http://localhost:${server.port}`);
}
