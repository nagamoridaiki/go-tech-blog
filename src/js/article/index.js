'use strict';

// DOM Tree の構築が完了したら処理を開始します。
document.addEventListener('DOMContentLoaded', () => {
  // DOM API を利用して HTML 要素を取得します。
  const deleteBtns = document.querySelectorAll('.articles__item-delete');

  // CSRF トークンを取得します。
  const csrfToken = document.getElementsByName('csrf')[0].content;

  // 記事を削除する関数を定義します。
  const deleteArticle = id => {
    let statusCode;

    // Fetch API を利用して削除リクエストを送信します。
    fetch(`/${id}`, {
      method: 'DELETE',
      headers: { 'X-CSRF-Token': csrfToken }
    })
      .then(res => {
        statusCode = res.status;
        return res.json();
      })
      .then(data => {
        console.log(JSON.stringify(data));
        if (statusCode == 200) {
          // 削除に成功したら画面から記事の HTML 要素を削除します。
          document.querySelector(`.articles__item-${id}`).remove();
        }
      })
      .catch(err => console.error(err));
  };

  // 削除ボタンそれぞれに対してイベントリスナーを設定します。
  for (let elm of deleteBtns) {
    elm.addEventListener('click', event => {
      event.preventDefault();

      // 削除ボタンのカスタムデータ属性からIDを取得して引数に渡します。
      deleteArticle(elm.dataset.id);
    });
  }
});
