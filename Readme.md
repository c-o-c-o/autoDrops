# CDrops
プロファイルに従ってコマンドライン形式でaviutl上にファイルをドロップするアプリです  
ごちゃまぜドロップスAPIを使用しているため、できるだけ新しいバージョンを使用してください  

# 使い方
autoDrop.exe と同じ階層に [cdrop.exe](https://github.com/c-o-c-o/cdrops) を置いてください
```

autoDrop.exe [option] [filelist]
option
  -t テキストファイルのパス、プロファイルで Target: Text を利用しない場合は他のファイルでも可
  -p Profileのパス [Profile.yml]
filelist
  スペース区切りのドロップするファイルパスのリスト
```

```
example
  autoDrop.exe -t セリフ.txt セリフ.wav セリフ.txt  
```

# Licence
This software is released under the MIT License, see LICENSE.  

ごちゃまぜドロップス  
https://github.com/oov/aviutl_gcmzdrops