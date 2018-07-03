# 概要
rangeで処理するパターン。  

## ポイント

下部、

    for v := range valueChan {
        fmt.Println(v.value)
    }
    for v := range errChan { ←ここは届かないのでdaedlock
        fmt.Println(v.value)
    }

rangeの場合、上記のように実装すると、errChanのrange部分に到達することができないので、daedlockとなる。  
また、rangeの場合、明示的にclose(chan)を呼び出して上げないと、これもdaedlock。  
複数の戻り値がある場合、structでまとめて返却してあげればいい感じに返却できる。  
