module.exports = {
  content: [
    './views/**/*.tmpl',  // 根据需要修改路径
    './public/js/**/*.js', // 根据需要修改路径
      './views/**/*.html',  // 根据需要修改路径
  ],
  darkMode: 'class', // 启用暗黑模式
  theme: {
    extend: {
      colors: {
        primary: '#1d4ed8', // 自定义颜色
      },
    },
  },
  plugins: [],
}

