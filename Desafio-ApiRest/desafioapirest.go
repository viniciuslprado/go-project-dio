package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
)

// Estruturas dos dados
type Person struct {
	ID         string      `json:"id,omitempty"`
	Firstname  string      `json:"firstname,omitempty"`
	Lastname   string      `json:"lastname,omitempty"`
	Age        int         `json:"age,omitempty"`
	Address    *Address    `json:"address,omitempty"`
	SocialMedia *SocialMedia `json:"social_media,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

type SocialMedia struct {
	Instagram string `json:"instagram,omitempty"`
}

// Estrutura auxiliar para carregar do JSON
type PeopleData struct {
	People []Person `json:"people"`
}

var people []Person

// Função para carregar os dados do JSON
func Load(filename string) error {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("erro ao abrir o arquivo JSON: %v", err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return fmt.Errorf("erro ao ler o arquivo JSON: %v", err)
	}

	var data PeopleData
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return fmt.Errorf("erro ao decodificar JSON: %v", err)
	}

	people = data.People
	fmt.Println("Dados carregados com sucesso do arquivo", filename)

	// Apenas para debug: Exibir os dados carregados no terminal
	for _, person := range people {
		fmt.Printf("Carregado: %+v\n", person)
	}

	return nil
}

// Rotas da API
func GetPeople(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recebida requisição GET para listar todos os contatos")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Printf("Recebida requisição GET para buscar contato com ID: %s\n", params["id"])
	w.Header().Set("Content-Type", "application/json")

	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Contato não encontrado"})
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	person.ID = params["id"]
	people = append(people, person)
	fmt.Printf("Criado novo contato: %+v\n", person)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(person)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Printf("Recebida requisição DELETE para contato com ID: %s\n", params["id"])
	w.Header().Set("Content-Type", "application/json")

	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			fmt.Printf("Contato com ID %s removido com sucesso\n", params["id"])
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Contato não encontrado"})
}

// Função principal para iniciar a API
func main() {
	// Carregar os dados do JSON
	err := Load("people.json")
	if err != nil {
		log.Fatalf("Erro ao carregar os dados: %v", err)
	}

	// Configuração do roteador
	router := mux.NewRouter()
	router.HandleFunc("/contato", GetPeople).Methods("GET")
	router.HandleFunc("/contato/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/contato/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/contato/{id}", DeletePerson).Methods("DELETE")

	// Inicia o servidor
	fmt.Println("Servidor rodando na porta 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
