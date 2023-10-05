package route

import (
	"net/http"
	"uts/database"
	"uts/models"

	"github.com/gofiber/fiber/v2"
)

// buatlah endpoint Insert Data sesuai dengan Colection Postman
func InsertData(c *fiber.Ctx) error {

	//Tulis Jawaban Code di Sini :))
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user := models.User{
		Nama:     data["nama"],
		Username: data["username"],
		Email:    data["email"],
		Password: data["password"],
	}

	database.DB.Create(&user)

	return c.JSON(fiber.Map{
		"pesan": "Data telah berhasil di tambahkan",
		"code":  http.StatusOK,
		"data":  user,
	})
}

// Lengkapi Code Berikut untuk untuk Mengambil data untuk semua user - user
func GetAllData(c *fiber.Ctx) error {

	users := []models.User{}

	database.DB.Find(&users)

	return c.JSON(fiber.Map{
		"pesan": "Berhasil mengambil semua data user",
		"code":  http.StatusOK,
		"data":  users,
	})

}

//Lengkapi Code berikut Untuk Mengambil data dari id_user berdasarkan Parameter

func GetUserByid(c *fiber.Ctx) error {
	user := models.User{}
	id_user := c.Params("id_user")
	database.DB.Where("id_user = ?", id_user).Find(&user)

	return c.JSON(fiber.Map{
		"pesan": "Berhasil mengambil data user dengan id = " + id_user,
		"code":  http.StatusOK,
		"data":  user,
	})
}

func Delete(c *fiber.Ctx) error {

	var user models.User

	id_user := c.Params("id_user")

	database.DB.Where("id_user = ?", id_user).Delete(user)

	return c.JSON(fiber.Map{
		"pesan": "Data telah di hapus",
		"code":  http.StatusOK,
		"data":  nil,
	})
}

// mengupdate data user berdasarkan id_user yang di tempatkan di parameter
func Update(c *fiber.Ctx) error {

	id_user := c.Params("id_user")

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	//data yang di ubah
	//membuat variable user berdasarkan model user
	var user models.User

	update := models.User{
		Nama:     data["nama"],
		Username: data["username"],
		Email:    data["email"],
		Password: data["password"],
	}
	//mengambil database untuk di update

	database.DB.Model(&user).Where("id_user = ?", id_user).Updates(update)

	return c.JSON(fiber.Map{
		"pesan": "Data User telah di Update",
		"code":  http.StatusOK,
		"data":  update,
	})
}
