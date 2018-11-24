## 建立自己的 image

1. 撰寫 Dockerfile
2. `From` 指定 base image
3. `RUN` 當在建立 image 時會執行的指令，通常用來安裝 package 及設定 image 相關事項
4. `CMD` 當 container 被 start 時會執行之指令