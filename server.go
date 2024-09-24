package main

import (
	"log"
	"os/exec"
	"sync"

	"github.com/valyala/fasthttp"
)

var (
	mu           sync.Mutex
	cachedOutput []byte
	cached       bool
)

func runScript() ([]byte, error) {
	mu.Lock()
	defer mu.Unlock()

	if cached {
		return cachedOutput, nil
	}

	cmd := exec.Command("python3", "script.py")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	cachedOutput = output
	cached = true
	return cachedOutput, nil
}

func RunScript(ctx *fasthttp.RequestCtx) {
	output, err := runScript()
	if err != nil {
		log.Printf("スクリプトの実行に失敗しました: %v", err)
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.WriteString("Internal Server Error")
		return
	}
	ctx.Write(output)
}

func main() {
	// サーバー起動時にスクリプトを実行
	output, err := runScript()
	if err != nil {
		log.Fatalf("サーバー起動時にスクリプトの実行に失敗しました: %v", err)
	}
	log.Println(string(output))

	// HTTPサーバーの設定
	if err := fasthttp.ListenAndServe(":3000", RunScript); err != nil {
		log.Fatalf("サーバーの起動に失敗しました: %v", err)
	}
}
