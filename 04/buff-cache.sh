#!/binbash

echo "ファイル作成前のシステム全体のメモリ使用量を表示"
free -m

echo "1GBのファイルを新規作成して、1GBのページキャッシュ領域を獲得"
dd if=/dev/zero of=testfile bs=1M count=1K

echo "ページキャッシュ獲得後のシステム全体のメモリ使用量を表示"
free -m

echo "ファイル削除後のシステム全体のメモリ使用量を表示"
rm testfile
free -m