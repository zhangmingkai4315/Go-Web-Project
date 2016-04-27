package main

import (
	"encoding/json"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	private_key = "secret/app.rsa"
	public_key  = "secret/app.rsa.pub"
)

var (
	verifyKey, signKey []byte
)

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func init() {
	var err error
	signKey, err = ioutil.ReadFile(private_key)
	if err != nil {
		log.Fatal("Error reading private_key file")
		return
	}
	verifyKey, err = ioutil.ReadFile(public_key)
	if err != nil {
		log.Fatal("Error reading public_key file")
		return
	}
}
func authHandler(w http.ResponseWriter, r *http.Request) {
	token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintln(w, "Token Expired, got a now one")
				return
			default:
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, "Error while parsing token")
				log.Println("Error validation:%v", vErr.Errors)
				return
			}
		default:
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Error while parsing token:%v", err)
			log.Println("Error validation")
			return
		}
	}

	if token.Valid {
		response := Response{"Authorized to the system"}
		jsonResponse(response, w)
	} else {
		response := Response{"Invalid token"}
		jsonResponse(response, w)
	}
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	// body, err := ioutil.ReadAll(r.Body)
	// fmt.Println(body)
	// // err := json.NewDecoder(r.Body).Decode(&user)
	// post ==> { "username": "mike", "password": "password" }
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error in request body")
		return
	}
	if user.UserName != "mike" && user.Password != "password" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Error in user info")
		return
	}
	t := jwt.New(jwt.GetSigningMethod("RS256"))
	t.Claims["iss"] = "admin"
	t.Claims["Userinfo"] = struct {
		Name string
		Role string
	}{user.UserName, "Member"}
	t.Claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	tokenString, err := t.SignedString(signKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error in token gen")
		return
	}
	response := Token{tokenString}
	jsonResponse(response, w)
}

type Response struct {
	Text string `json:"token"`
}
type Token struct {
	Token string `json:"token"`
}

func jsonResponse(response interface{}, w http.ResponseWriter) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/auth", authHandler).Methods("POST")
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("Listening")
	server.ListenAndServe()
}
