package main

// func CreateUser(w http.ResponseWriter, r *http.Request) {
// 	var usr *User
// 	err := readJSON(r, &usr)
// 	if err != nil {
// 		errorJSON(w, err, http.StatusBadRequest)
// 		return
// 	}

// 	err = validateStruct(*usr)
// 	if err != nil {
// 		errorJSON(w, err, http.StatusBadRequest)
// 		return
// 	}

// 	writeJSON(w, http.StatusCreated, "Usuário criado com sucesso!")
// }
