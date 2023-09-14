package main

import (
	"os"

	"github.com/hanio89/golang-blockchain/cli"
)

func main() {
	defer os.Exit(0)

	cmd := cli.CommandLine{}
	cmd.Run()

	// http.HandleFunc("/run-command", func(w http.ResponseWriter, r *http.Request) {
	// 	// Đọc lệnh từ yêu cầu POST
	// 	// Thực hiện lệnh và trả về kết quả
	// 	command := r.FormValue("command")
	// 	cmd := exec.Command("cmd", "/c", command)
	// 	output, err := cmd.CombinedOutput()
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	w.Write(output)
	// })

	// http.ListenAndServe(":3000", nil)
}
