package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Storage interface {
    GetRataRata() (float64, error)
    GetRataRataSuliet() (float64, error)
    GetMahasiswaMaxIpk() (*Mahasiswa, error)
    GetMahasiswaMinIpk() (*Mahasiswa, error)
    GetMahasiswaMaxSuliet() (*Mahasiswa, error)
    GetMahasiswaMinSuliet() (*Mahasiswa, error)
    GetPujianSangatMemuaskanMemuaskan() (int, int, int, error)
    GetAgregatMasaStudi() (map[string]int, error)
    GetDataUmurMahasiswa() ([]*DataUmur, error)
}

type MysqlStore struct {
    db *sql.DB
}

func NewMysqlStore() (*MysqlStore, error) {
    dbUsername := os.Getenv("DBUSERNAME")
    dbPassword := os.Getenv("DBPASS")
    dbHost := os.Getenv("DBHOST")
    dbPort := os.Getenv("DBPORT")
    dbName := os.Getenv("DBNAME")

    dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)

    db, err := sql.Open("mysql", dataSourceName)
    if err != nil {
      log.Fatal("Error connecting to database")
    }

    if err := db.Ping(); err != nil {
        return nil, err
    }

    return &MysqlStore{
        db: db,
    }, nil
}

func (s *MysqlStore) GetRataRata() (float64, error) {
    query := "SELECT AVG(ipk) FROM ds_wisuda_tibil"
    var rata float64

    err := s.db.QueryRow(query).Scan(&rata)
    if err != nil {
        return 0, err
    }

    return rata, nil
}

func (s *MysqlStore) GetRataRataSuliet() (float64, error) {
    query := "SELECT AVG(suliet) FROM ds_wisuda_tibil"
    var rataSuliet float64

    err := s.db.QueryRow(query).Scan(&rataSuliet)
    if err != nil {
        return 0, err
    }

    return rataSuliet, nil
}

func (s *MysqlStore) GetMahasiswaMaxIpk() (*Mahasiswa, error) {
    query := `
        SELECT * FROM ds_wisuda_tibil 
        WHERE ipk = (SELECT MAX(ipk) FROM ds_wisuda_tibil)
    ` 
    row, err := s.db.Query(query)
    if err != nil {
        return nil, err
    }
    
    for row.Next() {
        return scanIntoMahasiswa(row)
    }

    return nil, fmt.Errorf("Database error")
}

func (s *MysqlStore) GetMahasiswaMinIpk() (*Mahasiswa, error) {
    query := `
        SELECT * FROM ds_wisuda_tibil 
        WHERE ipk = (SELECT MIN(ipk) FROM ds_wisuda_tibil)
    ` 

    row, err := s.db.Query(query)
    if err != nil {
        return nil, err
    }
    
    for row.Next() {
        return scanIntoMahasiswa(row)
    }

    return nil, fmt.Errorf("Database error")
}

func (s *MysqlStore) GetMahasiswaMaxSuliet() (*Mahasiswa, error) {
    query := `
        SELECT * FROM ds_wisuda_tibil 
        WHERE suliet = (SELECT MAX(suliet) FROM ds_wisuda_tibil)
    ` 

    row, err := s.db.Query(query)
    if err != nil {
        return nil, err
    }
    
    for row.Next() {
        return scanIntoMahasiswa(row)
    }

    return nil, fmt.Errorf("Database error")
}

func (s *MysqlStore) GetMahasiswaMinSuliet() (*Mahasiswa, error) {
    query := `
        SELECT * FROM ds_wisuda_tibil 
        WHERE suliet = (SELECT MIN(suliet) FROM ds_wisuda_tibil)
    ` 

    row, err := s.db.Query(query)
    if err != nil {
        return nil, err
    }
    
    for row.Next() {
        return scanIntoMahasiswa(row)
    }

    return nil, fmt.Errorf("Database error")
}

func (s *MysqlStore) GetPujianSangatMemuaskanMemuaskan() (int, int, int, error) {
    query := `
        SELECT predikat, COUNT(*)
        FROM ds_wisuda_tibil
        GROUP BY predikat
    `

    rows, err := s.db.Query(query)
    if err != nil {
        return 0, 0, 0, err
    }

    counts := map[string]int{
        "Dengan Pujian":     0,
        "Sangat Memuaskan":  0,
        "Memuaskan":         0,
    }

    for rows.Next() {
        var predikat string
        var count int

        if err := rows.Scan(&predikat, &count); err != nil {
            return 0, 0, 0, err
        }

        counts[predikat] = count
    }

    if err = rows.Err(); err != nil {
        return 0, 0, 0, err
    }

    return counts["Dengan Pujian"], counts["Sangat Memuaskan"], counts["Memuaskan"], nil
}

func (s *MysqlStore) GetAgregatMasaStudi() (map[string]int, error) {
    query := `
    SELECT masastudi
    FROM ds_wisuda_tibil;
    `

    rows, err := s.db.Query(query)
    if err != nil {
        return nil, err
    }

    counts := make(map[string]int)

    for rows.Next() {
        var masastudi string
        if err := rows.Scan(&masastudi); err != nil {
            return nil, err
        }

        counts[masastudi]++
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return counts, nil
}

func (s *MysqlStore) GetDataUmurMahasiswa() ([]*DataUmur, error) {
    query := `
        SELECT CONCAT('NIM : ', nim, ' – ', nama, ' – ', TIMESTAMPDIFF(YEAR, tgllahir, CURDATE()), ' Tahun') AS description
            FROM ds_wisuda_tibil
    `
    rows, err := s.db.Query(query)
    if err != nil {
        return nil, err
    }

    dataUmurs := []*DataUmur{}

    for rows.Next() {
        dataUmur, err := scanIntoDataUmur(rows)
        if err != nil {
            return nil, err
        }
        dataUmurs = append(dataUmurs, dataUmur)
    }

    return dataUmurs, nil
}

