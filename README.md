# 396gi Database

- 過去のライブのセトリ等の情報を管理するデータベース（まだ開発中）
- 将来的にはセトリ考案とか特設ページ制作とかで活用していけると良いなぁ…という展望

## 環境構築

```
$ docker-compose build
$ docker-compose up
```

参考記事：

- [【React&Golang&MySql】Docker で開発環境構築 - Qiita](https://qiita.com/takuya911/items/2447c97525d4c48b72a2)

<details>
<summary>やったこと</summary>

### Frontend

```
$ docker-compose run --rm frontend create-react-app react-app
```

参考記事：

- [reactjs — Node / Docker のビルド時に uid / gid を取得できませんでした](https://www.it-mure.jp.net/ja/reactjs/node-docker%E3%81%AE%E3%83%93%E3%83%AB%E3%83%89%E6%99%82%E3%81%ABuid-gid%E3%82%92%E5%8F%96%E5%BE%97%E3%81%A7%E3%81%8D%E3%81%BE%E3%81%9B%E3%82%93%E3%81%A7%E3%81%97%E3%81%9F/807482822/)

### Backend

参考記事：

- [Docker で Go の開発環境を構築する - Qiita](https://qiita.com/uji_/items/8c9eda89526abe0ba900)

</details>
