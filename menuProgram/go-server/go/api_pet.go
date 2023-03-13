/*
 * Swagger Petstore - OpenAPI 3.0
 *
 * This is a sample Pet Store Server based on the OpenAPI 3.0 specification.  You can find out more about Swagger at [https://swagger.io](https://swagger.io). In the third iteration of the pet store, we've switched to the design first approach! You can now help us improve the API whether it's by making changes to the definition itself or to the code. That way, with time, we can improve the API in general, and expose some of the new features in OAS3.  _If you're looking for the Swagger 2.0/OAS 2.0 version of Petstore, then click [here](https://editor.swagger.io/?url=https://petstore.swagger.io/v2/swagger.yaml). Alternatively, you can load via the `Edit > Load Petstore OAS 2.0` menu option!_  Some useful links: - [The Pet Store repository](https://github.com/swagger-api/swagger-petstore) - [The source API definition for the Pet Store](https://github.com/swagger-api/swagger-petstore/blob/master/src/main/resources/openapi.yaml)
 *
 * API version: 1.0.11
 * Contact: apiteam@swagger.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var pets map[int64]Pet
func AddPet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	data,err := ioutil.ReadAll(r.Body)
	if(err!=nil){
		log.Print(err.Error())
		return
	}//read err

	var newpet Pet 
	err = json.Unmarshal(data,&newpet)
	if err!=nil {
		log.Print(err.Error())
		return
	}//parsing err

	if newpet.Id==0 || newpet.Category.Id==0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		//log.Print(err.Error())
		return
	}//input err

	_,ok := pets[newpet.Id]
	if ok {
		w.WriteHeader(http.StatusBadRequest)
		//fmt.Fprint(w,"pet already exist")
		return
	}
	pets[newpet.Id] = newpet
	//pets = append(pets, newpet)

	fmt.Fprint(w, string(data))
	w.WriteHeader(http.StatusCreated)
}

func DeletePet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r) // HTTP요청의 url 경로에서 변수 추출에 사용
	petId, err := strconv.Atoi(vars["petId"]) // 문자열과 기본 자료형간의 형변환 패키지
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	for i,v := range pets{
		if v.Id == int64(petId){
			delete(pets,i)
			//pets = append(pets[:i], pets[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	
	w.WriteHeader(http.StatusBadRequest)
	// data, _ := json.Marshal(pet)
	// fmt.Fprint(w, string(data))
}

func FindPetsByStatus(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	status := values.Get("status")
	//api_key := r.Header.Get("Api_key")

	StatusPets := []Pet{}

	for _,v := range pets{
		if v.Status == status{
			StatusPets = append(StatusPets, v)
			//fmt.Fprintf(w,"Response:\n")
		}
	}

	if len(StatusPets)!=0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		data, _ := json.Marshal(StatusPets) //Serialize
		fmt.Fprint(w, string(data))	
	} else{
		w.WriteHeader(http.StatusNotFound) // no data
	}
}

func FindPetsByTags(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetPetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r) // HTTP요청의 url 경로에서 변수 추출에 사용
	petId, err := strconv.Atoi(vars["petId"]) // 문자열과 기본 자료형간의 형변환 패키지
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}


	pet,ok := pets[int64(petId)]
	if !ok{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(pet)
	if err==nil{
		fmt.Fprint(w, string(data))
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusBadRequest)	
}

func UpdatePet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UpdatePetWithForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetPetList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if len(pets)!=0 {
		data, _ := json.Marshal(pets)
		fmt.Fprint(w, string(data))	
		w.WriteHeader(http.StatusOK)
	} else{
		w.WriteHeader(http.StatusBadRequest)
	}
}

func ClearPet(w http.ResponseWriter, r *http.Request) {
	pets = make(map[int64]Pet)
	w.WriteHeader(http.StatusOK)
}

