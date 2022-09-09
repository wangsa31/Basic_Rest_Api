package controller

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Pepole struct {
	Id         int
	Frisrtname string
	Lastname   string
	Address    string
	No_card    int
}

// insert data

type InsertPepole struct {
	Frisrtname string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Address    string `json:"address"`
	No_card    int    `json:"no_card"`
}



func Connection() {
	var err error
	db, err = sql.Open("mysql", "root:@/latihan2")

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

}

func ShowData(c *gin.Context) {
	ress, err := db.Query("select * from pepole")

	if err != nil {
		log.Fatal(err)
	}

	for ress.Next() {
		var pepoles Pepole
		var pepoless []Pepole
		err = ress.Scan(&pepoles.Id, &pepoles.Frisrtname, &pepoles.Lastname, &pepoles.Address, &pepoles.No_card)

		pepoless = append(pepoless, pepoles)
		if err != nil {
			c.JSON(404, gin.H{
				"message": "ERROR SHOW TABLE",
			})
			// log.Fatal(err)
		} else {
			c.JSON(200, gin.H{
				"data": pepoless,
			})

		}

	}

}

func ShowFindId(c *gin.Context) {
	id := c.Param("id")
	ress, err := db.Query("select * from pepole where id= ?", id)

	if err != nil {
		log.Fatal(err)
	}

	for ress.Next() {
		var pepoles Pepole
		var pepoless []Pepole
		err = ress.Scan(&pepoles.Id, &pepoles.Frisrtname, &pepoles.Lastname, &pepoles.Address, &pepoles.No_card)

		pepoless = append(pepoless, pepoles)
		if err != nil {
			c.JSON(404, gin.H{
				"message": "ERROR SHOW TABLE",
			})
			// log.Fatal(err)
		} else {
			c.JSON(200, gin.H{
				"data": pepoless,
			})

		}

	}
}

func InsertData(c *gin.Context) {
	Inserts := InsertPepole{}

	err := c.ShouldBindJSON(&Inserts)

	if err != nil {
		c.JSON(404, gin.H{
			"message": err,
		})
	}

	_, err = db.Exec("INSERT INTO pepole values (?,?,?,?,?)", "", Inserts.Frisrtname, Inserts.Lastname, Inserts.Address, Inserts.No_card)

	if err != nil {
		c.JSON(404, gin.H{
			"message": err,
		})
	}

	c.JSON(200, gin.H{
		"message": "Insert data successfully",
	})

}

func EditData(c *gin.Context) {
	id := c.Param("id")
	c.JSON(300, gin.H{
		"test": id,
	})

	var UpdateData InsertPepole

	err := c.ShouldBindJSON(&UpdateData)

	if err != nil {
		c.JSON(404, gin.H{
			"message": err,
		})
	}

	_, err = db.Exec("UPDATE pepole SET firstname= ?, lastname=?, address=?, no_card=? WHERE id=? ", UpdateData.Frisrtname, UpdateData.Lastname, UpdateData.Address, UpdateData.No_card, id)

	if err != nil {
		log.Fatal(err)
	} else {
		c.JSON(200, gin.H{
			"message": "Success update data",
		})
	}

}

func Delete(c *gin.Context) {
	id := c.Param("id")
	_, err := db.Exec("DELETE FROM pepole WHERE id = ? ", id)

	if err != nil {
		c.JSON(200, gin.H{
			"message": err,
		})
	}
	c.JSON(200, gin.H{
		"message": "Delete Succesfully",
	})
}
