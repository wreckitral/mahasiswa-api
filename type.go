package main

type Mahasiswa struct {
    Nama      string `json:"nama"`
    NIM       int `json:"nim"`
    TmptLahir string `json:"tmptLahir"`
    TglLahir  string `json:"tglLahir"`
    IPK       float64 `json:"ipk"`
    Predikat  string `json:"predikat"`
    Prodi     string `json:"prodi"`
    MasaStudi string `json:"masaStudi"`
    Suliet    int `json:"suliet"`
}

