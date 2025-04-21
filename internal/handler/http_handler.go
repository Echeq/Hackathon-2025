package handler

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

// HandleHTTP maneja las solicitudes HTTP.
func HandleHTTP(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/api/UserService/GetUser" && r.Method == "POST" {
        // Leer el cuerpo de la solicitud
        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
            http.Error(w, "Error al leer el cuerpo", http.StatusBadRequest)
            return
        }

        // Parsear el JSON
        var params map[string]interface{}
        err = json.Unmarshal(body, &params)
        if err != nil {
            http.Error(w, "Error al parsear JSON", http.StatusBadRequest)
            return
        }

        // Extraer el userID
        userID, ok := params["userID"].(float64)
        if !ok {
            http.Error(w, "userID no válido", http.StatusBadRequest)
            return
        }

        // Simular la respuesta del servicio
        result := fmt.Sprintf("Usuario con ID %d", int64(userID))

        // Devolver la respuesta en formato JSON
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{"result": result})
    } else {
        http.Error(w, "Ruta no encontrada", http.StatusNotFound)
    }
}