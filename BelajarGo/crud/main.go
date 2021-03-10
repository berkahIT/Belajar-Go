package main

import (
	"crud/model"
	"fmt"
	
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
 
  func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/db_curd?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil{   
		fmt.Println("Koneksi Gagal")
	}

	// fmt.Println("Koneksi Berhasil")
 
	//dosen
	//insert
	// data := model.Dosen{
	// 	Nama: "Edwin Farid",
	// 	Nidn: "2019820938",
	// 	Ttl: "2002/09/22",
	// 	Alamat: "Selong",
	// }
	// db.Create(&data)

	//update
	// db.Model(&data).Where("id_dosen =?",1).Update("nidn","200602071")

	///delete
	// db.Where("id_dosen =?",1).Delete(&data)

	//read
	// var dosen []model.Dosen
	// db.Find(&dosen)
	// for _, m := range dosen{
	// 	fmt.Println("Nama = "+m.Nama)
	// 	fmt.Println("Nidn = "+m.Nidn)
	// 	fmt.Println("Ttl = "+m.Ttl)
	// 	fmt.Println("Alamat = "+m.Alamat)
	// 	fmt.Println("==============")
	// }

	//jalanin server
	r := gin.Default()
	r.GET("/getDosen", func(c *gin.Context) {
		var dosen []model.Dosen
		db.Find(&dosen)
		c.JSON(http.StatusOK, &dosen)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// -------------------------------------------------------------------------------------------------------
	// Mahasiswa
	// insert
	// data := model.Mahasiswa{
	// 	Nama: "farid", 
	// 	Nim: "200602008", 
	// 	Alamat: "pancor",
	// }

	// db.Create(&data) 

	// update
	// db.Model(&data).Where("id =?",1).Update("nama", "Edwin Farid") 

	// delete
	// db.Where("id =?",2).Delete(&model.Mahasiswa{}) 

	//read
	// var mahasiswa []model.Mahasiswa
	// db.Find(&mahasiswa)
	// for _, m := range mahasiswa{
	// 	fmt.Println("Nama = "+m.Nama)
	// 	fmt.Println("Nim = "+m.Nim)
	// 	fmt.Println("Alamat = "+m.Alamat)
	// 	fmt.Println("==============")
	// }
	


  }