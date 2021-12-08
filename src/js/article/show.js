'use strict';

// DOM Tree の構築が完了したら処理を開始します。
document.addEventListener('DOMContentLoaded', function() {
  // DOM API を利用して HTML 要素を取得します。
  const elm = document.getElementById('article-body');

  // カスタムデータ属性から Markdown 形式のテキストを取得し、
  // Remarkable で HTML に変換して要素に追加します。
  elm.innerHTML = md.render(elm.dataset.markdown);
});
