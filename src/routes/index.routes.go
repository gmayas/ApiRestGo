package routes

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Api Rest Go working 🙂👍🖐😎, ©2025 Copyright: GMayaS C:\\>_Desarrollo en Sistemas."))
}
