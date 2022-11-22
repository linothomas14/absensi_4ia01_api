# ABSENSI API 

API ini ditujukan untuk Project Bisnis Informatika pada Aplikasi Absensi 4IA01

## Table of Contents

* [Setup](#setup)
* [Routes](#routes)
* [API Documentation](#api-documentation)
* [Contributor](#contributor)

## Setup

To run this project, you must have Golang and postgreSQL

- go run main.go

## Routes

| HTTP METHOD | POST            | GET       | PUT         | DELETE |
| ----------- | :-------: | :------:  | :------:  | :------: |
| /api/mahasiswa       | - | List of Mahasiswa | - | - |
| /api/mahasiswa/`<string:npm>`       | - | Detail of Mahasiswa | - | - |
| /api/presensi | Add Presensi | List of Presensi | - | Delete Presensi |


## API Documentation 
### List of Endpoints
* [Mahasiswa](#Mahasiswa)
    * [Get All Mahasiswa](#get-all-mahasiswa)
    * [Get Mahasiswa by NPM](#get-mahasiswa-by-npm)
* [Presensi](#presensi)
    * [Get Presensi](#get-presensi)
    * [Add Presensi](#add-presensi)
    * [Delete Presensi](#delete-presensi)

## Mahasiswa
### Get All Mahasiswa
* Method : GET
* URL : `/api/mahasiswa` 
* Request body : -   
* Response body  :
```json
{
    "message": "OK",
    "data": [
        {
            "id": 1,
            "npm": "50419135",
            "nama": "ADELIYAADHIMI RISKY NURROCHMAH" 
        },

        {
            "id": 2,
            "npm": "50419339",
            "nama": "AHMAD NAUFAL FADHIL" 
        },

        {
            "id": 3,
            "npm": "50419515",
            "nama": "ALIFA NUR RIZQILLAH" 
        }
    ]
}
```
### Get Mahasiswa by Npm
* Method : GET
* URL : `/api/mahasiswa/<string:npm>`    
* Request body : -
* Response body  :
```json 
{
    "message": "OK",
    "data": {
        "id": 37,
        "npm": "56419777",
        "nama": "YULYANO THOMAS DJAYA" 
    }
}
```

## Presensi

### Add Presensi
* Method : POST
* URL : `/api/presensi`    
* Request body :
```json
{
    "npm": "56419764",
    "minggu": 1, 
    "matkul":"Bisnis Informatika"
}
```
* Response body :
```json
{
    "message": "OK",
    "data": {
        "id": 1,
        "npm": "56419764",
        "matkul": "Bisnis Informatika",
        "minggu": 1 
    }
}
```
### Get Presensi
* Method : GET
* URL : `/api/presensi` 
* Request body :
```json
{
    "minggu": 1, 
    "matkul":"Bisnis Informatika"
}
```   
* Response body  :
```json
{
    "message": "OK",
    "data": {
        "matkul": "Bisnis Informatika",
        "minggu": 1,
        "mahasiswa" : [
            {
                "nama" : "YULYANO THOMAS DJAYA",
                "npm" : "56419764"
            }
        ]
    }
}
```
### Delete Presensi
* Method : DELETE
* URL : `/api/presensi`    
* Request body :
```json
{
    "npm": "56419764",
    "minggu": 1, 
    "matkul":"Bisnis Informatika"
}
```  
* Response body:
```json
{
    "message": "Data berhasil dihapus",
    "data": ""
}
```

## Contributor
- Dwi Pertiwi Ani - 51419933
- Muhammad Arief Rubbyansyah - 54419032
- Shafiah Qonita - 56419004
- Yulyano Thomas Djaya - 56419764
